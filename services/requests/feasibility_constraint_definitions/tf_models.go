package feasibility_constraint_definitions

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/leanspace/terraform-provider-leanspace/helper/general_objects"
	"github.com/leanspace/terraform-provider-leanspace/services/activities/activity_definitions"
)

type FeasibilityConstraintDefinitionTF struct {
	general_objects.AuditModelTF
	Name                types.String                                `tfsdk:"name"`
	Description         types.String                                `tfsdk:"description"`
	ArgumentDefinitions []activity_definitions.ArgumentDefinitionTF `tfsdk:"argument_definitions"`
}

func (x *FeasibilityConstraintDefinition) ToTF() interface{} {
	argDefs := make([]activity_definitions.ArgumentDefinitionTF, len(x.ArgumentDefinitions))
	for i, a := range x.ArgumentDefinitions {
		attr := general_objects.DefinitionAttributeToTF(&a.Attributes)
		argDefs[i] = activity_definitions.ArgumentDefinitionTF{
			Name:        types.StringValue(a.Name),
			Description: types.StringPointerValue(a.Description),
			Attributes:  &attr,
		}
	}

	return &FeasibilityConstraintDefinitionTF{
		AuditModelTF:        general_objects.AuditModelToTF(&x.AuditModel),
		Name:                types.StringValue(x.Name),
		Description:         types.StringPointerValue(x.Description),
		ArgumentDefinitions: argDefs,
	}
}

func (tf *FeasibilityConstraintDefinitionTF) ToAPI() interface{} {
	argDefs := make([]activity_definitions.ArgumentDefinition[any], len(tf.ArgumentDefinitions))
	for i, a := range tf.ArgumentDefinitions {
		argDefs[i] = activity_definitions.ArgumentDefinition[any]{
			Name:        a.Name.ValueString(),
			Description: a.Description.ValueStringPointer(),
		}
		if a.Attributes != nil {
			argDefs[i].Attributes = general_objects.DefinitionAttributeFromTF(*a.Attributes)
		}
	}

	return &FeasibilityConstraintDefinition{
		AuditModel:          general_objects.AuditModelFromTF(tf.AuditModelTF),
		Name:                tf.Name.ValueString(),
		Description:         tf.Description.ValueStringPointer(),
		ArgumentDefinitions: argDefs,
	}
}
