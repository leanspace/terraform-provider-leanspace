package resources

import (
	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	datasourceschema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	resourceschema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/float64default"
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
		Computed:      true,
		Validators:    helper.ValidUUID(),
		PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown(), stringplanmodifier.RequiresReplace()},
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
		Computed: true,
		Default:  float64default.StaticFloat64(0.0),
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
		Computed: true,
		Default:  booldefault.StaticBool(false),
	},
	"value": resourceschema.Float64Attribute{
		Required: true,
	},
}

var dataSourceFilterSchema = general_objects.AuditFilterFieldsWithTags(map[string]datasourceschema.Attribute{
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
	"default_level": datasourceschema.Float64Attribute{
		Optional:    true,
		Description: "The default level of the resource.",
	},
})
