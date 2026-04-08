package monitors

import (
	"github.com/hashicorp/terraform-plugin-framework-validators/float64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/setvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	datasourceschema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	resourceschema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/leanspace/terraform-provider-leanspace/helper"
	"github.com/leanspace/terraform-provider-leanspace/helper/general_objects"
	"github.com/leanspace/terraform-provider-leanspace/services/monitors/action_templates"
)

var validComparisonOperators = []string{
	"GREATER_THAN",
	"LESSER_THAN",
	"GREATER_THAN_OR_EQUAL_TO",
	"LESSER_THAN_OR_EQUAL_TO",
	"EQUAL_TO",
	"NOT_EQUAL_TO",
}

var actionTemplateSchema = action_templates.MakeActionTemplateSchema(true)

var monitorSchema = general_objects.ResourceSchemaWith(map[string]resourceschema.Attribute{
	"name": resourceschema.StringAttribute{
		Required: true,
	},
	"description": resourceschema.StringAttribute{
		Optional: true,
	},
	"status": resourceschema.StringAttribute{
		Computed: true,
	},
	"metric_id": resourceschema.StringAttribute{
		Required:   true,
		Validators: helper.ValidUUID(),
	},
	"node_id": resourceschema.StringAttribute{
		Computed: true,
	},
	"rule": resourceschema.SingleNestedAttribute{
		Required:   true,
		Attributes: ruleSchema,
	},
	"action_templates": resourceschema.SetNestedAttribute{
		Computed: true,
		NestedObject: resourceschema.NestedAttributeObject{
			Attributes: actionTemplateSchema,
		},
	},
	"action_template_links": resourceschema.SetNestedAttribute{
		Optional: true,
		NestedObject: resourceschema.NestedAttributeObject{
			Attributes: actionTemplateLinkSchema,
		},
	},
	"tags": general_objects.KeyValuesSchema,
	"type": resourceschema.StringAttribute{
		Computed:    true,
		Description: "Represent the type of the monitor. This field is deprecated and it will be removed soon. Please use only this type: REALTIME.",
	},
})

var ruleSchema = map[string]resourceschema.Attribute{ // ruleSchema
	"comparison_operator": resourceschema.StringAttribute{
		Required:    true,
		Description: helper.AllowedValuesToDescription(validComparisonOperators),
		Validators:  []validator.String{stringvalidator.OneOf(validComparisonOperators...)},
	},
	"comparison_value": resourceschema.Float64Attribute{
		Required: true,
	},
	"tolerance": resourceschema.Float64Attribute{
		Optional:    true,
		Description: "Only valid for EQUAL_TO or NOT_EQUAL_TO comparison operator",
		Validators:  []validator.Float64{float64validator.AtLeast(0)},
	},
}

var actionTemplateLinkSchema = map[string]resourceschema.Attribute{
	"id": resourceschema.StringAttribute{
		Required:    true,
		Description: "Identifier of the Action Template",
		Validators:  helper.ValidUUID(),
	},
	"triggered_on": resourceschema.SetAttribute{
		ElementType: types.StringType,
		Optional:    true,
		Validators:  []validator.Set{setvalidator.ValueStringsAre(stringvalidator.OneOf(action_templates.ValidTriggeredOn...))},
	},
}

var dataSourceFilterSchema = map[string]datasourceschema.Attribute{
	"metric_ids": datasourceschema.ListAttribute{
		ElementType: types.StringType,
		Optional:    true,
		Validators:  []validator.List{listvalidator.ValueStringsAre(helper.ValidUUID()...)},
	},
	"name": datasourceschema.StringAttribute{
		Optional: true,
	},
	"node_ids": datasourceschema.ListAttribute{
		ElementType: types.StringType,
		Optional:    true,
		Validators:  []validator.List{listvalidator.ValueStringsAre(helper.ValidUUID()...)},
	},
	"statuses": datasourceschema.ListAttribute{
		ElementType: types.StringType,
		Optional:    true,
	},
	"tags": datasourceschema.ListAttribute{
		ElementType: types.StringType,
		Optional:    true,
	},
	"action_template_ids": datasourceschema.ListAttribute{
		ElementType: types.StringType,
		Optional:    true,
		Validators:  []validator.List{listvalidator.ValueStringsAre(helper.ValidUUID()...)},
	},
}
