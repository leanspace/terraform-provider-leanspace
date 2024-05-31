package records

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/leanspace/terraform-provider-leanspace/helper/general_objects"
)

var validRecordConstraintTypes = []string{"UNKNOWN", "PASSED", "FAILED"}

var recordSchema = map[string]*schema.Schema{
	"id": {
		Type:     schema.TypeString,
		Computed: true,
	},
	"record_template_id": {
		Type:     schema.TypeString,
		Required: true,
	},
	"name": {
		Type:     schema.TypeString,
		Required: true,
	},
	"state": {
		Type:         schema.TypeString,
		Optional:     true,
		Default:      "UNKNOWN",
		ValidateFunc: validation.StringInSlice(validRecordConstraintTypes, false),
	},
	"processing_status": {
		Type:     schema.TypeString,
		Computed: true,
	},
	"start_date_time": {
		Type:     schema.TypeString,
		Required: true,
	},
	"stop_date_time": {
		Type:     schema.TypeString,
		Optional: true,
	},
	"stream_id": {
		Type:     schema.TypeString,
		Computed: true,
	},
	"node_ids": {
		Type:     schema.TypeSet,
		Computed: true,
		Elem: &schema.Schema{
			Type:         schema.TypeString,
			ValidateFunc: validation.IsUUID,
		},
	},
	"metric_ids": {
		Type:     schema.TypeSet,
		Computed: true,
		Elem: &schema.Schema{
			Type:         schema.TypeString,
			ValidateFunc: validation.IsUUID,
		},
	},
	"properties": {
		Type:     schema.TypeSet,
		Optional: true,
		Elem: &schema.Resource{
			Schema: recordPropertySchema,
		},
	},
	"command_definition_ids": {
		Type:     schema.TypeSet,
		Computed: true,
		Elem: &schema.Schema{
			Type:         schema.TypeString,
			ValidateFunc: validation.IsUUID,
		},
	},
	"tags": general_objects.KeyValuesSchema,
	"comments": {
		Type:     schema.TypeSet,
		Computed: true,
		Elem: &schema.Schema{
			Type: schema.TypeString,
		},
	},
	"errors": {
		Type:     schema.TypeSet,
		Computed: true,
		Elem: &schema.Resource{
			Schema: errorSchema,
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

var recordPropertySchema = map[string]*schema.Schema{
	"name": {
		Type:     schema.TypeString,
		Required: true,
	},
	"attributes": {
		Type:     schema.TypeList,
		Required: true,
		MinItems: 1,
		MaxItems: 1,
		Elem: &schema.Resource{
			Schema: general_objects.ValueAttributeSchema,
		},
	},
}

var dataSourceFilterSchema = map[string]*schema.Schema{
	"record_template_ids": {
		Type:     schema.TypeList,
		Optional: true,
		Elem: &schema.Schema{
			Type:         schema.TypeString,
			ValidateFunc: validation.IsUUID,
		},
	},
	"names": {
		Type:     schema.TypeList,
		Optional: true,
		Elem: &schema.Schema{
			Type: schema.TypeString,
		},
	},
	"tags": {
		Type:     schema.TypeList,
		Optional: true,
		Elem: &schema.Schema{
			Type: schema.TypeString,
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
