package nodes

import (
	"regexp"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"

	"leanspace-terraform-provider/helper/general_objects"
)

var nodeSchema = makeNodeSchema(nil)            // no sub nodes
var rootNodeSchema = makeNodeSchema(nodeSchema) // max depth 1

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
			Type:     schema.TypeString,
			Computed: true,
		},
		"created_by": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"last_modified_at": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"last_modified_by": {
			Type:     schema.TypeString,
			Computed: true,
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
			ValidateFunc: validation.StringInSlice([]string{"ASSET", "GROUP", "COMPONENT"}, false),
		},
		"kind": {
			Type:     schema.TypeString,
			Optional: true,
			ForceNew: true,
		},
		"tags": general_objects.TagsSchema,
		"norad_id": {
			Type:         schema.TypeString,
			Optional:     true,
			ValidateFunc: validation.StringMatch(regexp.MustCompile(`^\d{5}$`), "It must be 5 digits"),
		},
		"international_designator": {
			Type:         schema.TypeString,
			Optional:     true,
			ValidateFunc: validation.StringMatch(regexp.MustCompile(`^(\d{4}-|\d{2})[0-9]{3}[A-Za-z]{0,3}$`), ""),
		},
		"tle": {
			Type:     schema.TypeList,
			MaxItems: 2,
			MinItems: 2,
			Optional: true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
		},
		// These fields are *required* when kind = GROUND_STATION
		// However currently I don't think there is a way to have conditionally required fields
		// A solution would be creating a new "ground station" schema, if we want to ensure type safety
		// For now this does the job!
		"latitude": {
			Type:     schema.TypeFloat,
			Optional: true,
		},
		"longitude": {
			Type:     schema.TypeFloat,
			Optional: true,
		},
		"elevation": {
			Type:     schema.TypeFloat,
			Optional: true,
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
