package request_definitions

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/leanspace/terraform-provider-leanspace/helper"
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
		Optional: true,
		MinItems: 1,
		MaxItems: 499,
		Elem: &schema.Resource{
			Schema: argumentDefinitionSchema,
		},
	},
	"configuration_argument_mappings": {
		Type:     schema.TypeSet,
		Optional: true,
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
		Required: true,
	},
	"name": {
		Type:     schema.TypeString,
		Computed: true,
	},
	"description": {
		Type:     schema.TypeString,
		Computed: true,
	},
	"required": {
		Type:     schema.TypeBool,
		Required: true,
	},
	"argument_definitions": {
		Type:     schema.TypeSet,
		Computed: true,
		Elem: &schema.Resource{
			Schema: computedArgumentDefinitionSchema,
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
		MaxItems: 499,
		Elem: &schema.Resource{
			Schema: general_objects.DefinitionAttributeSchema(
				[]string{"BINARY", "BOOLEAN", "ENUM", "TIMESTAMP", "DATE", "ARRAY", "STRUCTURE", "TLE"}, // attribute types not allowed in command definition attributes
				nil,   // All fields are used
				false, // Does not force recreation if the type changes
			),
		},
	},
}

var computedArgumentDefinitionSchema = map[string]*schema.Schema{
	"name": {
		Type:     schema.TypeString,
		Computed: true,
	},
	"description": {
		Type:     schema.TypeString,
		Computed: true,
	},
	"attributes": {
		Type:     schema.TypeList,
		Computed: true,
		Elem: &schema.Resource{
			Schema: general_objects.DefinitionAttributeSchema(
				[]string{"BINARY", "BOOLEAN", "ENUM", "TIMESTAMP", "DATE", "ARRAY", "STRUCTURE", "TLE"}, // attribute types not allowed in command definition attributes
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

var requestDefinitionFilterSchema = map[string]*schema.Schema{
	"feasibility_constraint_definition_ids": {
		Type:     schema.TypeList,
		Optional: true,
		Elem: &schema.Schema{
			Type:         schema.TypeString,
			ValidateFunc: validation.IsUUID,
		},
	},
	"plan_template_ids": {
		Type:     schema.TypeList,
		Optional: true,
		Elem: &schema.Schema{
			Type:         schema.TypeString,
			ValidateFunc: validation.IsUUID,
		},
	},
	"created_bys": {
		Type:     schema.TypeList,
		Optional: true,
		Elem: &schema.Schema{
			Type:         schema.TypeString,
			ValidateFunc: validation.IsUUID,
		},
	},
	"to_created_at": {
		Type:         schema.TypeString,
		Optional:     true,
		ValidateFunc: helper.IsValidTimeDateOrTimestamp,
	},
	"last_modified_bys": {
		Type:     schema.TypeList,
		Optional: true,
		Elem: &schema.Schema{
			Type:         schema.TypeString,
			ValidateFunc: validation.IsUUID,
		},
	},
	"from_last_modified_at": {
		Type:         schema.TypeString,
		Optional:     true,
		ValidateFunc: helper.IsValidTimeDateOrTimestamp,
	},
	"to_last_modified_at": {
		Type:         schema.TypeString,
		Optional:     true,
		ValidateFunc: helper.IsValidTimeDateOrTimestamp,
	},
}
