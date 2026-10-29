package metrics

import (
	"github.com/leanspace/terraform-provider-leanspace/helper"
	"github.com/leanspace/terraform-provider-leanspace/helper/general_objects"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func (metric *Metric[T]) ToMap() map[string]any {
	metricMap := make(map[string]any)
	metricMap["id"] = metric.ID
	metricMap["node_id"] = metric.NodeId
	metricMap["name"] = metric.Name
	metricMap["description"] = metric.Description
	metricMap["created_at"] = metric.CreatedAt
	metricMap["created_by"] = metric.CreatedBy
	metricMap["last_modified_at"] = metric.LastModifiedAt
	metricMap["last_modified_by"] = metric.LastModifiedBy
	metricMap["attributes"] = []map[string]any{metric.Attributes.ToMap()}
	if metric.Calibration != nil {
		metricMap["calibration"] = []map[string]any{metric.Calibration.ToMap()}
	}
	metricMap["tags"] = helper.ParseToMaps(metric.Tags)
	return metricMap
}

func (metric *Metric[T]) FromMap(metricMap map[string]any) error {
	metric.ID = metricMap["id"].(string)
	metric.NodeId = metricMap["node_id"].(string)
	metric.Name = metricMap["name"].(string)
	metric.Description = metricMap["description"].(string)
	metric.CreatedAt = metricMap["created_at"].(string)
	metric.CreatedBy = metricMap["created_by"].(string)
	metric.LastModifiedAt = metricMap["last_modified_at"].(string)
	metric.LastModifiedBy = metricMap["last_modified_by"].(string)
	if len(metricMap["attributes"].([]any)) > 0 {
		if err := metric.Attributes.FromMap(metricMap["attributes"].([]any)[0].(map[string]any)); err != nil {
			return err
		}
	}
	if len(metricMap["calibration"].([]any)) > 0 && metricMap["calibration"].([]any)[0] != nil {
		metric.Calibration = new(Calibration)
		if err := metric.Calibration.FromMap(metricMap["calibration"].([]any)[0].(map[string]any)); err != nil {
			return err
		}
	}
	if tags, err := helper.ParseFromMaps[general_objects.KeyValue](metricMap["tags"].(*schema.Set).List()); err != nil {
		return err
	} else {
		metric.Tags = tags
	}
	return nil
}

func (calibration *Calibration) ToMap() map[string]any {
	calibrationMap := make(map[string]any)
	calibrationMap["type"] = string(calibration.Type)
	calibrationMap["description"] = calibration.Description
	calibrationMap["coefficients"] = calibration.Coefficients
	calibrationMap["base"] = calibration.Base
	if calibration.Capping != nil {
		calibrationMap["capping"] = []map[string]any{calibration.Capping.ToMap()}
	}
	calibrationMap["min"] = boundToList(calibration.Min)
	calibrationMap["max"] = boundToList(calibration.Max)
	calibrationMap["ranges"] = helper.ParseToMaps(calibration.Ranges)
	calibrationMap["mappings"] = calibration.Mappings
	return calibrationMap
}

func (calibration *Calibration) FromMap(calibrationMap map[string]any) error {
	calibration.Type = CalibrationType(calibrationMap["type"].(string))
	calibration.Description = calibrationMap["description"].(string)
	calibration.Coefficients = make([]float64, len(calibrationMap["coefficients"].([]any)))
	for index, coefficient := range calibrationMap["coefficients"].([]any) {
		calibration.Coefficients[index] = coefficient.(float64)
	}
	calibration.Base = calibrationMap["base"].(float64)
	if len(calibrationMap["capping"].([]any)) > 0 && calibrationMap["capping"].([]any)[0] != nil {
		calibration.Capping = new(CalibrationBounds)
		if err := calibration.Capping.FromMap(calibrationMap["capping"].([]any)[0].(map[string]any)); err != nil {
			return err
		}
	}
	calibration.Min = boundFromList(calibrationMap["min"])
	calibration.Max = boundFromList(calibrationMap["max"])
	if ranges, err := helper.ParseFromMaps[CalibrationRange](calibrationMap["ranges"].([]any)); err != nil {
		return err
	} else {
		calibration.Ranges = ranges
	}
	calibration.Mappings = make(map[string]string, len(calibrationMap["mappings"].(map[string]any)))
	for key, value := range calibrationMap["mappings"].(map[string]any) {
		calibration.Mappings[key] = value.(string)
	}
	return nil
}

func (bounds *CalibrationBounds) ToMap() map[string]any {
	boundsMap := make(map[string]any)
	boundsMap["min"] = boundToList(bounds.Min)
	boundsMap["max"] = boundToList(bounds.Max)
	return boundsMap
}

func (bounds *CalibrationBounds) FromMap(boundsMap map[string]any) error {
	bounds.Min = boundFromList(boundsMap["min"])
	bounds.Max = boundFromList(boundsMap["max"])
	return nil
}

func (calibrationRange *CalibrationRange) ToMap() map[string]any {
	rangeMap := make(map[string]any)
	rangeMap["min"] = boundToList(calibrationRange.Min)
	rangeMap["max"] = boundToList(calibrationRange.Max)
	rangeMap["value"] = calibrationRange.Value
	return rangeMap
}

func (calibrationRange *CalibrationRange) FromMap(rangeMap map[string]any) error {
	calibrationRange.Min = boundFromList(rangeMap["min"])
	calibrationRange.Max = boundFromList(rangeMap["max"])
	calibrationRange.Value = rangeMap["value"].(float64)
	return nil
}

// Bounds are stored as a list of at most one value, an empty list being no bound.
func boundToList(bound *float64) []any {
	if bound == nil {
		return []any{}
	}
	return []any{*bound}
}

func boundFromList(raw any) *float64 {
	if list, ok := raw.([]any); ok && len(list) > 0 {
		if value, ok := list[0].(float64); ok {
			return &value
		}
	}
	return nil
}
