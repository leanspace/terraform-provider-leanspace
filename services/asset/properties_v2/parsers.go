package properties_v2

import (
	"github.com/leanspace/terraform-provider-leanspace/helper"
	"github.com/leanspace/terraform-provider-leanspace/helper/general_objects"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func (property *Property[T]) ToMap() map[string]any {
	propertyMap := make(map[string]any)
	propertyMap["id"] = property.ID
	propertyMap["name"] = property.Name
	propertyMap["description"] = property.Description
	propertyMap["built_in"] = property.IsBuiltIn
	propertyMap["node_id"] = property.NodeId
	propertyMap["created_at"] = property.CreatedAt
	propertyMap["created_by"] = property.CreatedBy
	propertyMap["last_modified_at"] = property.LastModifiedAt
	propertyMap["last_modified_by"] = property.LastModifiedBy
	propertyMap["tags"] = helper.ParseToMaps(property.Tags)
	propertyMap["attributes"] = []any{property.Attributes.ToMap()}
	return propertyMap
}

func (property *Property[T]) FromMap(propertyMap map[string]any) error {
	property.Name = propertyMap["name"].(string)
	property.Description = propertyMap["description"].(string)
	property.NodeId = propertyMap["node_id"].(string)
	property.CreatedAt = propertyMap["created_at"].(string)
	property.CreatedBy = propertyMap["created_by"].(string)
	property.LastModifiedAt = propertyMap["last_modified_at"].(string)
	property.LastModifiedBy = propertyMap["last_modified_by"].(string)
	property.IsBuiltIn = propertyMap["built_in"].(bool)
	if tags, err := helper.ParseFromMaps[general_objects.Tag](propertyMap["tags"].(*schema.Set).List()); err != nil {
		return err
	} else {
		property.Tags = tags
	}
	attributeMap := propertyMap["attributes"].([]any)[0].(map[string]any)
	err := property.Attributes.FromMap(attributeMap)
	return err
}
