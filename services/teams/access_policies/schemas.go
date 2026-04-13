package access_policies

import (
	"regexp"

	"github.com/hashicorp/terraform-plugin-framework-validators/setvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	datasourceschema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	resourceschema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/leanspace/terraform-provider-leanspace/helper/general_objects"
)

var accessPolicySchema = general_objects.ResourceSchemaWith(map[string]resourceschema.Attribute{
	"name": resourceschema.StringAttribute{
		Required: true,
	},
	"description": resourceschema.StringAttribute{
		Optional: true,
	},
	"read_only": resourceschema.BoolAttribute{
		Computed: true,
	},
	"statements": resourceschema.SetNestedAttribute{
		Optional: true,
		NestedObject: resourceschema.NestedAttributeObject{
			Attributes: statementSchema,
		},
	},
	"tags": general_objects.KeyValuesSchema,
})

var actionRegex = regexp.MustCompile(`^([a-zA-Z0-9*]+):([a-zA-Z0-9*]+)$`)

var statementSchema = map[string]resourceschema.Attribute{
	"name": resourceschema.StringAttribute{
		Required: true,
	},
	"actions": resourceschema.SetAttribute{
		ElementType: types.StringType,
		Required:    true,
		Validators:  []validator.Set{setvalidator.ValueStringsAre(stringvalidator.RegexMatches(actionRegex, "must match pattern: <service>:<action> (e.g. nodes:READ)"))},
	},
}

var dataSourceFilterSchema = map[string]datasourceschema.Attribute{
	"actions": datasourceschema.ListAttribute{
		ElementType: types.StringType,
		Optional:    true,
	},
	"tags": datasourceschema.ListAttribute{
		ElementType: types.StringType,
		Optional:    true,
	},
}
