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
				"name":        types.StringValue(a.Name),
				"description": types.StringPointerValue(a.Description),
				"attributes":  buildDefinitionAttributeAttrValue(&defAttr),
			})
			argDefElems[j] = argDefObj
		}
		argDefList := types.ListValueMust(types.ObjectType{AttrTypes: rdArgDefAttrTypes}, argDefElems)

		fcds[i] = FeasibilityConstraintDefinitionTF{
			ID:                  types.StringValue(fcd.AuditModel.ID),
			CreatedAt:           types.StringValue(fcd.AuditModel.CreatedAt),
			CreatedBy:           types.StringValue(fcd.AuditModel.CreatedBy),
			LastModifiedAt:      types.StringValue(fcd.AuditModel.LastModifiedAt),
			LastModifiedBy:      types.StringValue(fcd.AuditModel.LastModifiedBy),
			Name:                types.StringValue(fcd.Name),
			Description:         types.StringPointerValue(fcd.Description),
			Required:            types.BoolValue(fcd.Required),
			ArgumentDefinitions: argDefList,
		}
	}

	configArgDefs := make([]activity_definitions.ArgumentDefinitionTF, len(x.ConfigurationArgumentDefinitions))
	for i, a := range x.ConfigurationArgumentDefinitions {
		attrVal := general_objects.DefinitionAttributeToTF(&a.Attributes)
		configArgDefs[i] = activity_definitions.ArgumentDefinitionTF{
			Name:        types.StringValue(a.Name),
			Description: types.StringPointerValue(a.Description),
			Attributes:  &attrVal,
		}
	}

	mappings := make([]ArgumentMappingTF, len(x.ConfigurationArgumentMappings))
	for i, m := range x.ConfigurationArgumentMappings {
		mappings[i] = ArgumentMappingTF{
			PlanTemplateId:                           types.StringValue(m.PlanTemplateId),
			ActivityDefinitionPosition:               helper.TFInt64Value(m.ActivityDefinitionPosition),
			ConfigurationArgumentDefinitionName:      types.StringValue(m.ConfigurationArgumentDefinitionName),
			ActivityDefinitionArgumentDefinitionName: types.StringValue(m.ActivityDefinitionArgumentDefinitionName),
		}
	}

	return &RequestDefinitionTF{
		AuditModelTF:                     general_objects.AuditModelToTF(&x.AuditModel),
		Name:                             types.StringValue(x.Name),
		Description:                      types.StringPointerValue(x.Description),
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
				ID:             fcd.ID.ValueString(),
				CreatedAt:      fcd.CreatedAt.ValueString(),
				CreatedBy:      fcd.CreatedBy.ValueString(),
				LastModifiedAt: fcd.LastModifiedAt.ValueString(),
				LastModifiedBy: fcd.LastModifiedBy.ValueString(),
			},
			Name:        fcd.Name.ValueString(),
			Description: fcd.Description.ValueStringPointer(),
			Required:    fcd.Required.ValueBool(),
		}
	}

	configArgDefs := make([]activity_definitions.ArgumentDefinition[any], len(tf.ConfigurationArgumentDefinitions))
	for i, a := range tf.ConfigurationArgumentDefinitions {
		configArgDefs[i] = activity_definitions.ArgumentDefinition[any]{
			Name:        a.Name.ValueString(),
			Description: a.Description.ValueStringPointer(),
		}
		if a.Attributes != nil {
			configArgDefs[i].Attributes = general_objects.DefinitionAttributeFromTF(*a.Attributes)
		}
	}

	mappings := make([]ArgumentMapping, len(tf.ConfigurationArgumentMappings))
	for i, m := range tf.ConfigurationArgumentMappings {
		mappings[i] = ArgumentMapping{
			PlanTemplateId:                           m.PlanTemplateId.ValueString(),
			ActivityDefinitionPosition:               helper.FromTFInt64(m.ActivityDefinitionPosition),
			ConfigurationArgumentDefinitionName:      m.ConfigurationArgumentDefinitionName.ValueString(),
			ActivityDefinitionArgumentDefinitionName: m.ActivityDefinitionArgumentDefinitionName.ValueString(),
		}
	}

	return &RequestDefinition{
		AuditModel:                       general_objects.AuditModelFromTF(tf.AuditModelTF),
		Name:                             tf.Name.ValueString(),
		Description:                      tf.Description.ValueStringPointer(),
		PlanTemplateIds:                  helper.FromTFStrings(tf.PlanTemplateIds),
		FeasibilityConstraintDefinitions: fcds,
		ConfigurationArgumentDefinitions: configArgDefs,
		ConfigurationArgumentMappings:    mappings,
	}
}
