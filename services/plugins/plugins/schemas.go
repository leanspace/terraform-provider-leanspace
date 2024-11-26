package plugins

import (
	"regexp"

	"github.com/leanspace/terraform-provider-leanspace/helper"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

var validPluginTypes = []string{
	"COMMANDS_COMMAND_TRANSFORMER_PLUGIN_TYPE",
	"COMMANDS_PROTOCOL_TRANSFORMER_PLUGIN_TYPE",
}

var classNameRegex = regexp.MustCompile(`^([a-z]+\.)+([A-Z][a-zA-Z0-9]+)$`)

var pluginSchema = map[string]*schema.Schema{
	"id": {
		Type:     schema.TypeString,
		Computed: true,
	},
	"type": {
		Type:         schema.TypeString,
		Required:     true,
		ValidateFunc: validation.StringInSlice(validPluginTypes, false),
		Description:  helper.AllowedValuesToDescription(validPluginTypes),
	},
	"implementation_class_name": {
		Type:     schema.TypeString,
		Required: true,
		ValidateFunc: validation.StringMatch(
			classNameRegex,
			"'implementation_class_name' must be a valid java class path",
		),
		Description: "It must be a valid java class path",
	},
	"name": {
		Type:     schema.TypeString,
		Required: true,
	},
	"description": {
		Type:     schema.TypeString,
		Optional: true,
	},
	"source_code_file_download_authorized": {
		Type:     schema.TypeBool,
		Optional: true,
		Default:  true,
	},
	"file_path": {
		Type:     schema.TypeString,
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
	"sdk_version": {
		Type:         schema.TypeString,
		Optional:     true,
		ValidateFunc: isValidSemVerForPlugins,
		Description:  "SDK version in the semantic version format with major versions 1 or 2.",
	},
	"sdk_version_family": {
		Type:        schema.TypeString,
		Computed:    true,
		Description: "SDK family that indicates the major version.",
	},
	"status": {
		Type:        schema.TypeString,
		Computed:    true,
		Description: "Plugin status. Can be ACTIVE, PENDING or FAILED",
	},
	"file_sha": {
		Type:        schema.TypeString,
		Computed:    true,
		Description: "Unique identifier of the plugin file",
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
	"types": {
		Type:     schema.TypeList,
		Optional: true,
		Elem: &schema.Schema{
			Type:         schema.TypeString,
			ValidateFunc: validation.StringInSlice(validPluginTypes, false),
			Description:  helper.AllowedValuesToDescription(validPluginTypes),
		},
	},
	"query": {
		Type:     schema.TypeString,
		Optional: true,
	},
}
