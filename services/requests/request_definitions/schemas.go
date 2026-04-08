package request_definitions

import (
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/setvalidator"
	datasourceschema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	resourceschema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/leanspace/terraform-provider-leanspace/helper"
	"github.com/leanspace/terraform-provider-leanspace/helper/general_objects"
)

var requestDefinitionSchema = general_objects.ResourceSchemaWith(map[string]resourceschema.Attribute{
	"name": resourceschema.StringAttribute{
		Required: true,
	},
	"description": resourceschema.StringAttribute{
		Optional: true,
	},
	"plan_template_ids": resourceschema.SetAttribute{
		ElementType: types.StringType,
		Required:    true,
		Validators:  []validator.Set{setvalidator.SizeAtMost(499), setvalidator.ValueStringsAre(helper.ValidUUID()...)},
	},
	"feasibility_constraint_definitions": resourceschema.SetNestedAttribute{
		Required: true,
		NestedObject: resourceschema.NestedAttributeObject{
			Attributes: feasibilityConstraintDefinitionSchema,
		},
		Validators: []validator.Set{setvalidator.SizeAtMost(499)},
	},
	"configuration_argument_definitions": resourceschema.SetNestedAttribute{
		Optional: true,
		NestedObject: resourceschema.NestedAttributeObject{
			Attributes: argumentDefinitionSchema,
		},
		Validators: []validator.Set{setvalidator.SizeAtMost(499)},
	},
	"configuration_argument_mappings": resourceschema.SetNestedAttribute{
		Optional: true,
		NestedObject: resourceschema.NestedAttributeObject{
			Attributes: argumentMappingSchema,
		},
		Validators: []validator.Set{setvalidator.SizeAtMost(499)},
	},
})

var feasibilityConstraintDefinitionSchema = map[string]resourceschema.Attribute{
	"id": resourceschema.StringAttribute{
		Required: true,
	},
	"name": resourceschema.StringAttribute{
		Computed: true,
	},
	"description": resourceschema.StringAttribute{
		Computed: true,
	},
	"required": resourceschema.BoolAttribute{
		Required: true,
	},
	"argument_definitions": resourceschema.SetNestedAttribute{
		Computed: true,
		NestedObject: resourceschema.NestedAttributeObject{
			Attributes: computedArgumentDefinitionSchema,
		},
	},
	"created_at": resourceschema.StringAttribute{
		Computed: true,
	},
	"created_by": resourceschema.StringAttribute{
		Computed: true,
	},
	"last_modified_at": resourceschema.StringAttribute{
		Computed: true,
	},
	"last_modified_by": resourceschema.StringAttribute{
		Computed: true,
	},
}

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

var computedArgumentDefinitionSchema = map[string]resourceschema.Attribute{
	"name": resourceschema.StringAttribute{
		Computed: true,
	},
	"description": resourceschema.StringAttribute{
		Computed: true,
	},
	"attributes": resourceschema.SingleNestedAttribute{
		Computed: true,
		Attributes: general_objects.DefinitionAttributeSchema(
			[]string{"BINARY", "BOOLEAN", "ENUM", "DATE", "ARRAY", "STRUCTURE", "TLE"}, // attribute types not allowed in command definition attributes
			nil,   // All fields are used
			false, // Does not force recreation if the type changes
		),
	},
}

var argumentMappingSchema = map[string]resourceschema.Attribute{
	"plan_template_id": resourceschema.StringAttribute{
		Required:   true,
		Validators: helper.ValidUUID(),
	},
	"activity_definition_position": resourceschema.Int64Attribute{
		Required:   true,
		Validators: []validator.Int64{int64validator.Between(0, 499)},
	},
	"configuration_argument_definition_name": resourceschema.StringAttribute{
		Required: true,
	},
	"activity_definition_argument_definition_name": resourceschema.StringAttribute{
		Required: true,
	},
}

var requestDefinitionFilterSchema = map[string]datasourceschema.Attribute{
	"feasibility_constraint_definition_ids": datasourceschema.ListAttribute{
		ElementType: types.StringType,
		Optional:    true,
		Validators:  []validator.List{listvalidator.ValueStringsAre(helper.ValidUUID()...)},
	},
	"plan_template_ids": datasourceschema.ListAttribute{
		ElementType: types.StringType,
		Optional:    true,
		Validators:  []validator.List{listvalidator.ValueStringsAre(helper.ValidUUID()...)},
	},
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
