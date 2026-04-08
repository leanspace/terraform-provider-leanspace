package event_definitions

import (
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	datasourceschema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	resourceschema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/leanspace/terraform-provider-leanspace/helper"
	"github.com/leanspace/terraform-provider-leanspace/helper/general_objects"
)

var validOperator = []string{"EQUAL_TO"}
var source = []string{"COMMAND_STATE_CHANGED", "MONITOR_TRIGGERED", "PASS_AOS", "PASS_LOS", "STREAM_DECODED", "CUSTOM", "FILE_UPLOADED"}
var state = []string{"ACTIVE", "INACTIVE"}
var validMetadataTypes = []string{
	"NUMERIC", "BOOLEAN", "TEXT",
}

var eventsDefinitions = general_objects.ResourceSchemaWith(map[string]resourceschema.Attribute{
	"description": resourceschema.StringAttribute{
		Optional: true,
	},
	"name": resourceschema.StringAttribute{
		Required: true,
	},
	"criticality": resourceschema.StringAttribute{
		Optional: true,
		Computed: true,
		Default:  stringdefault.StaticString("NORMAL"),
	},
	"rules": resourceschema.SetNestedAttribute{
		Optional: true,
		NestedObject: resourceschema.NestedAttributeObject{
			Attributes: ruleSchema,
		},
	},
	"source": resourceschema.StringAttribute{
		Required:    true,
		Description: helper.AllowedValuesToDescription(source),
		Validators:  []validator.String{stringvalidator.OneOf(source...)},
	},
	"state": resourceschema.StringAttribute{
		Required:    true,
		Description: helper.AllowedValuesToDescription(state),
		Validators:  []validator.String{stringvalidator.OneOf(state...)},
	},
	"tags": general_objects.KeyValuesSchema,
})

var ruleSchema = map[string]resourceschema.Attribute{
	"path": resourceschema.StringAttribute{
		Required: true,
	},
	"operator": resourceschema.StringAttribute{
		Required:    true,
		Description: helper.AllowedValuesToDescription(validOperator),
		Validators:  []validator.String{stringvalidator.OneOf(validOperator...)},
	},
	"comparison_value": resourceschema.SingleNestedAttribute{
		Optional:   true,
		Attributes: ComparisonValueAttributeSchema,
	},
}

var ComparisonValueAttributeSchema = map[string]resourceschema.Attribute{
	"value": resourceschema.StringAttribute{
		Optional: true,
	},
	"type": resourceschema.StringAttribute{
		Required:    true,
		Description: helper.AllowedValuesToDescription(validMetadataTypes),
		Validators:  []validator.String{stringvalidator.OneOf(validMetadataTypes...)},
	},
}

var dataSourceFilterSchema = map[string]datasourceschema.Attribute{
	"tags": datasourceschema.ListAttribute{
		ElementType: types.StringType,
		Optional:    true,
	},
}
