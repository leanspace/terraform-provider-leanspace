package request_definitions

import (
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/leanspace/terraform-provider-leanspace/helper"
	"github.com/leanspace/terraform-provider-leanspace/helper/general_objects"
	"github.com/leanspace/terraform-provider-leanspace/services/activities/activity_definitions"
)

// attr.Type definitions for the computed argument_definitions nested structure.
var rdFieldDefAttrTypes = map[string]attr.Type{
	"default_value": types.StringType,
	"min":           types.Float64Type,
	"max":           types.Float64Type,
	"scale":         types.Int64Type,
	"precision":     types.Int64Type,
	"unit_id":       types.StringType,
}

var rdFieldsDefAttrTypes = map[string]attr.Type{
	"elevation": types.ObjectType{AttrTypes: rdFieldDefAttrTypes},
	"latitude":  types.ObjectType{AttrTypes: rdFieldDefAttrTypes},
	"longitude": types.ObjectType{AttrTypes: rdFieldDefAttrTypes},
}

var rdArrayConstraintAttrTypes = map[string]attr.Type{
	"type":       types.StringType,
	"required":   types.BoolType,
	"min_length": types.Int64Type,
	"max_length": types.Int64Type,
	"pattern":    types.StringType,
	"min":        types.Float64Type,
	"max":        types.Float64Type,
	"scale":      types.Int64Type,
	"precision":  types.Int64Type,
	"unit_id":    types.StringType,
	"before":     types.StringType,
	"after":      types.StringType,
	"options":    types.MapType{ElemType: types.StringType},
}

var rdDefAttrAttrTypes = map[string]attr.Type{
	"type":          types.StringType,
	"required":      types.BoolType,
	"default_value": types.StringType,
	"min_length":    types.Int64Type,
	"max_length":    types.Int64Type,
	"pattern":       types.StringType,
	"min":           types.Float64Type,
	"max":           types.Float64Type,
	"scale":         types.Int64Type,
	"precision":     types.Int64Type,
	"unit_id":       types.StringType,
	"before":        types.StringType,
	"after":         types.StringType,
	"options":       types.MapType{ElemType: types.StringType},
	"fields":        types.ObjectType{AttrTypes: rdFieldsDefAttrTypes},
	"min_size":      types.Int64Type,
	"max_size":      types.Int64Type,
	"unique":        types.BoolType,
	"constraint":    types.ObjectType{AttrTypes: rdArrayConstraintAttrTypes},
}

var rdArgDefAttrTypes = map[string]attr.Type{
	"name":        types.StringType,
	"description": types.StringType,
	"attributes":  types.ObjectType{AttrTypes: rdDefAttrAttrTypes},
}

// buildFieldDefAttrValue converts a *FieldDefTF to an attr.Value (ObjectValue or ObjectNull).
func buildFieldDefAttrValue(f *general_objects.FieldDefTF) attr.Value {
	if f == nil {
		return types.ObjectNull(rdFieldDefAttrTypes)
	}
	obj, _ := types.ObjectValue(rdFieldDefAttrTypes, map[string]attr.Value{
		"default_value": f.DefaultValue,
		"min":           f.Min,
		"max":           f.Max,
		"scale":         f.Scale,
		"precision":     f.Precision,
		"unit_id":       f.UnitId,
	})
	return obj
}

// buildDefinitionAttributeAttrValue converts a *DefinitionAttributeTF to an attr.Value.
func buildDefinitionAttributeAttrValue(tf *general_objects.DefinitionAttributeTF) attr.Value {
	if tf == nil {
		return types.ObjectNull(rdDefAttrAttrTypes)
	}

	// Build options map
	optionsMap := map[string]attr.Value{}
	for k, v := range tf.Options {
		optionsMap[k] = v
	}
	optionsVal, _ := types.MapValue(types.StringType, optionsMap)

	// Build fields object
	var fieldsVal attr.Value
	if tf.Fields == nil {
		fieldsVal = types.ObjectNull(rdFieldsDefAttrTypes)
	} else {
		fieldsObj, _ := types.ObjectValue(rdFieldsDefAttrTypes, map[string]attr.Value{
			"elevation": buildFieldDefAttrValue(tf.Fields.Elevation),
			"latitude":  buildFieldDefAttrValue(tf.Fields.Latitude),
			"longitude": buildFieldDefAttrValue(tf.Fields.Longitude),
		})
		fieldsVal = fieldsObj
	}

	// Build constraint object
	var constraintVal attr.Value
	if tf.Constraint == nil {
		constraintVal = types.ObjectNull(rdArrayConstraintAttrTypes)
	} else {
		constraintOptions := map[string]attr.Value{}
		for k, v := range tf.Constraint.Options {
			constraintOptions[k] = v
		}
		constraintOptionsVal, _ := types.MapValue(types.StringType, constraintOptions)
		constraintObj, _ := types.ObjectValue(rdArrayConstraintAttrTypes, map[string]attr.Value{
			"type":       tf.Constraint.Type,
			"required":   tf.Constraint.Required,
			"min_length": tf.Constraint.MinLength,
			"max_length": tf.Constraint.MaxLength,
			"pattern":    tf.Constraint.Pattern,
			"min":        tf.Constraint.Min,
			"max":        tf.Constraint.Max,
			"scale":      tf.Constraint.Scale,
			"precision":  tf.Constraint.Precision,
			"unit_id":    tf.Constraint.UnitId,
			"before":     tf.Constraint.Before,
			"after":      tf.Constraint.After,
			"options":    constraintOptionsVal,
		})
		constraintVal = constraintObj
	}

	obj, _ := types.ObjectValue(rdDefAttrAttrTypes, map[string]attr.Value{
		"type":          tf.Type,
		"required":      tf.Required,
		"default_value": tf.DefaultValue,
		"min_length":    tf.MinLength,
		"max_length":    tf.MaxLength,
		"pattern":       tf.Pattern,
		"min":           tf.Min,
		"max":           tf.Max,
		"scale":         tf.Scale,
		"precision":     tf.Precision,
		"unit_id":       tf.UnitId,
		"before":        tf.Before,
		"after":         tf.After,
		"options":       optionsVal,
		"fields":        fieldsVal,
		"min_size":      tf.MinSize,
		"max_size":      tf.MaxSize,
		"unique":        tf.Unique,
		"constraint":    constraintVal,
	})
	return obj
}

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
	ID                  types.String `tfsdk:"id"`
	CreatedAt           types.String `tfsdk:"created_at"`
	CreatedBy           types.String `tfsdk:"created_by"`
	LastModifiedAt      types.String `tfsdk:"last_modified_at"`
	LastModifiedBy      types.String `tfsdk:"last_modified_by"`
	Name                types.String `tfsdk:"name"`
	Description         types.String `tfsdk:"description"`
	Required            types.Bool   `tfsdk:"required"`
	ArgumentDefinitions types.List   `tfsdk:"argument_definitions"`
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
		argDefElems := make([]attr.Value, len(fcd.ArgumentDefinitions))
		for j, a := range fcd.ArgumentDefinitions {
			defAttr := general_objects.DefinitionAttributeToTF(&a.Attributes)
			argDefObj, _ := types.ObjectValue(rdArgDefAttrTypes, map[string]attr.Value{
				"name":        helper.TFStringValue(a.Name),
				"description": helper.TFStringPtrValue(a.Description),
				"attributes":  buildDefinitionAttributeAttrValue(&defAttr),
			})
			argDefElems[j] = argDefObj
		}
		argDefList := types.ListValueMust(types.ObjectType{AttrTypes: rdArgDefAttrTypes}, argDefElems)

		fcds[i] = FeasibilityConstraintDefinitionTF{
			ID:                  helper.TFStringValue(fcd.AuditModel.ID),
			CreatedAt:           helper.TFStringValue(fcd.AuditModel.CreatedAt),
			CreatedBy:           helper.TFStringValue(fcd.AuditModel.CreatedBy),
			LastModifiedAt:      helper.TFStringValue(fcd.AuditModel.LastModifiedAt),
			LastModifiedBy:      helper.TFStringValue(fcd.AuditModel.LastModifiedBy),
			Name:                helper.TFStringValue(fcd.Name),
			Description:         helper.TFStringPtrValue(fcd.Description),
			Required:            helper.TFBoolValue(fcd.Required),
			ArgumentDefinitions: argDefList,
		}
	}

	configArgDefs := make([]activity_definitions.ArgumentDefinitionTF, len(x.ConfigurationArgumentDefinitions))
	for i, a := range x.ConfigurationArgumentDefinitions {
		attrVal := general_objects.DefinitionAttributeToTF(&a.Attributes)
		configArgDefs[i] = activity_definitions.ArgumentDefinitionTF{
			Name:        helper.TFStringValue(a.Name),
			Description: helper.TFStringPtrValue(a.Description),
			Attributes:  &attrVal,
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
		Description:                      helper.TFStringPtrValue(x.Description),
		PlanTemplateIds:                  helper.TFStringsValue(x.PlanTemplateIds),
		FeasibilityConstraintDefinitions: fcds,
		ConfigurationArgumentDefinitions: configArgDefs,
		ConfigurationArgumentMappings:    mappings,
	}
}

func (tf *RequestDefinitionTF) ToAPI() interface{} {
	fcds := make([]FeasibilityConstraintDefinition, len(tf.FeasibilityConstraintDefinitions))
	for i, fcd := range tf.FeasibilityConstraintDefinitions {
		fcds[i] = FeasibilityConstraintDefinition{
			AuditModel: general_objects.AuditModel{
				ID:             helper.FromTFString(fcd.ID),
				CreatedAt:      helper.FromTFString(fcd.CreatedAt),
				CreatedBy:      helper.FromTFString(fcd.CreatedBy),
				LastModifiedAt: helper.FromTFString(fcd.LastModifiedAt),
				LastModifiedBy: helper.FromTFString(fcd.LastModifiedBy),
			},
			Name:        helper.FromTFString(fcd.Name),
			Description: helper.FromTFStringPtr(fcd.Description),
			Required:    helper.FromTFBool(fcd.Required),
		}
	}

	configArgDefs := make([]activity_definitions.ArgumentDefinition[any], len(tf.ConfigurationArgumentDefinitions))
	for i, a := range tf.ConfigurationArgumentDefinitions {
		configArgDefs[i] = activity_definitions.ArgumentDefinition[any]{
			Name:        helper.FromTFString(a.Name),
			Description: helper.FromTFStringPtr(a.Description),
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
		Description:                      helper.FromTFStringPtr(tf.Description),
		PlanTemplateIds:                  helper.FromTFStrings(tf.PlanTemplateIds),
		FeasibilityConstraintDefinitions: fcds,
		ConfigurationArgumentDefinitions: configArgDefs,
		ConfigurationArgumentMappings:    mappings,
	}
}
