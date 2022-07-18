package plugins

import (
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
	},
	"implementation_class_name": {
		Type:     schema.TypeString,
		Required: true,
		ValidateFunc: validation.StringMatch(
			classNameRegex,
			"'implementation_class_name' must be a valid java class path",
		),
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
		Required: true,
	},
	"file_path": {
		Type:     schema.TypeString,
		Required: true,
		ValidateFunc: validation.StringMatch(
			pathToJarFileRegex,
			"'file_path' must be a valid path to a .jar file",
		),
	},
	"created_at": {
		Type:     schema.TypeString,
		Computed: true,
	},
	"created_by": {
		Type:     schema.TypeString,
		Computed: true,
	},
	"last_modified_at": {
		Type:     schema.TypeString,
		Computed: true,
	},
	"last_modified_by": {
		Type:     schema.TypeString,
		Computed: true,
	},
}
