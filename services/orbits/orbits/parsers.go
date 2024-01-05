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
