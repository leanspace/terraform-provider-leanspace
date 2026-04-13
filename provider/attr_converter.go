package provider

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	datasourceschema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	resourceschema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// MapToObjectValue converts a map[string]any (from ToMap()) into a types.Object
// using the given attribute type map for type information.
func MapToObjectValue(ctx context.Context, attrTypes map[string]attr.Type, data map[string]any) (types.Object, diag.Diagnostics) {
	if data == nil {
		return types.ObjectNull(attrTypes), nil
	}
	attrValues := make(map[string]attr.Value)
	var diags diag.Diagnostics
	for key, attrType := range attrTypes {
		val, d := nativeToAttrValue(ctx, attrType, data[key])
		diags.Append(d...)
		attrValues[key] = val
	}
	obj, d := types.ObjectValue(attrTypes, attrValues)
	diags.Append(d...)
	return obj, diags
}

// MapToAttrValues converts a map[string]any (from ToMap()) into a map of attr.Value
// suitable for setting framework state. Uses resource schema attributes for type info.
// Empty strings are treated as null to mirror SDK v2 TypeString behaviour.
// SetNestedAttribute that was converted to ListNestedBlock produces types.List;
// Computed-only SetNestedAttribute (stays as attribute) produces types.Set.
func MapToAttrValues(ctx context.Context, schemaAttrs map[string]resourceschema.Attribute, data map[string]any) (map[string]attr.Value, diag.Diagnostics) {
	result := make(map[string]attr.Value)
	var diags diag.Diagnostics
	for key, schemaAttr := range schemaAttrs {
		rawVal := data[key]

		if sa, ok := schemaAttr.(resourceschema.SetNestedAttribute); ok {
			var val attr.Value
			var d diag.Diagnostics
			setType := sa.GetType().(basetypes.SetType)
			if sa.Computed && !sa.Optional && !sa.Required {
				// Stays as SetNestedAttribute in schema — produce types.Set
				val, d = nativeToAttrValue(ctx, setType, rawVal)
			} else //if hasComputedField(sa.NestedObject.Attributes) {
			{
				// Has server-computed element fields — converted to ListNestedBlock — produce types.List
				listType := basetypes.ListType{ElemType: setType.ElemType}
				val, d = nativeToAttrValue(ctx, listType, rawVal)
			} /*else {
				// No Computed element fields — converted to SetNestedBlock — produce types.Set
				val, d = nativeToAttrValue(ctx, setType, rawVal)
			}*/
			diags.Append(d...)
			result[key] = val
			continue
		}

		val, d := nativeToAttrValue(ctx, normalizeAttrType(schemaAttr.GetType()), rawVal)
		diags.Append(d...)
		result[key] = val
	}
	return result, diags
}

// MapToAttrValuesDatasource is the same as MapToAttrValues but for datasource schemas.
func MapToAttrValuesDatasource(ctx context.Context, schemaAttrs map[string]datasourceschema.Attribute, data map[string]any) (map[string]attr.Value, diag.Diagnostics) {
	result := make(map[string]attr.Value)
	var diags diag.Diagnostics
	for key, schemaAttr := range schemaAttrs {
		attrType := schemaAttr.GetType()
		val, d := nativeToAttrValue(ctx, attrType, data[key])
		diags.Append(d...)
		result[key] = val
	}
	return result, diags
}

// AttrValuesToMap converts framework attr.Values back into a map[string]any
// suitable for FromMap(). Uses resource schema attributes for type info.
func AttrValuesToMap(ctx context.Context, schemaAttrs map[string]resourceschema.Attribute, values map[string]attr.Value) (map[string]any, diag.Diagnostics) {
	result := make(map[string]any)
	var diags diag.Diagnostics
	for key, schemaAttr := range schemaAttrs {
		attrType := schemaAttr.GetType()
		val, d := attrValueToNative(ctx, attrType, values[key])
		diags.Append(d...)
		result[key] = val
	}
	return result, diags
}

// nativeToAttrValue converts a Go native value (from ToMap()) to an attr.Value
// based on the target attr.Type.
func nativeToAttrValue(ctx context.Context, attrType attr.Type, value any) (attr.Value, diag.Diagnostics) {
	var diags diag.Diagnostics

	if value == nil {
		return nullValueForType(attrType), nil
	}

	switch t := attrType.(type) {
	case basetypes.StringType:
		s := fmt.Sprint(value)
		return types.StringValue(s), nil

	case basetypes.BoolType:
		if b, ok := value.(bool); ok {
			return types.BoolValue(b), nil
		}
		return types.BoolNull(), nil

	case basetypes.Int64Type:
		switch v := value.(type) {
		case int:
			return types.Int64Value(int64(v)), nil
		case int64:
			return types.Int64Value(v), nil
		case float64:
			return types.Int64Value(int64(v)), nil
		}
		return types.Int64Null(), nil

	case basetypes.Float64Type:
		switch v := value.(type) {
		case float64:
			return types.Float64Value(v), nil
		case int:
			return types.Float64Value(float64(v)), nil
		}
		return types.Float64Null(), nil

	case basetypes.ListType:
		return nativeToListValue(ctx, t, value)

	case basetypes.SetType:
		return nativeToSetValue(ctx, t, value)

	case basetypes.MapType:
		return nativeToMapValue(ctx, t, value)

	case basetypes.ObjectType:
		return nativeToObjectValue(ctx, t, value)

	default:
		diags.AddError("Unsupported attr type", fmt.Sprintf("Cannot convert to attr type %T", attrType))
		return nullValueForType(attrType), diags
	}
}

func nativeToListValue(ctx context.Context, listType basetypes.ListType, value any) (attr.Value, diag.Diagnostics) {
	var diags diag.Diagnostics
	elemType := listType.ElemType

	slice, ok := toSlice(value)
	if !ok {
		return types.ListNull(elemType), nil
	}

	if len(slice) == 0 {
		return types.ListValueMust(elemType, []attr.Value{}), nil
	}

	elems := make([]attr.Value, len(slice))
	for i, item := range slice {
		val, d := nativeToAttrValue(ctx, elemType, item)
		diags.Append(d...)
		elems[i] = val
	}
	list, d := types.ListValue(elemType, elems)
	diags.Append(d...)
	return list, diags
}

func nativeToSetValue(ctx context.Context, setType basetypes.SetType, value any) (attr.Value, diag.Diagnostics) {
	var diags diag.Diagnostics
	elemType := setType.ElemType

	slice, ok := toSlice(value)
	if !ok {
		return types.SetNull(elemType), nil
	}

	if len(slice) == 0 {
		return types.SetValueMust(elemType, []attr.Value{}), nil
	}

	elems := make([]attr.Value, len(slice))
	for i, item := range slice {
		val, d := nativeToAttrValue(ctx, elemType, item)
		diags.Append(d...)
		elems[i] = val
	}
	set, d := types.SetValue(elemType, elems)
	diags.Append(d...)
	return set, diags
}

func nativeToMapValue(ctx context.Context, mapType basetypes.MapType, value any) (attr.Value, diag.Diagnostics) {
	var diags diag.Diagnostics
	elemType := mapType.ElemType

	// Normalise typed maps (e.g. map[string]string from model fields) to map[string]any
	var m map[string]any
	switch v := value.(type) {
	case map[string]any:
		m = v
	case map[string]string:
		m = make(map[string]any, len(v))
		for k, s := range v {
			m[k] = s
		}
	default:
		return types.MapNull(elemType), nil
	}

	if len(m) == 0 {
		return types.MapValueMust(elemType, map[string]attr.Value{}), nil
	}

	elems := make(map[string]attr.Value)
	for k, v := range m {
		val, d := nativeToAttrValue(ctx, elemType, v)
		diags.Append(d...)
		elems[k] = val
	}
	mapVal, d := types.MapValue(elemType, elems)
	diags.Append(d...)
	return mapVal, diags
}

func nativeToObjectValue(ctx context.Context, objType basetypes.ObjectType, value any) (attr.Value, diag.Diagnostics) {
	var diags diag.Diagnostics

	m, ok := value.(map[string]any)
	if !ok {
		return types.ObjectNull(objType.AttrTypes), nil
	}

	attrValues := make(map[string]attr.Value)
	for key, childType := range objType.AttrTypes {
		val, d := nativeToAttrValue(ctx, childType, m[key])
		diags.Append(d...)
		attrValues[key] = val
	}
	obj, d := types.ObjectValue(objType.AttrTypes, attrValues)
	diags.Append(d...)
	return obj, diags
}

// attrValueToNative converts an attr.Value back to a Go native value for FromMap().
// Null and unknown values return nil. Parsers must use helper.CastXxx helpers instead
// of bare type assertions to safely absorb nil without panicking.
func attrValueToNative(ctx context.Context, attrType attr.Type, value attr.Value) (any, diag.Diagnostics) {
	var diags diag.Diagnostics

	if value == nil || value.IsNull() || value.IsUnknown() {
		return nil, nil
	}

	switch attrType.(type) {
	case basetypes.StringType:
		if sv, ok := value.(types.String); ok {
			return sv.ValueString(), nil
		}

	case basetypes.BoolType:
		if bv, ok := value.(types.Bool); ok {
			return bv.ValueBool(), nil
		}

	case basetypes.Int64Type:
		if iv, ok := value.(types.Int64); ok {
			return int(iv.ValueInt64()), nil
		}

	case basetypes.Float64Type:
		if fv, ok := value.(types.Float64); ok {
			return fv.ValueFloat64(), nil
		}

	case basetypes.ListType:
		return listValueToNative(ctx, value)

	case basetypes.SetType:
		// SetNestedAttribute may be stored as types.List (converted to ListNestedBlock)
		// or as types.Set (Computed-only, stayed as attribute). Dispatch on runtime type.
		if _, isList := value.(types.List); isList {
			return listValueToNative(ctx, value)
		}
		return setValueToNative(ctx, value)

	case basetypes.MapType:
		return mapValueToNative(ctx, attrType.(basetypes.MapType), value)

	case basetypes.ObjectType:
		return objectValueToNative(ctx, attrType.(basetypes.ObjectType), value)
	}

	return nil, diags
}

func listValueToNative(ctx context.Context, value attr.Value) (any, diag.Diagnostics) {
	var diags diag.Diagnostics
	lv, ok := value.(types.List)
	if !ok {
		return nil, nil
	}
	elems := lv.Elements()
	elemType := lv.ElementType(ctx)

	// If the element type is an object type, return []any of maps
	if _, isObj := elemType.(types.ObjectType); isObj {
		result := make([]any, len(elems))
		for i, elem := range elems {
			val, d := attrValueToNative(ctx, elemType, elem)
			diags.Append(d...)
			result[i] = val
		}
		return result, diags
	}

	// For primitive list elements, return []any
	result := make([]any, len(elems))
	for i, elem := range elems {
		val, d := attrValueToNative(ctx, elemType, elem)
		diags.Append(d...)
		result[i] = val
	}
	return result, diags
}

func setValueToNative(ctx context.Context, value attr.Value) (any, diag.Diagnostics) {
	var diags diag.Diagnostics
	sv, ok := value.(types.Set)
	if !ok {
		return nil, nil
	}
	elems := sv.Elements()
	elemType := sv.ElementType(ctx)

	result := make([]any, len(elems))
	for i, elem := range elems {
		val, d := attrValueToNative(ctx, elemType, elem)
		diags.Append(d...)
		result[i] = val
	}
	return result, diags
}

func mapValueToNative(ctx context.Context, mapType basetypes.MapType, value attr.Value) (any, diag.Diagnostics) {
	var diags diag.Diagnostics
	mv, ok := value.(types.Map)
	if !ok {
		return nil, nil
	}
	elems := mv.Elements()
	result := make(map[string]any)
	for k, v := range elems {
		val, d := attrValueToNative(ctx, mapType.ElemType, v)
		diags.Append(d...)
		result[k] = val
	}
	return result, diags
}

func objectValueToNative(ctx context.Context, objType basetypes.ObjectType, value attr.Value) (any, diag.Diagnostics) {
	var diags diag.Diagnostics
	ov, ok := value.(types.Object)
	if !ok {
		return nil, nil
	}
	attrs := ov.Attributes()
	result := make(map[string]any)
	for key, childType := range objType.AttrTypes {
		val, d := attrValueToNative(ctx, childType, attrs[key])
		diags.Append(d...)
		result[key] = val
	}
	return result, diags
}

// normalizeAttrType recursively replaces SetType-of-objects with ListType throughout an attr.Type tree.
// Only nested Set attributes (SetNestedAttribute, whose ElemType is ObjectType) are converted,
// because SplitResourceSchemaBlocks converts those to ListNestedBlock at runtime.
// Primitive Set attributes (e.g. SetAttribute{ElementType: StringType}) are left as SetType.
func normalizeAttrType(t attr.Type) attr.Type {
	switch v := t.(type) {
	case basetypes.SetType:
		normalizedElem := normalizeAttrType(v.ElemType)
		// Only nested Sets (element is an Object) become List (ListNestedBlock).
		if _, isObj := normalizedElem.(basetypes.ObjectType); isObj {
			return basetypes.ListType{ElemType: normalizedElem}
		}
		return basetypes.SetType{ElemType: normalizedElem}
	case basetypes.ListType:
		return basetypes.ListType{ElemType: normalizeAttrType(v.ElemType)}
	case basetypes.MapType:
		return basetypes.MapType{ElemType: normalizeAttrType(v.ElemType)}
	case basetypes.ObjectType:
		normalized := make(map[string]attr.Type, len(v.AttrTypes))
		for k, child := range v.AttrTypes {
			normalized[k] = normalizeAttrType(child)
		}
		return basetypes.ObjectType{AttrTypes: normalized}
	default:
		return t
	}
}

// normalizeAttrTypes applies normalizeAttrType to every entry in an attr.Type map.
func normalizeAttrTypes(attrTypes map[string]attr.Type) map[string]attr.Type {
	result := make(map[string]attr.Type, len(attrTypes))
	for k, t := range attrTypes {
		result[k] = normalizeAttrType(t)
	}
	return result
}

// nullValueForType returns a typed null value for the given attr.Type.
func nullValueForType(attrType attr.Type) attr.Value {
	switch t := attrType.(type) {
	case basetypes.StringType:
		return types.StringNull()
	case basetypes.BoolType:
		return types.BoolNull()
	case basetypes.Int64Type:
		return types.Int64Null()
	case basetypes.Float64Type:
		return types.Float64Null()
	case basetypes.ListType:
		return types.ListNull(t.ElemType)
	case basetypes.SetType:
		// After normalizeAttrType, only primitive Sets remain here.
		return types.SetNull(t.ElemType)
	case basetypes.MapType:
		return types.MapNull(t.ElemType)
	case basetypes.ObjectType:
		return types.ObjectNull(t.AttrTypes)
	default:
		return types.StringNull()
	}
}

// toSlice converts various slice-like values to []any.
// Handles []any, []map[string]any, and other slice types.
func toSlice(value any) ([]any, bool) {
	switch v := value.(type) {
	case []any:
		return v, true
	case []map[string]any:
		result := make([]any, len(v))
		for i, m := range v {
			result[i] = m
		}
		return result, true
	case []string:
		result := make([]any, len(v))
		for i, s := range v {
			result[i] = s
		}
		return result, true
	default:
		return nil, false
	}
}

// ResourceSchemaAttrTypes extracts the attr.Type map from resource schema attributes.
func ResourceSchemaAttrTypes(attrs map[string]resourceschema.Attribute) map[string]attr.Type {
	result := make(map[string]attr.Type)
	for key, a := range attrs {
		result[key] = a.GetType()
	}
	return result
}

// DatasourceSchemaAttrTypes extracts the attr.Type map from datasource schema attributes.
func DatasourceSchemaAttrTypes(attrs map[string]datasourceschema.Attribute) map[string]attr.Type {
	result := make(map[string]attr.Type)
	for key, a := range attrs {
		result[key] = a.GetType()
	}
	return result
}
