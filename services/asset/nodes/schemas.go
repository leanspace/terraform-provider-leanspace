package nodes

import (
	"regexp"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"

	"github.com/leanspace/terraform-provider-leanspace/helper"
	"github.com/leanspace/terraform-provider-leanspace/helper/general_objects"
)

var nodeSchema = makeNodeSchema(nil)            // no sub nodes
var rootNodeSchema = makeNodeSchema(nodeSchema) // max depth 1

var validNodeTypes = []string{"ASSET", "GROUP", "COMPONENT"}
var validNodeKinds = []string{"GENERIC", "SATELLITE", "GROUND_STATION"}

func makeNodeSchema(recursiveNodes map[string]*schema.Schema) map[string]*schema.Schema {
	baseSchema := map[string]*schema.Schema{
		"id": {
			Type:     schema.TypeString,
			Computed: true,
			ForceNew: true,
		},
		"name": {
			Type:     schema.TypeString,
			Required: true,
		},
		"description": {
			Type:     schema.TypeString,
			Optional: true,
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
		"parent_node_id": {
			Type:         schema.TypeString,
			Optional:     true,
			ValidateFunc: validation.IsUUID,
		},
		"type": {
			Type:         schema.TypeString,
			Required:     true,
			ForceNew:     true,
			ValidateFunc: validation.StringInSlice(validNodeTypes, false),
			Description:  helper.AllowedValuesToDescription(validNodeTypes),
		},
		"kind": {
			Type:         schema.TypeString,
			Optional:     true,
			ForceNew:     true,
			ValidateFunc: validation.StringInSlice(validNodeKinds, false),
			Description:  helper.AllowedValuesToDescription(validNodeKinds),
		},
		"tags": general_objects.TagsSchema,
		"norad_id": {
			Type:         schema.TypeString,
			Optional:     true,
			ValidateFunc: validation.StringMatch(regexp.MustCompile(`^\d{5}$`), "It must be 5 digits"),
			Description:  "It must be 5 digits. **NOTE**: The attribute `norad_id` is deprecated and will be removed in a future version. Consider using the node `properties` instead.", // description visible in the documentation
			Deprecated:   "The attribute norad_id is deprecated and will be removed in a future version. Consider using the node properties instead.",                                    // deprecation warning message on the terraform console
		},
		"international_designator": {
			Type:         schema.TypeString,
			Optional:     true,
			ValidateFunc: validation.StringMatch(regexp.MustCompile(`^(\d{4}-|\d{2})[0-9]{3}[A-Za-z]{0,3}$`), ""),
			Description:  "**NOTE**: The attribute `international_designator` is deprecated and will be removed in a future version. Consider using the node `properties` instead.", // description visible in the documentation
			Deprecated:   "The attribute international_designator is deprecated and will be removed in a future version. Consider using the node properties instead.",               // deprecation warning message on the terraform console
		},
		"tle": {
			Type:     schema.TypeList,
			MaxItems: 2,
			MinItems: 2,
			Optional: true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
			Description: "TLE composed of its 2 lines. **NOTE**: The attribute `tle` is deprecated and will be removed in a future version. Consider using the node `properties` instead.", // description visible in the documentation
			Deprecated:  "The attribute tle is deprecated and will be removed in a future version. Consider using the node properties instead.",                                            // deprecation warning message on the terraform console
		},
		"number_of_children": {
			Type:        schema.TypeInt,
			Computed:    true,
			Description: "Numeric only",
		},
		// These fields are *required* when kind = GROUND_STATION
		// However currently I don't think there is a way to have conditionally required fields
		// A solution would be creating a new "ground station" schema, if we want to ensure type safety
		// For now this does the job!
		"latitude": {
			Type:        schema.TypeFloat,
			Optional:    true,
			Description: "**NOTE**: The attribute `latitude` is deprecated and will be removed in a future version. Consider using the node `properties` instead.", // description visible in the documentation
			Deprecated:  "The attribute latitude is deprecated and will be removed in a future version. Consider using the node properties instead.",               // deprecation warning message on the terraform console
		},
		"longitude": {
			Type:        schema.TypeFloat,
			Optional:    true,
			Description: "**NOTE**: The attribute `longitude` is deprecated and will be removed in a future version. Consider using the node `properties` instead.", // description visible in the documentation
			Deprecated:  "The attribute longitude is deprecated and will be removed in a future version. Consider using the node properties instead.",               // deprecation warning message on the terraform console
		},
		"elevation": {
			Type:        schema.TypeFloat,
			Optional:    true,
			Description: "**NOTE**: The attribute `elevation` is deprecated and will be removed in a future version. Consider using the node `properties` instead.", // description visible in the documentation
			Deprecated:  "The attribute elevation is deprecated and will be removed in a future version. Consider using the node properties instead.",               // deprecation warning message on the terraform console
		},
	}

	if recursiveNodes != nil {
		baseSchema["nodes"] = &schema.Schema{
			Type:     schema.TypeSet,
			Computed: true,
			Elem: &schema.Resource{
				Schema: recursiveNodes,
			},
		}
	}

	return baseSchema
}

var dataSourceFilterSchema = map[string]*schema.Schema{
	"parent_node_ids": {
		Type:     schema.TypeList,
		Optional: true,
		Elem: &schema.Schema{
			Type:         schema.TypeString,
			ValidateFunc: validation.IsUUID,
		},
	},
	"property_ids": {
		Type:     schema.TypeList,
		Optional: true,
		Elem: &schema.Schema{
			Type:         schema.TypeString,
			ValidateFunc: validation.IsUUID,
		},
	},
	"metric_ids": {
		Type:     schema.TypeList,
		Optional: true,
		Elem: &schema.Schema{
			Type:         schema.TypeString,
			ValidateFunc: validation.IsUUID,
		},
	},
	"types": {
		Type:     schema.TypeList,
		Optional: true,
		Elem: &schema.Schema{
			Type:         schema.TypeString,
			ValidateFunc: validation.StringInSlice(validNodeTypes, false),
			Description:  helper.AllowedValuesToDescription(validNodeTypes),
		},
	},
	"kinds": {
		Type:     schema.TypeList,
		Optional: true,
		Elem: &schema.Schema{
			Type:         schema.TypeString,
			ValidateFunc: validation.StringInSlice(validNodeKinds, false),
			Description:  helper.AllowedValuesToDescription(validNodeKinds),
		},
	},
	"tags": {
		Type:     schema.TypeList,
		Optional: true,
		Elem: &schema.Schema{
			Type: schema.TypeString,
		},
	},
}
