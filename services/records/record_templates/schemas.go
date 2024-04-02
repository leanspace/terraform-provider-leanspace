package record_templates

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/leanspace/terraform-provider-leanspace/helper/general_objects"
)

var validRecordTemplatesConstraintTypes = []string{"UNKNOWN", "PASSED", "FAILED"}

var recordTemplateSchema = map[string]*schema.Schema{
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
	"record_state": {
		Type:         schema.TypeString,
		Optional:     true,
		ValidateFunc: validation.StringInSlice(validRecordTemplatesConstraintTypes, false),
	},
	"start_date_time": {
		Type:     schema.TypeString,
		Computed: true,
	},
	"stop_date_time": {
		Type:     schema.TypeString,
		Computed: true,
	},
	"default_parsers": {
		Type:     schema.TypeSet,
		Optional: true,
		Elem: &schema.Resource{
			Schema: recordTemplateDefaultParserSchema,
		},
	},
	"nodes": {
		Type:     schema.TypeSet,
		Optional: true,
		Elem: &schema.Resource{
			Schema: recordTemplateNodeSnapshotSchema,
		},
	},
	"properties": {
		Type:     schema.TypeSet,
		Optional: true,
		Elem: &schema.Resource{
			Schema: recordTemplatePropertySchema,
		},
	},
	"tags": general_objects.KeyValuesSchema,
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

var recordTemplateDefaultParserSchema = map[string]*schema.Schema{
	// TODO
}

var recordTemplateNodeSnapshotSchema = map[string]*schema.Schema{
	// TODO
}

var recordTemplatePropertySchema = map[string]*schema.Schema{
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
	"names": {
		Type:     schema.TypeList,
		Optional: true,
		Elem: &schema.Schema{
			Type: schema.TypeString,
		},
	},
	"related_asset_ids": {
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
		Type:        schema.TypeString,
		Optional:    true,
		Description: "Partial search by name. Allowed wildcard characters are `.*` and `%`",
	},
}
