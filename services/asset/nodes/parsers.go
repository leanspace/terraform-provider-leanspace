package nodes

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/leanspace/terraform-provider-leanspace/helper"
	"github.com/leanspace/terraform-provider-leanspace/helper/general_objects"
	"github.com/leanspace/terraform-provider-leanspace/services/asset/properties"
)

var NORAD_ID = "NORAD ID"
var INTERNATIONAL_DESIGNATOR = "International Designator"
var LOCATION_COORDINATES = "Location Coordinates"

func (node *Node) ToMap() map[string]any {
	nodeMap := make(map[string]any)

	nodeMap["id"] = node.ID
	nodeMap["name"] = node.Name
	nodeMap["description"] = node.Description
	nodeMap["created_at"] = node.CreatedAt
	nodeMap["created_by"] = node.CreatedBy
	nodeMap["parent_node_id"] = node.ParentNodeId
	nodeMap["last_modified_at"] = node.LastModifiedAt
	nodeMap["last_modified_by"] = node.LastModifiedBy
	nodeMap["type"] = node.Type
	nodeMap["kind"] = node.Kind
	nodeMap["number_of_children"] = node.NumberOfChildren
	nodeMap["tags"] = helper.ParseToMaps(node.Tags)

	if node.Kind == "SATELLITE" {
		nodeMap["norad_id"] = node.NoradId
		nodeMap["international_designator"] = node.InternationalDesignator
		nodeMap["tle"] = node.Tle
	}
	if node.Kind == "GROUND_STATION" {
		nodeMap["latitude"] = node.Latitude
		nodeMap["longitude"] = node.Longitude
		nodeMap["elevation"] = node.Elevation
	}

	return nodeMap
}

func (node *Node) FromMap(nodeMap map[string]any) error {
	node.Name = nodeMap["name"].(string)
	node.Description = nodeMap["description"].(string)
	node.CreatedAt = nodeMap["created_at"].(string)
	node.CreatedBy = nodeMap["created_by"].(string)
	node.ParentNodeId = nodeMap["parent_node_id"].(string)
	node.LastModifiedAt = nodeMap["last_modified_at"].(string)
	node.LastModifiedBy = nodeMap["last_modified_by"].(string)
	node.Type = nodeMap["type"].(string)
	node.Kind = nodeMap["kind"].(string)
	node.NumberOfChildren = nodeMap["number_of_children"].(int)
	if tags, err := helper.ParseFromMaps[general_objects.KeyValue](nodeMap["tags"].(*schema.Set).List()); err != nil {
		return err
	} else {
		node.Tags = tags
	}
	if nodeMap["nodes"] != nil {
		node.Nodes = make([]Node, nodeMap["nodes"].(*schema.Set).Len())
		for i, subNode := range nodeMap["nodes"].(*schema.Set).List() {
			err := node.Nodes[i].FromMap(subNode.(map[string]any))
			if err != nil {
				return err
			}
		}
	}
	var propertylist []properties.Property[any]
	if nodeMap["norad_id"] != "" {
		noradInfo := properties.Property[any]{}
		noradInfo.Attributes.Type = "TEXT"
		noradInfo.Attributes.Value = nodeMap["norad_id"].(string)
		noradInfo.Name = NORAD_ID
		propertylist = append(propertylist, noradInfo)
	}
	if nodeMap["international_designator"] != "" {
		internationalDesignatorInfo := properties.Property[any]{}
		internationalDesignatorInfo.Attributes.Type = "TEXT"
		internationalDesignatorInfo.Attributes.Value = nodeMap["international_designator"].(string)
		internationalDesignatorInfo.Name = INTERNATIONAL_DESIGNATOR
		propertylist = append(propertylist, internationalDesignatorInfo)
	}
	if nodeMap["tle"] != nil {
		tleInfo := properties.Property[any]{}
		tleInfo.Attributes.Type = "TLE"
		tleInfo.Name = "TLE"
		var stringTleValues = nodeMap["tle"].([]interface{})
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
		groundStationInfo.Attributes.Fields = &properties.Fields{}
		groundStationInfo.Attributes.Fields.Latitude.Value = nodeMap["latitude"]
		groundStationInfo.Attributes.Fields.Longitude.Value = nodeMap["longitude"]
		groundStationInfo.Attributes.Fields.Elevation.Value = nodeMap["elevation"]
		propertylist = append(propertylist, groundStationInfo)
	}
	node.PropertyList = propertylist

	return nil
}
