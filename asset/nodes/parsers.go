package nodes

import (
	"fmt"
	"regexp"

	"terraform-provider-asset/asset/general_objects"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func nodeStructToInterfaceBase(node Node) map[string]any {
	return nodeStructToInterface(&node, 0)
}

func nodeStructToInterface(node *Node, level int) map[string]any {
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
	nodeMap["tags"] = general_objects.TagsStructToMap(node.Tags)
	if node.Nodes != nil && level == 0 {
		nodes := make([]any, len(node.Nodes))
		for i, node := range node.Nodes {
			nodes[i] = nodeStructToInterface(&node, level+1)
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

	return nodeMap
}

var tle1stLine = `^1 (?P<noradId>[ 0-9]{5})[A-Z] [ 0-9]{5}[ A-Z]{3} [ 0-9]{5}[.][ 0-9]{8} (?:(?:[ 0+-][.][ 0-9]{8})|(?: [ +-][.][ 0-9]{7})) [ +-][ 0-9]{5}[+-][ 0-9] [ +-][ 0-9]{5}[+-][ 0-9] [ 0-9] [ 0-9]{4}[ 0-9]$`
var tle2ndLine = `^2 (?P<noradId>[ 0-9]{5}) [ 0-9]{3}[.][ 0-9]{4} [ 0-9]{3}[.][ 0-9]{4} [ 0-9]{7} [ 0-9]{3}[.][ 0-9]{4} [ 0-9]{3}[.][ 0-9]{4} [ 0-9]{2}[.][ 0-9]{13}[ 0-9]$`

func nodeInterfaceToStruct(node map[string]any) (Node, error) {
	nodeStruct := Node{}

	nodeStruct.Name = node["name"].(string)
	nodeStruct.Description = node["description"].(string)
	nodeStruct.CreatedAt = node["created_at"].(string)
	nodeStruct.CreatedBy = node["created_by"].(string)
	nodeStruct.ParentNodeId = node["parent_node_id"].(string)
	nodeStruct.LastModifiedAt = node["last_modified_at"].(string)
	nodeStruct.LastModifiedBy = node["last_modified_by"].(string)
	nodeStruct.Type = node["type"].(string)
	if nodeStruct.Type == "ASSET" && !(node["kind"] == "GENERIC" || node["kind"] == "SATELLITE" || node["kind"] == "GROUND_STATION") {
		return nodeStruct, fmt.Errorf("kind must be either GENERIC, SATELLITE ou GROUND_STATION, got: %q", node["kind"])
	}
	nodeStruct.Kind = node["kind"].(string)
	nodeStruct.Tags = general_objects.TagsInterfaceToStruct(node["tags"])
	if node["nodes"] != nil {
		nodeStruct.Nodes = make([]Node, node["nodes"].(*schema.Set).Len())
		for i, node := range node["nodes"].(*schema.Set).List() {
			childNodeStruct, err := nodeInterfaceToStruct(node.(map[string]any))
			if err != nil {
				return nodeStruct, err
			}
			nodeStruct.Nodes[i] = childNodeStruct
		}
	}
	nodeStruct.NoradId = node["norad_id"].(string)
	nodeStruct.InternationalDesignator = node["international_designator"].(string)
	if node["tle"] != nil && len(node["tle"].([]any)) == 2 {
		nodeStruct.Tle = make([]string, 2)
		matched, _ := regexp.MatchString(tle1stLine, node["tle"].([]any)[0].(string))
		if !matched {
			return nodeStruct, fmt.Errorf("TLE first line mutch match %q, got: %q", tle1stLine, node["tle"].([]any)[0].(string))
		}
		matched, _ = regexp.MatchString(tle2ndLine, node["tle"].([]any)[1].(string))
		if !matched {
			return nodeStruct, fmt.Errorf("TLE second line mutch match %q, got: %q", tle2ndLine, node["tle"].([]any)[1].(string))
		}
		for i, tle := range node["tle"].([]any) {
			nodeStruct.Tle[i] = tle.(string)
		}

	}

	return nodeStruct, nil
}
