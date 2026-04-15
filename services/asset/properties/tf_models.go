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
		Name:         helper.TFStringValue(x.Name),
		Description:  helper.TFStringPtrValue(x.Description),
		IsBuiltIn:    helper.TFBoolValue(x.IsBuiltIn),
		NodeId:       helper.TFStringValue(x.NodeId),
		Tags:         general_objects.KeyValuesToTF(x.Tags),
		Type:         helper.TFStringValue(x.Attributes.Type),
		Min:          helper.TFFloat64PtrValue(x.Attributes.Min),
		Max:          helper.TFFloat64PtrValue(x.Attributes.Max),
		Scale:        helper.TFIntPtrValue(x.Attributes.Scale),
		Precision:    helper.TFIntPtrValue(x.Attributes.Precision),
		UnitId:       helper.TFStringPtrValue(x.Attributes.UnitId),
		MinLength:    helper.TFIntPtrValue(x.Attributes.MinLength),
		MaxLength:    helper.TFIntPtrValue(x.Attributes.MaxLength),
		Pattern:      helper.TFStringPtrValue(x.Attributes.Pattern),
		Before:       helper.TFStringPtrValue(x.Attributes.Before),
		After:        helper.TFStringPtrValue(x.Attributes.After),
		Fields:       general_objects.FieldsToTF(x.Attributes.Fields),
	}
	// Value handling by type
	val := interface{}(x.Attributes.Value)
	switch x.Attributes.Type {
	case "NUMERIC", "ENUM":
		if val != nil {
			if f, ok := val.(float64); ok {
				tf.Value = helper.TFStringValue(helper.ParseFloat(f))
			}
		}
	case "BOOLEAN":
		if val != nil {
			if b, ok := val.(bool); ok {
				tf.Value = helper.TFStringValue(strconv.FormatBool(b))
			}
		}
	case "TEXT", "TIMESTAMP", "DATE", "TIME":
		if val != nil {
			tf.Value = helper.TFStringValue(fmt.Sprint(val))
		}
	case "TLE":
		if val != nil {
			if tleValues, ok := val.([]interface{}); ok {
				var tleValue string
				for _, value := range tleValues {
					tleValue = tleValue + "," + fmt.Sprint(value)
				}
				tf.Value = helper.TFStringValue(strings.TrimPrefix(tleValue, ","))
			}
		}
	}
	if x.Attributes.Options != nil {
		tf.Options = make(map[string]types.String, len(*x.Attributes.Options))
		for k, v := range *x.Attributes.Options {
			tf.Options[k] = helper.TFStringValue(fmt.Sprint(v))
		}
	}
	return tf
}

func (tf *PropertyTF) ToAPI() interface{} {
	p := &Property[interface{}]{
		AuditModel:  general_objects.AuditModelFromTF(tf.AuditModelTF),
		Name:        helper.FromTFString(tf.Name),
		Description: helper.FromTFStringPtr(tf.Description),
		IsBuiltIn:   helper.FromTFBool(tf.IsBuiltIn),
		NodeId:      helper.FromTFString(tf.NodeId),
		Tags:        general_objects.KeyValuesFromTF(tf.Tags),
	}
	p.Attributes.Type = helper.FromTFString(tf.Type)
	p.Attributes.Min = helper.FromTFFloat64Ptr(tf.Min)
	p.Attributes.Max = helper.FromTFFloat64Ptr(tf.Max)
	p.Attributes.Scale = helper.FromTFIntPtr(tf.Scale)
	p.Attributes.Precision = helper.FromTFIntPtr(tf.Precision)
	p.Attributes.UnitId = helper.FromTFStringPtr(tf.UnitId)
	p.Attributes.MinLength = helper.FromTFIntPtr(tf.MinLength)
	p.Attributes.MaxLength = helper.FromTFIntPtr(tf.MaxLength)
	p.Attributes.Pattern = helper.FromTFStringPtr(tf.Pattern)
	p.Attributes.Before = helper.FromTFStringPtr(tf.Before)
	p.Attributes.After = helper.FromTFStringPtr(tf.After)
	p.Attributes.Fields = general_objects.FieldsFromTF(tf.Fields)
	if !tf.Value.IsNull() && !tf.Value.IsUnknown() {
		p.Attributes.Value = tf.Value.ValueString()
	}
	if tf.Options != nil {
		opts := make(map[string]any, len(tf.Options))
		for k, v := range tf.Options {
			opts[k] = helper.FromTFString(v)
		}
		p.Attributes.Options = &opts
	}
	return p
}
