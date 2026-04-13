package orbits

import (
	"github.com/leanspace/terraform-provider-leanspace/helper"
	"github.com/leanspace/terraform-provider-leanspace/helper/general_objects"
)

func (orbit *Orbit) ToMap() map[string]any {
	orbitMap := orbit.ToAuditMap()
	orbitMap["satellite_id"] = helper.NilIfEmpty(orbit.SatelliteId)
	orbitMap["name"] = helper.NilIfEmpty(orbit.Name)
	if orbit.IdealOrbit != nil {
		orbitMap["ideal_orbit"] = orbit.IdealOrbit.ToMap()
	}
	if orbit.GpsConfiguration != nil {
		orbitMap["gps_configuration"] = orbit.GpsConfiguration.ToMap()
	}
	if orbit.SatelliteConfiguration != nil {
		orbitMap["satellite_configuration"] = orbit.SatelliteConfiguration.ToMap()
	}
	orbitMap["tags"] = helper.ParseToMaps(orbit.Tags)
	return orbitMap
}

func (idealOrbit *IdealOrbit) ToMap() map[string]any {
	idealOrbitMap := make(map[string]any)
	idealOrbitMap["type"] = helper.NilIfEmpty(idealOrbit.Type)
	idealOrbitMap["inclination"] = helper.NilIfEmpty(idealOrbit.Inclination)
	idealOrbitMap["right_ascension_of_ascending_node"] = helper.NilIfEmpty(idealOrbit.RightAscensionOfAscendingNode)
	idealOrbitMap["argument_of_perigee"] = helper.NilIfEmpty(idealOrbit.ArgumentOfPerigee)
	idealOrbitMap["altitude_in_meters"] = helper.NilIfEmpty(idealOrbit.AltitudeInMeters)
	idealOrbitMap["eccentricity"] = helper.NilIfEmpty(idealOrbit.Eccentricity)
	idealOrbitMap["perigee_altitude_in_meters"] = helper.NilIfEmpty(idealOrbit.PerigeeAltitudeInMeters)
	idealOrbitMap["apogee_altitude_in_meters"] = helper.NilIfEmpty(idealOrbit.ApogeeAltitudeInMeters)
	idealOrbitMap["semi_major_axis"] = helper.NilIfEmpty(idealOrbit.SemiMajorAxis)
	return idealOrbitMap
}

func (gpsConfiguration *GpsConfiguration) ToMap() map[string]any {
	gpsConfigurationMap := make(map[string]any)
	if gpsConfiguration.GpsMetrics != nil {
		gpsConfigurationMap["gps_metrics"] = gpsConfiguration.GpsMetrics.ToMap()
	}
	if gpsConfiguration.StandardDeviations != nil {
		gpsConfigurationMap["standard_deviations"] = gpsConfiguration.StandardDeviations.ToMap()
	}
	return gpsConfigurationMap
}

func (gpsMetrics *GpsMetrics) ToMap() map[string]any {
	gpsMetricsMap := make(map[string]any)
	gpsMetricsMap["metric_id_for_latitude"] = helper.NilIfEmpty(gpsMetrics.MetricIdForLatitude)
	gpsMetricsMap["metric_id_for_longitude"] = helper.NilIfEmpty(gpsMetrics.MetricIdForLongitude)
	gpsMetricsMap["metric_id_for_altitude"] = helper.NilIfEmpty(gpsMetrics.MetricIdForAltitude)
	gpsMetricsMap["metric_id_for_ground_speed"] = helper.NilIfEmpty(gpsMetrics.MetricIdForGroundSpeed)
	return gpsMetricsMap
}

func (standardDeviations *StandardDeviations) ToMap() map[string]any {
	standardDeviationsMap := make(map[string]any)
	standardDeviationsMap["latitude"] = helper.NilIfEmpty(standardDeviations.Latitude)
	standardDeviationsMap["longitude"] = helper.NilIfEmpty(standardDeviations.Longitude)
	standardDeviationsMap["altitude"] = helper.NilIfEmpty(standardDeviations.Altitude)
	standardDeviationsMap["ground_speed"] = helper.NilIfEmpty(standardDeviations.GroundSpeed)
	return standardDeviationsMap
}

func (satelliteConfiguration *SatelliteConfiguration) ToMap() map[string]any {
	satelliteConfigurationMap := make(map[string]any)
	satelliteConfigurationMap["drag_cross_section"] = helper.NilIfEmpty(satelliteConfiguration.DragCrossSection)
	satelliteConfigurationMap["radiation_cross_section"] = helper.NilIfEmpty(satelliteConfiguration.RadiationCrossSection)
	return satelliteConfigurationMap
}

func (orbit *Orbit) FromMap(orbitMap map[string]any) error {
	orbit.FromAuditMap(orbitMap)
	orbit.SatelliteId = helper.CastString(orbitMap, "satellite_id")
	orbit.Name = helper.CastString(orbitMap, "name")
	orbit.IdealOrbit = new(IdealOrbit)
	if err := orbit.IdealOrbit.FromMap(helper.CastMapAny(orbitMap, "ideal_orbit")); err != nil {
		return err
	}
	orbit.GpsConfiguration = new(GpsConfiguration)
	if err := orbit.GpsConfiguration.FromMap(helper.CastMapAny(orbitMap, "gps_configuration")); err != nil {
		return err
	}
	orbit.SatelliteConfiguration = new(SatelliteConfiguration)
	if err := orbit.SatelliteConfiguration.FromMap(helper.CastMapAny(orbitMap, "satellite_configuration")); err != nil {
		return err
	}
	if tags, err := helper.ParseFromMaps[general_objects.KeyValue](helper.CastSlice(orbitMap, "tags")); err != nil {
		return err
	} else {
		orbit.Tags = tags
	}
	return nil
}

func (idealOrbit *IdealOrbit) FromMap(idealOrbitMap map[string]any) error {
	idealOrbit.Type = helper.CastString(idealOrbitMap, "type")
	idealOrbit.Inclination = helper.CastFloat64(idealOrbitMap, "inclination")
	idealOrbit.RightAscensionOfAscendingNode = helper.CastFloat64(idealOrbitMap, "right_ascension_of_ascending_node")
	idealOrbit.ArgumentOfPerigee = helper.CastFloat64(idealOrbitMap, "argument_of_perigee")
	idealOrbit.AltitudeInMeters = helper.CastFloat64(idealOrbitMap, "altitude_in_meters")
	idealOrbit.Eccentricity = helper.CastFloat64(idealOrbitMap, "eccentricity")
	idealOrbit.PerigeeAltitudeInMeters = helper.CastFloat64(idealOrbitMap, "perigee_altitude_in_meters")
	idealOrbit.ApogeeAltitudeInMeters = helper.CastFloat64(idealOrbitMap, "apogee_altitude_in_meters")
	idealOrbit.SemiMajorAxis = helper.CastFloat64(idealOrbitMap, "semi_major_axis")
	return nil
}

func (gpsConfiguration *GpsConfiguration) FromMap(gpsConfigurationMap map[string]any) error {
	gpsConfiguration.GpsMetrics = new(GpsMetrics)
	if err := gpsConfiguration.GpsMetrics.FromMap(helper.CastMapAny(gpsConfigurationMap, "gps_metrics")); err != nil {
		return err
	}

	gpsConfiguration.StandardDeviations = new(StandardDeviations)
	if err := gpsConfiguration.StandardDeviations.FromMap(helper.CastMapAny(gpsConfigurationMap, "standard_deviations")); err != nil {
		return err
	}

	return nil
}

func (gpsMetrics *GpsMetrics) FromMap(gpsMetricsMap map[string]any) error {
	gpsMetrics.MetricIdForLatitude = helper.CastString(gpsMetricsMap, "metric_id_for_latitude")
	gpsMetrics.MetricIdForLongitude = helper.CastString(gpsMetricsMap, "metric_id_for_longitude")
	gpsMetrics.MetricIdForAltitude = helper.CastString(gpsMetricsMap, "metric_id_for_altitude")
	gpsMetrics.MetricIdForGroundSpeed = helper.CastString(gpsMetricsMap, "metric_id_for_ground_speed")
	return nil
}

func (standardDeviations *StandardDeviations) FromMap(standardDeviationsMap map[string]any) error {
	standardDeviations.Latitude = helper.CastFloat64(standardDeviationsMap, "latitude")
	standardDeviations.Longitude = helper.CastFloat64(standardDeviationsMap, "longitude")
	standardDeviations.Altitude = helper.CastFloat64(standardDeviationsMap, "altitude")
	standardDeviations.GroundSpeed = helper.CastFloat64(standardDeviationsMap, "ground_speed")
	return nil
}

func (satelliteConfiguration *SatelliteConfiguration) FromMap(satelliteConfigurationMap map[string]any) error {
	satelliteConfiguration.DragCrossSection = helper.CastFloat64(satelliteConfigurationMap, "drag_cross_section")
	satelliteConfiguration.RadiationCrossSection = helper.CastFloat64(satelliteConfigurationMap, "radiation_cross_section")
	return nil
}
