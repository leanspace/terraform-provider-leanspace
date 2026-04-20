package metrics

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/leanspace/terraform-provider-leanspace/helper"
	"github.com/leanspace/terraform-provider-leanspace/helper/general_objects"
)

type MetricTF struct {
	general_objects.AuditModelTF
	Name            types.String                 `tfsdk:"name"`
	Description     types.String                 `tfsdk:"description"`
	NodeId          types.String                 `tfsdk:"node_id"`
	AncestorAssetId types.String                 `tfsdk:"ancestor_asset_id"`
	Tags            []general_objects.KeyValueTF `tfsdk:"tags"`
	Attributes      *MetricAttributeTF           `tfsdk:"attributes"`
}

// MetricAttributeTF is a local variant of DefinitionAttributeTF
// without "required" and "default_value" (excluded from the metric schema).
type MetricAttributeTF struct {
	Type       types.String                       `tfsdk:"type"`
	MinLength  types.Int64                        `tfsdk:"min_length"`
	MaxLength  types.Int64                        `tfsdk:"max_length"`
	Pattern    types.String                       `tfsdk:"pattern"`
	Min        types.Float64                      `tfsdk:"min"`
	Max        types.Float64                      `tfsdk:"max"`
	Scale      types.Int64                        `tfsdk:"scale"`
	Precision  types.Int64                        `tfsdk:"precision"`
	UnitId     types.String                       `tfsdk:"unit_id"`
	Before     types.String                       `tfsdk:"before"`
	After      types.String                       `tfsdk:"after"`
	Options    map[string]types.String            `tfsdk:"options"`
	Fields     *general_objects.FieldsDefTF       `tfsdk:"fields"`
	MinSize    types.Int64                        `tfsdk:"min_size"`
	MaxSize    types.Int64                        `tfsdk:"max_size"`
	Unique     types.Bool                         `tfsdk:"unique"`
	Constraint *general_objects.ArrayConstraintTF `tfsdk:"constraint"`
}

func metricAttributeToTF(a *general_objects.DefinitionAttribute[any]) *MetricAttributeTF {
	tf := &MetricAttributeTF{
		Type:      types.StringValue(a.Type),
		MinLength: helper.TFIntPtrValue(a.MinLength),
		MaxLength: helper.TFIntPtrValue(a.MaxLength),
		Pattern:   helper.TFStringPtrValue(a.Pattern),
		Min:       helper.TFFloat64PtrValue(a.Min),
		Max:       helper.TFFloat64PtrValue(a.Max),
		Scale:     helper.TFIntPtrValue(a.Scale),
		Precision: helper.TFIntPtrValue(a.Precision),
		UnitId:    helper.TFStringPtrValue(a.UnitId),
		Before:    helper.TFStringPtrValue(a.Before),
		After:     helper.TFStringPtrValue(a.After),
		Fields:    general_objects.FieldsDefToTF(a.Fields),
		MinSize:   helper.TFIntPtrValue(a.MinSize),
		MaxSize:   helper.TFIntPtrValue(a.MaxSize),
		Unique:    types.BoolValue(a.Unique),
	}
	if a.Options != nil {
		tf.Options = make(map[string]types.String, len(*a.Options))
		for k, v := range *a.Options {
			tf.Options[k] = types.StringValue(fmt.Sprint(v))
		}
	}
	tf.Constraint = general_objects.ArrayConstraintToTF(&a.Constraint)
	return tf
}

func metricAttributeFromTF(tf *MetricAttributeTF) general_objects.DefinitionAttribute[any] {
	if tf == nil {
		return general_objects.DefinitionAttribute[any]{}
	}
	a := general_objects.DefinitionAttribute[any]{
		Type:      helper.FromTFString(tf.Type),
		MinLength: helper.FromTFIntPtr(tf.MinLength),
		MaxLength: helper.FromTFIntPtr(tf.MaxLength),
		Pattern:   helper.FromTFStringPtr(tf.Pattern),
		Min:       helper.FromTFFloat64Ptr(tf.Min),
		Max:       helper.FromTFFloat64Ptr(tf.Max),
		Scale:     helper.FromTFIntPtr(tf.Scale),
		Precision: helper.FromTFIntPtr(tf.Precision),
		UnitId:    helper.FromTFStringPtr(tf.UnitId),
		Before:    helper.FromTFStringPtr(tf.Before),
		After:     helper.FromTFStringPtr(tf.After),
		Fields:    general_objects.FieldsDefFromTF(tf.Fields),
		MinSize:   helper.FromTFIntPtr(tf.MinSize),
		MaxSize:   helper.FromTFIntPtr(tf.MaxSize),
		Unique:    helper.FromTFBool(tf.Unique),
	}
	if tf.Options != nil {
		opts := make(map[string]any, len(tf.Options))
		for k, v := range tf.Options {
			opts[k] = helper.FromTFString(v)
		}
		a.Options = &opts
	}
	a.Constraint = general_objects.ArrayConstraintFromTF(tf.Constraint)
	return a
}

func (x *Metric[T]) ToTF() interface{} {
	genAttr := general_objects.DefinitionAttribute[interface{}]{
		Type: x.Attributes.Type, Required: x.Attributes.Required, DefaultValue: interface{}(x.Attributes.DefaultValue),
		MinLength: x.Attributes.MinLength, MaxLength: x.Attributes.MaxLength, Pattern: x.Attributes.Pattern,
		Min: x.Attributes.Min, Max: x.Attributes.Max, Scale: x.Attributes.Scale, Precision: x.Attributes.Precision,
		UnitId: x.Attributes.UnitId, Before: x.Attributes.Before, After: x.Attributes.After,
		Options: x.Attributes.Options, Fields: x.Attributes.Fields,
		MinSize: x.Attributes.MinSize, MaxSize: x.Attributes.MaxSize, Unique: x.Attributes.Unique,
	}
	return &MetricTF{
		AuditModelTF:    general_objects.AuditModelToTF(&x.AuditModel),
		Name:            types.StringValue(x.Name),
		Description:     helper.TFStringPtrValue(x.Description),
		NodeId:          types.StringValue(x.NodeId),
		AncestorAssetId: helper.TFStringPtrValue(x.AncestorAssetId),
		Tags:            general_objects.KeyValuesToTF(x.Tags),
		Attributes:      metricAttributeToTF(&genAttr),
	}
}

func (tf *MetricTF) ToAPI() interface{} {
	return &Metric[interface{}]{
		AuditModel:      general_objects.AuditModelFromTF(tf.AuditModelTF),
		Name:            helper.FromTFString(tf.Name),
		Description:     helper.FromTFStringPtr(tf.Description),
		NodeId:          helper.FromTFString(tf.NodeId),
		AncestorAssetId: helper.FromTFStringPtr(tf.AncestorAssetId),
		Tags:            general_objects.KeyValuesFromTF(tf.Tags),
		Attributes:      metricAttributeFromTF(tf.Attributes),
	}
}
