package resources

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/leanspace/terraform-provider-leanspace/helper"
	"github.com/leanspace/terraform-provider-leanspace/helper/general_objects"
)

type ResourceTF struct {
	general_objects.AuditModelTF
	AssetId      types.String                `tfsdk:"asset_id"`
	UnitId       types.String                `tfsdk:"unit_id"`
	MetricId     types.String                `tfsdk:"metric_id"`
	Name         types.String                `tfsdk:"name"`
	Description  types.String                `tfsdk:"description"`
	DefaultLevel types.Float64               `tfsdk:"default_level"`
	LowerLimit   types.Float64               `tfsdk:"lower_limit"`
	UpperLimit   types.Float64               `tfsdk:"upper_limit"`
	Thresholds   []ResourceThresholdTF       `tfsdk:"thresholds"`
	Tags         []general_objects.KeyValueTF `tfsdk:"tags"`
}

type ResourceThresholdTF struct {
	Kind                 types.String  `tfsdk:"kind"`
	Name                 types.String  `tfsdk:"name"`
	ViolationWhenReached types.Bool    `tfsdk:"violation_when_reached"`
	Value                types.Float64 `tfsdk:"value"`
}

func (x *Resource) ToTF() any {
	thresholds := make([]ResourceThresholdTF, len(x.Thresholds))
	for i, t := range x.Thresholds {
		thresholds[i] = ResourceThresholdTF{
			Kind:                 helper.TFStringValue(t.Kind),
			Name:                 helper.TFStringValue(t.Name),
			ViolationWhenReached: helper.TFBoolValue(t.ViolationWhenReached),
			Value:                helper.TFFloat64Value(t.Value),
		}
	}
	return &ResourceTF{
		AuditModelTF: general_objects.AuditModelToTF(&x.AuditModel),
		AssetId:      helper.TFStringValue(x.AssetId),
		UnitId:       helper.TFStringValue(x.UnitId),
		MetricId:     helper.TFStringValue(x.MetricId),
		Name:         helper.TFStringValue(x.Name),
		Description:  helper.TFStringValue(x.Description),
		DefaultLevel: helper.TFFloat64Value(x.DefaultLevel),
		LowerLimit:   helper.TFFloat64PtrValue(x.LowerLimit),
		UpperLimit:   helper.TFFloat64PtrValue(x.UpperLimit),
		Thresholds:   thresholds,
		Tags:         general_objects.KeyValuesToTF(x.Tags),
	}
}

func (tf *ResourceTF) ToAPI() any {
	thresholds := make([]ResourceThreshold, len(tf.Thresholds))
	for i, t := range tf.Thresholds {
		thresholds[i] = ResourceThreshold{
			Kind:                 helper.FromTFString(t.Kind),
			Name:                 helper.FromTFString(t.Name),
			ViolationWhenReached: helper.FromTFBool(t.ViolationWhenReached),
			Value:                helper.FromTFFloat64(t.Value),
		}
	}
	return &Resource{
		AuditModel:   general_objects.AuditModelFromTF(tf.AuditModelTF),
		AssetId:      helper.FromTFString(tf.AssetId),
		UnitId:       helper.FromTFString(tf.UnitId),
		MetricId:     helper.FromTFString(tf.MetricId),
		Name:         helper.FromTFString(tf.Name),
		Description:  helper.FromTFString(tf.Description),
		DefaultLevel: helper.FromTFFloat64(tf.DefaultLevel),
		LowerLimit:   helper.FromTFFloat64Ptr(tf.LowerLimit),
		UpperLimit:   helper.FromTFFloat64Ptr(tf.UpperLimit),
		Thresholds:   thresholds,
		Tags:         general_objects.KeyValuesFromTF(tf.Tags),
	}
}
