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

var tle1stLineRegex = regexp.MustCompile(`^1 (?P<noradId>[ 0-9]{5})[A-Z] [ 0-9]{5}[ A-Z]{3} [ 0-9]{5}[.][ 0-9]{8} (?:(?:[ 0+-][.][ 0-9]{8})|(?: [ +-][.][ 0-9]{7})) [ +-][ 0-9]{5}[+-][ 0-9] [ +-][ 0-9]{5}[+-][ 0-9] [ 0-9] [ 0-9]{4}[ 0-9]$`)
var tle2ndLineRegex = regexp.MustCompile(`^2 (?P<noradId>[ 0-9]{5}) [ 0-9]{3}[.][ 0-9]{4} [ 0-9]{3}[.][ 0-9]{4} [ 0-9]{7} [ 0-9]{3}[.][ 0-9]{4} [ 0-9]{3}[.][ 0-9]{4} [ 0-9]{2}[.][ 0-9]{13}[ 0-9]$`)

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
		"tags": general_objects.KeyValuesSchema,
		"number_of_children": {
			Type:        schema.TypeInt,
			Computed:    true,
			Description: "Numeric only",
		},
		// The following fields are part of V1 properties in the API that have been marked as deprecated for node updates.
		// In terraform, an update occurs when using `terraform apply` multiple times on the same resource with different field values.
		// When these fields are deleted in the API, we suggest to follow these steps during node updates :
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
	"created_by": {
		Type:         schema.TypeString,
		Optional:     true,
		ValidateFunc: validation.IsUUID,
		Description:  "Filter on the user who created the Node. If you have no wish to use this field as a filter, either provide a null value or remove the field.",
	},
	"from_created_at": {
		Type:         schema.TypeString,
		Optional:     true,
		ValidateFunc: helper.IsValidTimeDateOrTimestamp,
		Description:  "Filter on the Node creation date. Properties with a creation date greater or equals than the filter value will be selected (if they are not excluded by other filters). If you have no wish to use this field as a filter, either provide a null value or remove the field.",
	},
	"from_last_modified_at": {
		Type:         schema.TypeString,
		Optional:     true,
		ValidateFunc: helper.IsValidTimeDateOrTimestamp,
		Description:  "Filter on the Node last modification date. Nodes with a last modification date greater or equals than the filter value will be selected (if they are not excluded by other filters). If you have no wish to use this field as a filter, either provide a null value or remove the field.",
	},
	"last_modified_by": {
		Type:         schema.TypeString,
		Optional:     true,
		ValidateFunc: validation.IsUUID,
		Description:  "Filter on the user who modified last the Node. If you have no wish to use this field as a filter, either provide a null value or remove the field.",
	},
	"to_created_at": {
		Type:         schema.TypeString,
		Optional:     true,
		ValidateFunc: helper.IsValidTimeDateOrTimestamp,
		Description:  "Filter on the Node creation date. Nodes with a creation date lower or equals than the filter value will be selected (if they are not excluded by other filters). If you have no wish to use this field as a filter, either provide a null value or remove the field.",
	},
	"to_last_modified_at": {
		Type:         schema.TypeString,
		Optional:     true,
		ValidateFunc: helper.IsValidTimeDateOrTimestamp,
		Description:  "Filter on the Node last modification date. Nodes with a last modification date lower or equals than the filter value will be selected (if they are not excluded by other filters). If you have no wish to use this field as a filter, either provide a null value or remove the field.",
	},
	"ids": {
		Type:     schema.TypeList,
		Optional: true,
		Elem: &schema.Schema{
			Type:         schema.TypeString,
			ValidateFunc: validation.IsUUID,
		},
		Description: "Only returns node whose id matches one of the provided values.",
	},
	"is_root_node": {
		Type:        schema.TypeBool,
		Optional:    true,
		Description: "Show only Root Nodes, or hide only Root Nodes. true: select Root Nodes only - false: select Nodes with Parent only",
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
	"parent_node_ids": {
		Type:     schema.TypeList,
		Optional: true,
		Elem: &schema.Schema{
			Type:         schema.TypeString,
			ValidateFunc: validation.IsUUID,
		},
	},
	"query": {
		Type:        schema.TypeString,
		Optional:    true,
		Description: "Search by name or description",
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
	"tags": {
		Type:     schema.TypeList,
		Optional: true,
		Elem: &schema.Schema{
			Type: schema.TypeString,
		},
	},
}
