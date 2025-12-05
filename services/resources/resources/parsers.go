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

	if resource.UpperLimit != nil {
		resourceMap["upper_limit"] = []any{float64(*resource.UpperLimit)}
	}
	if resource.LowerLimit != nil {
		resourceMap["lower_limit"] = []any{float64(*resource.LowerLimit)}
	}

	if resource.Thresholds != nil {
		resourceMap["thresholds"] = helper.ParseToMaps(resource.Thresholds)
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
	constraintsMap["name"] = constraints.Name
	constraintsMap["value"] = constraints.Value
	return constraintsMap
}

func (thresholds *ResourceThreshold) ToMap() map[string]any {
	thresholdsMap := make(map[string]any)
	thresholdsMap["kind"] = thresholds.Kind
	thresholdsMap["value"] = thresholds.Value
	thresholdsMap["violation_when_reached"] = thresholds.ViolationWhenReached
	thresholdsMap["name"] = thresholds.Name
	return thresholdsMap
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
		constraints, err := helper.ParseFromMaps[ResourceConstraints](resourceMap["constraints"].(*schema.Set).List())
		if err != nil {
			return err
		}
		resource.Constraints = constraints
	}

	if v, ok := resourceMap["lower_limit"]; ok && v != nil {
		if list, ok := v.([]interface{}); ok && len(list) > 0 {
			if floatVal, ok := list[0].(float64); ok {
				resource.LowerLimit = &floatVal
			}
		}
	}
	if v, ok := resourceMap["upper_limit"]; ok && v != nil {
		if list, ok := v.([]interface{}); ok && len(list) > 0 {
			if floatVal, ok := list[0].(float64); ok {
				resource.UpperLimit = &floatVal
			}
		}
	}

	if resourceMap["thresholds"] != nil {
		thresholds, err := helper.ParseFromMaps[ResourceThreshold](resourceMap["thresholds"].(*schema.Set).List())
		if err != nil {
			return err
		}
		resource.Thresholds = thresholds
	}

	tags, err := helper.ParseFromMaps[general_objects.KeyValue](resourceMap["tags"].(*schema.Set).List())
	if err != nil {
		return err
	}
	resource.Tags = tags

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
	constraints.Name = constraintsMap["name"].(string)
	return nil
}

func (thresholds *ResourceThreshold) FromMap(thresholdsMap map[string]any) error {
	thresholds.Kind = thresholdsMap["kind"].(string)
	thresholds.Value = thresholdsMap["value"].(float64)
	thresholds.Name = thresholdsMap["name"].(string)
	thresholds.ViolationWhenReached = thresholdsMap["violation_when_reached"].(bool)
	return nil
}
