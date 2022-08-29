package nodes

import (
	"github.com/leanspace/terraform-provider-leanspace/helper"
	"github.com/leanspace/terraform-provider-leanspace/helper/general_objects"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func (node *Node) ToMap() map[string]any {
	return node.toMapRecursive(0)
}

func (node *Node) toMapRecursive(level int) map[string]any {
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
	nodeMap["tags"] = helper.ParseToMaps(node.Tags)
	if node.Nodes != nil && level == 0 {
		nodes := make([]any, len(node.Nodes))
		for i, subNode := range node.Nodes {
			nodes[i] = (&subNode).toMapRecursive(level + 1)
		}
		nodeMap["nodes"] = nodes
	}
	if len(node.NoradId) != 0 {
		nodeMap["norad_id"] = node.NoradId
	}
	if len(node.InternationalDesignator) != 0 {
		nodeMap["international_designator"] = node.InternationalDesignator
	}
	if len(node.Tle) == 2 {
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
	if tags, err := helper.ParseFromMaps[general_objects.Tag](nodeMap["tags"].(*schema.Set).List()); err != nil {
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
	node.NoradId = nodeMap["norad_id"].(string)
	node.InternationalDesignator = nodeMap["international_designator"].(string)
	if nodeMap["tle"] != nil {
		node.Tle = make([]string, len(nodeMap["tle"].([]any)))
		for i, tle := range nodeMap["tle"].([]any) {
			node.Tle[i] = tle.(string)
		}

	}
	if nodeMap["kind"] == "GROUND_STATION" {
		node.Latitude = helper.Ptr(nodeMap["latitude"].(float64))
		node.Longitude = helper.Ptr(nodeMap["longitude"].(float64))
		node.Elevation = helper.Ptr(nodeMap["elevation"].(float64))
	}

	return nil
}
