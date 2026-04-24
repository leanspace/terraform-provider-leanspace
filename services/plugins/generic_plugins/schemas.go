package generic_plugins

import (
	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	datasourceschema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	resourceschema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/leanspace/terraform-provider-leanspace/helper"
	"github.com/leanspace/terraform-provider-leanspace/helper/general_objects"
)

var validGenericPluginTypes = []string{
	"CHECKSUM_FUNCTION",
}

var validGenericPluginLanguages = []string{
	"JAVA",
}

var validGenericPluginStatuses = []string{
	"ACTIVE",
	"PENDING",
	"FAILED",
}

var genericPluginSchema = general_objects.ResourceSchemaWith(map[string]resourceschema.Attribute{
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
	"source_code_link": resourceschema.SingleNestedAttribute{
		Computed:   true,
		Attributes: sourceCodeLinkSchema,
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
})

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
		Validators:  []validator.List{listvalidator.ValueStringsAre(stringvalidator.OneOf(validGenericPluginStatuses...))},
	},
	"types": datasourceschema.ListAttribute{
		ElementType: types.StringType,
		Optional:    true,
		Validators:  []validator.List{listvalidator.ValueStringsAre(stringvalidator.OneOf(append([]string{"JOB"}, validGenericPluginTypes...)...))},
	},
}
