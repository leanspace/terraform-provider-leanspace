package streams

import (
	"leanspace-terraform-provider/helper"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

var validDataTypes = []string{
	"INTEGER", "UINTEGER", "DECIMAL", "TEXT", "BOOLEAN",
}

var validEndianness = []string{
	"BE", "LE",
}

var streamSchema = map[string]*schema.Schema{
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
			Schema: elementListSchema(streamComponentSchema(5)),
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
			Schema: elementListSchema(computationSchema),
		},
	},
	"valid": {
		Type:     schema.TypeBool,
		Computed: true,
	},
	"errors": {
		Type:     schema.TypeSet,
		Computed: true,
		Elem: &schema.Resource{
			Schema: errorSchema,
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
		"valid": {
			Type:     schema.TypeBool,
			Computed: true,
		},
		"errors": {
			Type:     schema.TypeSet,
			Computed: true,
			Elem: &schema.Resource{
				Schema: errorSchema,
			},
		},

		// Field only
		"length_in_bits": {
			Type:         schema.TypeInt,
			Optional:     true,
			ValidateFunc: validation.IntAtLeast(1),
		},
		"processor": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"data_type": {
			Type:         schema.TypeString,
			Optional:     true,
			ValidateFunc: validation.StringInSlice(validDataTypes, false),
			Description:  helper.AllowedValuesToDescription(validDataTypes),
		},
		"endianness": {
			Type:         schema.TypeString,
			Optional:     true,
			ValidateFunc: validation.StringInSlice(validEndianness, false),
			Description:  helper.AllowedValuesToDescription(validEndianness),
		},
	}

	if depth > 0 {
		// Switch only
		baseSchema["expression"] = &schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			MinItems: 1,
			MaxItems: 1,
			Elem: &schema.Resource{
				Schema: switchExpressionSchema,
			},
		}
		// Container and switch only
		baseSchema["elements"] = &schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			Elem: &schema.Resource{
				Schema: streamComponentSchema(depth - 1),
			},
		}
	}

	return baseSchema
}

var switchExpressionSchema = map[string]*schema.Schema{
	"switch_on": {
		Type:        schema.TypeString,
		Required:    true,
		Description: "Path of the field that the switch will use",
	},
	"options": {
		Type:     schema.TypeSet,
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
		ValidateFunc: validation.StringInSlice(validDataTypes, false),
		Description:  helper.AllowedValuesToDescription(validDataTypes),
	},
	"data": {
		Type:     schema.TypeString,
		Required: true,
	},
}

var metadataSchema = map[string]*schema.Schema{
	"packet_id": {
		Type:     schema.TypeList,
		Computed: true,
		Elem: &schema.Resource{
			Schema: elementStatusSchema,
		},
	},
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
	"valid": {
		Type:     schema.TypeBool,
		Computed: true,
	},
	"errors": {
		Type:     schema.TypeSet,
		Computed: true,
		Elem: &schema.Resource{
			Schema: errorSchema,
		},
	},
}

var elementStatusSchema = map[string]*schema.Schema{
	"valid": {
		Type:     schema.TypeBool,
		Computed: true,
	},
	"errors": {
		Type:     schema.TypeSet,
		Computed: true,
		Elem: &schema.Resource{
			Schema: errorSchema,
		},
	},
}

var timestampDefinitionSchema = map[string]*schema.Schema{
	"expression": {
		Type:     schema.TypeString,
		Required: true,
	},
	"valid": {
		Type:     schema.TypeBool,
		Computed: true,
	},
	"errors": {
		Type:     schema.TypeSet,
		Computed: true,
		Elem: &schema.Resource{
			Schema: errorSchema,
		},
	},
}

func elementListSchema(content map[string]*schema.Schema) map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"elements": {
			Type:     schema.TypeList,
			Optional: true,
			Elem: &schema.Resource{
				Schema: content,
			},
		},
		"valid": {
			Type:     schema.TypeBool,
			Computed: true,
		},
		"errors": {
			Type:     schema.TypeSet,
			Computed: true,
			Elem: &schema.Resource{
				Schema: errorSchema,
			},
		},
	}
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
		ValidateFunc: validation.StringInSlice(validDataTypes, false),
		Description:  helper.AllowedValuesToDescription(validDataTypes),
	},
	"expression": {
		Type:        schema.TypeString,
		Required:    true,
		Description: "i.e.: javascript function with 2 input parameters and a return value (ctx, raw) => ctx['metadata.received_at']",
	},
	"valid": {
		Type:     schema.TypeBool,
		Computed: true,
	},
	"errors": {
		Type:     schema.TypeSet,
		Computed: true,
		Elem: &schema.Resource{
			Schema: errorSchema,
		},
	},
}

var mappingSchema = map[string]*schema.Schema{
	"metric_id": {
		Type:         schema.TypeString,
		Required:     true,
		ValidateFunc: validation.IsUUID,
	},
	"component": {
		Type:     schema.TypeString,
		Required: true,
	},
}

var errorSchema = map[string]*schema.Schema{
	"code": {
		Type:     schema.TypeString,
		Computed: true,
	},
	"message": {
		Type:     schema.TypeString,
		Computed: true,
	},
}

var dataSourceFilterSchema = map[string]*schema.Schema{
	"asset_ids": {
		Type:     schema.TypeList,
		Optional: true,
		Elem: &schema.Schema{
			Type:         schema.TypeString,
			ValidateFunc: validation.IsUUID,
		},
	},
}
