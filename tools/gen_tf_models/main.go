// gen_tf_models generates a tf_models.go file from a models.go file.
//
// Invoked via go:generate directives in models.go files, e.g.:
//
//	//go:generate go run github.com/leanspace/terraform-provider-leanspace/tools/gen_tf_models -struct CommandQueue
//
// Supported API → TF field type pairs:
//
//	string  / *string          → types.String
//	bool    / *bool            → types.Bool
//	int     / *int             → types.Int64
//	float64 / *float64         → types.Float64
//	[]string                   → []types.String
//	[]general_objects.KeyValue → []general_objects.KeyValueTF
//	general_objects.AuditModel → general_objects.AuditModelTF  (embedded)
//	*LocalStruct               → *LocalStructTF   (struct defined in same file)
//	[]LocalStruct              → []LocalStructTF  (struct defined in same file)
//
// When nested local struct fields are present the generator emits full ToTF/ToAPI
// bodies and nil-safe helper functions. Nested structs themselves must be flat.
// Generic structs are rejected.
package main

import (
	"bytes"
	"flag"
	"fmt"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"log"
	"os"
	"strings"
	"unicode"
)

// kind constants for fieldInfo
const (
	kindEmbed      = "embed"
	kindString     = "string"
	kindPtrString  = "*string"
	kindBool       = "bool"
	kindPtrBool    = "*bool"
	kindInt        = "int"
	kindPtrInt     = "*int"
	kindFloat64    = "float64"
	kindPtrFloat64 = "*float64"
	kindStrings    = "[]string"
	kindKeyValues  = "[]KeyValue"
	kindPtrLocal   = "*local"
	kindSliceLocal = "[]local"
)

// fieldInfo holds derived metadata for a single TF struct field.
type fieldInfo struct {
	goName    string // Go field name (same in API and TF struct)
	tfType    string // TF type expression, e.g. "types.String", "*ResourceFunctionFormulaTF"
	tfsdk     string // tfsdk tag value
	kind      string // one of the kind* constants
	localName string // for *local / []local: the local struct Go name
	isEmbed   bool   // true for embedded AuditModelTF
}

func main() {
	structName := flag.String("struct", "", "Name of the API struct to generate TF model for (default: first exported struct)")
	outFile := flag.String("out", "tf_models.go", "Output file name")
	flag.Parse()

	gofile := os.Getenv("GOFILE")
	if gofile == "" {
		gofile = "models.go"
	}
	gopkg := os.Getenv("GOPACKAGE")

	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, gofile, nil, 0)
	if err != nil {
		log.Fatalf("parsing %s: %v", gofile, err)
	}
	if gopkg == "" {
		gopkg = f.Name.Name
	}

	target := *structName
	if target == "" {
		target = firstExportedStruct(f)
	}
	if target == "" {
		log.Fatalf("no exported struct found in %s", gofile)
	}

	src, err := generateTFModel(gopkg, f, target)
	if err != nil {
		log.Fatalf("generating TF model for %s: %v", target, err)
	}

	formatted, err := format.Source(src)
	if err != nil {
		log.Fatalf("formatting generated output:\n%s\nerror: %v", src, err)
	}

	if err := os.WriteFile(*outFile, formatted, 0644); err != nil {
		log.Fatalf("writing %s: %v", *outFile, err)
	}
	log.Printf("generated %sTF → %s", target, *outFile)
}

// firstExportedStruct returns the name of the first exported struct in f.
func firstExportedStruct(f *ast.File) string {
	for _, decl := range f.Decls {
		gd, ok := decl.(*ast.GenDecl)
		if !ok {
			continue
		}
		for _, spec := range gd.Specs {
			ts, ok := spec.(*ast.TypeSpec)
			if !ok {
				continue
			}
			if _, ok := ts.Type.(*ast.StructType); ok && ast.IsExported(ts.Name.Name) {
				return ts.Name.Name
			}
		}
	}
	return ""
}

// allLocalStructs returns all non-generic exported struct names and their AST nodes.
func allLocalStructs(f *ast.File) map[string]*ast.StructType {
	result := make(map[string]*ast.StructType)
	for _, decl := range f.Decls {
		gd, ok := decl.(*ast.GenDecl)
		if !ok {
			continue
		}
		for _, spec := range gd.Specs {
			ts, ok := spec.(*ast.TypeSpec)
			if !ok {
				continue
			}
			st, ok := ts.Type.(*ast.StructType)
			if !ok {
				continue
			}
			if ts.TypeParams != nil && len(ts.TypeParams.List) > 0 {
				continue // skip generics
			}
			result[ts.Name.Name] = st
		}
	}
	return result
}

func generateTFModel(pkg string, f *ast.File, structName string) ([]byte, error) {
	locals := allLocalStructs(f)

	if _, exists := locals[structName]; !exists {
		return nil, fmt.Errorf("struct %s not found", structName)
	}

	// Reject generics for the target struct explicitly
	for _, decl := range f.Decls {
		gd, ok := decl.(*ast.GenDecl)
		if !ok {
			continue
		}
		for _, spec := range gd.Specs {
			ts, ok := spec.(*ast.TypeSpec)
			if !ok || ts.Name.Name != structName {
				continue
			}
			if ts.TypeParams != nil && len(ts.TypeParams.List) > 0 {
				return nil, fmt.Errorf("%s is generic; write tf_models.go by hand", structName)
			}
		}
	}

	rootST := locals[structName]
	fields, err := parseFields(rootST, locals)
	if err != nil {
		return nil, err
	}

	// Collect nested local struct names referenced by root
	nestedNames := collectNestedNames(fields)

	if len(nestedNames) == 0 {
		return renderFileFlat(pkg, structName, fields)
	}

	// Parse nested struct fields (they must be flat — no further local references)
	nestedFields := make(map[string][]fieldInfo)
	for name := range nestedNames {
		st, exists := locals[name]
		if !exists {
			return nil, fmt.Errorf("referenced local struct %s not found in file", name)
		}
		nf, err := parseFields(st, nil) // nil = no local struct references allowed in nested
		if err != nil {
			return nil, fmt.Errorf("nested struct %s: %w", name, err)
		}
		nestedFields[name] = nf
	}

	return renderFileFull(pkg, structName, fields, nestedFields)
}

func collectNestedNames(fields []fieldInfo) map[string]bool {
	result := make(map[string]bool)
	for _, f := range fields {
		if f.kind == kindPtrLocal || f.kind == kindSliceLocal {
			result[f.localName] = true
		}
	}
	return result
}

// parseFields parses all struct fields. locals is the set of known local struct
// names; pass nil to disallow local struct references (for nested structs).
func parseFields(st *ast.StructType, locals map[string]*ast.StructType) ([]fieldInfo, error) {
	var result []fieldInfo
	for _, field := range st.Fields.List {
		if len(field.Names) == 0 {
			fi, err := parseEmbedded(field)
			if err != nil {
				return nil, err
			}
			result = append(result, fi)
			continue
		}
		fi, err := parseNamedField(field, locals)
		if err != nil {
			return nil, fmt.Errorf("field %s: %w", field.Names[0].Name, err)
		}
		result = append(result, fi)
	}
	return result, nil
}

func parseEmbedded(field *ast.Field) (fieldInfo, error) {
	sel, ok := field.Type.(*ast.SelectorExpr)
	if !ok {
		return fieldInfo{}, fmt.Errorf("unsupported embedded type: %T", field.Type)
	}
	pkg, ok := sel.X.(*ast.Ident)
	if !ok {
		return fieldInfo{}, fmt.Errorf("unsupported embedded selector: %T", sel.X)
	}
	if pkg.Name == "general_objects" && sel.Sel.Name == "AuditModel" {
		return fieldInfo{isEmbed: true, kind: kindEmbed}, nil
	}
	return fieldInfo{}, fmt.Errorf("unsupported embedded type %s.%s (only general_objects.AuditModel supported)", pkg.Name, sel.Sel.Name)
}

func parseNamedField(field *ast.Field, locals map[string]*ast.StructType) (fieldInfo, error) {
	goName := field.Names[0].Name
	kind, localName, tfType, err := resolveType(field.Type, locals)
	if err != nil {
		return fieldInfo{}, err
	}
	tfsdk, err := extractTFSDKTag(goName, field.Tag)
	if err != nil {
		return fieldInfo{}, err
	}
	return fieldInfo{goName: goName, tfType: tfType, tfsdk: tfsdk, kind: kind, localName: localName}, nil
}

// resolveType maps an AST type expression to (kind, localName, tfTypeStr, error).
func resolveType(expr ast.Expr, locals map[string]*ast.StructType) (string, string, string, error) {
	switch t := expr.(type) {
	case *ast.Ident:
		switch t.Name {
		case "string":
			return kindString, "", "types.String", nil
		case "bool":
			return kindBool, "", "types.Bool", nil
		case "int":
			return kindInt, "", "types.Int64", nil
		case "float64":
			return kindFloat64, "", "types.Float64", nil
		}
	case *ast.StarExpr:
		inner, ok := t.X.(*ast.Ident)
		if !ok {
			return "", "", "", fmt.Errorf("unsupported pointer target: %T", t.X)
		}
		switch inner.Name {
		case "string":
			return kindPtrString, "", "types.String", nil
		case "bool":
			return kindPtrBool, "", "types.Bool", nil
		case "int":
			return kindPtrInt, "", "types.Int64", nil
		case "float64":
			return kindPtrFloat64, "", "types.Float64", nil
		}
		if locals != nil {
			if _, ok := locals[inner.Name]; ok {
				return kindPtrLocal, inner.Name, "*" + inner.Name + "TF", nil
			}
		}
	case *ast.ArrayType:
		switch el := t.Elt.(type) {
		case *ast.Ident:
			if el.Name == "string" {
				return kindStrings, "", "[]types.String", nil
			}
			if locals != nil {
				if _, ok := locals[el.Name]; ok {
					return kindSliceLocal, el.Name, "[]" + el.Name + "TF", nil
				}
			}
		case *ast.SelectorExpr:
			pkg, ok := el.X.(*ast.Ident)
			if !ok {
				break
			}
			if pkg.Name == "general_objects" && el.Sel.Name == "KeyValue" {
				return kindKeyValues, "", "[]general_objects.KeyValueTF", nil
			}
		}
	}
	return "", "", "", fmt.Errorf("unsupported type — write tf_models.go by hand or extend the generator")
}

// extractTFSDKTag derives the tfsdk tag from the json tag.
func extractTFSDKTag(goName string, tag *ast.BasicLit) (string, error) {
	if tag != nil {
		raw := strings.Trim(tag.Value, "`")
		for _, pair := range strings.Fields(raw) {
			if !strings.HasPrefix(pair, `json:"`) {
				continue
			}
			val := strings.TrimPrefix(pair, `json:"`)
			val = strings.TrimSuffix(val, `"`)
			val = strings.SplitN(val, ",", 2)[0]
			if val != "" && val != "-" {
				return camelToSnake(val), nil
			}
		}
	}
	return camelToSnake(goName), nil
}

// camelToSnake converts camelCase to snake_case.
func camelToSnake(s string) string {
	var b strings.Builder
	runes := []rune(s)
	for i, r := range runes {
		if i > 0 && unicode.IsUpper(r) && unicode.IsLower(runes[i-1]) {
			b.WriteByte('_')
		}
		b.WriteRune(unicode.ToLower(r))
	}
	return b.String()
}

func lcFirst(s string) string {
	if s == "" {
		return s
	}
	return strings.ToLower(s[:1]) + s[1:]
}

// ---- Flat (no nested local structs) ----

func renderFileFlat(pkg, structName string, fields []fieldInfo) ([]byte, error) {
	tfName := structName + "TF"

	needTypes := false
	for _, f := range fields {
		if !f.isEmbed {
			needTypes = true
			break
		}
	}

	var buf bytes.Buffer
	fmt.Fprintf(&buf, "// Code generated by gen_tf_models. DO NOT EDIT.\n")
	fmt.Fprintf(&buf, "package %s\n\n", pkg)
	fmt.Fprintf(&buf, "import (\n")
	if needTypes {
		fmt.Fprintf(&buf, "\t\"github.com/hashicorp/terraform-plugin-framework/types\"\n")
	}
	fmt.Fprintf(&buf, "\t\"github.com/leanspace/terraform-provider-leanspace/helper/general_objects\"\n")
	fmt.Fprintf(&buf, ")\n\n")

	writeStructDecl(&buf, tfName, fields)

	fmt.Fprintf(&buf, "func (x *%s) ToTF() any {\n", structName)
	fmt.Fprintf(&buf, "\treturn general_objects.ReflectToTF[%s](x)\n", tfName)
	fmt.Fprintf(&buf, "}\n\n")

	fmt.Fprintf(&buf, "func (tf *%s) ToAPI() any {\n", tfName)
	fmt.Fprintf(&buf, "\treturn general_objects.ReflectFromTF[%s](tf)\n", structName)
	fmt.Fprintf(&buf, "}\n")

	return buf.Bytes(), nil
}

// ---- Full (with nested local structs) ----

func renderFileFull(pkg, structName string, fields []fieldInfo, nestedFields map[string][]fieldInfo) ([]byte, error) {
	tfName := structName + "TF"

	var buf bytes.Buffer
	fmt.Fprintf(&buf, "// Code generated by gen_tf_models. DO NOT EDIT.\n")
	fmt.Fprintf(&buf, "package %s\n\n", pkg)
	fmt.Fprintf(&buf, "import (\n")
	fmt.Fprintf(&buf, "\t\"github.com/hashicorp/terraform-plugin-framework/types\"\n")
	fmt.Fprintf(&buf, "\t\"github.com/leanspace/terraform-provider-leanspace/helper\"\n")
	fmt.Fprintf(&buf, "\t\"github.com/leanspace/terraform-provider-leanspace/helper/general_objects\"\n")
	fmt.Fprintf(&buf, ")\n\n")

	// Nested TF struct declarations
	for name, nf := range nestedFields {
		writeStructDecl(&buf, name+"TF", nf)
	}

	// Root TF struct
	writeStructDecl(&buf, tfName, fields)

	// Root ToTF
	fmt.Fprintf(&buf, "func (x *%s) ToTF() any {\n", structName)
	fmt.Fprintf(&buf, "\treturn &%s{\n", tfName)
	for _, f := range fields {
		if f.isEmbed {
			fmt.Fprintf(&buf, "\t\tAuditModelTF: general_objects.AuditModelToTF(&x.AuditModel),\n")
		} else {
			fmt.Fprintf(&buf, "\t\t%s: %s,\n", f.goName, toTFExpr(f, "x"))
		}
	}
	fmt.Fprintf(&buf, "\t}\n}\n\n")

	// Root ToAPI
	fmt.Fprintf(&buf, "func (tf *%s) ToAPI() any {\n", tfName)
	fmt.Fprintf(&buf, "\treturn &%s{\n", structName)
	for _, f := range fields {
		if f.isEmbed {
			fmt.Fprintf(&buf, "\t\tAuditModel: general_objects.AuditModelFromTF(tf.AuditModelTF),\n")
		} else {
			fmt.Fprintf(&buf, "\t\t%s: %s,\n", f.goName, fromTFExpr(f, "tf"))
		}
	}
	fmt.Fprintf(&buf, "\t}\n}\n\n")

	// Helper functions for each nested local struct
	for name := range nestedFields {
		lc := lcFirst(name)

		fmt.Fprintf(&buf, "func %sToTF(x *%s) *%sTF {\n", lc, name, name)
		fmt.Fprintf(&buf, "\tif x == nil {\n\t\treturn nil\n\t}\n")
		fmt.Fprintf(&buf, "\treturn general_objects.ReflectToTF[%sTF](x)\n", name)
		fmt.Fprintf(&buf, "}\n\n")

		fmt.Fprintf(&buf, "func %sFromTF(tf *%sTF) *%s {\n", lc, name, name)
		fmt.Fprintf(&buf, "\tif tf == nil {\n\t\treturn nil\n\t}\n")
		fmt.Fprintf(&buf, "\treturn general_objects.ReflectFromTF[%s](tf)\n", name)
		fmt.Fprintf(&buf, "}\n\n")

		fmt.Fprintf(&buf, "func %sSliceToTF(xs []%s) []%sTF {\n", lc, name, name)
		fmt.Fprintf(&buf, "\tresult := make([]%sTF, len(xs))\n", name)
		fmt.Fprintf(&buf, "\tfor i := range xs {\n")
		fmt.Fprintf(&buf, "\t\tresult[i] = *general_objects.ReflectToTF[%sTF](&xs[i])\n", name)
		fmt.Fprintf(&buf, "\t}\n")
		fmt.Fprintf(&buf, "\treturn result\n}\n\n")

		fmt.Fprintf(&buf, "func %sSliceFromTF(tfs []%sTF) []%s {\n", lc, name, name)
		fmt.Fprintf(&buf, "\tresult := make([]%s, len(tfs))\n", name)
		fmt.Fprintf(&buf, "\tfor i := range tfs {\n")
		fmt.Fprintf(&buf, "\t\tresult[i] = *general_objects.ReflectFromTF[%s](&tfs[i])\n", name)
		fmt.Fprintf(&buf, "\t}\n")
		fmt.Fprintf(&buf, "\treturn result\n}\n\n")
	}

	return buf.Bytes(), nil
}

// writeStructDecl writes a TF struct declaration to buf.
func writeStructDecl(buf *bytes.Buffer, tfName string, fields []fieldInfo) {
	fmt.Fprintf(buf, "type %s struct {\n", tfName)
	for _, f := range fields {
		if f.isEmbed {
			fmt.Fprintf(buf, "\tgeneral_objects.AuditModelTF\n")
		} else {
			fmt.Fprintf(buf, "\t%s %s `tfsdk:\"%s\"`\n", f.goName, f.tfType, f.tfsdk)
		}
	}
	fmt.Fprintf(buf, "}\n\n")
}

// toTFExpr returns the expression that converts an API field to its TF value.
func toTFExpr(fi fieldInfo, src string) string {
	ref := src + "." + fi.goName
	switch fi.kind {
	case kindString:
		return "helper.TFStringValue(" + ref + ")"
	case kindPtrString:
		return "helper.TFStringPtrValue(" + ref + ")"
	case kindBool:
		return "helper.TFBoolValue(" + ref + ")"
	case kindPtrBool:
		return "helper.TFBoolPtrValue(" + ref + ")"
	case kindInt:
		return "helper.TFInt64Value(" + ref + ")"
	case kindPtrInt:
		return "helper.TFIntPtrValue(" + ref + ")"
	case kindFloat64:
		return "helper.TFFloat64Value(" + ref + ")"
	case kindPtrFloat64:
		return "helper.TFFloat64PtrValue(" + ref + ")"
	case kindStrings:
		return "helper.TFStringsValue(" + ref + ")"
	case kindKeyValues:
		return "general_objects.KeyValuesToTF(" + ref + ")"
	case kindPtrLocal:
		return lcFirst(fi.localName) + "ToTF(" + ref + ")"
	case kindSliceLocal:
		return lcFirst(fi.localName) + "SliceToTF(" + ref + ")"
	}
	return "/* unknown */"
}

// fromTFExpr returns the expression that converts a TF field back to its API value.
func fromTFExpr(fi fieldInfo, src string) string {
	ref := src + "." + fi.goName
	switch fi.kind {
	case kindString:
		return "helper.FromTFString(" + ref + ")"
	case kindPtrString:
		return "helper.FromTFStringPtr(" + ref + ")"
	case kindBool:
		return "helper.FromTFBool(" + ref + ")"
	case kindPtrBool:
		return "helper.FromTFBoolPtr(" + ref + ")"
	case kindInt:
		return "helper.FromTFInt64(" + ref + ")"
	case kindPtrInt:
		return "helper.FromTFIntPtr(" + ref + ")"
	case kindFloat64:
		return "helper.FromTFFloat64(" + ref + ")"
	case kindPtrFloat64:
		return "helper.FromTFFloat64Ptr(" + ref + ")"
	case kindStrings:
		return "helper.FromTFStrings(" + ref + ")"
	case kindKeyValues:
		return "general_objects.KeyValuesFromTF(" + ref + ")"
	case kindPtrLocal:
		return lcFirst(fi.localName) + "FromTF(" + ref + ")"
	case kindSliceLocal:
		return lcFirst(fi.localName) + "SliceFromTF(" + ref + ")"
	}
	return "/* unknown */"
}
