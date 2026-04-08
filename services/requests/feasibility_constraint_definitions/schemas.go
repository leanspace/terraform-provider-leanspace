package feasibility_constraint_definitions

import (
	"github.com/hashicorp/terraform-plugin-framework-validators/setvalidator"
	datasourceschema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	resourceschema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/leanspace/terraform-provider-leanspace/helper/general_objects"
)

var feasibilityConstraintDefinitionSchema = general_objects.ResourceSchemaWith(map[string]resourceschema.Attribute{
	"name": resourceschema.StringAttribute{
		Required: true,
	},
	"description": resourceschema.StringAttribute{
		Optional: true,
	},
	"argument_definitions": resourceschema.SetNestedAttribute{
		Optional: true,
		NestedObject: resourceschema.NestedAttributeObject{
			Attributes: argumentDefinitionSchema,
		},
		Validators: []validator.Set{setvalidator.SizeAtMost(499)},
	},
})

var argumentDefinitionSchema = map[string]resourceschema.Attribute{
	"name": resourceschema.StringAttribute{
		Required: true,
	},
	"description": resourceschema.StringAttribute{
		Optional: true,
	},
	"attributes": resourceschema.SingleNestedAttribute{
		Required: true,
		Attributes: general_objects.DefinitionAttributeSchema(
			[]string{"BINARY", "BOOLEAN", "ENUM", "DATE", "ARRAY", "STRUCTURE", "TLE"}, // attribute types not allowed in command definition attributes
			nil,   // All fields are used
			false, // Does not force recreation if the type changes
		),
	},
}

var feasibilityConstraintDefinitionFilterSchema = map[string]datasourceschema.Attribute{
	"created_bys": datasourceschema.ListAttribute{
		ElementType: types.StringType,
		Optional:    true,
	},
	"to_created_at": datasourceschema.StringAttribute{
		Optional: true,
	},
	"last_modified_bys": datasourceschema.ListAttribute{
		ElementType: types.StringType,
		Optional:    true,
	},
	"from_last_modified_at": datasourceschema.StringAttribute{
		Optional: true,
	},
	"to_last_modified_at": datasourceschema.StringAttribute{
		Optional: true,
	},
}
