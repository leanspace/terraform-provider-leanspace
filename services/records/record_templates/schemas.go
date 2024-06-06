package record_templates

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/leanspace/terraform-provider-leanspace/helper/general_objects"
)

var recordTemplateSchema = map[string]*schema.Schema{
	"id": {
		Type:     schema.TypeString,
		Computed: true,
	},
	"name": {
		Type:     schema.TypeString,
		Required: true,
		ForceNew: true,
	},
	"description": {
		Type:     schema.TypeString,
		Optional: true,
		ForceNew: true,
	},
	"state": {
		Type:     schema.TypeString,
		Computed: true,
	},
	"start_date_time": {
		Type:     schema.TypeString,
		Computed: true,
	},
	"stop_date_time": {
		Type:     schema.TypeString,
		Computed: true,
	},
	"stream_id": {
		Type:     schema.TypeString,
		Optional: true,
		ForceNew: true,
	},
	"default_parsers": {
		Type:     schema.TypeSet,
		Computed: true,
		Elem: &schema.Resource{
			Schema: recordTemplateDefaultParserSchema,
		},
	},
	"node_ids": {
		Type:     schema.TypeSet,
		Optional: true,
		ForceNew: true,
		Elem: &schema.Schema{
			Type:         schema.TypeString,
			ValidateFunc: validation.IsUUID,
		},
	},
	"metric_ids": {
		Type:     schema.TypeSet,
		Optional: true,
		ForceNew: true,
		Elem: &schema.Schema{
			Type:         schema.TypeString,
			ValidateFunc: validation.IsUUID,
		},
	},
	"command_definition_ids": {
		Type:     schema.TypeSet,
		Optional: true,
		ForceNew: true,
		Elem: &schema.Schema{
			Type:         schema.TypeString,
			ValidateFunc: validation.IsUUID,
		},
	},
	"properties": {
		Type:     schema.TypeSet,
		Optional: true,
		ForceNew: true,
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
	"id": {
		Type:     schema.TypeString,
		Computed: true,
	},
	"file_type": {
		Type:     schema.TypeString,
		Computed: true,
	},
}

var recordTemplatePropertySchema = map[string]*schema.Schema{
	"name": {
		Type:     schema.TypeString,
		Required: true,
		ForceNew: true,
	},
	"attributes": {
		Type:     schema.TypeList,
		Required: true,
		ForceNew: true,
		MinItems: 1,
		MaxItems: 1,
		Elem: &schema.Resource{
			Schema: general_objects.DefinitionAttributeSchema(
				[]string{"BINARY", "ENUM", "TIMESTAMP", "DATE", "TIME", "TLE", "ARRAY", "GEOPOINT", "TUPLE"}, // Attribute types not allowed in attributes
				nil,   // All fields are used
				false, // Does not force recreation if the type changes
			),
		},
	},
}

var dataSourceFilterSchema = map[string]*schema.Schema{
	"names": {
		Type:     schema.TypeList,
		Optional: true,
		ForceNew: true,
		Elem: &schema.Schema{
			Type: schema.TypeString,
		},
	},
	"related_asset_ids": {
		Type:     schema.TypeList,
		Optional: true,
		ForceNew: true,
		Elem: &schema.Schema{
			Type:         schema.TypeString,
			ValidateFunc: validation.IsUUID,
		},
	},
	"tags": {
		Type:     schema.TypeList,
		Optional: true,
		ForceNew: true,
		Elem: &schema.Schema{
			Type: schema.TypeString,
		},
	},
}
