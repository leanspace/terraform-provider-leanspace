package orbits

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/leanspace/terraform-provider-leanspace/helper"
	"github.com/leanspace/terraform-provider-leanspace/helper/general_objects"
)

type OrbitTF struct {
	general_objects.AuditModelTF
	SatelliteId            types.String                 `tfsdk:"satellite_id"`
	Name                   types.String                 `tfsdk:"name"`
	IdealOrbit             *IdealOrbitTF                `tfsdk:"ideal_orbit"`
	GpsConfiguration       *GpsConfigurationTF          `tfsdk:"gps_configuration"`
	SatelliteConfiguration *SatelliteConfigurationTF    `tfsdk:"satellite_configuration"`
	Tags                   []general_objects.KeyValueTF `tfsdk:"tags"`
}

type IdealOrbitTF struct {
	Type                          types.String  `tfsdk:"type"`
	Inclination                   types.Float64 `tfsdk:"inclination"`
	RightAscensionOfAscendingNode types.Float64 `tfsdk:"right_ascension_of_ascending_node"`
	ArgumentOfPerigee             types.Float64 `tfsdk:"argument_of_perigee"`
	AltitudeInMeters              types.Float64 `tfsdk:"altitude_in_meters"`
	Eccentricity                  types.Float64 `tfsdk:"eccentricity"`
	PerigeeAltitudeInMeters       types.Float64 `tfsdk:"perigee_altitude_in_meters"`
	ApogeeAltitudeInMeters        types.Float64 `tfsdk:"apogee_altitude_in_meters"`
	SemiMajorAxis                 types.Float64 `tfsdk:"semi_major_axis"`
}

type GpsConfigurationTF struct {
	GpsMetrics         *GpsMetricsTF         `tfsdk:"gps_metrics"`
	StandardDeviations *StandardDeviationsTF `tfsdk:"standard_deviations"`
}

type GpsMetricsTF struct {
	MetricIdForLatitude    types.String `tfsdk:"metric_id_for_latitude"`
	MetricIdForLongitude   types.String `tfsdk:"metric_id_for_longitude"`
	MetricIdForAltitude    types.String `tfsdk:"metric_id_for_altitude"`
	MetricIdForGroundSpeed types.String `tfsdk:"metric_id_for_ground_speed"`
}

type StandardDeviationsTF struct {
	Latitude    types.Float64 `tfsdk:"latitude"`
	Longitude   types.Float64 `tfsdk:"longitude"`
	Altitude    types.Float64 `tfsdk:"altitude"`
	GroundSpeed types.Float64 `tfsdk:"ground_speed"`
}

type SatelliteConfigurationTF struct {
	DragCrossSection      types.Float64 `tfsdk:"drag_cross_section"`
	RadiationCrossSection types.Float64 `tfsdk:"radiation_cross_section"`
}

func (x *Orbit) ToTF() any {
	var idealOrbit *IdealOrbitTF
	if x.IdealOrbit != nil {
		idealOrbit = &IdealOrbitTF{
			Type:                          helper.TFStringValue(x.IdealOrbit.Type),
			Inclination:                   helper.TFFloat64Value(x.IdealOrbit.Inclination),
			RightAscensionOfAscendingNode: helper.TFFloat64Value(x.IdealOrbit.RightAscensionOfAscendingNode),
			ArgumentOfPerigee:             helper.TFFloat64Value(x.IdealOrbit.ArgumentOfPerigee),
			AltitudeInMeters:              helper.TFFloat64Value(x.IdealOrbit.AltitudeInMeters),
			Eccentricity:                  helper.TFFloat64Value(x.IdealOrbit.Eccentricity),
			PerigeeAltitudeInMeters:       helper.TFFloat64Value(x.IdealOrbit.PerigeeAltitudeInMeters),
			ApogeeAltitudeInMeters:        helper.TFFloat64Value(x.IdealOrbit.ApogeeAltitudeInMeters),
			SemiMajorAxis:                 helper.TFFloat64Value(x.IdealOrbit.SemiMajorAxis),
		}
	}

	var gpsConfiguration *GpsConfigurationTF
	if x.GpsConfiguration != nil {
		var gpsMetrics *GpsMetricsTF
		if x.GpsConfiguration.GpsMetrics != nil {
			gpsMetrics = &GpsMetricsTF{
				MetricIdForLatitude:    helper.TFStringValue(x.GpsConfiguration.GpsMetrics.MetricIdForLatitude),
				MetricIdForLongitude:   helper.TFStringValue(x.GpsConfiguration.GpsMetrics.MetricIdForLongitude),
				MetricIdForAltitude:    helper.TFStringValue(x.GpsConfiguration.GpsMetrics.MetricIdForAltitude),
				MetricIdForGroundSpeed: helper.TFStringValue(x.GpsConfiguration.GpsMetrics.MetricIdForGroundSpeed),
			}
		}
		var standardDeviations *StandardDeviationsTF
		if x.GpsConfiguration.StandardDeviations != nil {
			standardDeviations = &StandardDeviationsTF{
				Latitude:    helper.TFFloat64Value(x.GpsConfiguration.StandardDeviations.Latitude),
				Longitude:   helper.TFFloat64Value(x.GpsConfiguration.StandardDeviations.Longitude),
				Altitude:    helper.TFFloat64Value(x.GpsConfiguration.StandardDeviations.Altitude),
				GroundSpeed: helper.TFFloat64Value(x.GpsConfiguration.StandardDeviations.GroundSpeed),
			}
		}
		gpsConfiguration = &GpsConfigurationTF{
			GpsMetrics:         gpsMetrics,
			StandardDeviations: standardDeviations,
		}
	}

	var satelliteConfiguration *SatelliteConfigurationTF
	if x.SatelliteConfiguration != nil {
		satelliteConfiguration = &SatelliteConfigurationTF{
			DragCrossSection:      helper.TFFloat64Value(x.SatelliteConfiguration.DragCrossSection),
			RadiationCrossSection: helper.TFFloat64Value(x.SatelliteConfiguration.RadiationCrossSection),
		}
	}

	return &OrbitTF{
		AuditModelTF:           general_objects.AuditModelToTF(&x.AuditModel),
		SatelliteId:            helper.TFStringValue(x.SatelliteId),
		Name:                   helper.TFStringValue(x.Name),
		IdealOrbit:             idealOrbit,
		GpsConfiguration:       gpsConfiguration,
		SatelliteConfiguration: satelliteConfiguration,
		Tags:                   general_objects.KeyValuesToTF(x.Tags),
	}
}

func (tf *OrbitTF) ToAPI() any {
	var idealOrbit *IdealOrbit
	if tf.IdealOrbit != nil {
		idealOrbit = &IdealOrbit{
			Type:                          helper.FromTFString(tf.IdealOrbit.Type),
			Inclination:                   helper.FromTFFloat64(tf.IdealOrbit.Inclination),
			RightAscensionOfAscendingNode: helper.FromTFFloat64(tf.IdealOrbit.RightAscensionOfAscendingNode),
			ArgumentOfPerigee:             helper.FromTFFloat64(tf.IdealOrbit.ArgumentOfPerigee),
			AltitudeInMeters:              helper.FromTFFloat64(tf.IdealOrbit.AltitudeInMeters),
			Eccentricity:                  helper.FromTFFloat64(tf.IdealOrbit.Eccentricity),
			PerigeeAltitudeInMeters:       helper.FromTFFloat64(tf.IdealOrbit.PerigeeAltitudeInMeters),
			ApogeeAltitudeInMeters:        helper.FromTFFloat64(tf.IdealOrbit.ApogeeAltitudeInMeters),
			SemiMajorAxis:                 helper.FromTFFloat64(tf.IdealOrbit.SemiMajorAxis),
		}
	}

	var gpsConfiguration *GpsConfiguration
	if tf.GpsConfiguration != nil {
		var gpsMetrics *GpsMetrics
		if tf.GpsConfiguration.GpsMetrics != nil {
			gpsMetrics = &GpsMetrics{
				MetricIdForLatitude:    helper.FromTFString(tf.GpsConfiguration.GpsMetrics.MetricIdForLatitude),
				MetricIdForLongitude:   helper.FromTFString(tf.GpsConfiguration.GpsMetrics.MetricIdForLongitude),
				MetricIdForAltitude:    helper.FromTFString(tf.GpsConfiguration.GpsMetrics.MetricIdForAltitude),
				MetricIdForGroundSpeed: helper.FromTFString(tf.GpsConfiguration.GpsMetrics.MetricIdForGroundSpeed),
			}
		}
		var standardDeviations *StandardDeviations
		if tf.GpsConfiguration.StandardDeviations != nil {
			standardDeviations = &StandardDeviations{
				Latitude:    helper.FromTFFloat64(tf.GpsConfiguration.StandardDeviations.Latitude),
				Longitude:   helper.FromTFFloat64(tf.GpsConfiguration.StandardDeviations.Longitude),
				Altitude:    helper.FromTFFloat64(tf.GpsConfiguration.StandardDeviations.Altitude),
				GroundSpeed: helper.FromTFFloat64(tf.GpsConfiguration.StandardDeviations.GroundSpeed),
			}
		}
		gpsConfiguration = &GpsConfiguration{
			GpsMetrics:         gpsMetrics,
			StandardDeviations: standardDeviations,
		}
	}

	var satelliteConfiguration *SatelliteConfiguration
	if tf.SatelliteConfiguration != nil {
		satelliteConfiguration = &SatelliteConfiguration{
			DragCrossSection:      helper.FromTFFloat64(tf.SatelliteConfiguration.DragCrossSection),
			RadiationCrossSection: helper.FromTFFloat64(tf.SatelliteConfiguration.RadiationCrossSection),
		}
	}

	return &Orbit{
		AuditModel:             general_objects.AuditModelFromTF(tf.AuditModelTF),
		SatelliteId:            helper.FromTFString(tf.SatelliteId),
		Name:                   helper.FromTFString(tf.Name),
		IdealOrbit:             idealOrbit,
		GpsConfiguration:       gpsConfiguration,
		SatelliteConfiguration: satelliteConfiguration,
		Tags:                   general_objects.KeyValuesFromTF(tf.Tags),
	}
}
