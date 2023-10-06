package processors

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/leanspace/terraform-provider-leanspace/helper"
)

var processorSchema = map[string]*schema.Schema{
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
	"version": {
		Type:     schema.TypeString,
		ForceNew: true,
		Required: true,
	},
	"type": {
		Type:     schema.TypeString,
		Computed: true,
	},
	"file_path": {
		Type:     schema.TypeString,
		ForceNew: true,
		Required: true,
		ValidateFunc: validation.StringMatch(
			helper.PathToJarFileRegex,
			"'file_path' must be a valid path to a .jar file",
		),
		Description: "It must be a valid path to a .jar file",
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
	"file_sha": {
		Type:        schema.TypeString,
		Computed:    true,
		Description: "Unique identifier of the processor file",
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
	"query": {
		Type:     schema.TypeString,
		Optional: true,
	},
}
