package nodes

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/leanspace/terraform-provider-leanspace/provider"
)

type apiShiftNodeInfo struct {
	TargetParentNodeId string `json:"targetParentNodeId"`
}

func (node *Node) toAPIFormat() ([]byte, error) {
	shiftNode := apiShiftNodeInfo{
		TargetParentNodeId: node.ParentNodeId,
	}
	return json.Marshal(shiftNode)
}

func nodeAction(action string, nodeId string, node *Node, client *provider.Client) error {
	data, err := node.toAPIFormat()
	if err != nil {
		return err
	}
	path := fmt.Sprintf("%s/%s/%s/shift", client.HostURL, NodeDataType.Path, nodeId)
	req, err := http.NewRequest(action, path, strings.NewReader(string(data)))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	_, err, _ = client.DoRequest(req, &(client).Token)
	if err != nil {
		return err
	}
	return nil
}

func shiftNode(nodeId string, node *Node, client *provider.Client) error {
	return nodeAction("PUT", nodeId, node, client)
}

func currentProperties(client *provider.Client, nodeId string) ([]any, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/asset-repository/properties/v2?category=BUILT_IN_PROPERTIES_ONLY&nodeIds=%s", client.HostURL, nodeId), nil)
	if err != nil {
		return nil, err
	}
	body, err, _ := client.DoRequest(req, &(client).Token)
	if err != nil {
		return nil, err
	}

	dataMap := make(map[string]any)
	err = json.Unmarshal(body, &dataMap)
	if err != nil {
		return nil, err
	}
	builtInProperties := dataMap["content"].([]any)
	return builtInProperties, nil
}

func (node *Node) setPropertiesFromAttributes() (err error) {
	for _, property := range node.PropertyList {
		if property.Name == NORAD_ID {
			if property.Attributes.Value != nil {
				node.NoradId = property.Attributes.Value.(string)
			}
		}
		if property.Name == "TLE" {
			var strList []string
			if property.Attributes.Value != nil {
				for _, v := range property.Attributes.Value.([]interface{}) {
					str, ok := v.(string)
					if !ok {
						fmt.Println("Failed to convert interface{} to string")
						return nil
					}
					strList = append(strList, str)
				}
				node.Tle = strList
			}
		}
		if property.Name == INTERNATIONAL_DESIGNATOR {
			if property.Attributes.Value != nil {
				node.InternationalDesignator = property.Attributes.Value.(string)
			}
		}
		if property.Name == LOCATION_COORDINATES {
			field := property.Attributes.Fields
			if field.Latitude.Value != nil {
				node.Latitude = field.Latitude.Value.(float64)
			}
			if field.Longitude.Value != nil {
				node.Longitude = field.Longitude.Value.(float64)
			}
			if field.Elevation.Value != nil {
				node.Elevation = field.Elevation.Value.(float64)
			}
		}
	}
	return nil
}

func updateProperties(client *provider.Client, builtInProperties []any, node *Node) error {
	for _, property := range builtInProperties {
		if property.(map[string]any)["name"] == NORAD_ID {
			attributeProperites := property.(map[string]any)["attributes"].(map[string]any)
			if node.NoradId != attributeProperites["value"] {
				attributeProperites["value"] = node.NoradId
				if err := updateProperty(client, property.(map[string]any)["id"].(string), property); err != nil {
					return err
				}
			}
		}
		if property.(map[string]any)["name"] == "TLE" {
			attributeProperites := property.(map[string]any)["attributes"].(map[string]any)
			var strList []string
			if attributeProperites["value"] != nil {
				for _, v := range attributeProperites["value"].([]interface{}) {
					str, ok := v.(string)
					if !ok {
						fmt.Println("Failed to convert interface{} to string")
						return nil
					}
					strList = append(strList, str)
				}
			}
			if len(node.Tle) != len(strList) || len(node.Tle) == 2 && len(strList) == 2 && (node.Tle[0] != strList[0] || node.Tle[1] != strList[1]) {
				attributeProperites["value"] = node.Tle
				if err := updateProperty(client, property.(map[string]any)["id"].(string), property); err != nil {
					return err
				}
			}
		}
		if property.(map[string]any)["name"] == INTERNATIONAL_DESIGNATOR {
			attributeProperites := property.(map[string]any)["attributes"].(map[string]any)
			if node.InternationalDesignator != attributeProperites["value"] {
				attributeProperites["value"] = node.InternationalDesignator
				if err := updateProperty(client, property.(map[string]any)["id"].(string), property); err != nil {
					return err
				}
			}
		}
		if property.(map[string]any)["name"] == LOCATION_COORDINATES {
			attributeProperites := property.(map[string]any)["attributes"].(map[string]any)
			field := attributeProperites["fields"].(map[string]any)
			shouldUpdate := false
			if node.Latitude != field["latitude"].(map[string]any)["value"] {
				field["latitude"].(map[string]any)["value"] = node.Latitude
				shouldUpdate = true
			}
			if node.Longitude != field["longitude"].(map[string]any)["value"] {
				field["longitude"].(map[string]any)["value"] = node.Longitude
				shouldUpdate = true
			}
			if node.Elevation != field["elevation"].(map[string]any)["value"] {
				field["elevation"].(map[string]any)["value"] = node.Elevation
				shouldUpdate = true
			}
			if shouldUpdate {
				if err := updateProperty(client, property.(map[string]any)["id"].(string), property); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func updateProperty(client *provider.Client, propertyId string, data any) error {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}
	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/asset-repository/properties/v2/%s", client.HostURL, propertyId), strings.NewReader(string(jsonData)))
	req.Header.Set("Content-Type", "application/json")
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
	return nil
}

func (node *Node) PostUpdateProcess(client *provider.Client, updated any) error {
	updatedNode := updated.(*Node)
	if err := shiftNode(updatedNode.ID, node, client); err != nil {
		return err
	}
	builtInProperties, err := currentProperties(client, updatedNode.ID)
	if err != nil {
		return err
	}
	if err := node.setPropertiesFromAttributes(); err != nil {
		return err
	}
	if err := updateProperties(client, builtInProperties, node); err != nil {
		return err
	}
	return nil
}

func (node *Node) PostReadProcess(client *provider.Client, destNodeRaw any) error {
	createdNode := destNodeRaw.(*Node)
	builtInProperties, err := currentProperties(client, createdNode.ID)
	if err != nil {
		return err
	}
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
			if attributeProperites["value"] != nil {
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
			if field["latitude"].(map[string]any)["value"] != nil {
				createdNode.Latitude = field["latitude"].(map[string]any)["value"].(float64)
			}
			if field["longitude"].(map[string]any)["value"] != nil {
				createdNode.Longitude = field["longitude"].(map[string]any)["value"].(float64)
			}
			if field["elevation"].(map[string]any)["value"] != nil {
				createdNode.Elevation = field["elevation"].(map[string]any)["value"].(float64)
			}
		}
	}

	return nil
}
