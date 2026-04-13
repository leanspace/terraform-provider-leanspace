package resources

import (
	"fmt"

	"github.com/leanspace/terraform-provider-leanspace/helper"
	"github.com/leanspace/terraform-provider-leanspace/helper/general_objects"
	"github.com/leanspace/terraform-provider-leanspace/provider"
)

func (resource *Resource) ToMap() map[string]any {
	resourceMap := resource.ToAuditMap()
	resourceMap["asset_id"] = helper.NilIfEmpty(resource.AssetId)
	resourceMap["unit_id"] = helper.NilIfEmpty(resource.UnitId)
	resourceMap["metric_id"] = helper.NilIfEmpty(resource.MetricId)
	resourceMap["name"] = helper.NilIfEmpty(resource.Name)
	resourceMap["description"] = helper.NilIfEmpty(resource.Description)
	resourceMap["default_level"] = helper.NilIfEmpty(resource.DefaultLevel)

	if resource.UpperLimit != nil {
		resourceMap["upper_limit"] = float64(*resource.UpperLimit)
	}
	if resource.LowerLimit != nil {
		resourceMap["lower_limit"] = float64(*resource.LowerLimit)
	}

	if resource.Thresholds != nil {
		resourceMap["thresholds"] = helper.ParseToMaps(resource.Thresholds)
	}
	resourceMap["tags"] = helper.ParseToMaps(resource.Tags)

	return resourceMap
}

func (thresholds *ResourceThreshold) ToMap() map[string]any {
	thresholdsMap := make(map[string]any)
	thresholdsMap["kind"] = helper.NilIfEmpty(thresholds.Kind)
	thresholdsMap["value"] = helper.NilIfEmpty(thresholds.Value)
	thresholdsMap["violation_when_reached"] = helper.NilIfEmpty(thresholds.ViolationWhenReached)
	thresholdsMap["name"] = helper.NilIfEmpty(thresholds.Name)
	return thresholdsMap
}

func (resource *Resource) FromMap(resourceMap map[string]any) error {
	resource.FromAuditMap(resourceMap)
	resource.AssetId = helper.CastString(resourceMap, "asset_id")
	resource.UnitId = helper.CastString(resourceMap, "unit_id")
	resource.MetricId = helper.CastString(resourceMap, "metric_id")
	resource.Name = helper.CastString(resourceMap, "name")
	resource.Description = helper.CastString(resourceMap, "description")
	resource.DefaultLevel = helper.CastFloat64(resourceMap, "default_level")

	if v, ok := resourceMap["lower_limit"]; ok && v != nil {
		if floatVal, ok := v.(float64); ok {
			resource.LowerLimit = &floatVal
		}
	}
	if v, ok := resourceMap["upper_limit"]; ok && v != nil {
		if floatVal, ok := v.(float64); ok {
			resource.UpperLimit = &floatVal
		}
	}

	if resourceMap["thresholds"] != nil {
		thresholds, err := helper.ParseFromMaps[ResourceThreshold](helper.CastSlice(resourceMap, "thresholds"))
		if err != nil {
			return err
		}
		resource.Thresholds = thresholds
	}

	tags, err := helper.ParseFromMaps[general_objects.KeyValue](helper.CastSlice(resourceMap, "tags"))
	if err != nil {
		return err
	}
	resource.Tags = tags

	return nil
}

func (thresholds *ResourceThreshold) FromMap(thresholdsMap map[string]any) error {
	thresholds.Kind = helper.CastString(thresholdsMap, "kind")
	thresholds.Value = helper.CastFloat64(thresholdsMap, "value")
	thresholds.Name = helper.CastString(thresholdsMap, "name")
	thresholds.ViolationWhenReached = helper.CastBool(thresholdsMap, "violation_when_reached")
	return nil
}

// PostReadProcess reorders the API response's thresholds to match the plan/state
// order (matched by kind+value), preventing perpetual diffs when the API returns
// thresholds in a different order than they were configured.
func (resource *Resource) PostReadProcess(_ *provider.Client, newValue any) error {
	newResource, ok := newValue.(*Resource)
	if !ok || newResource == nil {
		return nil
	}
	newResource.Thresholds = helper.ReorderByKey(resource.Thresholds, newResource.Thresholds,
		func(t ResourceThreshold) string { return fmt.Sprintf("%s:%v", t.Kind, t.Value) })
	return nil
}
