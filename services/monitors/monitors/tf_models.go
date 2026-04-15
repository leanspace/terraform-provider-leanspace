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
		ComparisonOperator: helper.TFStringValue(x.Rule.ComparisonOperator),
		ComparisonValue:    helper.TFFloat64Value(x.Rule.ComparisonValue),
		Tolerance:          helper.TFFloat64PtrValue(x.Rule.Tolerance),
	}

	atElems := make([]attr.Value, len(x.ActionTemplates))
	for i, at := range x.ActionTemplates {
		headers := map[string]attr.Value{}
		for k, v := range at.Headers {
			headers[k] = helper.TFStringValue(v)
		}
		headersMap := types.MapValueMust(types.StringType, headers)

		triggeredOn := make([]attr.Value, len(at.TriggeredOn))
		for j, t := range at.TriggeredOn {
			triggeredOn[j] = helper.TFStringValue(t)
		}
		triggeredOnSet := types.SetValueMust(types.StringType, triggeredOn)

		am := general_objects.AuditModelToTF(&at.AuditModel)
		atObj, _ := types.ObjectValue(actionTemplateAttrTypes, map[string]attr.Value{
			"id":               am.ID,
			"created_at":       am.CreatedAt,
			"created_by":       am.CreatedBy,
			"last_modified_at": am.LastModifiedAt,
			"last_modified_by": am.LastModifiedBy,
			"name":             helper.TFStringValue(at.Name),
			"type":             helper.TFStringValue(at.Type),
			"url":              helper.TFStringValue(at.URL),
			"payload":          helper.TFStringValue(at.Payload),
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
			ID:          helper.TFStringValue(atl.ID),
			TriggeredOn: helper.TFStringsValue(triggeredOn),
		}
	}

	return &MonitorTF{
		AuditModelTF:        general_objects.AuditModelToTF(&x.AuditModel),
		Name:                helper.TFStringValue(x.Name),
		Description:         helper.TFStringPtrValue(x.Description),
		Status:              helper.TFStringValue(x.Status),
		MetricId:            helper.TFStringValue(x.MetricId),
		NodeId:              helper.TFStringValue(x.NodeId),
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
			ComparisonOperator: helper.FromTFString(tf.Rule.ComparisonOperator),
			ComparisonValue:    helper.FromTFFloat64(tf.Rule.ComparisonValue),
			Tolerance:          helper.FromTFFloat64Ptr(tf.Rule.Tolerance),
		}
	}

	actionTemplateLinks := make([]ActionTemplateLink, len(tf.ActionTemplateLinks))
	for i, atl := range tf.ActionTemplateLinks {
		actionTemplateLinks[i] = ActionTemplateLink{
			ID:          helper.FromTFString(atl.ID),
			TriggeredOn: helper.FromTFStrings(atl.TriggeredOn),
		}
	}

	return &Monitor{
		AuditModel:          general_objects.AuditModelFromTF(tf.AuditModelTF),
		Name:                helper.FromTFString(tf.Name),
		Description:         helper.FromTFStringPtr(tf.Description),
		Status:              helper.FromTFString(tf.Status),
		MetricId:            helper.FromTFString(tf.MetricId),
		NodeId:              helper.FromTFString(tf.NodeId),
		Rule:                rule,
		ActionTemplateLinks: actionTemplateLinks,
		Tags:                general_objects.KeyValuesFromTF(tf.Tags),
	}
}
