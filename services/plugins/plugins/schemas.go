package plugins

import (
	"regexp"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	datasourceschema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	resourceschema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/leanspace/terraform-provider-leanspace/helper"
)

var validPluginTypes = []string{
	"COMMANDS_COMMAND_TRANSFORMER_PLUGIN_TYPE",
	"COMMANDS_PROTOCOL_TRANSFORMER_PLUGIN_TYPE",
}

var classNameRegex = regexp.MustCompile(`^([a-z]+\.)+([A-Z][a-zA-Z0-9]+)$`)

var pluginSchema = map[string]resourceschema.Attribute{
	"id": resourceschema.StringAttribute{
		Computed: true,
	},
	"type": resourceschema.StringAttribute{
		Required:    true,
		Description: helper.AllowedValuesToDescription(validPluginTypes),
		Validators:  []validator.String{stringvalidator.OneOf(validPluginTypes...)},
	},
	"implementation_class_name": resourceschema.StringAttribute{
		Required:    true,
		Description: "It must be a valid java class path",
		Validators:  []validator.String{stringvalidator.RegexMatches(classNameRegex, "Must be a valid java class path")},
	},
	"name": resourceschema.StringAttribute{
		Required: true,
	},
	"description": resourceschema.StringAttribute{
		Optional: true,
	},
	"source_code_file_download_authorized": resourceschema.BoolAttribute{
		Optional: true,
		Default:  booldefault.StaticBool(true),
	},
	"file_path": resourceschema.StringAttribute{
		Required:    true,
		Description: "It must be a valid path to a .jar file",
		Validators:  []validator.String{stringvalidator.RegexMatches(helper.PathToJarFileRegex, "Must be a valid file path to a .jar file")},
	},
	"created_at": resourceschema.StringAttribute{
		Computed:    true,
		Description: "When the plugin was created",
	},
	"created_by": resourceschema.StringAttribute{
		Computed:    true,
		Description: "Who created the plugin",
	},
	"last_modified_by": resourceschema.StringAttribute{
		Computed:    true,
		Description: "Who modified the plugin the last",
	},
	"last_modified_at": resourceschema.StringAttribute{
		Computed:    true,
		Description: "When the plugin was last modified",
	},
	"sdk_version": resourceschema.StringAttribute{
		Optional:    true,
		Description: "SDK version in the semantic version format with major versions 1 or 2.",
		Validators:  helper.IsValidSemVer(),
	},
	"sdk_version_family": resourceschema.StringAttribute{
		Computed:    true,
		Description: "SDK family that indicates the major version.",
	},
	"status": resourceschema.StringAttribute{
		Computed:    true,
		Description: "Plugin status. Can be ACTIVE, PENDING or FAILED",
	},
	"file_sha": resourceschema.StringAttribute{
		Computed:    true,
		Description: "Unique identifier of the plugin file",
	},
}

var dataSourceFilterSchema = map[string]datasourceschema.Attribute{
	"types": datasourceschema.ListAttribute{
		ElementType: types.StringType,
		Optional:    true,
	},
}
