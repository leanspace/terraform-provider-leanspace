package nodes

import (
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/leanspace/terraform-provider-leanspace/helper"
	"github.com/leanspace/terraform-provider-leanspace/helper/general_objects"
	"github.com/leanspace/terraform-provider-leanspace/services/asset/properties"
)

// nodeObjectAttrTypes is the attr.Type map for a node object, matching nodeSchema.
// The nodes computed field uses a set of these objects.
var nodeTagAttrTypes = map[string]attr.Type{
	"key":   types.StringType,
	"value": types.StringType,
}

var nodeObjectAttrTypes = map[string]attr.Type{
	"id":                       types.StringType,
	"created_at":               types.StringType,
	"created_by":               types.StringType,
	"last_modified_at":         types.StringType,
	"last_modified_by":         types.StringType,
	"name":                     types.StringType,
	"description":              types.StringType,
	"parent_node_id":           types.StringType,
	"type":                     types.StringType,
	"kind":                     types.StringType,
	"tags":                     types.SetType{ElemType: types.ObjectType{AttrTypes: nodeTagAttrTypes}},
	"number_of_children":       types.Int64Type,
	"norad_id":                 types.StringType,
	"international_designator": types.StringType,
	"tle":                      types.ListType{ElemType: types.StringType},
	"latitude":                 types.Float64Type,
	"longitude":                types.Float64Type,
	"elevation":                types.Float64Type,
}

type NodeTF struct {
	general_objects.AuditModelTF
	Name                    types.String                 `tfsdk:"name"`
	Description             types.String                 `tfsdk:"description"`
	ParentNodeId            types.String                 `tfsdk:"parent_node_id"`
	Type                    types.String                 `tfsdk:"type"`
	Kind                    types.String                 `tfsdk:"kind"`
	Tags                    []general_objects.KeyValueTF `tfsdk:"tags"`
	NumberOfChildren        types.Int64                  `tfsdk:"number_of_children"`
	NoradId                 types.String                 `tfsdk:"norad_id"`
	InternationalDesignator types.String                 `tfsdk:"international_designator"`
	Tle                     []types.String               `tfsdk:"tle"`
	Latitude                types.Float64                `tfsdk:"latitude"`
	Longitude               types.Float64                `tfsdk:"longitude"`
	Elevation               types.Float64                `tfsdk:"elevation"`
	// Computed-only: holds nested node objects (max 1 level deep). Uses types.Set
	// to match the SetNestedAttribute schema without needing a concrete element type.
	Nodes types.Set `tfsdk:"nodes"`
}

func (x *Node) ToTF() any {
	tf := &NodeTF{
		AuditModelTF:     general_objects.AuditModelToTF(&x.AuditModel),
		Name:             helper.TFStringValue(x.Name),
		Description:      helper.TFStringPtrValue(x.Description),
		ParentNodeId:     helper.TFStringPtrValue(x.ParentNodeId),
		Type:             helper.TFStringValue(x.Type),
		Kind:             helper.TFStringPtrValue(x.Kind),
		Tags:             general_objects.KeyValuesToTF(x.Tags),
		NumberOfChildren: helper.TFInt64Value(x.NumberOfChildren),
		// nodes is Computed-only; set empty set so the framework does not see a null→value diff.
		Nodes: types.SetValueMust(types.ObjectType{AttrTypes: nodeObjectAttrTypes}, []attr.Value{}),
	}

	if x.Kind != nil && *x.Kind == "SATELLITE" {
		tf.NoradId = helper.TFStringPtrValue(x.NoradId)
		tf.InternationalDesignator = helper.TFStringPtrValue(x.InternationalDesignator)
		tf.Tle = helper.TFStringsValue(x.Tle)
	}

	if x.Kind != nil && *x.Kind == "GROUND_STATION" {
		tf.Latitude = helper.TFFloat64Value(x.Latitude)
		tf.Longitude = helper.TFFloat64Value(x.Longitude)
		tf.Elevation = helper.TFFloat64Value(x.Elevation)
	}

	return tf
}

func (tf *NodeTF) ToAPI() any {
	node := &Node{
		AuditModel:              general_objects.AuditModelFromTF(tf.AuditModelTF),
		Name:                    helper.FromTFString(tf.Name),
		Description:             helper.FromTFStringPtr(tf.Description),
		ParentNodeId:            helper.FromTFStringPtr(tf.ParentNodeId),
		Type:                    helper.FromTFString(tf.Type),
		Kind:                    helper.FromTFStringPtr(tf.Kind),
		Tags:                    general_objects.KeyValuesFromTF(tf.Tags),
		NumberOfChildren:        helper.FromTFInt64(tf.NumberOfChildren),
		NoradId:                 helper.FromTFStringPtr(tf.NoradId),
		InternationalDesignator: helper.FromTFStringPtr(tf.InternationalDesignator),
		Tle:                     helper.FromTFStrings(tf.Tle),
		Latitude:                helper.FromTFFloat64(tf.Latitude),
		Longitude:               helper.FromTFFloat64(tf.Longitude),
		Elevation:               helper.FromTFFloat64(tf.Elevation),
	}

	// Build PropertyList from the flattened fields (same logic as FromMap)
	var propertyList []properties.Property[any]
	if node.NoradId != nil {
		noradInfo := properties.Property[any]{}
		noradInfo.Attributes.Type = "TEXT"
		noradInfo.Attributes.Value = *node.NoradId
		noradInfo.Name = NORAD_ID
		propertyList = append(propertyList, noradInfo)
	}
	if node.InternationalDesignator != nil {
		intlDesig := properties.Property[any]{}
		intlDesig.Attributes.Type = "TEXT"
		intlDesig.Attributes.Value = *node.InternationalDesignator
		intlDesig.Name = INTERNATIONAL_DESIGNATOR
		propertyList = append(propertyList, intlDesig)
	}
	if len(node.Tle) == 2 {
		tleInfo := properties.Property[any]{}
		tleInfo.Attributes.Type = "TLE"
		tleInfo.Name = "TLE"
		var interfaceValues []interface{}
		for _, s := range node.Tle {
			interfaceValues = append(interfaceValues, s)
		}
		tleInfo.Attributes.Value = interfaceValues
		propertyList = append(propertyList, tleInfo)
	}
	if node.Kind != nil && *node.Kind == "GROUND_STATION" {
		gsInfo := properties.Property[any]{}
		gsInfo.Attributes.Type = "GEOPOINT"
		gsInfo.Name = LOCATION_COORDINATES
		gsInfo.Attributes.Fields = &general_objects.Fields{}
		gsInfo.Attributes.Fields.Latitude.Value = node.Latitude
		gsInfo.Attributes.Fields.Longitude.Value = node.Longitude
		gsInfo.Attributes.Fields.Elevation.Value = node.Elevation
		propertyList = append(propertyList, gsInfo)
	}
	node.PropertyList = propertyList

	return node
}
