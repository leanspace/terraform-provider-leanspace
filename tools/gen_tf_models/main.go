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
//	map[string]string          → map[string]types.String
//	general_objects.AuditModel → general_objects.AuditModelTF  (embedded)
//	*LocalStruct               → *LocalStructTF   (struct defined in same file)
//	[]LocalStruct              → []LocalStructTF  (struct defined in same file)
//	LocalStruct (tf:"object")  → types.Object     (for Computed-only SingleNestedAttribute)
//	[]LocalStruct (tf:"list")  → types.List       (for Computed-only ListNestedAttribute)
//
// When a local struct field is annotated with tf:"object" (e.g.
//
//	Credentials Credentials `json:"credentials" tf:"object"`
//
// ) the generator emits types.Object instead of *CredentialsTF. This is required
// when the corresponding schema attribute is Computed-only (no Required/Optional),
// because Terraform sets purely-computed attributes to unknown during plan and
// only types.Object can represent unknown object values.
//
// Similarly, tf:"list" on a []LocalStruct field emits types.List instead of
// []LocalStructTF, for the same reason (Computed-only ListNestedAttribute).
//
// When nested local struct fields are present the generator emits full ToTF/ToAPI
// bodies and nil-safe helper functions. Nesting is resolved recursively: a nested
// struct may itself reference further local structs. Generic structs are rejected.
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
	kindEmbed           = "embed"
	kindString          = "string"
	kindPtrString       = "*string"
	kindBool            = "bool"
	kindPtrBool         = "*bool"
	kindInt             = "int"
	kindPtrInt          = "*int"
	kindFloat64         = "float64"
	kindPtrFloat64      = "*float64"
	kindStrings         = "[]string"
	kindKeyValues       = "[]KeyValue"
	kindMapStringString = "map[string]string"
	kindLocal           = "local"
	kindPtrLocal        = "*local"
	kindSliceLocal      = "[]local"
	kindObjectLocal     = "object_local" // local struct annotated with tf:"object"
	kindListLocal       = "list_local"   // []local struct annotated with tf:"list"
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

	// Collect object-local struct names (tf:"object" annotated, handled as types.Object).
	objectLocalNames := collectObjectLocalNames(fields)

	// Transitively collect all nested local struct references (recursive).
	// kindObjectLocal and kindListLocal fields are excluded from this traversal.
	nestedOrder, nestedFields, nestedComplex, err := collectAllNested(fields, locals)
	if err != nil {
		return nil, err
	}

	// Collect list-local struct names (tf:"list" annotated, handled as types.List).
	// Must happen after collectAllNested so nested struct fields are available.
	listLocalNames := collectAllListLocalNames(fields, nestedFields)

	if len(nestedOrder) == 0 && len(objectLocalNames) == 0 && len(listLocalNames) == 0 && !hasNonReflectFields(fields) {
		return renderFileFlat(pkg, structName, fields)
	}

	return renderFileFull(pkg, structName, fields, nestedOrder, nestedFields, nestedComplex, objectLocalNames, listLocalNames, locals)
}

func collectNestedNames(fields []fieldInfo) map[string]bool {
	result := make(map[string]bool)
	for _, f := range fields {
		if f.kind == kindLocal || f.kind == kindPtrLocal || f.kind == kindSliceLocal {
			result[f.localName] = true
		}
	}
	return result
}

// hasNonReflectFields reports whether any field requires an explicit conversion
// that ReflectToTF/ReflectFromTF cannot handle (e.g. map[string]string).
func hasNonReflectFields(fields []fieldInfo) bool {
	for _, f := range fields {
		if f.kind == kindMapStringString || f.kind == kindLocal || f.kind == kindObjectLocal || f.kind == kindListLocal {
			return true
		}
	}
	return false
}

// collectObjectLocalNames returns the names of all local structs referenced via kindObjectLocal fields.
func collectObjectLocalNames(fields []fieldInfo) map[string]bool {
	result := make(map[string]bool)
	for _, f := range fields {
		if f.kind == kindObjectLocal {
			result[f.localName] = true
		}
	}
	return result
}

// collectAllListLocalNames collects names of all local structs referenced via kindListLocal
// from both root fields and all nested struct fields.
func collectAllListLocalNames(rootFields []fieldInfo, nestedFields map[string][]fieldInfo) map[string]bool {
	result := make(map[string]bool)
	for _, f := range rootFields {
		if f.kind == kindListLocal {
			result[f.localName] = true
		}
	}
	for _, fields := range nestedFields {
		for _, f := range fields {
			if f.kind == kindListLocal {
				result[f.localName] = true
			}
		}
	}
	return result
}

// collectAllNested transitively collects all local struct names referenced from
// rootFields via BFS. Returns:
//   - topoOrder: topological order (innermost/leaf structs first)
//   - fieldsMap: parsed fields for each nested struct
//   - complexMap: whether a nested struct itself references further local structs
func collectAllNested(rootFields []fieldInfo, locals map[string]*ast.StructType) (topoOrder []string, fieldsMap map[string][]fieldInfo, complexMap map[string]bool, err error) {
	fieldsMap = make(map[string][]fieldInfo)
	complexMap = make(map[string]bool)
	seen := map[string]bool{}
	var bfsOrder []string

	// Seed queue from root-level fields.
	var queue []string
	for name := range collectNestedNames(rootFields) {
		if !seen[name] {
			queue = append(queue, name)
			seen[name] = true
		}
	}

	for len(queue) > 0 {
		name := queue[0]
		queue = queue[1:]
		bfsOrder = append(bfsOrder, name)

		st, exists := locals[name]
		if !exists {
			err = fmt.Errorf("referenced local struct %s not found in file", name)
			return
		}
		var nf []fieldInfo
		if nf, err = parseFields(st, locals); err != nil {
			err = fmt.Errorf("nested struct %s: %w", name, err)
			return
		}
		fieldsMap[name] = nf

		childNames := collectNestedNames(nf)
		complexMap[name] = len(childNames) > 0 || hasNonReflectFields(nf)
		for child := range childNames {
			if !seen[child] {
				queue = append(queue, child)
				seen[child] = true
			}
		}
	}

	// Reverse BFS order for topological order (leaf structs first).
	topoOrder = make([]string, len(bfsOrder))
	for i, name := range bfsOrder {
		topoOrder[len(bfsOrder)-1-i] = name
	}
	return
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
	tfTag := extractTFTag(field.Tag)
	kind, localName, tfType, err := resolveType(field.Type, locals, tfTag)
	if err != nil {
		return fieldInfo{}, err
	}
	tfsdk, err := extractTFSDKTag(goName, field.Tag)
	if err != nil {
		return fieldInfo{}, err
	}
	return fieldInfo{goName: goName, tfType: tfType, tfsdk: tfsdk, kind: kind, localName: localName}, nil
}

// extractTFTag returns the value of the tf struct tag (e.g. tf:"object" → "object").
func extractTFTag(tag *ast.BasicLit) string {
	if tag == nil {
		return ""
	}
	raw := strings.Trim(tag.Value, "`")
	for _, pair := range strings.Fields(raw) {
		if !strings.HasPrefix(pair, `tf:"`) {
			continue
		}
		val := strings.TrimPrefix(pair, `tf:"`)
		val = strings.TrimSuffix(val, `"`)
		return val
	}
	return ""
}

// resolveType maps an AST type expression to (kind, localName, tfTypeStr, error).
// tfTag is the value of the `tf:"..."` struct tag on the field (empty if absent).
func resolveType(expr ast.Expr, locals map[string]*ast.StructType, tfTag string) (string, string, string, error) {
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
		if locals != nil {
			if _, ok := locals[t.Name]; ok {
				if tfTag == "object" {
					return kindObjectLocal, t.Name, "types.Object", nil
				}
				return kindLocal, t.Name, "*" + t.Name + "TF", nil
			}
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
					if tfTag == "list" {
						return kindListLocal, el.Name, "types.List", nil
					}
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
	case *ast.MapType:
		key, ok1 := t.Key.(*ast.Ident)
		val, ok2 := t.Value.(*ast.Ident)
		if ok1 && ok2 && key.Name == "string" && val.Name == "string" {
			return kindMapStringString, "", "map[string]types.String", nil
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

// anyFieldsNeedHelper reports whether any of the provided field sets contain a
// field whose generated code still calls the helper package (int/[]string/map kinds).
func anyFieldsNeedHelper(sets ...[]fieldInfo) bool {
	needHelperKinds := map[string]bool{
		kindInt:             true,
		kindPtrInt:          true,
		kindStrings:         true,
		kindMapStringString: true,
	}
	for _, fields := range sets {
		for _, f := range fields {
			if needHelperKinds[f.kind] {
				return true
			}
		}
	}
	return false
}

func renderFileFull(pkg, structName string, fields []fieldInfo, nestedOrder []string, nestedFields map[string][]fieldInfo, nestedComplex map[string]bool, objectLocalNames map[string]bool, listLocalNames map[string]bool, locals map[string]*ast.StructType) ([]byte, error) {
	tfName := structName + "TF"

	var buf bytes.Buffer
	fmt.Fprintf(&buf, "// Code generated by gen_tf_models. DO NOT EDIT.\n")
	fmt.Fprintf(&buf, "package %s\n\n", pkg)
	fmt.Fprintf(&buf, "import (\n")
	needAttr := len(objectLocalNames) > 0 || len(listLocalNames) > 0
	if needAttr {
		fmt.Fprintf(&buf, "\t\"github.com/hashicorp/terraform-plugin-framework/attr\"\n")
	}
	fmt.Fprintf(&buf, "\t\"github.com/hashicorp/terraform-plugin-framework/types\"\n")
	// Collect all field sets that emit explicit helper calls to determine if helper is needed.
	// Non-complex nested structs use ReflectToTF/ReflectFromTF and never call helper directly.
	allFieldSets := [][]fieldInfo{fields}
	for name, nf := range nestedFields {
		if nestedComplex[name] {
			allFieldSets = append(allFieldSets, nf)
		}
	}
	for name := range objectLocalNames {
		if st, ok := locals[name]; ok {
			if nf, err2 := parseFields(st, locals); err2 == nil {
				allFieldSets = append(allFieldSets, nf)
			}
		}
	}
	for name := range listLocalNames {
		if st, ok := locals[name]; ok {
			if nf, err2 := parseFields(st, locals); err2 == nil {
				allFieldSets = append(allFieldSets, nf)
			}
		}
	}
	if anyFieldsNeedHelper(allFieldSets...) {
		fmt.Fprintf(&buf, "\t\"github.com/leanspace/terraform-provider-leanspace/helper\"\n")
	}
	fmt.Fprintf(&buf, "\t\"github.com/leanspace/terraform-provider-leanspace/helper/general_objects\"\n")
	fmt.Fprintf(&buf, ")\n\n")

	// Nested TF struct declarations in topological order (innermost first).
	for _, name := range nestedOrder {
		writeStructDecl(&buf, name+"TF", nestedFields[name])
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

	// Helper functions for each nested local struct.
	for _, name := range nestedOrder {
		lc := lcFirst(name)
		nf := nestedFields[name]
		isComplex := nestedComplex[name]

		if isComplex {
			// Complex nested struct: emit explicit field-level bodies.
			fmt.Fprintf(&buf, "func %sToTF(x *%s) *%sTF {\n", lc, name, name)
			fmt.Fprintf(&buf, "\tif x == nil {\n\t\treturn nil\n\t}\n")
			fmt.Fprintf(&buf, "\treturn &%sTF{\n", name)
			for _, f := range nf {
				if f.isEmbed {
					fmt.Fprintf(&buf, "\t\tAuditModelTF: general_objects.AuditModelToTF(&x.AuditModel),\n")
				} else {
					fmt.Fprintf(&buf, "\t\t%s: %s,\n", f.goName, toTFExpr(f, "x"))
				}
			}
			fmt.Fprintf(&buf, "\t}\n}\n\n")

			fmt.Fprintf(&buf, "func %sFromTF(tf *%sTF) *%s {\n", lc, name, name)
			fmt.Fprintf(&buf, "\tif tf == nil {\n\t\treturn nil\n\t}\n")
			fmt.Fprintf(&buf, "\treturn &%s{\n", name)
			for _, f := range nf {
				if f.isEmbed {
					fmt.Fprintf(&buf, "\t\tAuditModel: general_objects.AuditModelFromTF(tf.AuditModelTF),\n")
				} else {
					fmt.Fprintf(&buf, "\t\t%s: %s,\n", f.goName, fromTFExpr(f, "tf"))
				}
			}
			fmt.Fprintf(&buf, "\t}\n}\n\n")
		} else {
			// Flat nested struct: delegate to ReflectToTF.
			fmt.Fprintf(&buf, "func %sToTF(x *%s) *%sTF {\n", lc, name, name)
			fmt.Fprintf(&buf, "\tif x == nil {\n\t\treturn nil\n\t}\n")
			fmt.Fprintf(&buf, "\treturn general_objects.ReflectToTF[%sTF](x)\n", name)
			fmt.Fprintf(&buf, "}\n\n")

			fmt.Fprintf(&buf, "func %sFromTF(tf *%sTF) *%s {\n", lc, name, name)
			fmt.Fprintf(&buf, "\tif tf == nil {\n\t\treturn nil\n\t}\n")
			fmt.Fprintf(&buf, "\treturn general_objects.ReflectFromTF[%s](tf)\n", name)
			fmt.Fprintf(&buf, "}\n\n")
		}

		// Slice helpers reuse the pointer helpers above.
		fmt.Fprintf(&buf, "func %sValueFromTF(tf *%sTF) %s {\n", lc, name, name)
		fmt.Fprintf(&buf, "\tif v := %sFromTF(tf); v != nil {\n", lc)
		fmt.Fprintf(&buf, "\t\treturn *v\n\t}\n")
		fmt.Fprintf(&buf, "\treturn %s{}\n}\n\n", name)

		// Slice helpers reuse the pointer helpers above.
		fmt.Fprintf(&buf, "func %sSliceToTF(xs []%s) []%sTF {\n", lc, name, name)
		fmt.Fprintf(&buf, "\tresult := make([]%sTF, len(xs))\n", name)
		fmt.Fprintf(&buf, "\tfor i := range xs {\n")
		if isComplex {
			fmt.Fprintf(&buf, "\t\tresult[i] = *%sToTF(&xs[i])\n", lc)
		} else {
			fmt.Fprintf(&buf, "\t\tresult[i] = *general_objects.ReflectToTF[%sTF](&xs[i])\n", name)
		}
		fmt.Fprintf(&buf, "\t}\n")
		fmt.Fprintf(&buf, "\treturn result\n}\n\n")

		fmt.Fprintf(&buf, "func %sSliceFromTF(tfs []%sTF) []%s {\n", lc, name, name)
		fmt.Fprintf(&buf, "\tresult := make([]%s, len(tfs))\n", name)
		fmt.Fprintf(&buf, "\tfor i := range tfs {\n")
		if isComplex {
			fmt.Fprintf(&buf, "\t\tresult[i] = *%sFromTF(&tfs[i])\n", lc)
		} else {
			fmt.Fprintf(&buf, "\t\tresult[i] = *general_objects.ReflectFromTF[%s](&tfs[i])\n", name)
		}
		fmt.Fprintf(&buf, "\t}\n")
		fmt.Fprintf(&buf, "\treturn result\n}\n\n")
	}

	// Helper functions for object-local structs (tf:"object" annotated).
	// These emit attrTypes vars and toObject/fromObject helpers instead of a TF struct.
	for name := range objectLocalNames {
		lc := lcFirst(name)
		st := locals[name]
		nf, err := parseFields(st, locals)
		if err != nil {
			return nil, fmt.Errorf("object-local struct %s: %w", name, err)
		}

		// attrTypes var
		fmt.Fprintf(&buf, "var %sAttrTypes = map[string]attr.Type{\n", lc)
		for _, f := range nf {
			fmt.Fprintf(&buf, "\t%q: %s,\n", f.tfsdk, primitiveAttrType(f.kind))
		}
		fmt.Fprintf(&buf, "}\n\n")

		// nameToObject
		fmt.Fprintf(&buf, "func %sToObject(x *%s) types.Object {\n", lc, name)
		fmt.Fprintf(&buf, "\tif x == nil {\n\t\treturn types.ObjectNull(%sAttrTypes)\n\t}\n", lc)
		fmt.Fprintf(&buf, "\treturn types.ObjectValueMust(%sAttrTypes, map[string]attr.Value{\n", lc)
		for _, f := range nf {
			fmt.Fprintf(&buf, "\t\t%q: %s,\n", f.tfsdk, toTFExpr(f, "x"))
		}
		fmt.Fprintf(&buf, "\t})\n}\n\n")

		// nameFromObject
		fmt.Fprintf(&buf, "func %sFromObject(obj types.Object) %s {\n", lc, name)
		fmt.Fprintf(&buf, "\tif obj.IsNull() || obj.IsUnknown() {\n\t\treturn %s{}\n\t}\n", name)
		fmt.Fprintf(&buf, "\tattrs := obj.Attributes()\n")
		fmt.Fprintf(&buf, "\treturn %s{\n", name)
		for _, f := range nf {
			fmt.Fprintf(&buf, "\t\t%s: %s,\n", f.goName, fromObjectAttrExpr(f))
		}
		fmt.Fprintf(&buf, "\t}\n}\n\n")
	}

	// Helper functions for list-local structs (tf:"list" annotated).
	// These emit attrTypes vars and xToList/xFromList helpers instead of a TF struct.
	for name := range listLocalNames {
		lc := lcFirst(name)
		st := locals[name]
		nf, err := parseFields(st, locals)
		if err != nil {
			return nil, fmt.Errorf("list-local struct %s: %w", name, err)
		}

		// attrTypes var
		fmt.Fprintf(&buf, "var %sAttrTypes = map[string]attr.Type{\n", lc)
		for _, f := range nf {
			fmt.Fprintf(&buf, "\t%q: %s,\n", f.tfsdk, primitiveAttrType(f.kind))
		}
		fmt.Fprintf(&buf, "}\n\n")

		// xToList
		fmt.Fprintf(&buf, "func %sToList(xs []%s) types.List {\n", lc, name)
		fmt.Fprintf(&buf, "\telems := make([]attr.Value, len(xs))\n")
		fmt.Fprintf(&buf, "\tfor i := range xs {\n")
		fmt.Fprintf(&buf, "\t\telems[i] = types.ObjectValueMust(%sAttrTypes, map[string]attr.Value{\n", lc)
		for _, f := range nf {
			fmt.Fprintf(&buf, "\t\t\t%q: %s,\n", f.tfsdk, toTFExpr(f, "xs[i]"))
		}
		fmt.Fprintf(&buf, "\t\t})\n\t}\n")
		fmt.Fprintf(&buf, "\treturn types.ListValueMust(types.ObjectType{AttrTypes: %sAttrTypes}, elems)\n}\n\n", lc)

		// xFromList
		fmt.Fprintf(&buf, "func %sFromList(list types.List) []%s {\n", lc, name)
		fmt.Fprintf(&buf, "\tif list.IsNull() || list.IsUnknown() {\n\t\treturn nil\n\t}\n")
		fmt.Fprintf(&buf, "\tresult := make([]%s, len(list.Elements()))\n", name)
		fmt.Fprintf(&buf, "\tfor i, elem := range list.Elements() {\n")
		fmt.Fprintf(&buf, "\t\tobj := elem.(types.Object)\n")
		fmt.Fprintf(&buf, "\t\tattrs := obj.Attributes()\n")
		fmt.Fprintf(&buf, "\t\tresult[i] = %s{\n", name)
		for _, f := range nf {
			fmt.Fprintf(&buf, "\t\t\t%s: %s,\n", f.goName, fromObjectAttrExpr(f))
		}
		fmt.Fprintf(&buf, "\t\t}\n\t}\n")
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

// primitiveAttrType maps a primitive kind to its attr.Type expression.
// Used when building attrTypes maps for object-local structs.
func primitiveAttrType(kind string) string {
	switch kind {
	case kindString, kindPtrString:
		return "types.StringType"
	case kindBool, kindPtrBool:
		return "types.BoolType"
	case kindInt, kindPtrInt:
		return "types.Int64Type"
	case kindFloat64, kindPtrFloat64:
		return "types.Float64Type"
	}
	return "/* unsupported attr.Type */"
}

// fromObjectAttrExpr returns the expression to extract a field from attrs["tfsdk"] for object-local structs.
func fromObjectAttrExpr(fi fieldInfo) string {
	ref := `attrs["` + fi.tfsdk + `"]`
	switch fi.kind {
	case kindString:
		return ref + ".(types.String).ValueString()"
	case kindPtrString:
		return ref + ".(types.String).ValueStringPointer()"
	case kindBool:
		return ref + ".(types.Bool).ValueBool()"
	case kindPtrBool:
		return ref + ".(types.Bool).ValueBoolPointer()"
	case kindInt:
		return "helper.FromTFInt64(" + ref + ".(types.Int64))"
	case kindPtrInt:
		return "helper.FromTFIntPtr(" + ref + ".(types.Int64))"
	case kindFloat64:
		return ref + ".(types.Float64).ValueFloat64()"
	case kindPtrFloat64:
		return ref + ".(types.Float64).ValueFloat64Pointer()"
	}
	return "/* unsupported field type for object-local struct */"
}

// toTFExpr returns the expression that converts an API field to its TF value.
func toTFExpr(fi fieldInfo, src string) string {
	ref := src + "." + fi.goName
	switch fi.kind {
	case kindString:
		return "types.StringValue(" + ref + ")"
	case kindPtrString:
		return "types.StringPointerValue(" + ref + ")"
	case kindBool:
		return "types.BoolValue(" + ref + ")"
	case kindPtrBool:
		return "types.BoolPointerValue(" + ref + ")"
	case kindInt:
		return "helper.TFInt64Value(" + ref + ")"
	case kindPtrInt:
		return "helper.TFIntPtrValue(" + ref + ")"
	case kindFloat64:
		return "types.Float64Value(" + ref + ")"
	case kindPtrFloat64:
		return "types.Float64PointerValue(" + ref + ")"
	case kindStrings:
		return "helper.TFStringsValue(" + ref + ")"
	case kindKeyValues:
		return "general_objects.KeyValuesToTF(" + ref + ")"
	case kindMapStringString:
		return "helper.MapStringToTF(" + ref + ")"
	case kindLocal:
		return lcFirst(fi.localName) + "ToTF(&" + ref + ")"
	case kindPtrLocal:
		return lcFirst(fi.localName) + "ToTF(" + ref + ")"
	case kindSliceLocal:
		return lcFirst(fi.localName) + "SliceToTF(" + ref + ")"
	case kindObjectLocal:
		return lcFirst(fi.localName) + "ToObject(&" + ref + ")"
	case kindListLocal:
		return lcFirst(fi.localName) + "ToList(" + ref + ")"
	}
	return "/* unknown */"
}

// fromTFExpr returns the expression that converts a TF field back to its API value.
func fromTFExpr(fi fieldInfo, src string) string {
	ref := src + "." + fi.goName
	switch fi.kind {
	case kindString:
		return ref + ".ValueString()"
	case kindPtrString:
		return ref + ".ValueStringPointer()"
	case kindBool:
		return ref + ".ValueBool()"
	case kindPtrBool:
		return ref + ".ValueBoolPointer()"
	case kindInt:
		return "helper.FromTFInt64(" + ref + ")"
	case kindPtrInt:
		return "helper.FromTFIntPtr(" + ref + ")"
	case kindFloat64:
		return ref + ".ValueFloat64()"
	case kindPtrFloat64:
		return ref + ".ValueFloat64Pointer()"
	case kindStrings:
		return "helper.FromTFStrings(" + ref + ")"
	case kindKeyValues:
		return "general_objects.KeyValuesFromTF(" + ref + ")"
	case kindMapStringString:
		return "helper.MapStringFromTF(" + ref + ")"
	case kindLocal:
		return lcFirst(fi.localName) + "ValueFromTF(" + ref + ")"
	case kindPtrLocal:
		return lcFirst(fi.localName) + "FromTF(" + ref + ")"
	case kindSliceLocal:
		return lcFirst(fi.localName) + "SliceFromTF(" + ref + ")"
	case kindObjectLocal:
		return lcFirst(fi.localName) + "FromObject(" + ref + ")"
	case kindListLocal:
		return lcFirst(fi.localName) + "FromList(" + ref + ")"
	}
	return "/* unknown */"
}
