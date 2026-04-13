package nodes

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/leanspace/terraform-provider-leanspace/helper"
	"github.com/leanspace/terraform-provider-leanspace/helper/general_objects"
	"github.com/leanspace/terraform-provider-leanspace/services/asset/properties"
)

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
}

func (x *Node) ToTF() any {
	tf := &NodeTF{
		AuditModelTF:     general_objects.AuditModelToTF(&x.AuditModel),
		Name:             helper.TFStringValue(x.Name),
		Description:      helper.TFStringValue(x.Description),
		ParentNodeId:     helper.TFStringValue(x.ParentNodeId),
		Type:             helper.TFStringValue(x.Type),
		Kind:             helper.TFStringValue(x.Kind),
		Tags:             general_objects.KeyValuesToTF(x.Tags),
		NumberOfChildren: helper.TFInt64Value(x.NumberOfChildren),
	}

	if x.Kind == "SATELLITE" {
		tf.NoradId = helper.TFStringValue(x.NoradId)
		tf.InternationalDesignator = helper.TFStringValue(x.InternationalDesignator)
		tf.Tle = helper.TFStringsValue(x.Tle)
	}

	if x.Kind == "GROUND_STATION" {
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
		Description:             helper.FromTFString(tf.Description),
		ParentNodeId:            helper.FromTFString(tf.ParentNodeId),
		Type:                    helper.FromTFString(tf.Type),
		Kind:                    helper.FromTFString(tf.Kind),
		Tags:                    general_objects.KeyValuesFromTF(tf.Tags),
		NumberOfChildren:        helper.FromTFInt64(tf.NumberOfChildren),
		NoradId:                 helper.FromTFString(tf.NoradId),
		InternationalDesignator: helper.FromTFString(tf.InternationalDesignator),
		Tle:                     helper.FromTFStrings(tf.Tle),
		Latitude:                helper.FromTFFloat64(tf.Latitude),
		Longitude:               helper.FromTFFloat64(tf.Longitude),
		Elevation:               helper.FromTFFloat64(tf.Elevation),
	}

	// Build PropertyList from the flattened fields (same logic as FromMap)
	var propertyList []properties.Property[any]
	if node.NoradId != "" {
		noradInfo := properties.Property[any]{}
		noradInfo.Attributes.Type = "TEXT"
		noradInfo.Attributes.Value = node.NoradId
		noradInfo.Name = NORAD_ID
		propertyList = append(propertyList, noradInfo)
	}
	if node.InternationalDesignator != "" {
		intlDesig := properties.Property[any]{}
		intlDesig.Attributes.Type = "TEXT"
		intlDesig.Attributes.Value = node.InternationalDesignator
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
	if node.Kind == "GROUND_STATION" {
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
