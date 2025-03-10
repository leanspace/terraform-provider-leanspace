package streams

import (
	"github.com/leanspace/terraform-provider-leanspace/helper"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
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

var StreamSchema = map[string]*schema.Schema{
	"id": {
		Type:     schema.TypeString,
		Computed: true,
	},
	"version": {
		Type:        schema.TypeInt,
		Computed:    true,
		Description: "Version of the stream, this is incremented each time the stream is updated",
	},
	"name": {
		Type:     schema.TypeString,
		Required: true,
	},
	"description": {
		Type:     schema.TypeString,
		Optional: true,
	},
	"asset_id": {
		Type:         schema.TypeString,
		Required:     true,
		ForceNew:     true,
		ValidateFunc: validation.IsUUID,
	},
	"configuration": {
		Type:     schema.TypeList,
		Required: true,
		MaxItems: 1,
		MinItems: 1,
		Elem: &schema.Resource{
			Schema: configurationSchema,
		},
	},
	"mappings": {
		Type:     schema.TypeSet,
		Optional: true,
		Elem: &schema.Resource{
			Schema: mappingSchema,
		},
	},
	"created_at": {
		Type:        schema.TypeString,
		Computed:    true,
		Description: "When it was created",
	},
	"created_by": {
		Type:        schema.TypeString,
		Computed:    true,
		Description: "Who created it",
	},
	"last_modified_at": {
		Type:        schema.TypeString,
		Computed:    true,
		Description: "When it was last modified",
	},
	"last_modified_by": {
		Type:        schema.TypeString,
		Computed:    true,
		Description: "Who modified it the last",
	},
}

var configurationSchema = map[string]*schema.Schema{
	"endianness": {
		Type:         schema.TypeString,
		Required:     true,
		ValidateFunc: validation.StringInSlice(validEndianness, false),
		Description:  "Endianness of the stream, " + helper.AllowedValuesToDescription(validEndianness),
	},
	"structure": {
		Type:     schema.TypeList,
		Required: true,
		MaxItems: 1,
		MinItems: 1,
		Elem: &schema.Resource{
			// Arbitrary depth
			Schema: elementListSchema(streamComponentSchema(14), false),
		},
	},
	"metadata": {
		Type:     schema.TypeList,
		Required: true,
		MaxItems: 1,
		MinItems: 1,
		Elem: &schema.Resource{
			Schema: metadataSchema,
		},
	},
	"computations": {
		Type:     schema.TypeList,
		Required: true,
		MaxItems: 1,
		MinItems: 1,
		Elem: &schema.Resource{
			Schema: elementListSchema(computationSchema, true),
		},
	},
}

func streamComponentSchema(depth int) map[string]*schema.Schema {
	baseSchema := map[string]*schema.Schema{
		// Common
		"name": {
			Type:     schema.TypeString,
			Required: true,
		},
		"order": {
			Type:        schema.TypeInt,
			Computed:    true,
			Description: "Position of this component in the current context",
		},
		"path": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "Path of this component in the current context",
		},
		"type": {
			Type:         schema.TypeString,
			Required:     true,
			ValidateFunc: validation.StringInSlice([]string{"CONTAINER", "FIELD", "SWITCH"}, false),
		},
		"repetitive": {
			Type:     schema.TypeList,
			Optional: true,
			MinItems: 1,
			MaxItems: 1,
			Elem: &schema.Resource{
				Schema: repetitiveSchema,
			},
		},

		// Field only
		"length": {
			Type:     schema.TypeList,
			Optional: true,
			MinItems: 1,
			MaxItems: 1,
			Elem: &schema.Resource{
				Schema: lengthSchema,
			},
		},
		"processor": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Only required for fields",
		},
		"data_type": {
			Type:         schema.TypeString,
			Optional:     true,
			ValidateFunc: validation.StringInSlice(validFieldDataTypes, false),
			Description:  "Only required for fields, " + helper.AllowedValuesToDescription(validFieldDataTypes),
		},
		"endianness": {
			Type:         schema.TypeString,
			Optional:     true,
			ValidateFunc: validation.StringInSlice(validEndianness, false),
			Description:  "Only required for fields, " + helper.AllowedValuesToDescription(validEndianness),
		},
	}

	if depth > 0 {
		// Switch only
		baseSchema["expression"] = &schema.Schema{
			Type:        schema.TypeList,
			Optional:    true,
			MinItems:    1,
			MaxItems:    1,
			Description: "Only required for switches",
			Elem: &schema.Resource{
				Schema: switchExpressionSchema,
			},
		}
		// Container and switch only
		baseSchema["elements"] = &schema.Schema{
			Type:        schema.TypeList,
			Optional:    true,
			Description: "Only required for switches and containers",
			Elem: &schema.Resource{
				Schema: streamComponentSchema(depth - 1),
			},
		}
	}

	return baseSchema
}

var repetitiveSchema = map[string]*schema.Schema{
	"value": {
		Type:     schema.TypeInt,
		Optional: true,
	},
	"path": {
		Type:     schema.TypeString,
		Optional: true,
	},
}

var lengthSchema = map[string]*schema.Schema{
	"type": {
		Type:         schema.TypeString,
		Required:     true,
		ValidateFunc: validation.StringInSlice(validLengthTypes, false),
		Description:  "Type of the length, " + helper.AllowedValuesToDescription(validLengthTypes),
	},
	"unit": {
		Type:         schema.TypeString,
		Required:     true,
		ValidateFunc: validation.StringInSlice(validLengthUnits, false),
		Description:  "Unit of the length, " + helper.AllowedValuesToDescription(validLengthUnits),
	},
	"value": {
		Type:     schema.TypeInt,
		Optional: true,
	},
	"path": {
		Type:     schema.TypeString,
		Optional: true,
	},
}

var switchExpressionSchema = map[string]*schema.Schema{
	"switch_on": {
		Type:        schema.TypeString,
		Required:    true,
		Description: "Path of the field that the switch will use",
	},
	"options": {
		Type:     schema.TypeList,
		MinItems: 1,
		Required: true,
		Elem: &schema.Resource{
			Schema: switchOptionSchema,
		},
	},
}

var switchOptionSchema = map[string]*schema.Schema{
	"component": {
		Type:     schema.TypeString,
		Required: true,
	},
	"value": {
		Type:     schema.TypeList,
		Required: true,
		MinItems: 1,
		MaxItems: 1,
		Elem: &schema.Resource{
			Schema: switchValueSchema,
		},
	},
}

var switchValueSchema = map[string]*schema.Schema{
	"data_type": {
		Type:         schema.TypeString,
		Required:     true,
		ValidateFunc: validation.StringInSlice(validFieldDataTypes, false),
		Description:  helper.AllowedValuesToDescription(validFieldDataTypes),
	},
	"data": {
		Type:     schema.TypeString,
		Required: true,
	},
}

var metadataSchema = map[string]*schema.Schema{
	"timestamp": {
		Type:     schema.TypeList,
		Optional: true,
		Computed: true,
		MinItems: 1,
		MaxItems: 1,
		Elem: &schema.Resource{
			Schema: timestampDefinitionSchema,
		},
	},
}

var timestampDefinitionSchema = map[string]*schema.Schema{
	"expression": {
		Type:     schema.TypeString,
		Required: true,
	},
}

func elementListSchema(content map[string]*schema.Schema, valid bool) map[string]*schema.Schema {
	element := map[string]*schema.Schema{
		"elements": {
			Type:     schema.TypeList,
			Optional: true,
			Elem: &schema.Resource{
				Schema: content,
			},
		},
	}
	if valid {
		element["valid"] = &schema.Schema{
			Type:     schema.TypeBool,
			Computed: true,
		}
	}
	return element
}

var computationSchema = map[string]*schema.Schema{
	"name": {
		Type:     schema.TypeString,
		Required: true,
	},
	"order": {
		Type:     schema.TypeInt,
		Computed: true,
	},
	"type": {
		Type:     schema.TypeString,
		Computed: true,
	},
	"data_type": {
		Type:         schema.TypeString,
		Required:     true,
		ValidateFunc: validation.StringInSlice(validComputationDataTypes, false),
		Description:  helper.AllowedValuesToDescription(validComputationDataTypes),
	},
	"expression": {
		Type:        schema.TypeString,
		Required:    true,
		Description: "i.e.: javascript function with 2 input parameters and a return value (ctx, raw) => ctx.metadata.received_at",
	},
}

var mappingSchema = map[string]*schema.Schema{
	"metric_id": {
		Type:         schema.TypeString,
		Required:     true,
		ValidateFunc: validation.IsUUID,
	},
	"expression": {
		Type:     schema.TypeString,
		Required: true,
	},
}

var DataSourceFilterSchema = map[string]*schema.Schema{
	"asset_ids": {
		Type:     schema.TypeList,
		Optional: true,
		Elem: &schema.Schema{
			Type:         schema.TypeString,
			ValidateFunc: validation.IsUUID,
		},
	},
}
