package resources

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

var validResourceConstraintKinds = []string{"UPPER", "LOWER"}

var resourceSchema = general_objects.ResourceSchemaWith(map[string]resourceschema.Attribute{
	"asset_id": resourceschema.StringAttribute{
		Required:      true,
		Validators:    helper.ValidUUID(),
		PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
	},
	"unit_id": resourceschema.StringAttribute{
		Optional:      true,
		Validators:    helper.ValidUUID(),
		PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
	},
	"metric_id": resourceschema.StringAttribute{
		Optional:      true,
		Validators:    helper.ValidUUID(),
		PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
	},
	"name": resourceschema.StringAttribute{
		Required: true,
	},
	"description": resourceschema.StringAttribute{
		Optional: true,
	},
	"default_level": resourceschema.Float64Attribute{
		Optional: true,
	},
	"lower_limit": resourceschema.Float64Attribute{
		Optional: true,
	},
	"upper_limit": resourceschema.Float64Attribute{
		Optional: true,
	},
	"thresholds": resourceschema.SetNestedAttribute{
		Optional:    true,
		Description: "Currently, at most three LOWER and three UPPER thresholds can be set",
		NestedObject: resourceschema.NestedAttributeObject{
			Attributes: resourceThresholdSchema,
		},
	},
	"tags": general_objects.KeyValuesSchema,
})

var resourceThresholdSchema = map[string]resourceschema.Attribute{
	"kind": resourceschema.StringAttribute{
		Required:    true,
		Description: helper.AllowedValuesToDescription(validResourceConstraintKinds),
		Validators:  []validator.String{stringvalidator.OneOf(validResourceConstraintKinds...)},
	},
	"name": resourceschema.StringAttribute{
		Optional:   true,
		Validators: helper.ValidName(),
	},
	"violation_when_reached": resourceschema.BoolAttribute{
		Optional: true,
	},
	"value": resourceschema.Float64Attribute{
		Required: true,
	},
}

var dataSourceFilterSchema = map[string]datasourceschema.Attribute{
	"asset_ids": datasourceschema.ListAttribute{
		ElementType: types.StringType,
		Optional:    true,
		Validators:  []validator.List{listvalidator.ValueStringsAre(helper.ValidUUID()...)},
	},
	"unit_ids": datasourceschema.ListAttribute{
		ElementType: types.StringType,
		Optional:    true,
		Computed:    true,
		Validators:  []validator.List{listvalidator.ValueStringsAre(helper.ValidUUID()...)},
	},
	"metric_ids": datasourceschema.ListAttribute{
		ElementType: types.StringType,
		Optional:    true,
		Validators:  []validator.List{listvalidator.ValueStringsAre(helper.ValidUUID()...)},
	},
	"tags": datasourceschema.ListAttribute{
		ElementType: types.StringType,
		Optional:    true,
	},
	"created_bys": datasourceschema.ListAttribute{
		ElementType: types.StringType,
		Optional:    true,
		Description: "Filter on the user who created the Resource. If you have no wish to use this field as a filter, either provide a null value or remove the field.",
	},
	"last_modified_bys": datasourceschema.ListAttribute{
		ElementType: types.StringType,
		Optional:    true,
		Description: "Filter on the user who last modified the Resource. If you have no wish to use this field as a filter, either provide a null value or remove the field.",
	},
	"from_created_at": datasourceschema.StringAttribute{
		Optional:    true,
		Description: "Filter on the Resource creation date. Resources with a creation date greater or equals than the filter value will be selected (if they are not excluded by other filters). If you have no wish to use this field as a filter, either provide a null value or remove the field.",
	},
	"from_last_modified_at": datasourceschema.StringAttribute{
		Optional:    true,
		Description: "Filter on the Resource last modification date. Resources with a last modification date greater or equals than the filter value will be selected (if they are not excluded by other filters). If you have no wish to use this field as a filter, either provide a null value or remove the field.",
	},
	"to_created_at": datasourceschema.StringAttribute{
		Optional:    true,
		Description: "Filter on the Resource creation date. Resources with a creation date lower or equals than the filter value will be selected (if they are not excluded by other filters). If you have no wish to use this field as a filter, either provide a null value or remove the field.",
	},
	"to_last_modified_at": datasourceschema.StringAttribute{
		Optional:    true,
		Description: "Filter on the Resource last modification date. Resources with a last modification date lower or equals than the filter value will be selected (if they are not excluded by other filters). If you have no wish to use this field as a filter, either provide a null value or remove the field.",
	},
}
