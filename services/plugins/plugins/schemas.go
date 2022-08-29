package plugins

import (
	"leanspace-terraform-provider/helper"
	"regexp"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

var validPluginTypes = []string{
	"STRING_IDENTITY_PLUGIN_TYPE",
	"JSON_IDENTITY_PLUGIN_TYPE",
	"COMMANDS_COMMAND_TRANSFORMER_PLUGIN_TYPE",
	"COMMANDS_PROTOCOL_TRANSFORMER_PLUGIN_TYPE",
	"SIMULATIONS_ANALYTICAL_NOMINAL_PROPAGATION_PLUGIN_TYPE",
}

var classNameRegex = regexp.MustCompile(`^([a-z]+\.)+([A-Z][a-zA-Z0-9]+)$`)
var pathToJarFileRegex = regexp.MustCompile(`^.*\.jar$`)

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
			pathToJarFileRegex,
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
}

var dataSourceFilterSchema = map[string]*schema.Schema{
	"types": {
		Type:     schema.TypeList,
		Optional: true,
		Elem: &schema.Schema{
			Type:         schema.TypeString,
			ValidateFunc: validation.StringInSlice(validPluginTypes, false),
			Description:  helper.AllowedValuesToDescription(validPluginTypes),
		},
	},
}
