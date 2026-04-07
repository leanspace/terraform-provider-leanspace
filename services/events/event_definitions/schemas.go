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

var eventsDefinitions = map[string]resourceschema.Attribute{
	"id": resourceschema.StringAttribute{
		Computed: true,
	},
	"description": resourceschema.StringAttribute{
		Optional: true,
	},
	"name": resourceschema.StringAttribute{
		Required: true,
	},
	"criticality": resourceschema.StringAttribute{
		Optional: true,
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
	"tags": general_objects.KeyValuesSchema,
}

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
