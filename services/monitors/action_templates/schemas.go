package action_templates

import (
	"regexp"

	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	datasourceschema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	resourceschema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/mapdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/leanspace/terraform-provider-leanspace/helper"
)

var validTypes = []string{"WEBHOOK", "LEANSPACE_EVENT"}

var ValidTriggeredOn = []string{
	"TRIGGERED",
	"OK",
}

var baseActionTemplateSchema = MakeActionTemplateSchema(false)

func MakeActionTemplateSchema(includeTriggeredOn bool) map[string]resourceschema.Attribute {
	baseSchema := map[string]resourceschema.Attribute{
		"id": resourceschema.StringAttribute{
			Computed: true,
		},
		"name": resourceschema.StringAttribute{
			Required: true,
		},
		"type": resourceschema.StringAttribute{
			Optional:    true,
			Default:     stringdefault.StaticString("WEBHOOK"),
			Description: helper.AllowedValuesToDescription(validTypes),
			Validators:  []validator.String{stringvalidator.OneOf(validTypes...)},
		},
		"url": resourceschema.StringAttribute{
			Optional: true,
			Validators: []validator.String{
				stringvalidator.RegexMatches(regexp.MustCompile(`^https?://`), "must be a valid URL starting with http:// or https://"),
			},
		},
		"payload": resourceschema.StringAttribute{
			Optional: true,
		},
		"content": resourceschema.StringAttribute{
			Computed: true,
		},
		"headers": resourceschema.MapAttribute{
			ElementType: types.StringType,
			Optional:    true,
			Default:     mapdefault.StaticValue(types.MapValueMust(types.StringType, map[string]attr.Value{})),
		},
		"created_at": resourceschema.StringAttribute{
			Computed:    true,
			Description: "When it was created",
		},
		"created_by": resourceschema.StringAttribute{
			Computed:    true,
			Description: "Who created it",
		},
		"last_modified_at": resourceschema.StringAttribute{
			Computed:    true,
			Description: "When it was last modified",
		},
		"last_modified_by": resourceschema.StringAttribute{
			Computed:    true,
			Description: "Who modified it the last",
		},
	}

	if includeTriggeredOn {
		baseSchema["triggered_on"] = resourceschema.SetAttribute{
			Optional:    true,
			ElementType: types.StringType,
			Description: helper.AllowedValuesToDescription(ValidTriggeredOn),
		}
	}

	return baseSchema
}

var dataSourceFilterSchema = map[string]datasourceschema.Attribute{
	"types": datasourceschema.ListAttribute{
		ElementType: types.StringType,
		Optional:    true,
		Validators:  []validator.List{listvalidator.ValueStringsAre(stringvalidator.OneOf(validTypes...))},
	},
	"monitor_ids": datasourceschema.ListAttribute{
		ElementType: types.StringType,
		Optional:    true,
		Validators:  []validator.List{listvalidator.ValueStringsAre(helper.ValidUUID()...)},
	},
}
