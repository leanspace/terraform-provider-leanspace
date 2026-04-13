package request_definitions

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/leanspace/terraform-provider-leanspace/helper"
	"github.com/leanspace/terraform-provider-leanspace/helper/general_objects"
	"github.com/leanspace/terraform-provider-leanspace/services/activities/activity_definitions"
)

type RequestDefinitionTF struct {
	general_objects.AuditModelTF
	Name                             types.String                                `tfsdk:"name"`
	Description                      types.String                                `tfsdk:"description"`
	PlanTemplateIds                  []types.String                              `tfsdk:"plan_template_ids"`
	FeasibilityConstraintDefinitions []FeasibilityConstraintDefinitionTF         `tfsdk:"feasibility_constraint_definitions"`
	ConfigurationArgumentDefinitions []activity_definitions.ArgumentDefinitionTF `tfsdk:"configuration_argument_definitions"`
	ConfigurationArgumentMappings    []ArgumentMappingTF                         `tfsdk:"configuration_argument_mappings"`
}

type FeasibilityConstraintDefinitionTF struct {
	ID             types.String                                `tfsdk:"id"`
	CreatedAt      types.String                                `tfsdk:"created_at"`
	CreatedBy      types.String                                `tfsdk:"created_by"`
	LastModifiedAt types.String                                `tfsdk:"last_modified_at"`
	LastModifiedBy types.String                                `tfsdk:"last_modified_by"`
	Name           types.String                                `tfsdk:"name"`
	Description    types.String                                `tfsdk:"description"`
	Required       types.Bool                                  `tfsdk:"required"`
	ArgumentDefinitions []activity_definitions.ArgumentDefinitionTF `tfsdk:"argument_definitions"`
}

type ArgumentMappingTF struct {
	PlanTemplateId                           types.String `tfsdk:"plan_template_id"`
	ActivityDefinitionPosition               types.Int64  `tfsdk:"activity_definition_position"`
	ConfigurationArgumentDefinitionName      types.String `tfsdk:"configuration_argument_definition_name"`
	ActivityDefinitionArgumentDefinitionName types.String `tfsdk:"activity_definition_argument_definition_name"`
}

func (x *RequestDefinition) ToTF() interface{} {
	fcds := make([]FeasibilityConstraintDefinitionTF, len(x.FeasibilityConstraintDefinitions))
	for i, fcd := range x.FeasibilityConstraintDefinitions {
		argDefs := make([]activity_definitions.ArgumentDefinitionTF, len(fcd.ArgumentDefinitions))
		for j, a := range fcd.ArgumentDefinitions {
			attr := general_objects.DefinitionAttributeToTF(&a.Attributes)
			argDefs[j] = activity_definitions.ArgumentDefinitionTF{
				Name:        helper.TFStringValue(a.Name),
				Description: helper.TFStringValue(a.Description),
				Attributes:  &attr,
			}
		}
		fcds[i] = FeasibilityConstraintDefinitionTF{
			ID:                  helper.TFStringValue(fcd.AuditModel.ID),
			CreatedAt:           helper.TFStringValue(fcd.AuditModel.CreatedAt),
			CreatedBy:           helper.TFStringValue(fcd.AuditModel.CreatedBy),
			LastModifiedAt:      helper.TFStringValue(fcd.AuditModel.LastModifiedAt),
			LastModifiedBy:      helper.TFStringValue(fcd.AuditModel.LastModifiedBy),
			Name:                helper.TFStringValue(fcd.Name),
			Description:        helper.TFStringValue(fcd.Description),
			Required:           helper.TFBoolValue(fcd.Required),
			ArgumentDefinitions: argDefs,
		}
	}

	configArgDefs := make([]activity_definitions.ArgumentDefinitionTF, len(x.ConfigurationArgumentDefinitions))
	for i, a := range x.ConfigurationArgumentDefinitions {
		attr := general_objects.DefinitionAttributeToTF(&a.Attributes)
		configArgDefs[i] = activity_definitions.ArgumentDefinitionTF{
			Name:        helper.TFStringValue(a.Name),
			Description: helper.TFStringValue(a.Description),
			Attributes:  &attr,
		}
	}

	mappings := make([]ArgumentMappingTF, len(x.ConfigurationArgumentMappings))
	for i, m := range x.ConfigurationArgumentMappings {
		mappings[i] = ArgumentMappingTF{
			PlanTemplateId:                           helper.TFStringValue(m.PlanTemplateId),
			ActivityDefinitionPosition:               helper.TFInt64Value(m.ActivityDefinitionPosition),
			ConfigurationArgumentDefinitionName:      helper.TFStringValue(m.ConfigurationArgumentDefinitionName),
			ActivityDefinitionArgumentDefinitionName: helper.TFStringValue(m.ActivityDefinitionArgumentDefinitionName),
		}
	}

	return &RequestDefinitionTF{
		AuditModelTF:                     general_objects.AuditModelToTF(&x.AuditModel),
		Name:                             helper.TFStringValue(x.Name),
		Description:                      helper.TFStringValue(x.Description),
		PlanTemplateIds:                  helper.TFStringsValue(x.PlanTemplateIds),
		FeasibilityConstraintDefinitions: fcds,
		ConfigurationArgumentDefinitions: configArgDefs,
		ConfigurationArgumentMappings:    mappings,
	}
}

func (tf *RequestDefinitionTF) ToAPI() interface{} {
	fcds := make([]FeasibilityConstraintDefinition, len(tf.FeasibilityConstraintDefinitions))
	for i, fcd := range tf.FeasibilityConstraintDefinitions {
		argDefs := make([]activity_definitions.ArgumentDefinition[any], len(fcd.ArgumentDefinitions))
		for j, a := range fcd.ArgumentDefinitions {
			argDefs[j] = activity_definitions.ArgumentDefinition[any]{
				Name:        helper.FromTFString(a.Name),
				Description: helper.FromTFString(a.Description),
			}
			if a.Attributes != nil {
				argDefs[j].Attributes = general_objects.DefinitionAttributeFromTF(*a.Attributes)
			}
		}
		fcds[i] = FeasibilityConstraintDefinition{
			AuditModel: general_objects.AuditModel{
				ID:             helper.FromTFString(fcd.ID),
				CreatedAt:      helper.FromTFString(fcd.CreatedAt),
				CreatedBy:      helper.FromTFString(fcd.CreatedBy),
				LastModifiedAt: helper.FromTFString(fcd.LastModifiedAt),
				LastModifiedBy: helper.FromTFString(fcd.LastModifiedBy),
			},
			Name:                helper.FromTFString(fcd.Name),
			Description:         helper.FromTFString(fcd.Description),
			Required:            helper.FromTFBool(fcd.Required),
			ArgumentDefinitions: argDefs,
		}
	}

	configArgDefs := make([]activity_definitions.ArgumentDefinition[any], len(tf.ConfigurationArgumentDefinitions))
	for i, a := range tf.ConfigurationArgumentDefinitions {
		configArgDefs[i] = activity_definitions.ArgumentDefinition[any]{
			Name:        helper.FromTFString(a.Name),
			Description: helper.FromTFString(a.Description),
		}
		if a.Attributes != nil {
			configArgDefs[i].Attributes = general_objects.DefinitionAttributeFromTF(*a.Attributes)
		}
	}

	mappings := make([]ArgumentMapping, len(tf.ConfigurationArgumentMappings))
	for i, m := range tf.ConfigurationArgumentMappings {
		mappings[i] = ArgumentMapping{
			PlanTemplateId:                           helper.FromTFString(m.PlanTemplateId),
			ActivityDefinitionPosition:               helper.FromTFInt64(m.ActivityDefinitionPosition),
			ConfigurationArgumentDefinitionName:      helper.FromTFString(m.ConfigurationArgumentDefinitionName),
			ActivityDefinitionArgumentDefinitionName: helper.FromTFString(m.ActivityDefinitionArgumentDefinitionName),
		}
	}

	return &RequestDefinition{
		AuditModel:                       general_objects.AuditModelFromTF(tf.AuditModelTF),
		Name:                             helper.FromTFString(tf.Name),
		Description:                      helper.FromTFString(tf.Description),
		PlanTemplateIds:                  helper.FromTFStrings(tf.PlanTemplateIds),
		FeasibilityConstraintDefinitions: fcds,
		ConfigurationArgumentDefinitions: configArgDefs,
		ConfigurationArgumentMappings:    mappings,
	}
}
