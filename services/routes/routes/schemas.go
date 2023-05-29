package routes

import (
	"github.com/leanspace/terraform-provider-leanspace/helper"
	"github.com/leanspace/terraform-provider-leanspace/helper/general_objects"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

var validLogLevels = []string{"INFO", "DEBUG", "TRACE", "WARN", "ERROR"}

var routeSchema = map[string]*schema.Schema{
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
	"tags": general_objects.KeyValuesSchema,

	"definition": {
		Type:     schema.TypeList,
		Required: true,
		MinItems: 1,
		MaxItems: 1,
		Elem: &schema.Resource{
			Schema: definitionSchema,
		},
	},

	"route_instances": {
		Type:     schema.TypeList,
		Computed: true,
		Elem: &schema.Resource{
			Schema: routeInstanceSchema,
		},
	},

	"processor_ids": {
		Type:     schema.TypeList,
		Optional: true,
		Elem: &schema.Schema{
			Type:         schema.TypeString,
			ValidateFunc: validation.IsUUID,
		},
	},

	"created_at": {
		Type:        schema.TypeString,
		Computed:    true,
		Description: "When it was created",
	},
	"created_by": {
		Type:        schema.TypeString,
		Computed:    true,
		Description: "Who created it",
	},
	"last_modified_at": {
		Type:        schema.TypeString,
		Computed:    true,
		Description: "When it was last modified",
	},
	"last_modified_by": {
		Type:        schema.TypeString,
		Computed:    true,
		Description: "Who modified it the last",
	},
}

var definitionSchema = map[string]*schema.Schema{
	"configuration": {
		Type:     schema.TypeString,
		Optional: true,
	},
	"log_level": {
		Type:         schema.TypeString,
		Required:     true,
		ValidateFunc: validation.StringInSlice(validLogLevels, false),
		Description:  helper.AllowedValuesToDescription(validLogLevels),
	},
	"valid": {
		Type:     schema.TypeBool,
		Computed: true,
	},
	"service_account_id": {
		Type:         schema.TypeString,
		Optional:     true,
		ValidateFunc: validation.IsUUID,
	},
	"errors": {
		Type:     schema.TypeSet,
		Computed: true,
		Elem: &schema.Resource{
			Schema: errorSchema,
		},
	},
}

var errorSchema = map[string]*schema.Schema{
	"code": {
		Type:     schema.TypeString,
		Computed: true,
	},
	"message": {
		Type:     schema.TypeString,
		Computed: true,
	},
}

var routeInstanceSchema = map[string]*schema.Schema{
	"status": {
		Type:     schema.TypeString,
		Required: true,
	},
	"last_status_at": {
		Type:     schema.TypeString,
		Optional: true,
	},
	"container_id": {
		Type:         schema.TypeString,
		Required:     true,
		ValidateFunc: validation.IsUUID,
	},
	"last_message_start_process_at": {
		Type:     schema.TypeString,
		Optional: true,
	},
	"last_message_end_process_at": {
		Type:     schema.TypeString,
		Optional: true,
	},
	"number_of_messages_processed": {
		Type:     schema.TypeInt,
		Optional: true,
	},
	"camel_route_id": {
		Type:         schema.TypeString,
		Optional:     true,
		ValidateFunc: validation.IsUUID,
	},
}

var dataSourceFilterSchema = map[string]*schema.Schema{
	"ids": {
		Type:     schema.TypeList,
		Optional: true,
		Elem: &schema.Schema{
			Type:         schema.TypeString,
			ValidateFunc: validation.IsUUID,
		},
	},
	"tags": {
		Type:     schema.TypeList,
		Optional: true,
		Elem: &schema.Schema{
			Type: schema.TypeString,
		},
	},
	"query": {
		Type:     schema.TypeString,
		Optional: true,
	},
}
