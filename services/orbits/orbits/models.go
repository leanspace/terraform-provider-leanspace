package orbits

import (
	"github.com/leanspace/terraform-provider-leanspace/helper/general_objects"
)

type Orbit struct {
	ID                   string                     `json:"id"`
	SatelliteId          string                     `json:"satelliteId"`
	Name                 string                     `json:"name"`
	IdealOrbit           *IdealOrbit                `json:"idealOrbit,omitempty"`
	Tags                 []general_objects.KeyValue `json:"tags,omitempty"`
	CreatedAt            string                     `json:"createdAt"`
	CreatedBy            string                     `json:"createdBy"`
	LastModifiedAt       string                     `json:"lastModifiedAt"`
	LastModifiedBy       string                     `json:"lastModifiedBy"`
}

func (orbit *Orbit) GetID() string { return orbit.ID }

type IdealOrbit struct {
	Type string `json:"type"`
	Inclination float64 `json:"inclination"`
	RightAscensionOfAscendingNode float64 `json:"rightAscensionOfAscendingNode"`
	ArgumentOfPerigee float64 `json:"argumentOfPerigee"`
	AltitudeInMeters float64 `json:"altitudeInMeters"`
	Eccentricity float64 `json:"eccentricity"`
	PerigeeAltitudeInMeters float64 `json:"perigeeAltitudeInMeters"`
	ApogeeAltitudeInMeters float64 `json:"apogeeAltitudeInMeters"`
	SemiMajorAxis float64 `json:"semiMajorAxis"`
}
