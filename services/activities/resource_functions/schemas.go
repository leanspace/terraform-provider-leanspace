package resource_functions

import (
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

var validResourceFunctionTimeUnits = []string{"SECONDS", "MINUTES", "HOURS", "DAYS"}
var validFormulaTypes = []string{"LINEAR", "RECTANGULAR"}

var resourceFunctionSchema = general_objects.ResourceSchemaWith(map[string]resourceschema.Attribute{
	"activity_definition_id": resourceschema.StringAttribute{
		Required:      true,
		Validators:    helper.ValidUUID(),
		PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
	},
	"resource_id": resourceschema.StringAttribute{
		Required:      true,
		Validators:    helper.ValidUUID(),
		PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
	},
	"name": resourceschema.StringAttribute{
		Optional: true,
	},
	"formula": resourceschema.SingleNestedAttribute{
		Required:   true,
		Attributes: formulaSchema,
	},
})

var formulaSchema = map[string]resourceschema.Attribute{
	"type": resourceschema.StringAttribute{
		Required:    true,
		Description: helper.AllowedValuesToDescription(validFormulaTypes),
		Validators:  []validator.String{stringvalidator.OneOf(validFormulaTypes...)},
	},
	"amplitude": resourceschema.Float64Attribute{
		Optional: true,
	},
	"constant": resourceschema.Float64Attribute{
		Optional: true,
	},
	"rate": resourceschema.Float64Attribute{
		Optional: true,
	},
	"time_unit": resourceschema.StringAttribute{
		Optional:    true,
		Description: helper.AllowedValuesToDescription(validResourceFunctionTimeUnits),
		Validators:  []validator.String{stringvalidator.OneOf(validResourceFunctionTimeUnits...)},
	},
}

var dataSourceFilterSchema = map[string]datasourceschema.Attribute{
	"activity_definition_ids": datasourceschema.ListAttribute{
		ElementType: types.StringType,
		Optional:    true,
	},
	"resource_ids": datasourceschema.ListAttribute{
		ElementType: types.StringType,
		Optional:    true,
	},
	"tags": datasourceschema.ListAttribute{
		ElementType: types.StringType,
		Optional:    true,
	},
	"created_bys": datasourceschema.ListAttribute{
		ElementType: types.StringType,
		Optional:    true,
		Description: "Filter on the user who created the Resource Function. If you have no wish to use this field as a filter, either provide a null value or remove the field.",
	},
	"last_modified_bys": datasourceschema.ListAttribute{
		ElementType: types.StringType,
		Optional:    true,
		Description: "Filter on the user who last modified the Resource Function. If you have no wish to use this field as a filter, either provide a null value or remove the field.",
	},
	"from_created_at": datasourceschema.StringAttribute{
		Optional:    true,
		Description: "Filter on the Resource Function creation date. Resource Functions with a creation date greater or equals than the filter value will be selected (if they are not excluded by other filters). If you have no wish to use this field as a filter, either provide a null value or remove the field.",
	},
	"from_last_modified_at": datasourceschema.StringAttribute{
		Optional:    true,
		Description: "Filter on the Resource Function last modification date. Resource Functions with a last modification date greater or equals than the filter value will be selected (if they are not excluded by other filters). If you have no wish to use this field as a filter, either provide a null value or remove the field.",
	},
	"to_created_at": datasourceschema.StringAttribute{
		Optional:    true,
		Description: "Filter on the Resource Function creation date. Resource Functions with a creation date lower or equals than the filter value will be selected (if they are not excluded by other filters). If you have no wish to use this field as a filter, either provide a null value or remove the field.",
	},
	"to_last_modified_at": datasourceschema.StringAttribute{
		Optional:    true,
		Description: "Filter on the Resource Function last modification date. Resource Functions with a last modification date lower or equals than the filter value will be selected (if they are not excluded by other filters). If you have no wish to use this field as a filter, either provide a null value or remove the field.",
	},
}
