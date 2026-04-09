package nodes

import (
	"regexp"

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

var nodeSchema = makeNodeSchema(nil)            // no sub nodes
var rootNodeSchema = makeNodeSchema(nodeSchema) // max depth 1

var validNodeTypes = []string{"ASSET", "GROUP", "COMPONENT"}
var validNodeKinds = []string{"GENERIC", "SATELLITE", "GROUND_STATION"}

var tle1stLineRegex = regexp.MustCompile(`^1 (?P<noradId>[ 0-9]{5})[A-Z] [ 0-9]{5}[ A-Z]{3} [ 0-9]{5}[.][ 0-9]{8} (?:(?:[ 0+-][.][ 0-9]{8})|(?: [ +-][.][ 0-9]{7})) [ +-][ 0-9]{5}[+-][ 0-9] [ +-][ 0-9]{5}[+-][ 0-9] [ 0-9] [ 0-9]{4}[ 0-9]$`)
var tle2ndLineRegex = regexp.MustCompile(`^2 (?P<noradId>[ 0-9]{5}) [ 0-9]{3}[.][ 0-9]{4} [ 0-9]{3}[.][ 0-9]{4} [ 0-9]{7} [ 0-9]{3}[.][ 0-9]{4} [ 0-9]{3}[.][ 0-9]{4} [ 0-9]{2}[.][ 0-9]{13}[ 0-9]$`)

func makeNodeSchema(recursiveNodes map[string]resourceschema.Attribute) map[string]resourceschema.Attribute {
	baseSchema := general_objects.ResourceSchemaWith(map[string]resourceschema.Attribute{
		"id": resourceschema.StringAttribute{ // redefine id here to add the plan modifier
			Computed:      true,
			PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
		},
		"name": resourceschema.StringAttribute{
			Required: true,
		},
		"description": resourceschema.StringAttribute{
			Optional: true,
		},
		"parent_node_id": resourceschema.StringAttribute{
			Optional:   true,
			Validators: helper.ValidUUID(),
		},
		"type": resourceschema.StringAttribute{
			Required:      true,
			Description:   helper.AllowedValuesToDescription(validNodeTypes),
			Validators:    []validator.String{stringvalidator.OneOf(validNodeTypes...)},
			PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
		},
		"kind": resourceschema.StringAttribute{
			Optional:      true,
			Description:   helper.AllowedValuesToDescription(validNodeKinds),
			Validators:    []validator.String{stringvalidator.OneOf(validNodeKinds...)},
			PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
		},
		"tags": general_objects.KeyValuesSchema,
		"number_of_children": resourceschema.Int64Attribute{
			Computed: true,
		},
		// The following fields were part of V1 properties in the API.
		// In terraform, an update occurs when using `terraform apply` multiple times on the same resource with different field values.
		// When these fields are deleted in the API, we suggest to follow these steps during node updates :
		// 1- Do not change this schema so that the user is not impacted by this deprecation
		// 2- Update the built-in properties :
		// 		- Call the endpoint https://api.develop.leanspace.io/asset-repository/properties/v2 to retrieve all the built-in properties.
		//		- For each built-in property, call the endpoint https://api.develop.leanspace.io/asset-repository/properties/v2/{propertyId} to update the property
		//		Hint: you can create a request.go file with a PostUpdateProcess function
		"norad_id": resourceschema.StringAttribute{
			Optional:    true,
			Description: "It must be 5 digits.",
			Validators:  []validator.String{stringvalidator.RegexMatches(regexp.MustCompile(`^\d{5}$`), "It must be 5 digits")},
		},
		"international_designator": resourceschema.StringAttribute{
			Optional:   true,
			Validators: []validator.String{stringvalidator.RegexMatches(regexp.MustCompile(`^(\d{4}-|\d{2})[0-9]{3}[A-Za-z]{0,3}$`), "")},
		},
		"tle": resourceschema.ListAttribute{
			ElementType: types.StringType,
			Optional:    true,
			Description: "TLE composed of its 2 lines.",
			Validators:  []validator.List{listvalidator.SizeBetween(2, 2)},
		},
		"latitude": resourceschema.Float64Attribute{
			Optional:    true,
			Description: "Only for ground stations",
		},
		"longitude": resourceschema.Float64Attribute{
			Optional:    true,
			Description: "Only for ground stations",
		},
		"elevation": resourceschema.Float64Attribute{
			Optional:    true,
			Description: "Only for ground stations",
		},
	})

	if recursiveNodes != nil {
		baseSchema["nodes"] = resourceschema.SetNestedAttribute{
			Computed: true,
			NestedObject: resourceschema.NestedAttributeObject{
				Attributes: recursiveNodes,
			},
		}
	}

	return baseSchema
}

var dataSourceFilterSchema = general_objects.AuditFilterFieldsWithTagsAndSingularBy(map[string]datasourceschema.Attribute{
	"is_root_node": datasourceschema.BoolAttribute{
		Optional:    true,
		Description: "Show only Root Nodes, or hide only Root Nodes. true: select Root Nodes only - false: select Nodes with Parent only",
	},
	"kinds": datasourceschema.ListAttribute{
		ElementType: types.StringType,
		Optional:    true,
	},
	"parent_node_ids": datasourceschema.ListAttribute{
		ElementType: types.StringType,
		Optional:    true,
	},
	"types": datasourceschema.ListAttribute{
		ElementType: types.StringType,
		Optional:    true,
	},
})
