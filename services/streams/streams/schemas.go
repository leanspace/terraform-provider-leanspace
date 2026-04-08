package streams

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

var validFieldDataTypes = []string{
	"INTEGER", "UINTEGER", "DECIMAL", "TEXT", "BOOLEAN", "BINARY",
}

var validComputationDataTypes = []string{
	"INTEGER", "UINTEGER", "DECIMAL", "TEXT", "BOOLEAN", "BINARY", "TIMESTAMP", "DATE",
}

var validEndianness = []string{
	"BE", "LE",
}

var validLengthUnits = []string{
	"BITS", "BYTES",
}

var validLengthTypes = []string{
	"FIXED", "DYNAMIC",
}

var StreamSchema = general_objects.ResourceSchemaWith(map[string]resourceschema.Attribute{
	"version": resourceschema.Int64Attribute{
		Computed:    true,
		Description: "Version of the stream, this is incremented each time the stream is updated",
	},
	"name": resourceschema.StringAttribute{
		Required: true,
	},
	"description": resourceschema.StringAttribute{
		Optional: true,
	},
	"asset_id": resourceschema.StringAttribute{
		Required:      true,
		Validators:    helper.ValidUUID(),
		PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
	},
	"tags": general_objects.KeyValuesSchema,
	"configuration": resourceschema.SingleNestedAttribute{
		Required:   true,
		Attributes: configurationSchema,
	},
	"mappings": resourceschema.SetNestedAttribute{
		Optional: true,
		NestedObject: resourceschema.NestedAttributeObject{
			Attributes: mappingSchema,
		},
	},
})

var configurationSchema = map[string]resourceschema.Attribute{
	"endianness": resourceschema.StringAttribute{
		Required:    true,
		Description: "Endianness of the stream, " + helper.AllowedValuesToDescription(validEndianness),
		Validators:  []validator.String{stringvalidator.OneOf(validEndianness...)},
	},
	"structure": resourceschema.SingleNestedAttribute{
		Required:   true,
		Attributes: elementListSchema(streamComponentSchema(14), false),
	},
	"metadata": resourceschema.SingleNestedAttribute{
		Required:   true,
		Attributes: metadataSchema,
	},
	"computations": resourceschema.SingleNestedAttribute{
		Required:   true,
		Attributes: elementListSchema(computationSchema, true),
	},
}

func streamComponentSchema(depth int) map[string]resourceschema.Attribute {
	baseSchema := map[string]resourceschema.Attribute{
		// Common
		"name": resourceschema.StringAttribute{
			Required: true,
		},
		"order": resourceschema.Int64Attribute{
			Computed:    true,
			Description: "Position of this component in the current context",
		},
		"path": resourceschema.StringAttribute{
			Computed:    true,
			Description: "Path of this component in the current context",
		},
		"type": resourceschema.StringAttribute{
			Required:   true,
			Validators: []validator.String{stringvalidator.OneOf("CONTAINER", "FIELD", "SWITCH")},
		},
		"repetitive": resourceschema.SingleNestedAttribute{
			Optional:   true,
			Attributes: repetitiveSchema,
		},

		// Field only
		"length": resourceschema.SingleNestedAttribute{
			Optional:   true,
			Attributes: lengthSchema,
		},
		"processor": resourceschema.StringAttribute{
			Optional:    true,
			Description: "Only required for fields",
		},
		"data_type": resourceschema.StringAttribute{
			Optional:    true,
			Description: "Only required for fields, " + helper.AllowedValuesToDescription(validFieldDataTypes),
			Validators:  []validator.String{stringvalidator.OneOf(validFieldDataTypes...)},
		},
		"endianness": resourceschema.StringAttribute{
			Optional:    true,
			Description: "Only required for fields, " + helper.AllowedValuesToDescription(validEndianness),
			Validators:  []validator.String{stringvalidator.OneOf(validEndianness...)},
		},
	}

	if depth > 0 {
		// Switch only
		baseSchema["expression"] = resourceschema.SingleNestedAttribute{
			Optional:    true,
			Description: "Only required for switches",
			Attributes:  switchExpressionSchema,
		}
		// Container and switch only
		baseSchema["elements"] = resourceschema.ListNestedAttribute{
			Optional:    true,
			Description: "Only required for switches and containers",
			NestedObject: resourceschema.NestedAttributeObject{
				Attributes: streamComponentSchema(depth - 1),
			},
		}
	}

	return baseSchema
}

var repetitiveSchema = map[string]resourceschema.Attribute{
	"value": resourceschema.Int64Attribute{
		Optional: true,
	},
	"path": resourceschema.StringAttribute{
		Optional: true,
	},
}

var lengthSchema = map[string]resourceschema.Attribute{
	"type": resourceschema.StringAttribute{
		Optional:    true,
		Description: "Type of the length, " + helper.AllowedValuesToDescription(validLengthTypes),
		Validators: []validator.String{
			stringvalidator.OneOf(validLengthTypes...),
			helper.RequiredIfParentConfigured(),
		},
	},
	"unit": resourceschema.StringAttribute{
		Optional:    true,
		Description: "Unit of the length, " + helper.AllowedValuesToDescription(validLengthUnits),
		Validators: []validator.String{
			stringvalidator.OneOf(validLengthUnits...),
			helper.RequiredIfParentConfigured(),
		},
	},
	"value": resourceschema.Int64Attribute{
		Optional: true,
	},
	"path": resourceschema.StringAttribute{
		Optional: true,
	},
}

var switchExpressionSchema = map[string]resourceschema.Attribute{
	"switch_on": resourceschema.StringAttribute{
		Optional:    true,
		Description: "Path of the field that the switch will use",
		Validators:  []validator.String{helper.RequiredIfParentConfigured()},
	},
	"options": resourceschema.ListNestedAttribute{
		Optional: true,
		NestedObject: resourceschema.NestedAttributeObject{
			Attributes: switchOptionSchema,
		},
	},
}

var switchOptionSchema = map[string]resourceschema.Attribute{
	"component": resourceschema.StringAttribute{
		Required: true,
	},
	"value": resourceschema.SingleNestedAttribute{
		Required:   true,
		Attributes: switchValueSchema,
	},
}

var switchValueSchema = map[string]resourceschema.Attribute{
	"data_type": resourceschema.StringAttribute{
		Required:    true,
		Description: helper.AllowedValuesToDescription(validFieldDataTypes),
		Validators:  []validator.String{stringvalidator.OneOf(validFieldDataTypes...)},
	},
	"data": resourceschema.StringAttribute{
		Required: true,
	},
}

var metadataSchema = map[string]resourceschema.Attribute{
	"timestamp": resourceschema.SingleNestedAttribute{
		Optional:   true,
		Computed:   true,
		Attributes: timestampDefinitionSchema,
	},
}

var timestampDefinitionSchema = map[string]resourceschema.Attribute{
	"expression": resourceschema.StringAttribute{
		Required: true,
	},
}

func elementListSchema(content map[string]resourceschema.Attribute, valid bool) map[string]resourceschema.Attribute {
	element := map[string]resourceschema.Attribute{
		"elements": resourceschema.ListNestedAttribute{
			Optional: true,
			NestedObject: resourceschema.NestedAttributeObject{
				Attributes: content,
			},
		},
	}
	if valid {
		element["valid"] = resourceschema.BoolAttribute{
			Computed: true,
		}
	}
	return element
}

var computationSchema = map[string]resourceschema.Attribute{
	"name": resourceschema.StringAttribute{
		Required: true,
	},
	"order": resourceschema.Int64Attribute{
		Computed: true,
	},
	"type": resourceschema.StringAttribute{
		Computed: true,
	},
	"data_type": resourceschema.StringAttribute{
		Required:    true,
		Description: helper.AllowedValuesToDescription(validComputationDataTypes),
		Validators:  []validator.String{stringvalidator.OneOf(validComputationDataTypes...)},
	},
	"expression": resourceschema.StringAttribute{
		Required:    true,
		Description: "i.e.: javascript function with 2 input parameters and a return value (ctx, raw) => ctx.metadata.received_at",
	},
}

var mappingSchema = map[string]resourceschema.Attribute{
	"metric_id": resourceschema.StringAttribute{
		Required:   true,
		Validators: helper.ValidUUID(),
	},
	"expression": resourceschema.StringAttribute{
		Required: true,
	},
}

var DataSourceFilterSchema = map[string]datasourceschema.Attribute{
	"asset_ids": datasourceschema.ListAttribute{
		ElementType: types.StringType,
		Optional:    true,
		Validators:  []validator.List{listvalidator.ValueStringsAre(helper.ValidUUID()...)},
	},
}
