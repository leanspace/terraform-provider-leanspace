package analysis_definitions

import (
	"leanspace-terraform-provider/helper"
	"regexp"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

var validFieldTypes = []string{
	"STRUCTURE",
	"ARRAY",
	"NUMERIC",
	"TEXT",
	"BOOLEAN",
	"ENUM",
	"TIMESTAMP",
	"DATE",
	"TIME",
	"TLE",
}
var validFieldSources = []string{
	"REFERENCE",
	"STATIC",
}
var lrnRegex = regexp.MustCompile(`lrn::leanspace::[\w]+::[\w]+::[\w]+::[a-zA-Z0-9-]+\/[\w]+`)

var analysisDefinitionSchema = map[string]*schema.Schema{
	"id": {
		Type:     schema.TypeString,
		Computed: true,
	},
	"name": {
		Type:     schema.TypeString,
		Required: true,
	},
	"description": {
		Type:     schema.TypeString,
		Optional: true,
	},
	"framework": {
		Type:        schema.TypeString,
		Computed:    true,
		Description: "The framework type used for the simulation. Currently only \"ANALYTICAL_NOMINAL_PROPAGATION\" is supported.",
	},
	"model_id": {
		Type:         schema.TypeString,
		Required:     true,
		ForceNew:     true,
		ValidateFunc: validation.IsUUID,
		Description:  "The UUID of the model to run the simulation with.",
	},
	"node_id": {
		Type:         schema.TypeString,
		Required:     true,
		ForceNew:     true,
		ValidateFunc: validation.IsUUID,
		Description:  "The UUID of the node to run the simulation for.",
	},
	"statistics": {
		Type:        schema.TypeList,
		Computed:    true,
		Description: "Statistics about the simulation.",
		Elem: &schema.Resource{
			Schema: statisticsSchema,
		},
	},
	"inputs": {
		Type:        schema.TypeList,
		Required:    true,
		Description: "Inputs used to configure the simulation.",
		Elem: &schema.Resource{
			Schema: analysisFieldSchema(5, false),
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

var statisticsSchema = map[string]*schema.Schema{
	"number_of_executions": {
		Type:        schema.TypeInt,
		Computed:    true,
		Description: "How many times the simulation was executed.",
	},
	"last_executed_at": {
		Type:        schema.TypeString,
		Computed:    true,
		Description: "Last execution time in ISO 8601 (may be null).",
	},
}

func analysisFieldSchema(depth int, nameRequired bool) map[string]*schema.Schema {
	var validTypes = validFieldTypes
	if depth == 0 {
		// STRUCTURE is the first element - we skip it.
		validTypes = validTypes[1:]
	}
	var baseSchema = map[string]*schema.Schema{
		"type": {
			Type:         schema.TypeString,
			Required:     true,
			ValidateFunc: validation.StringInSlice(validTypes, false),
			Description:  helper.AllowedValuesToDescription(validTypes),
		},
		"source": {
			Type:         schema.TypeString,
			Optional:     true,
			ValidateFunc: validation.StringInSlice(validFieldSources, false),
			Description:  "Must be set if type isn't STRUCTURE or ARRAY, " + helper.AllowedValuesToDescription(validFieldSources),
		},
		"ref": {
			Type:         schema.TypeString,
			Optional:     true,
			ValidateFunc: validation.StringMatch(lrnRegex, "Must be a valid LRN"),
			Description:  "Must be set if source is \"REFERENCE\". An LRN, formatted as \"lrn::leanspace::<tenant>::<service>::<resource>::<resourceId>/<attribute>\".",
		},
		"value": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Must be set if source is \"STATIC\". Must match with type. For complex types, requires a JSON-encoding.",
		},
	}

	if nameRequired {
		baseSchema["name"] = &schema.Schema{
			Type:        schema.TypeString,
			Required:    true,
			Description: "The name of the field - must be set if field is child of structure field.",
		}
	}

	if depth > 0 {
		baseSchema["fields"] = &schema.Schema{
			Type:        schema.TypeSet,
			Optional:    true,
			Description: "Must be set if type is STRUCTURE. A map of fields.",
			Elem: &schema.Resource{
				Schema: analysisFieldSchema(depth-1, true),
			},
		}
		baseSchema["items"] = &schema.Schema{
			Type:        schema.TypeList,
			Optional:    true,
			Description: "Must be set if type is ARRAY. A list of fields.",
			Elem: &schema.Resource{
				Schema: analysisFieldSchema(depth-1, false),
			},
		}
	}

	return baseSchema
}

var dataSourceFilterSchema = map[string]*schema.Schema{
	"model_ids": {
		Type:     schema.TypeList,
		Optional: true,
		Elem: &schema.Schema{
			Type:         schema.TypeString,
			ValidateFunc: validation.IsUUID,
		},
	},
	"node_ids": {
		Type:     schema.TypeList,
		Optional: true,
		Elem: &schema.Schema{
			Type:         schema.TypeString,
			ValidateFunc: validation.IsUUID,
		},
	},
	"frameworks": {
		Type:     schema.TypeList,
		Optional: true,
		Elem: &schema.Schema{
			Type: schema.TypeString,
		},
	},
}
