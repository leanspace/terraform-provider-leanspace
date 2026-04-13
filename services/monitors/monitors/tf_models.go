package monitors

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/leanspace/terraform-provider-leanspace/helper"
	"github.com/leanspace/terraform-provider-leanspace/helper/general_objects"
)

type MonitorTF struct {
	general_objects.AuditModelTF
	Name                types.String              `tfsdk:"name"`
	Description         types.String              `tfsdk:"description"`
	Status              types.String              `tfsdk:"status"`
	MetricId            types.String              `tfsdk:"metric_id"`
	NodeId              types.String              `tfsdk:"node_id"`
	Rule                *RuleTF                   `tfsdk:"rule"`
	ActionTemplates     []ActionTemplateTF        `tfsdk:"action_templates"`
	ActionTemplateLinks []ActionTemplateLinkTF    `tfsdk:"action_template_links"`
	Tags                []general_objects.KeyValueTF `tfsdk:"tags"`
	Type                types.String              `tfsdk:"type"`
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

	actionTemplates := make([]ActionTemplateTF, len(x.ActionTemplates))
	for i, at := range x.ActionTemplates {
		var headers map[string]types.String
		if at.Headers != nil {
			headers = make(map[string]types.String, len(at.Headers))
			for k, v := range at.Headers {
				headers[k] = helper.TFStringValue(v)
			}
		}
		actionTemplates[i] = ActionTemplateTF{
			AuditModelTF: general_objects.AuditModelToTF(&at.AuditModel),
			Name:         helper.TFStringValue(at.Name),
			Type:         helper.TFStringValue(at.Type),
			URL:          helper.TFStringValue(at.URL),
			Payload:      helper.TFStringValue(at.Payload),
			Headers:      headers,
			TriggeredOn:  helper.TFStringsValue(at.TriggeredOn),
		}
	}

	actionTemplateLinks := make([]ActionTemplateLinkTF, len(x.ActionTemplateLinks))
	for i, atl := range x.ActionTemplateLinks {
		actionTemplateLinks[i] = ActionTemplateLinkTF{
			ID:          helper.TFStringValue(atl.ID),
			TriggeredOn: helper.TFStringsValue(atl.TriggeredOn),
		}
	}

	return &MonitorTF{
		AuditModelTF:        general_objects.AuditModelToTF(&x.AuditModel),
		Name:                helper.TFStringValue(x.Name),
		Description:         helper.TFStringValue(x.Description),
		Status:              helper.TFStringValue(x.Status),
		MetricId:            helper.TFStringValue(x.MetricId),
		NodeId:              helper.TFStringValue(x.NodeId),
		Rule:                rule,
		ActionTemplates:     actionTemplates,
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

	actionTemplates := make([]ActionTemplate, len(tf.ActionTemplates))
	for i, at := range tf.ActionTemplates {
		var headers map[string]string
		if at.Headers != nil {
			headers = make(map[string]string, len(at.Headers))
			for k, v := range at.Headers {
				headers[k] = helper.FromTFString(v)
			}
		}
		actionTemplates[i] = ActionTemplate{
			AuditModel:  general_objects.AuditModelFromTF(at.AuditModelTF),
			Name:        helper.FromTFString(at.Name),
			Type:        helper.FromTFString(at.Type),
			URL:         helper.FromTFString(at.URL),
			Payload:     helper.FromTFString(at.Payload),
			Headers:     headers,
			TriggeredOn: helper.FromTFStrings(at.TriggeredOn),
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
		Description:         helper.FromTFString(tf.Description),
		Status:              helper.FromTFString(tf.Status),
		MetricId:            helper.FromTFString(tf.MetricId),
		NodeId:              helper.FromTFString(tf.NodeId),
		Rule:                rule,
		ActionTemplates:     actionTemplates,
		ActionTemplateLinks: actionTemplateLinks,
		Tags:                general_objects.KeyValuesFromTF(tf.Tags),
	}
}
