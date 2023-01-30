package nodes

import (
	"regexp"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"

	"github.com/leanspace/terraform-provider-leanspace/helper"
	"github.com/leanspace/terraform-provider-leanspace/helper/general_objects"
	"github.com/leanspace/terraform-provider-leanspace/services/asset/properties"
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
		"number_of_children": {
			Type:        schema.TypeInt,
			Computed:    true,
			Description: "Numeric only",
		},
		"properties": {
			Type:        schema.TypeSet,
			Optional:    true,
			Description: "Node properties",
			Elem: &schema.Resource{
				Schema: properties.PropertySchema,
			},
		},
		// The following fields are part of V1 properties in the API that have been marked as deprecated for node updates.
		// In terraform, an update occurs when using `terraform apply` multiple times on the same resource with different field values.
		// When these fields are deleted in the API, we suggest to follow these steps :
		// 1- Do not change this schema so that the user is not impacted by this deprecation
		// 2- Update the built-in properties :
		// 		- Call the endpoint https://api.develop.leanspace.io/asset-repository/properties/v2 to retrieve all the built-in properties.
		//		- For each built-in property, call the endpoint https://api.develop.leanspace.io/asset-repository/properties/v2/{propertyId} to update the property
		//		Hint: you can create a request.go file with a PostUpdateProcess function
		"norad_id": {
			Type:         schema.TypeString,
			Optional:     true,
			ValidateFunc: validation.StringMatch(regexp.MustCompile(`^\d{5}$`), "It must be 5 digits"),
			Description:  "It must be 5 digits.",
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
			Description: "TLE composed of its 2 lines.",
		},
		"latitude": {
			Type:        schema.TypeFloat,
			Optional:    true,
			Description: "Only for ground stations",
		},
		"longitude": {
			Type:        schema.TypeFloat,
			Optional:    true,
			Description: "Only for ground stations",
		},
		"elevation": {
			Type:        schema.TypeFloat,
			Optional:    true,
			Description: "Only for ground stations",
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
