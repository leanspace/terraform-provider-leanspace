package passive_resource_functions

import (
	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	datasourceschema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	resourceschema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/leanspace/terraform-provider-leanspace/helper"
	"github.com/leanspace/terraform-provider-leanspace/helper/general_objects"
)

var validPassiveResourceFunctionTimeUnits = []string{"SECONDS", "MINUTES", "HOURS", "DAYS"}
var validPassiveResourceFunctionFormulaTypes = []string{"LINEAR"}

var passiveResourceFunctionSchema = general_objects.ResourceSchemaWith(map[string]resourceschema.Attribute{
	"resource_id": resourceschema.StringAttribute{
		Required:      true,
		Validators:    helper.ValidUUID(),
		PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
	},
	"name": resourceschema.StringAttribute{
		Optional: true,
	},
	"control_bound": resourceschema.Float64Attribute{
		Optional:    true,
		Description: "The function stops impacting the resource level once the controlBound is reached",
	},
	"formula": resourceschema.SingleNestedAttribute{
		Required:   true,
		Attributes: formulaSchema,
	},
	"tags": general_objects.KeyValuesSchema,
})

var formulaSchema = map[string]resourceschema.Attribute{
	"rate": resourceschema.Float64Attribute{
		Required: true,
	},
	"type": resourceschema.StringAttribute{
		Required:    true,
		Description: helper.AllowedValuesToDescription(validPassiveResourceFunctionFormulaTypes),
		Validators:  []validator.String{stringvalidator.OneOf(validPassiveResourceFunctionFormulaTypes...)},
	},
	"time_unit": resourceschema.StringAttribute{
		Required:    true,
		Description: helper.AllowedValuesToDescription(validPassiveResourceFunctionTimeUnits),
		Validators:  []validator.String{stringvalidator.OneOf(validPassiveResourceFunctionTimeUnits...)},
	},
}

var dataSourceFilterSchema = general_objects.AuditFilterFieldsWithTags(map[string]datasourceschema.Attribute{
	"resource_ids": datasourceschema.ListAttribute{
		ElementType: types.StringType,
		Optional:    true,
		Validators:  []validator.List{listvalidator.ValueStringsAre(helper.ValidUUID()...)},
	},
})
