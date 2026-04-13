package plan_templates

import (
	"regexp"

	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	resourceschema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"

	"github.com/leanspace/terraform-provider-leanspace/helper"
	"github.com/leanspace/terraform-provider-leanspace/helper/general_objects"
)

var nameRegex = regexp.MustCompile(`^[ a-zA-Z0-9_-]*$`)

var validResourceFunctionTimeUnits = []string{"SECONDS", "MINUTES", "HOURS", "DAYS"}
var validFormulaTypes = []string{"LINEAR", "RECTANGULAR"}

var planTemplateSchema = general_objects.ResourceSchemaWith(map[string]resourceschema.Attribute{
	"asset_id": resourceschema.StringAttribute{
		Required:      true,
		Validators:    helper.ValidUUID(),
		PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
	},
	"name": resourceschema.StringAttribute{
		Required:   true,
		Validators: []validator.String{stringvalidator.RegexMatches(nameRegex, "Must be a valid Plan Template name")},
	},
	"description": resourceschema.StringAttribute{
		Optional:   true,
		Validators: []validator.String{stringvalidator.LengthBetween(0, 2000)},
	},
	"integrity_status": resourceschema.StringAttribute{
		Computed: true,
	},
	"activity_configs": resourceschema.ListNestedAttribute{
		Optional: true,
		NestedObject: resourceschema.NestedAttributeObject{
			Attributes: activityConfigResultSchema,
		},
	},
	"estimated_duration_in_seconds": resourceschema.Int64Attribute{
		Computed: true,
	},
	"invalid_plan_template_reasons": resourceschema.ListNestedAttribute{
		Computed: true,
		NestedObject: resourceschema.NestedAttributeObject{
			Attributes: invalidPlanTemplateReasonSchema,
		},
	},
})

var activityConfigResultSchema = map[string]resourceschema.Attribute{
	"activity_definition_id": resourceschema.StringAttribute{
		Required:   true,
		Validators: helper.ValidUUID(),
	},
	"delay_reference_on_predecessor": resourceschema.StringAttribute{
		Optional: true,
	},

	"position": resourceschema.Int64Attribute{
		Required:   true,
		Validators: []validator.Int64{int64validator.Between(0, 499)},
	},

	"delay_in_seconds": resourceschema.Int64Attribute{
		Required:   true,
		Validators: []validator.Int64{int64validator.Between(0, 86400)},
	},

	"estimated_duration_in_seconds": resourceschema.Int64Attribute{
		Optional:   true,
		Validators: []validator.Int64{int64validator.Between(0, 86400)},
	},

	"name": resourceschema.StringAttribute{
		Optional:   true,
		Validators: []validator.String{stringvalidator.RegexMatches(nameRegex, "Must be a valid name")},
	},

	"arguments": resourceschema.ListNestedAttribute{
		Optional: true,
		NestedObject: resourceschema.NestedAttributeObject{
			Attributes: argumentSchema,
		},
	},

	"resource_function_formulas": resourceschema.ListNestedAttribute{
		Optional: true,
		NestedObject: resourceschema.NestedAttributeObject{
			Attributes: resourceFunctionFormulaOverloadSchema,
		},
	},

	"tags": general_objects.KeyValuesSchema,

	"definition_link_status": resourceschema.StringAttribute{
		Computed: true,
	},

	"invalid_definition_link_reasons": resourceschema.ListNestedAttribute{
		Computed: true,
		NestedObject: resourceschema.NestedAttributeObject{
			Attributes: invalidDefinitionLinkReasonSchema,
		},
	},
}

var invalidPlanTemplateReasonSchema = map[string]resourceschema.Attribute{
	"code": resourceschema.StringAttribute{
		Computed: true,
	},
	"message": resourceschema.StringAttribute{
		Computed: true,
	},
}

var argumentSchema = map[string]resourceschema.Attribute{
	"name": resourceschema.StringAttribute{
		Required:   true,
		Validators: []validator.String{stringvalidator.RegexMatches(nameRegex, "Must be a valid name")},
	},
	"attributes": resourceschema.SetNestedAttribute{
		Required: true,
		NestedObject: resourceschema.NestedAttributeObject{
			Attributes: general_objects.ValueAttributeSchema([]string{"TLE", "STRUCTURE"}),
		},
	},
}

var resourceFunctionFormulaOverloadSchema = map[string]resourceschema.Attribute{
	"resource_function_id": resourceschema.StringAttribute{
		Required:   true,
		Validators: helper.ValidUUID(),
	},
	"formula": resourceschema.SetNestedAttribute{
		Required: true,
		NestedObject: resourceschema.NestedAttributeObject{
			Attributes: resourceFunctionFormulaSchema,
		},
	},
}

var resourceFunctionFormulaSchema = map[string]resourceschema.Attribute{
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

var invalidDefinitionLinkReasonSchema = map[string]resourceschema.Attribute{
	"code": resourceschema.StringAttribute{
		Required: true,
	},
	"message": resourceschema.StringAttribute{
		Required: true,
	},
}
