package request_definitions

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/leanspace/terraform-provider-leanspace/helper/general_objects"
)

var requestDefinitionSchema = map[string]*schema.Schema{
	"id": {
		Type:     schema.TypeString,
		Computed: true,
	},
	"name": {
		Type:     schema.TypeString,
		Required: true,
	},
	"description": {
		Type:     schema.TypeString,
		Optional: true,
	},
	"plan_template_ids": {
		Type:     schema.TypeSet,
		Required: true,
		MinItems: 1,
		MaxItems: 499,
		Elem: &schema.Schema{
			Type:         schema.TypeString,
			ValidateFunc: validation.IsUUID,
		},
	},
	"feasibility_constraint_definitions": {
		Type:     schema.TypeSet,
		Required: true,
		MinItems: 1,
		MaxItems: 499,
		Elem: &schema.Resource{
			Schema: feasibilityConstraintDefinitionSchema,
		},
	},
	"configuration_argument_definitions": {
		Type:     schema.TypeSet,
		Required: true,
		MinItems: 1,
		MaxItems: 499,
		Elem: &schema.Resource{
			Schema: argumentDefinitionSchema,
		},
	},
	"configuration_argument_mappings": {
		Type:     schema.TypeSet,
		Required: true,
		MinItems: 1,
		MaxItems: 499,
		Elem: &schema.Resource{
			Schema: argumentMappingSchema,
		},
	},
	"created_at": {
		Type:     schema.TypeString,
		Computed: true,
	},
	"created_by": {
		Type:     schema.TypeString,
		Computed: true,
	},
	"last_modified_at": {
		Type:     schema.TypeString,
		Computed: true,
	},
	"last_modified_by": {
		Type:     schema.TypeString,
		Computed: true,
	},
}

var feasibilityConstraintDefinitionSchema = map[string]*schema.Schema{
	"id": {
		Type:     schema.TypeString,
		Computed: true,
	},
	"name": {
		Type:     schema.TypeString,
		Required: true,
	},
	"description": {
		Type:     schema.TypeString,
		Optional: true,
	},
	"cloned": {
		Type:     schema.TypeBool,
		Optional: true,
	},
	"argument_definitions": {
		Type:     schema.TypeSet,
		Optional: true,
		MaxItems: 499,
		Elem: &schema.Resource{
			Schema: argumentDefinitionSchema,
		},
	},
	"created_at": {
		Type:     schema.TypeString,
		Computed: true,
	},
	"created_by": {
		Type:     schema.TypeString,
		Computed: true,
	},
	"last_modified_at": {
		Type:     schema.TypeString,
		Computed: true,
	},
	"last_modified_by": {
		Type:     schema.TypeString,
		Computed: true,
	},
}

var argumentDefinitionSchema = map[string]*schema.Schema{
	"name": {
		Type:     schema.TypeString,
		Required: true,
	},
	"description": {
		Type:     schema.TypeString,
		Optional: true,
	},
	"attributes": {
		Type:     schema.TypeList,
		Required: true,
		MinItems: 1,
		MaxItems: 1,
		Elem: &schema.Resource{
			Schema: general_objects.DefinitionAttributeSchema(
				[]string{"TEXT", "BINARY", "BOOLEAN", "ENUM", "TIMESTAMP", "DATE", "ARRAY", "STRUCTURE", "TLE"}, // attribute types not allowed in command definition attributes
				nil,   // All fields are used
				false, // Does not force recreation if the type changes
			),
		},
	},
}

var argumentMappingSchema = map[string]*schema.Schema{
	"plan_template_id": {
		Type:         schema.TypeString,
		ValidateFunc: validation.IsUUID,
		Required:     true,
	},
	"activity_definition_position": {
		Type:         schema.TypeInt,
		Required:     true,
		ValidateFunc: validation.IntBetween(0, 499),
	},
	"configuration_argument_definition_name": {
		Type:     schema.TypeString,
		Required: true,
	},
	"activity_definition_argument_definition_name": {
		Type:     schema.TypeString,
		Required: true,
	},
}
