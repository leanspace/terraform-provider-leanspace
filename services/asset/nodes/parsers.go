package nodes

import (
	"encoding/json"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/leanspace/terraform-provider-leanspace/helper"
	"github.com/leanspace/terraform-provider-leanspace/helper/general_objects"
	"github.com/leanspace/terraform-provider-leanspace/provider"
	"github.com/leanspace/terraform-provider-leanspace/services/asset/properties"
	"net/http"
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
	if nodeMap["norad_id"] != nil {
		noradInfo := properties.Property[any]{}
		noradInfo.Attributes.Type = "TEXT"
		noradInfo.Attributes.Value = nodeMap["norad_id"].(string)
		noradInfo.Name = NORAD_ID
		propertylist = append(propertylist, noradInfo)
	}
	if nodeMap["international_designator"] != nil {
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

func (node *Node) PostReadProcess(client *provider.Client, destNodeRaw any) error {
	createdNode := destNodeRaw.(*Node)
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/asset-repository/properties/v2?category=BUILT_IN_PROPERTIES_ONLY&nodeIds=%s", client.HostURL, createdNode.ID), nil)
	if err != nil {
		return err
	}
	body, err, _ := client.DoRequest(req, &(client).Token)
	if err != nil {
		return err
	}

	dataMap := make(map[string]any)
	err = json.Unmarshal(body, &dataMap)
	if err != nil {
		return err
	}
	builtInProperties := dataMap["content"].([]any)
	for _, property := range builtInProperties {
		if property.(map[string]any)["name"] == NORAD_ID {
			attributeProperites := property.(map[string]any)["attributes"].(map[string]any)
			if attributeProperites["value"] != nil {
				createdNode.NoradId = attributeProperites["value"].(string)
			}
		}
		if property.(map[string]any)["name"] == "TLE" {
			attributeProperites := property.(map[string]any)["attributes"].(map[string]any)
			var strList []string
			for _, v := range attributeProperites["value"].([]interface{}) {
				str, ok := v.(string)
				if !ok {
					fmt.Println("Failed to convert interface{} to string")
					return nil
				}
				strList = append(strList, str)
			}
			createdNode.Tle = strList

		}
		if property.(map[string]any)["name"] == INTERNATIONAL_DESIGNATOR {
			attributeProperites := property.(map[string]any)["attributes"].(map[string]any)
			if attributeProperites["value"] != nil {
				createdNode.InternationalDesignator = attributeProperites["value"].(string)
			}
		}
		if property.(map[string]any)["name"] == LOCATION_COORDINATES {
			attributeProperites := property.(map[string]any)["attributes"].(map[string]any)
			field := attributeProperites["fields"].(map[string]any)
			createdNode.Latitude = field["latitude"].(map[string]any)["value"].(float64)
			createdNode.Longitude = field["longitude"].(map[string]any)["value"].(float64)
			createdNode.Elevation = field["elevation"].(map[string]any)["value"].(float64)
		}
	}
	return nil

}
