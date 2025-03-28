package orbits

import (
	"github.com/leanspace/terraform-provider-leanspace/helper"
	"github.com/leanspace/terraform-provider-leanspace/helper/general_objects"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func (orbit *Orbit) ToMap() map[string]any {
	orbitMap := make(map[string]any)
	orbitMap["id"] = orbit.ID
	orbitMap["satellite_id"] = orbit.SatelliteId
	orbitMap["name"] = orbit.Name
	if orbit.IdealOrbit != nil {
		orbitMap["ideal_orbit"] = []map[string]any{orbit.IdealOrbit.ToMap()}
	}
	if orbit.GpsConfiguration != nil {
		orbitMap["gps_configuration"] = []map[string]any{orbit.GpsConfiguration.ToMap()}
	}
	if orbit.SatelliteConfiguration != nil {
		orbitMap["satellite_configuration"] = []map[string]any{orbit.SatelliteConfiguration.ToMap()}
	}
	orbitMap["tags"] = helper.ParseToMaps(orbit.Tags)
	orbitMap["created_at"] = orbit.CreatedAt
	orbitMap["created_by"] = orbit.CreatedBy
	orbitMap["last_modified_at"] = orbit.LastModifiedAt
	orbitMap["last_modified_by"] = orbit.LastModifiedBy

	return orbitMap
}

func (idealOrbit *IdealOrbit) ToMap() map[string]any {
	idealOrbitMap := make(map[string]any)
	idealOrbitMap["type"] = idealOrbit.Type
	idealOrbitMap["inclination"] = idealOrbit.Inclination
	idealOrbitMap["right_ascension_of_ascending_node"] = idealOrbit.RightAscensionOfAscendingNode
	idealOrbitMap["argument_of_perigee"] = idealOrbit.ArgumentOfPerigee
	idealOrbitMap["altitude_in_meters"] = idealOrbit.AltitudeInMeters
	idealOrbitMap["eccentricity"] = idealOrbit.Eccentricity
	idealOrbitMap["perigee_altitude_in_meters"] = idealOrbit.PerigeeAltitudeInMeters
	idealOrbitMap["apogee_altitude_in_meters"] = idealOrbit.ApogeeAltitudeInMeters
	idealOrbitMap["semi_major_axis"] = idealOrbit.SemiMajorAxis
	return idealOrbitMap
}

func (gpsConfiguration *GpsConfiguration) ToMap() map[string]any {
	gpsConfigurationMap := make(map[string]any)
	if gpsConfiguration.GpsMetrics != nil {
		gpsConfigurationMap["gps_metrics"] = []map[string]any{gpsConfiguration.GpsMetrics.ToMap()}
	}
	if gpsConfiguration.StandardDeviations != nil {
		gpsConfigurationMap["standard_deviations"] = []map[string]any{gpsConfiguration.StandardDeviations.ToMap()}
	}
	return gpsConfigurationMap
}

func (gpsMetrics *GpsMetrics) ToMap() map[string]any {
	gpsMetricsMap := make(map[string]any)
	gpsMetricsMap["metric_id_for_latitude"] = gpsMetrics.MetricIdForLatitude
	gpsMetricsMap["metric_id_for_longitude"] = gpsMetrics.MetricIdForLongitude
	gpsMetricsMap["metric_id_for_altitude"] = gpsMetrics.MetricIdForAltitude
	gpsMetricsMap["metric_id_for_ground_speed"] = gpsMetrics.MetricIdForGroundSpeed
	return gpsMetricsMap
}

func (standardDeviations *StandardDeviations) ToMap() map[string]any {
	standardDeviationsMap := make(map[string]any)
	standardDeviationsMap["latitude"] = standardDeviations.Latitude
	standardDeviationsMap["longitude"] = standardDeviations.Longitude
	standardDeviationsMap["altitude"] = standardDeviations.Altitude
	standardDeviationsMap["ground_speed"] = standardDeviations.GroundSpeed
	return standardDeviationsMap
}

func (satelliteConfiguration *SatelliteConfiguration) ToMap() map[string]any {
	satelliteConfigurationMap := make(map[string]any)
	satelliteConfigurationMap["drag_cross_section"] = satelliteConfiguration.DragCrossSection
	satelliteConfigurationMap["radiation_cross_section"] = satelliteConfiguration.RadiationCrossSection
	return satelliteConfigurationMap
}

func (orbit *Orbit) FromMap(orbitMap map[string]any) error {
	orbit.ID = orbitMap["id"].(string)
	orbit.SatelliteId = orbitMap["satellite_id"].(string)
	orbit.Name = orbitMap["name"].(string)
	if len(orbitMap["ideal_orbit"].([]any)) > 0 && orbitMap["ideal_orbit"].([]any)[0] != nil {
		orbit.IdealOrbit = new(IdealOrbit)
		if err := orbit.IdealOrbit.FromMap(orbitMap["ideal_orbit"].([]any)[0].(map[string]any)); err != nil {
			return err
		}
	}
	if len(orbitMap["gps_configuration"].([]any)) > 0 && orbitMap["gps_configuration"].([]any)[0] != nil {
		orbit.GpsConfiguration = new(GpsConfiguration)
		if err := orbit.GpsConfiguration.FromMap(orbitMap["gps_configuration"].([]any)[0].(map[string]any)); err != nil {
			return err
		}
	}
	if len(orbitMap["satellite_configuration"].([]any)) > 0 && orbitMap["satellite_configuration"].([]any)[0] != nil {
		orbit.SatelliteConfiguration = new(SatelliteConfiguration)
		if err := orbit.SatelliteConfiguration.FromMap(orbitMap["satellite_configuration"].([]any)[0].(map[string]any)); err != nil {
			return err
		}
	}
	if tags, err := helper.ParseFromMaps[general_objects.KeyValue](orbitMap["tags"].(*schema.Set).List()); err != nil {
		return err
	} else {
		orbit.Tags = tags
	}
	orbit.CreatedAt = orbitMap["created_at"].(string)
	orbit.CreatedBy = orbitMap["created_by"].(string)
	orbit.LastModifiedAt = orbitMap["last_modified_at"].(string)
	orbit.LastModifiedBy = orbitMap["last_modified_by"].(string)

	return nil
}

func (idealOrbit *IdealOrbit) FromMap(idealOrbitMap map[string]any) error {
	idealOrbit.Type = idealOrbitMap["type"].(string)
	idealOrbit.Inclination = idealOrbitMap["inclination"].(float64)
	idealOrbit.RightAscensionOfAscendingNode = idealOrbitMap["right_ascension_of_ascending_node"].(float64)
	idealOrbit.ArgumentOfPerigee = idealOrbitMap["argument_of_perigee"].(float64)
	idealOrbit.AltitudeInMeters = idealOrbitMap["altitude_in_meters"].(float64)
	idealOrbit.Eccentricity = idealOrbitMap["eccentricity"].(float64)
	idealOrbit.PerigeeAltitudeInMeters = idealOrbitMap["perigee_altitude_in_meters"].(float64)
	idealOrbit.ApogeeAltitudeInMeters = idealOrbitMap["apogee_altitude_in_meters"].(float64)
	idealOrbit.SemiMajorAxis = idealOrbitMap["semi_major_axis"].(float64)
	return nil
}

func (gpsConfiguration *GpsConfiguration) FromMap(gpsConfigurationMap map[string]any) error {
	if len(gpsConfigurationMap["gps_metrics"].([]any)) > 0 && gpsConfigurationMap["gps_metrics"].([]any)[0] != nil {
		gpsConfiguration.GpsMetrics = new(GpsMetrics)
		if err := gpsConfiguration.GpsMetrics.FromMap(gpsConfigurationMap["gps_metrics"].([]any)[0].(map[string]any)); err != nil {
			return err
		}
	}

	if len(gpsConfigurationMap["standard_deviations"].([]any)) > 0 && gpsConfigurationMap["standard_deviations"].([]any)[0] != nil {
		gpsConfiguration.StandardDeviations = new(StandardDeviations)
		if err := gpsConfiguration.StandardDeviations.FromMap(gpsConfigurationMap["standard_deviations"].([]any)[0].(map[string]any)); err != nil {
			return err
		}
	}

	return nil
}

func (gpsMetrics *GpsMetrics) FromMap(gpsMetricsMap map[string]any) error {
	gpsMetrics.MetricIdForLatitude = gpsMetricsMap["metric_id_for_latitude"].(string)
	gpsMetrics.MetricIdForLongitude = gpsMetricsMap["metric_id_for_longitude"].(string)
	gpsMetrics.MetricIdForAltitude = gpsMetricsMap["metric_id_for_altitude"].(string)
	gpsMetrics.MetricIdForGroundSpeed = gpsMetricsMap["metric_id_for_ground_speed"].(string)
	return nil
}

func (standardDeviations *StandardDeviations) FromMap(standardDeviationsMap map[string]any) error {
	standardDeviations.Latitude = standardDeviationsMap["latitude"].(float64)
	standardDeviations.Longitude = standardDeviationsMap["longitude"].(float64)
	standardDeviations.Altitude = standardDeviationsMap["altitude"].(float64)
	standardDeviations.GroundSpeed = standardDeviationsMap["ground_speed"].(float64)
	return nil
}

func (satelliteConfiguration *SatelliteConfiguration) FromMap(satelliteConfigurationMap map[string]any) error {
	satelliteConfiguration.DragCrossSection = satelliteConfigurationMap["drag_cross_section"].(float64)
	satelliteConfiguration.RadiationCrossSection = satelliteConfigurationMap["radiation_cross_section"].(float64)
	return nil
}
