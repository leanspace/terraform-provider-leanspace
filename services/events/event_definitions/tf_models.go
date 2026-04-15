package event_definitions

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/leanspace/terraform-provider-leanspace/helper"
	"github.com/leanspace/terraform-provider-leanspace/helper/general_objects"
)

type EventsDefinitionTF struct {
	general_objects.AuditModelTF
	Name        types.String                 `tfsdk:"name"`
	Source      types.String                 `tfsdk:"source"`
	State       types.String                 `tfsdk:"state"`
	Description types.String                 `tfsdk:"description"`
	Criticality types.String                 `tfsdk:"criticality"`
	Rules       []RulesTF                    `tfsdk:"rules"`
	Tags        []general_objects.KeyValueTF `tfsdk:"tags"`
}

type RulesTF struct {
	Operator        types.String       `tfsdk:"operator"`
	Path            types.String       `tfsdk:"path"`
	ComparisonValue *ComparisonValueTF `tfsdk:"comparison_value"`
}

type ComparisonValueTF struct {
	Value types.String `tfsdk:"value"`
	Type  types.String `tfsdk:"type"`
}

func (x *EventsDefinition) ToTF() interface{} {
	rules := make([]RulesTF, len(x.Rules))
	for i, r := range x.Rules {
		var cv *ComparisonValueTF
		if r.ComparisonValue != nil {
			cv = &ComparisonValueTF{
				Value: helper.TFStringValue(fmt.Sprint(r.ComparisonValue.Value)),
				Type:  helper.TFStringValue(r.ComparisonValue.Type),
			}
		}
		rules[i] = RulesTF{
			Operator:        helper.TFStringValue(r.Operator),
			Path:            helper.TFStringValue(r.Path),
			ComparisonValue: cv,
		}
	}

	return &EventsDefinitionTF{
		AuditModelTF: general_objects.AuditModelToTF(&x.AuditModel),
		Name:         helper.TFStringValue(x.Name),
		Source:       helper.TFStringValue(x.Source),
		State:        helper.TFStringValue(x.State),
		Description:  helper.TFStringPtrValue(x.Description),
		Criticality:  helper.TFStringPtrValue(x.Criticality),
		Rules:        rules,
		Tags:         general_objects.KeyValuesToTF(x.Tags),
	}
}

func (tf *EventsDefinitionTF) ToAPI() interface{} {
	rules := make([]Rules[any], len(tf.Rules))
	for i, r := range tf.Rules {
		rules[i] = Rules[any]{
			Operator: helper.FromTFString(r.Operator),
			Path:     helper.FromTFString(r.Path),
		}
		if r.ComparisonValue != nil {
			rules[i].ComparisonValue = &ComparisonValue[any]{
				Value: helper.FromTFString(r.ComparisonValue.Value),
				Type:  helper.FromTFString(r.ComparisonValue.Type),
			}
		}
	}

	return &EventsDefinition{
		AuditModel:  general_objects.AuditModelFromTF(tf.AuditModelTF),
		Name:        helper.FromTFString(tf.Name),
		Source:      helper.FromTFString(tf.Source),
		State:       helper.FromTFString(tf.State),
		Description:  helper.FromTFStringPtr(tf.Description),
		Criticality: helper.FromTFStringPtr(tf.Criticality),
		Rules:       rules,
		Tags:        general_objects.KeyValuesFromTF(tf.Tags),
	}
}
