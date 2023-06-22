package orbit_resources

import (
	"encoding/json"
)

func (orbitResource *OrbitResource) ToMap() map[string]any {
	orbitResourceMap := make(map[string]any)
	orbitResourceMap["id"] = orbitResource.ID
	orbitResourceMap["satellite_id"] = orbitResource.SatelliteId
	orbitResourceMap["name"] = orbitResource.Name
	orbitResourceMap["data_source"] = orbitResource.DataSource
	orbitResourceMap["automatic_tle_update"] = orbitResource.AutomaticTleUpdate
	orbitResourceMap["automatic_propagation"] = orbitResource.AutomaticPropagation
    orbitResourceMap["gps_metric_ids"] = []any{orbitResource.GpsMetricIds.ToMap()}
	orbitResourceMap["created_at"] = orbitResource.CreatedAt
	orbitResourceMap["created_by"] = orbitResource.CreatedBy
	orbitResourceMap["last_modified_at"] = orbitResource.LastModifiedAt
	orbitResourceMap["last_modified_by"] = orbitResource.LastModifiedBy

	return orbitResourceMap
}

func (gpsMetricIds GpsMetricIds) ToMap() map[string]any {
	gpsMetricIdsMap := make(map[string]any)
	gpsMetricIdsMap["metric_id_for_position_x"] = gpsMetricIds.MetricIdForPositionX
	gpsMetricIdsMap["metric_id_for_position_y"] = gpsMetricIds.MetricIdForPositionY
	gpsMetricIdsMap["metric_id_for_position_z"] = gpsMetricIds.MetricIdForPositionZ
	gpsMetricIdsMap["metric_id_for_velocity_x"] = gpsMetricIds.MetricIdForVelocityX
	gpsMetricIdsMap["metric_id_for_velocity_y"] = gpsMetricIds.MetricIdForVelocityY
	gpsMetricIdsMap["metric_id_for_velocity_z"] = gpsMetricIds.MetricIdForVelocityZ
	return gpsMetricIdsMap
}

func (orbitResource *OrbitResource) FromMap(orbitResourceMap map[string]any) error {
	orbitResource.ID = orbitResourceMap["id"].(string)
	orbitResource.SatelliteId = orbitResourceMap["satellite_id"].(string)
	orbitResource.Name = orbitResourceMap["name"].(string)
	orbitResource.DataSource = orbitResourceMap["data_source"].(string)
	orbitResource.AutomaticTleUpdate = orbitResourceMap["automatic_tle_update"].(bool)
	orbitResource.AutomaticPropagation = orbitResourceMap["automatic_propagation"].(bool)
	if len(orbitResourceMap["gps_metric_ids"].([]any)) > 0 {
        if err := orbitResource.GpsMetricIds.FromMap(orbitResourceMap["gps_metric_ids"].([]any)[0].(map[string]any)); err != nil {
            return err
        }
    }
	orbitResource.CreatedAt = orbitResourceMap["created_at"].(string)
	orbitResource.CreatedBy = orbitResourceMap["created_by"].(string)
	orbitResource.LastModifiedAt = orbitResourceMap["last_modified_at"].(string)
	orbitResource.LastModifiedBy = orbitResourceMap["last_modified_by"].(string)

	return nil
}

func (gpsMetricIds *GpsMetricIds) FromMap(gpsMetricIdsMap map[string]any) error {
	gpsMetricIds.MetricIdForPositionX = gpsMetricIdsMap["metric_id_for_position_x"].(string)
	gpsMetricIds.MetricIdForPositionY = gpsMetricIdsMap["metric_id_for_position_y"].(string)
	gpsMetricIds.MetricIdForPositionZ = gpsMetricIdsMap["metric_id_for_position_z"].(string)
	gpsMetricIds.MetricIdForVelocityX = gpsMetricIdsMap["metric_id_for_velocity_x"].(string)
	gpsMetricIds.MetricIdForVelocityY = gpsMetricIdsMap["metric_id_for_velocity_y"].(string)
	gpsMetricIds.MetricIdForVelocityZ = gpsMetricIdsMap["metric_id_for_velocity_z"].(string)
	return nil
}

// will be automatically called by json.Marshal (see generic_client.go)
//
// technical note: In Terraform, when using the optional flag for a list variable,
// it will still be represented as an empty list ([]) in the resulting JSON output
// if it is not provided in the configuration.
// Omitting it entirely from the JSON output is not directly supported.
func (orbitResource OrbitResource) MarshalJSON() ([]byte, error) {

    if orbitResource.DataSource == "GPS_METRIC" {

        var complexOrbitResource ComplexOrbitResource

        complexOrbitResource.ID = orbitResource.ID
        complexOrbitResource.SatelliteId = orbitResource.SatelliteId
        complexOrbitResource.Name = orbitResource.Name
        complexOrbitResource.DataSource = orbitResource.DataSource
        complexOrbitResource.GpsMetricIds = orbitResource.GpsMetricIds
        complexOrbitResource.AutomaticTleUpdate = orbitResource.AutomaticTleUpdate
        complexOrbitResource.AutomaticPropagation = orbitResource.AutomaticPropagation
        complexOrbitResource.CreatedAt = orbitResource.CreatedAt
        complexOrbitResource.CreatedBy = orbitResource.CreatedBy
        complexOrbitResource.LastModifiedAt = orbitResource.LastModifiedAt
        complexOrbitResource.LastModifiedBy = orbitResource.LastModifiedBy

        return json.Marshal(complexOrbitResource)

    } else {

        var simpleOrbitResource SimpleOrbitResource

        simpleOrbitResource.ID = orbitResource.ID
        simpleOrbitResource.SatelliteId = orbitResource.SatelliteId
        simpleOrbitResource.Name = orbitResource.Name
        simpleOrbitResource.DataSource = orbitResource.DataSource
        simpleOrbitResource.AutomaticTleUpdate = orbitResource.AutomaticTleUpdate
        simpleOrbitResource.AutomaticPropagation = orbitResource.AutomaticPropagation
        simpleOrbitResource.CreatedAt = orbitResource.CreatedAt
        simpleOrbitResource.CreatedBy = orbitResource.CreatedBy
        simpleOrbitResource.LastModifiedAt = orbitResource.LastModifiedAt
        simpleOrbitResource.LastModifiedBy = orbitResource.LastModifiedBy

        return json.Marshal(simpleOrbitResource)
    }
}