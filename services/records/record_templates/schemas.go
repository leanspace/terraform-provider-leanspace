package record_templates

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/leanspace/terraform-provider-leanspace/helper"
	"github.com/leanspace/terraform-provider-leanspace/helper/general_objects"
)

var validFileTypes = []string{"CSV"}

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
		Type:        schema.TypeString,
		Computed:    true,
		Description: helper.AllowedValuesToDescription(validFileTypes),
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
				[]string{"BINARY", "GEOPOINT", "TLE"}, // Attribute types not allowed in attributes
				nil,                                     // All fields are used
				false,                                   // Does not force recreation if the type changes
			),
		},
	},
}

var dataSourceFilterSchema = map[string]*schema.Schema{
	"names": {
		Type:        schema.TypeList,
		Optional:    true,
		Description: "Only returns Record Templates who's name matches one of the provided values.",
		Elem: &schema.Schema{
			Type: schema.TypeString,
		},
	},
	"node_ids": {
		Type:        schema.TypeList,
		Optional:    true,
		Description: "Only returns Record Templates with at least one nodeId that matches one of the provided values.",
		Elem: &schema.Schema{
			Type:         schema.TypeString,
			ValidateFunc: validation.IsUUID,
		},
	},
	"metric_ids": {
		Type:        schema.TypeList,
		Optional:    true,
		Description: "Only returns Record Templates with at least one metricId that matches one of the provided values.",
		Elem: &schema.Schema{
			Type:         schema.TypeString,
			ValidateFunc: validation.IsUUID,
		},
	},
	"created_by": {
		Type:     schema.TypeList,
		Optional: true,
		Elem: &schema.Schema{
			Type:         schema.TypeString,
			ValidateFunc: validation.IsUUID,
		},
		Description: "Filter on the user who created the RecordTemplate. If you have no wish to use this field as a filter, either provide a null value or remove the field.",
	},
	"last_modified_by": {
		Type:     schema.TypeList,
		Optional: true,
		Elem: &schema.Schema{
			Type:         schema.TypeString,
			ValidateFunc: validation.IsUUID,
		},
		Description: "Filter on the user who last modified the RecordTemplate. If you have no wish to use this field as a filter, either provide a null value or remove the field.",
	},
	"from_created_at": {
		Type:         schema.TypeString,
		Optional:     true,
		ValidateFunc: helper.IsValidTimeDateOrTimestamp,
		Description:  "Filter on the RecordTemplate creation date. RecordTemplates with a creation date greater or equals than the filter value will be selected (if they are not excluded by other filters). If you have no wish to use this field as a filter, either provide a null value or remove the field.",
	},
	"to_created_at": {
		Type:         schema.TypeString,
		Optional:     true,
		ValidateFunc: helper.IsValidTimeDateOrTimestamp,
		Description:  "Filter on the RecordTemplate creation date. RecordTemplates with a creation date lower or equals than the filter value will be selected (if they are not excluded by other filters). If you have no wish to use this field as a filter, either provide a null value or remove the field.",
	},
	"from_last_modified_at": {
		Type:         schema.TypeString,
		Optional:     true,
		ValidateFunc: helper.IsValidTimeDateOrTimestamp,
		Description:  "Filter on the RecordTemplate last modification date. RecordTemplates with a last modification date greater or equals than the filter value will be selected (if they are not excluded by other filters). If you have no wish to use this field as a filter, either provide a null value or remove the field.",
	},
	"to_last_modified_at": {
		Type:         schema.TypeString,
		Optional:     true,
		ValidateFunc: helper.IsValidTimeDateOrTimestamp,
		Description:  "Filter on the RecordTemplate last modification date. RecordTemplates with a last modification date lower or equals than the filter value will be selected (if they are not excluded by other filters). If you have no wish to use this field as a filter, either provide a null value or remove the field.",
	},
	"tags": {
		Type:     schema.TypeList,
		Optional: true,
		Elem: &schema.Schema{
			Type: schema.TypeString,
		},
	},
}
