package properties

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/leanspace/terraform-provider-leanspace/helper"
	"github.com/leanspace/terraform-provider-leanspace/helper/general_objects"
)

type PropertyTF struct {
	general_objects.AuditModelTF
	Name        types.String                 `tfsdk:"name"`
	Description types.String                 `tfsdk:"description"`
	IsBuiltIn   types.Bool                   `tfsdk:"built_in"`
	NodeId      types.String                 `tfsdk:"node_id"`
	Tags        []general_objects.KeyValueTF `tfsdk:"tags"`
	// Attribute fields (flat, at top level)
	Type      types.String              `tfsdk:"type"`
	Value     types.String              `tfsdk:"value"`
	Min       types.Float64             `tfsdk:"min"`
	Max       types.Float64             `tfsdk:"max"`
	Scale     types.Int64               `tfsdk:"scale"`
	Precision types.Int64               `tfsdk:"precision"`
	UnitId    types.String              `tfsdk:"unit_id"`
	MinLength types.Int64               `tfsdk:"min_length"`
	MaxLength types.Int64               `tfsdk:"max_length"`
	Pattern   types.String              `tfsdk:"pattern"`
	Options   map[string]types.String   `tfsdk:"options"`
	Before    types.String              `tfsdk:"before"`
	After     types.String              `tfsdk:"after"`
	Fields    *general_objects.FieldsTF `tfsdk:"fields"`
}

func (x *Property[T]) ToTF() interface{} {
	tf := &PropertyTF{
		AuditModelTF: general_objects.AuditModelToTF(&x.AuditModel),
		Name:         types.StringValue(x.Name),
		Description:  types.StringPointerValue(x.Description),
		IsBuiltIn:    types.BoolValue(x.IsBuiltIn),
		NodeId:       types.StringValue(x.NodeId),
		Tags:         general_objects.KeyValuesToTF(x.Tags),
		Type:         types.StringValue(x.Attributes.Type),
		Min:          types.Float64PointerValue(x.Attributes.Min),
		Max:          types.Float64PointerValue(x.Attributes.Max),
		Scale:        helper.TFIntPtrValue(x.Attributes.Scale),
		Precision:    helper.TFIntPtrValue(x.Attributes.Precision),
		UnitId:       types.StringPointerValue(x.Attributes.UnitId),
		MinLength:    helper.TFIntPtrValue(x.Attributes.MinLength),
		MaxLength:    helper.TFIntPtrValue(x.Attributes.MaxLength),
		Pattern:      types.StringPointerValue(x.Attributes.Pattern),
		Before:       types.StringPointerValue(x.Attributes.Before),
		After:        types.StringPointerValue(x.Attributes.After),
		Fields:       general_objects.FieldsToTF(x.Attributes.Fields),
	}
	// Value handling by type
	val := interface{}(x.Attributes.Value)
	switch x.Attributes.Type {
	case "NUMERIC", "ENUM":
		if val != nil {
			if f, ok := val.(float64); ok {
				tf.Value = types.StringValue(helper.ParseFloat(f))
			}
		}
	case "BOOLEAN":
		if val != nil {
			if b, ok := val.(bool); ok {
				tf.Value = types.StringValue(strconv.FormatBool(b))
			}
		}
	case "TEXT", "TIMESTAMP", "DATE", "TIME":
		if val != nil {
			tf.Value = types.StringValue(fmt.Sprint(val))
		}
	case "TLE":
		if val != nil {
			if tleValues, ok := val.([]interface{}); ok {
				var tleValue string
				for _, value := range tleValues {
					tleValue = tleValue + "," + fmt.Sprint(value)
				}
				tf.Value = types.StringValue(strings.TrimPrefix(tleValue, ","))
			}
		}
	}
	if x.Attributes.Options != nil {
		tf.Options = make(map[string]types.String, len(*x.Attributes.Options))
		for k, v := range *x.Attributes.Options {
			tf.Options[k] = types.StringValue(fmt.Sprint(v))
		}
	}
	return tf
}

func (tf *PropertyTF) ToAPI() interface{} {
	p := &Property[interface{}]{
		AuditModel:  general_objects.AuditModelFromTF(tf.AuditModelTF),
		Name:        tf.Name.ValueString(),
		Description: tf.Description.ValueStringPointer(),
		IsBuiltIn:   tf.IsBuiltIn.ValueBool(),
		NodeId:      tf.NodeId.ValueString(),
		Tags:        general_objects.KeyValuesFromTF(tf.Tags),
	}
	p.Attributes.Type = tf.Type.ValueString()
	p.Attributes.Min = tf.Min.ValueFloat64Pointer()
	p.Attributes.Max = tf.Max.ValueFloat64Pointer()
	p.Attributes.Scale = helper.FromTFIntPtr(tf.Scale)
	p.Attributes.Precision = helper.FromTFIntPtr(tf.Precision)
	p.Attributes.UnitId = tf.UnitId.ValueStringPointer()
	p.Attributes.MinLength = helper.FromTFIntPtr(tf.MinLength)
	p.Attributes.MaxLength = helper.FromTFIntPtr(tf.MaxLength)
	p.Attributes.Pattern = tf.Pattern.ValueStringPointer()
	p.Attributes.Before = tf.Before.ValueStringPointer()
	p.Attributes.After = tf.After.ValueStringPointer()
	p.Attributes.Fields = general_objects.FieldsFromTF(tf.Fields, false)
	if !tf.Value.IsNull() && !tf.Value.IsUnknown() {
		p.Attributes.Value = tf.Value.ValueString()
	}
	if tf.Options != nil {
		opts := make(map[string]any, len(tf.Options))
		for k, v := range tf.Options {
			opts[k] = v.ValueString()
		}
		p.Attributes.Options = &opts
	}
	return p
}
