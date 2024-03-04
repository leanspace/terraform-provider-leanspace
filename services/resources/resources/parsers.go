package resources

import (
	"github.com/leanspace/terraform-provider-leanspace/helper"
	"github.com/leanspace/terraform-provider-leanspace/helper/general_objects"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func (resource *Resource) ToMap() map[string]any {
	resourceMap := make(map[string]any)
	resourceMap["id"] = resource.ID
	resourceMap["asset_id"] = resource.AssetId
	resourceMap["unit_id"] = resource.UnitId
	resourceMap["metric_id"] = resource.MetricId
	resourceMap["name"] = resource.Name
	resourceMap["description"] = resource.Description
	resourceMap["default_level"] = resource.DefaultLevel
	if resource.Constraints != nil {
		resourceMap["constraints"] = helper.ParseToMaps(resource.Constraints)
	}
	resourceMap["tags"] = helper.ParseToMaps(resource.Tags)
	resourceMap["created_at"] = resource.CreatedAt
	resourceMap["created_by"] = resource.CreatedBy
	resourceMap["last_modified_at"] = resource.LastModifiedAt
	resourceMap["last_modified_by"] = resource.LastModifiedBy

	return resourceMap
}

func (constraints *ResourceConstraints) ToMap() map[string]any {
	constraintsMap := make(map[string]any)
	constraintsMap["type"] = constraints.Type
	constraintsMap["kind"] = constraints.Kind
	constraintsMap["value"] = constraints.Value
	return constraintsMap
}

func (resource *Resource) FromMap(resourceMap map[string]any) error {
	resource.ID = resourceMap["id"].(string)
	resource.AssetId = resourceMap["asset_id"].(string)
	resource.UnitId = resourceMap["unit_id"].(string)
	resource.MetricId = resourceMap["metric_id"].(string)
	resource.Name = resourceMap["name"].(string)
	resource.Description = resourceMap["description"].(string)
	resource.DefaultLevel = resourceMap["default_level"].(float64)
	if resourceMap["constraints"] != nil {
        if constraints, err := helper.ParseFromMaps[ResourceConstraints](
            resourceMap["constraints"].(*schema.Set).List(),
        ); err != nil {
            return err
        } else {
            resource.Constraints = constraints
        }
    }
	if tags, err := helper.ParseFromMaps[general_objects.KeyValue](resourceMap["tags"].(*schema.Set).List()); err != nil {
		return err
	} else {
		resource.Tags = tags
	}
	resource.CreatedAt = resourceMap["created_at"].(string)
	resource.CreatedBy = resourceMap["created_by"].(string)
	resource.LastModifiedAt = resourceMap["last_modified_at"].(string)
	resource.LastModifiedBy = resourceMap["last_modified_by"].(string)

	return nil
}

func (constraints *ResourceConstraints) FromMap(constraintsMap map[string]any) error {
	constraints.Type = constraintsMap["type"].(string)
	constraints.Kind = constraintsMap["kind"].(string)
	constraints.Value = constraintsMap["value"].(float64)
	return nil
}
