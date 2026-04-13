package nodes

import (
	"github.com/leanspace/terraform-provider-leanspace/helper"
	"github.com/leanspace/terraform-provider-leanspace/helper/general_objects"
	"github.com/leanspace/terraform-provider-leanspace/services/asset/properties"
)

var NORAD_ID = "NORAD ID"
var INTERNATIONAL_DESIGNATOR = "International Designator"
var LOCATION_COORDINATES = "Location Coordinates"

func (node *Node) ToMap() map[string]any {
	nodeMap := node.ToAuditMap()

	nodeMap["name"] = helper.NilIfEmpty(node.Name)
	nodeMap["description"] = helper.NilIfEmpty(node.Description)
	nodeMap["parent_node_id"] = helper.NilIfEmpty(node.ParentNodeId)
	nodeMap["type"] = helper.NilIfEmpty(node.Type)
	nodeMap["kind"] = helper.NilIfEmpty(node.Kind)
	nodeMap["number_of_children"] = helper.NilIfEmpty(node.NumberOfChildren)
	nodeMap["tags"] = helper.ParseToMaps(node.Tags)

	if node.Kind == "SATELLITE" {
		nodeMap["norad_id"] = helper.NilIfEmpty(node.NoradId)
		nodeMap["international_designator"] = helper.NilIfEmpty(node.InternationalDesignator)
		nodeMap["tle"] = helper.NilIfEmpty(node.Tle)
	}
	if node.Kind == "GROUND_STATION" {
		nodeMap["latitude"] = helper.NilIfEmpty(node.Latitude)
		nodeMap["longitude"] = helper.NilIfEmpty(node.Longitude)
		nodeMap["elevation"] = helper.NilIfEmpty(node.Elevation)
	}

	return nodeMap
}

func (node *Node) FromMap(nodeMap map[string]any) error {
	node.FromAuditMap(nodeMap)
	node.Name = helper.CastString(nodeMap, "name")
	node.Description = helper.CastString(nodeMap, "description")
	node.ParentNodeId = helper.CastString(nodeMap, "parent_node_id")
	node.Type = helper.CastString(nodeMap, "type")
	node.Kind = helper.CastString(nodeMap, "kind")
	node.NumberOfChildren = helper.CastInt(nodeMap, "number_of_children")
	if tags, err := helper.ParseFromMaps[general_objects.KeyValue](helper.CastSlice(nodeMap, "tags")); err != nil {
		return err
	} else {
		node.Tags = tags
	}
	if nodeMap["nodes"] != nil {
		node.Nodes = make([]Node, len(helper.CastSlice(nodeMap, "nodes")))
		for i, subNode := range helper.CastSlice(nodeMap, "nodes") {
			err := node.Nodes[i].FromMap(subNode.(map[string]any))
			if err != nil {
				return err
			}
		}
	}
	var propertylist []properties.Property[any]
	if helper.CastString(nodeMap, "norad_id") != "" {
		noradInfo := properties.Property[any]{}
		noradInfo.Attributes.Type = "TEXT"
		noradInfo.Attributes.Value = helper.CastString(nodeMap, "norad_id")
		noradInfo.Name = NORAD_ID
		propertylist = append(propertylist, noradInfo)
	}
	if helper.CastString(nodeMap, "international_designator") != "" {
		internationalDesignatorInfo := properties.Property[any]{}
		internationalDesignatorInfo.Attributes.Type = "TEXT"
		internationalDesignatorInfo.Attributes.Value = helper.CastString(nodeMap, "international_designator")
		internationalDesignatorInfo.Name = INTERNATIONAL_DESIGNATOR
		propertylist = append(propertylist, internationalDesignatorInfo)
	}
	if nodeMap["tle"] != nil {
		tleInfo := properties.Property[any]{}
		tleInfo.Attributes.Type = "TLE"
		tleInfo.Name = "TLE"
		var stringTleValues = helper.CastSlice(nodeMap, "tle")
		if len(stringTleValues) == 2 {
			var interfaceOfTleValues []interface{}
			for _, str := range stringTleValues {
				interfaceOfTleValues = append(interfaceOfTleValues, str)
			}
			tleInfo.Attributes.Value = interfaceOfTleValues

			propertylist = append(propertylist, tleInfo)
		}
	}

	if nodeMap["kind"] == "GROUND_STATION" {
		groundStationInfo := properties.Property[any]{}
		groundStationInfo.Attributes.Type = "GEOPOINT"
		groundStationInfo.Name = LOCATION_COORDINATES
		groundStationInfo.Attributes.Fields = &general_objects.Fields{}
		groundStationInfo.Attributes.Fields.Latitude.Value = nodeMap["latitude"]
		groundStationInfo.Attributes.Fields.Longitude.Value = nodeMap["longitude"]
		groundStationInfo.Attributes.Fields.Elevation.Value = nodeMap["elevation"]
		propertylist = append(propertylist, groundStationInfo)
	}
	node.PropertyList = propertylist

	return nil
}
