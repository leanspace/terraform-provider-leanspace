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
	"record_state": {
		Type:         schema.TypeString,
		Optional:     true,
		Default:      "UNKNOWN",
		ValidateFunc: validation.StringInSlice(validRecordConstraintTypes, false),
	},
	"start_date_time": {
		Type:     schema.TypeString,
		Required: true,
	},
	"stop_date_time": {
		Type:     schema.TypeString,
		Optional: true,
	},
	"properties": {
		Type:     schema.TypeSet,
		Optional: true,
		Elem: &schema.Resource{
			Schema: recordPropertySchema,
		},
	},
	"tags": general_objects.KeyValuesSchema,
	"comment": {
		Type:     schema.TypeString,
		Optional: true,
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
	// TODO
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
	"query": {
		Type:        schema.TypeString,
		Optional:    true,
		Description: "Partial search by name. Allowed wildcard characters are `.*` and `%`",
	},
}
