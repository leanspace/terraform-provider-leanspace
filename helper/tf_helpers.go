package helper

import "github.com/hashicorp/terraform-plugin-framework/types"

// Go → TF (API model field to TF model field)

func TFStringValue(s string) types.String {
	return types.StringValue(s)
}

func TFStringPtrValue(s *string) types.String {
	if s == nil {
		return types.StringNull()
	}
	return types.StringValue(*s)
}

func TFBoolValue(b bool) types.Bool {
	return types.BoolValue(b)
}

func TFInt64Value(i int) types.Int64 {
	return types.Int64Value(int64(i))
}

func TFFloat64Value(f float64) types.Float64 {
	return types.Float64Value(f)
}

func TFFloat64PtrValue(f *float64) types.Float64 {
	if f == nil {
		return types.Float64Null()
	}
	return types.Float64Value(*f)
}

func TFIntPtrValue(i *int) types.Int64 {
	if i == nil {
		return types.Int64Null()
	}
	return types.Int64Value(int64(*i))
}

func TFBoolPtrValue(b *bool) types.Bool {
	if b == nil {
		return types.BoolNull()
	}
	return types.BoolValue(*b)
}

func TFStringsValue(ss []string) []types.String {
	if ss == nil {
		return nil
	}
	result := make([]types.String, len(ss))
	for i, s := range ss {
		result[i] = types.StringValue(s)
	}
	return result
}

// TF → Go (TF model field to API model field)

func FromTFString(s types.String) string {
	if s.IsNull() || s.IsUnknown() {
		return ""
	}
	return s.ValueString()
}

func FromTFStringPtr(s types.String) *string {
	if s.IsNull() || s.IsUnknown() {
		return nil
	}
	v := s.ValueString()
	return &v
}

func FromTFBool(b types.Bool) bool {
	if b.IsNull() || b.IsUnknown() {
		return false
	}
	return b.ValueBool()
}

func FromTFInt64(i types.Int64) int {
	if i.IsNull() || i.IsUnknown() {
		return 0
	}
	return int(i.ValueInt64())
}

func FromTFFloat64(f types.Float64) float64 {
	if f.IsNull() || f.IsUnknown() {
		return 0.0
	}
	return f.ValueFloat64()
}

func FromTFFloat64Ptr(f types.Float64) *float64 {
	if f.IsNull() || f.IsUnknown() {
		return nil
	}
	v := f.ValueFloat64()
	return &v
}

func FromTFIntPtr(i types.Int64) *int {
	if i.IsNull() || i.IsUnknown() {
		return nil
	}
	v := int(i.ValueInt64())
	return &v
}

func FromTFBoolPtr(b types.Bool) *bool {
	if b.IsNull() || b.IsUnknown() {
		return nil
	}
	v := b.ValueBool()
	return &v
}

func FromTFStrings(ts []types.String) []string {
	if ts == nil {
		return nil
	}
	result := make([]string, len(ts))
	for i, t := range ts {
		result[i] = t.ValueString()
	}
	return result
}

func MapStringToTF(m map[string]string) map[string]types.String {
	if m == nil {
		return nil
	}
	result := make(map[string]types.String, len(m))
	for k, v := range m {
		result[k] = types.StringValue(v)
	}
	return result
}

func MapStringFromTF(m map[string]types.String) map[string]string {
	if m == nil {
		return nil
	}
	result := make(map[string]string, len(m))
	for k, v := range m {
		result[k] = v.ValueString()
	}
	return result
}
