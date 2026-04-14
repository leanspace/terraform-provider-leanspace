package action_templates

import (
	"regexp"

	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	datasourceschema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	resourceschema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/mapdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/leanspace/terraform-provider-leanspace/helper"
	"github.com/leanspace/terraform-provider-leanspace/helper/general_objects"
)

var validTypes = []string{"WEBHOOK", "LEANSPACE_EVENT"}

var ValidTriggeredOn = []string{
	"TRIGGERED",
	"OK",
}

var baseActionTemplateSchema = MakeActionTemplateSchema(false)

func MakeActionTemplateSchema(includeTriggeredOn bool) map[string]resourceschema.Attribute {
	baseSchema := general_objects.ResourceSchemaWith(map[string]resourceschema.Attribute{
		"name": resourceschema.StringAttribute{
			Required: !includeTriggeredOn,
			Computed: includeTriggeredOn,
		},
		"type": resourceschema.StringAttribute{
			Optional:    !includeTriggeredOn,
			Computed:    true,
			Default:     stringdefault.StaticString("WEBHOOK"),
			Description: helper.AllowedValuesToDescription(validTypes),
			Validators:  []validator.String{stringvalidator.OneOf(validTypes...)},
		},
		"url": resourceschema.StringAttribute{
			Optional: !includeTriggeredOn,
			Computed: includeTriggeredOn,
			Validators: []validator.String{
				stringvalidator.RegexMatches(regexp.MustCompile(`^https?://`), "must be a valid URL starting with http:// or https://"),
			},
		},
		"payload": resourceschema.StringAttribute{
			Optional: !includeTriggeredOn,
			Computed: includeTriggeredOn,
		},
		"content": resourceschema.StringAttribute{
			Computed: true,
		},
		"headers": resourceschema.MapAttribute{
			ElementType: types.StringType,
			Optional:    !includeTriggeredOn,
			Computed:    true,
			Default:     mapdefault.StaticValue(types.MapValueMust(types.StringType, map[string]attr.Value{})),
		},
	})

	if includeTriggeredOn {
		baseSchema["triggered_on"] = resourceschema.SetAttribute{
			Computed:    true,
			ElementType: types.StringType,
			Description: helper.AllowedValuesToDescription(ValidTriggeredOn),
		}
		baseSchema["last_modified_at"] = resourceschema.StringAttribute{
			Computed:      true,
			Description:   "When it was last modified",
			PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
		}
		baseSchema["last_modified_by"] = resourceschema.StringAttribute{
			Computed:      true,
			Description:   "Who modified it the last",
			PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
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
