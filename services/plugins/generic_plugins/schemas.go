package generic_plugins

import (
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	datasourceschema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	resourceschema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/leanspace/terraform-provider-leanspace/helper"
)

var validGenericPluginTypes = []string{
	"CHECKSUM_FUNCTION",
}

var validGenericPluginLanguages = []string{
	"JAVA",
}

var genericPluginSchema = map[string]resourceschema.Attribute{
	"id": resourceschema.StringAttribute{
		Computed: true,
	},
	"name": resourceschema.StringAttribute{
		Required: true,
	},
	"description": resourceschema.StringAttribute{
		Optional: true,
	},
	"type": resourceschema.StringAttribute{
		Required:    true,
		Description: helper.AllowedValuesToDescription(validGenericPluginTypes),
		Validators:  []validator.String{stringvalidator.OneOf(validGenericPluginTypes...)},
	},
	"language": resourceschema.StringAttribute{
		Required:    true,
		Description: helper.AllowedValuesToDescription(validGenericPluginLanguages),
		Validators:  []validator.String{stringvalidator.OneOf(validGenericPluginLanguages...)},
	},
	"source_code_link": resourceschema.ListNestedAttribute{
		Computed: true,
		NestedObject: resourceschema.NestedAttributeObject{
			Attributes: sourceCodeLinkSchema,
		},
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
	"status": resourceschema.StringAttribute{
		Computed:    true,
		Description: "Generic Plugin status. Can be ACTIVE, PENDING or FAILED",
	},
	"source_code_path": resourceschema.StringAttribute{
		Required:    true,
		Description: "It must be a valid path to a .jar file",
		Validators:  []validator.String{stringvalidator.RegexMatches(helper.PathToJarFileRegex, "Must be a valid file path to a .jar file")},
	},
	"source_code_sha": resourceschema.StringAttribute{
		Computed:    true,
		Description: "Unique identifier of the generic plugin file",
	},
}

var sourceCodeLinkSchema = map[string]resourceschema.Attribute{
	"expiration_time": resourceschema.StringAttribute{
		Computed:    true,
		Description: "When the source code link expires",
	},
	"source_code_id": resourceschema.StringAttribute{
		Computed:    true,
		Description: "Unique identifier of the source code",
	},
	"url": resourceschema.StringAttribute{
		Computed:    true,
		Description: "URL to download the source code",
	},
}

var dataSourceFilterSchema = map[string]datasourceschema.Attribute{
	"statuses": datasourceschema.ListAttribute{
		ElementType: types.StringType,
		Optional:    true,
	},
	"types": datasourceschema.ListAttribute{
		ElementType: types.StringType,
		Optional:    true,
	},
}
