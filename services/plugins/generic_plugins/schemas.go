package generic_plugins

import (
	"github.com/leanspace/terraform-provider-leanspace/helper"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

var validGenericPluginTypes = []string{
	"CHECKSUM_FUNCTION",
}

var validGenericPluginLanguages = []string{
	"JAVA",
}

var genericPluginSchema = map[string]*schema.Schema{
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
	"type": {
		Type:         schema.TypeString,
		Required:     true,
		ValidateFunc: validation.StringInSlice(validGenericPluginTypes, false),
		Description:  helper.AllowedValuesToDescription(validGenericPluginTypes),
	},
	"language": {
		Type:         schema.TypeString,
		Required:     true,
		ValidateFunc: validation.StringInSlice(validGenericPluginLanguages, false),
		Description:  helper.AllowedValuesToDescription(validGenericPluginLanguages),
	},
	"source_code_link": {
		Type:     schema.TypeList,
		Computed: true,
		Elem: &schema.Resource{
			Schema: sourceCodeLinkSchema,
		},
	},
	"created_at": {
		Type:        schema.TypeString,
		Computed:    true,
		Description: "When the plugin was created",
	},
	"created_by": {
		Type:        schema.TypeString,
		Computed:    true,
		Description: "Who created the plugin",
	},
	"last_modified_by": {
		Type:        schema.TypeString,
		Computed:    true,
		Description: "Who modified the plugin the last",
	},
	"last_modified_at": {
		Type:        schema.TypeString,
		Computed:    true,
		Description: "When the plugin was last modified",
	},
	"status": {
		Type:        schema.TypeString,
		Computed:    true,
		Description: "Generic Plugin status. Can be ACTIVE, PENDING or FAILED",
	},
	"source_code_path": {
		Type:     schema.TypeString,
		Required: true,
		ValidateFunc: validation.StringMatch(
			helper.PathToJarFileRegex,
			"'source_code_path' must be a valid path to a .jar file",
		),
		Description: "It must be a valid path to a .jar file",
	},
	"source_code_sha": {
		Type:        schema.TypeString,
		Computed:    true,
		Description: "Unique identifier of the generic plugin file",
	},
}

var sourceCodeLinkSchema = map[string]*schema.Schema{
	"expiration_time": {
		Type:        schema.TypeString,
		Computed:    true,
		Description: "When the source code link expires",
	},
	"source_code_id": {
		Type:        schema.TypeString,
		Computed:    true,
		Description: "Unique identifier of the source code",
	},
	"url": {
		Type:        schema.TypeString,
		Computed:    true,
		Description: "URL to download the source code",
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
		Description: "Only returns plugin whose id matches one of the provided values.",
	},
	"statuses": {
		Type:     schema.TypeList,
		Optional: true,
		Elem: &schema.Schema{
			Type:         schema.TypeString,
			ValidateFunc: validation.StringInSlice([]string{"ACTIVE", "PENDING", "FAILED"}, false),
			Description:  helper.AllowedValuesToDescription([]string{"ACTIVE", "PENDING", "FAILED"}),
		},
	},
	"types": {
		Type:     schema.TypeList,
		Optional: true,
		Elem: &schema.Schema{
			Type:         schema.TypeString,
			ValidateFunc: validation.StringInSlice(validGenericPluginTypes, false),
			Description:  helper.AllowedValuesToDescription(validGenericPluginTypes),
		},
	},
	"query": {
		Type:     schema.TypeString,
		Optional: true,
	},
}
