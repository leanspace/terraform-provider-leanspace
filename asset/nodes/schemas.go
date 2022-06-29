package nodes

import (
	"regexp"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"

	"terraform-provider-asset/asset/general_objects"
)

var nodeSchema = makeNodeSchema(nil)            // not recursive
var rootNodeSchema = makeNodeSchema(nodeSchema) // recursive!

func makeNodeSchema(recursiveNodes map[string]*schema.Schema) map[string]*schema.Schema {
	baseSchema := map[string]*schema.Schema{
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
			Computed:     true,
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
	}

	if recursiveNodes != nil {
		baseSchema["nodes"] = &schema.Schema{
			Type:     schema.TypeSet,
			Optional: true,
			Elem: &schema.Resource{
				Schema: recursiveNodes,
			},
		}
	}

	return baseSchema
}
