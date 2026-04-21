package monitors

import (
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/leanspace/terraform-provider-leanspace/helper"
	"github.com/leanspace/terraform-provider-leanspace/helper/general_objects"
)

var actionTemplateAttrTypes = map[string]attr.Type{
	"id":               types.StringType,
	"created_at":       types.StringType,
	"created_by":       types.StringType,
	"last_modified_at": types.StringType,
	"last_modified_by": types.StringType,
	"name":             types.StringType,
	"type":             types.StringType,
	"url":              types.StringType,
	"payload":          types.StringType,
	"content":          types.StringType,
	"headers":          types.MapType{ElemType: types.StringType},
	"triggered_on":     types.SetType{ElemType: types.StringType},
}

type MonitorTF struct {
	general_objects.AuditModelTF
	Name                types.String                 `tfsdk:"name"`
	Description         types.String                 `tfsdk:"description"`
	Status              types.String                 `tfsdk:"status"`
	MetricId            types.String                 `tfsdk:"metric_id"`
	NodeId              types.String                 `tfsdk:"node_id"`
	Rule                *RuleTF                      `tfsdk:"rule"`
	ActionTemplates     types.Set                    `tfsdk:"action_templates"`
	ActionTemplateLinks []ActionTemplateLinkTF       `tfsdk:"action_template_links"`
	Tags                []general_objects.KeyValueTF `tfsdk:"tags"`
	Type                types.String                 `tfsdk:"type"`
}

type RuleTF struct {
	ComparisonOperator types.String  `tfsdk:"comparison_operator"`
	ComparisonValue    types.Float64 `tfsdk:"comparison_value"`
	Tolerance          types.Float64 `tfsdk:"tolerance"`
}

type ActionTemplateLinkTF struct {
	ID          types.String   `tfsdk:"id"`
	TriggeredOn []types.String `tfsdk:"triggered_on"`
}

type ActionTemplateTF struct {
	general_objects.AuditModelTF
	Name        types.String            `tfsdk:"name"`
	Type        types.String            `tfsdk:"type"`
	URL         types.String            `tfsdk:"url"`
	Payload     types.String            `tfsdk:"payload"`
	Content     types.String            `tfsdk:"content"`
	Headers     map[string]types.String `tfsdk:"headers"`
	TriggeredOn []types.String          `tfsdk:"triggered_on"`
}

func (x *Monitor) ToTF() any {
	var rule *RuleTF
	rule = &RuleTF{
		ComparisonOperator: types.StringValue(x.Rule.ComparisonOperator),
		ComparisonValue:    types.Float64Value(x.Rule.ComparisonValue),
		Tolerance:          types.Float64PointerValue(x.Rule.Tolerance),
	}

	atElems := make([]attr.Value, len(x.ActionTemplates))
	for i, at := range x.ActionTemplates {
		headers := map[string]attr.Value{}
		for k, v := range at.Headers {
			headers[k] = types.StringValue(v)
		}
		headersMap := types.MapValueMust(types.StringType, headers)

		triggeredOn := make([]attr.Value, len(at.TriggeredOn))
		for j, t := range at.TriggeredOn {
			triggeredOn[j] = types.StringValue(t)
		}
		triggeredOnSet := types.SetValueMust(types.StringType, triggeredOn)

		am := general_objects.AuditModelToTF(&at.AuditModel)
		atObj, _ := types.ObjectValue(actionTemplateAttrTypes, map[string]attr.Value{
			"id":               am.ID,
			"created_at":       am.CreatedAt,
			"created_by":       am.CreatedBy,
			"last_modified_at": am.LastModifiedAt,
			"last_modified_by": am.LastModifiedBy,
			"name":             types.StringValue(at.Name),
			"type":             types.StringValue(at.Type),
			"url":              types.StringValue(at.URL),
			"payload":          types.StringValue(at.Payload),
			"content":          types.StringNull(),
			"headers":          headersMap,
			"triggered_on":     triggeredOnSet,
		})
		atElems[i] = atObj
	}
	actionTemplatesSet := types.SetValueMust(types.ObjectType{AttrTypes: actionTemplateAttrTypes}, atElems)

	actionTemplateLinks := make([]ActionTemplateLinkTF, len(x.ActionTemplateLinks))
	for i, atl := range x.ActionTemplateLinks {
		var triggeredOn []string
		if len(atl.TriggeredOn) > 0 {
			triggeredOn = atl.TriggeredOn
		}
		actionTemplateLinks[i] = ActionTemplateLinkTF{
			ID:          types.StringValue(atl.ID),
			TriggeredOn: helper.TFStringsValue(triggeredOn),
		}
	}

	return &MonitorTF{
		AuditModelTF:        general_objects.AuditModelToTF(&x.AuditModel),
		Name:                types.StringValue(x.Name),
		Description:         types.StringPointerValue(x.Description),
		Status:              types.StringValue(x.Status),
		MetricId:            types.StringValue(x.MetricId),
		NodeId:              types.StringValue(x.NodeId),
		Rule:                rule,
		ActionTemplates:     actionTemplatesSet,
		ActionTemplateLinks: actionTemplateLinks,
		Tags:                general_objects.KeyValuesToTF(x.Tags),
	}
}

func (tf *MonitorTF) ToAPI() any {
	var rule Rule
	if tf.Rule != nil {
		rule = Rule{
			ComparisonOperator: tf.Rule.ComparisonOperator.ValueString(),
			ComparisonValue:    tf.Rule.ComparisonValue.ValueFloat64(),
			Tolerance:          tf.Rule.Tolerance.ValueFloat64Pointer(),
		}
	}

	actionTemplateLinks := make([]ActionTemplateLink, len(tf.ActionTemplateLinks))
	for i, atl := range tf.ActionTemplateLinks {
		actionTemplateLinks[i] = ActionTemplateLink{
			ID:          atl.ID.ValueString(),
			TriggeredOn: helper.FromTFStrings(atl.TriggeredOn),
		}
	}

	return &Monitor{
		AuditModel:          general_objects.AuditModelFromTF(tf.AuditModelTF),
		Name:                tf.Name.ValueString(),
		Description:         tf.Description.ValueStringPointer(),
		Status:              tf.Status.ValueString(),
		MetricId:            tf.MetricId.ValueString(),
		NodeId:              tf.NodeId.ValueString(),
		Rule:                rule,
		ActionTemplateLinks: actionTemplateLinks,
		Tags:                general_objects.KeyValuesFromTF(tf.Tags),
	}
}
