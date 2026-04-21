package helper

import "github.com/hashicorp/terraform-plugin-framework/types"

// Go → TF (API model field to TF model field)

func TFInt64Value(i int) types.Int64 {
	return types.Int64Value(int64(i))
}

func TFIntPtrValue(i *int) types.Int64 {
	if i == nil {
		return types.Int64Null()
	}
	return types.Int64Value(int64(*i))
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

func FromTFInt64(i types.Int64) int {
	if i.IsNull() || i.IsUnknown() {
		return 0
	}
	return int(i.ValueInt64())
}

func FromTFIntPtr(i types.Int64) *int {
	if i.IsNull() || i.IsUnknown() {
		return nil
	}
	v := int(i.ValueInt64())
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
