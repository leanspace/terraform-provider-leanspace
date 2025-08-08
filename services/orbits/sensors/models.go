package sensors

import "github.com/leanspace/terraform-provider-leanspace/helper/general_objects"

type Sensor struct {
	ID             string                     `json:"id"`
	SatelliteID    string                     `json:"satelliteId"`
	Name           string                     `json:"name"`
	ApertureShape  *ApertureShape             `json:"apertureShape"`
	Tags           []general_objects.KeyValue `json:"tags,omitempty"`
	CreatedAt      string                     `json:"createdAt"`
	CreatedBy      string                     `json:"createdBy"`
	LastModifiedAt string                     `json:"lastModifiedAt"`
	LastModifiedBy string                     `json:"lastModifiedBy"`
}

func (areaOfInterest *Sensor) GetID() string { return areaOfInterest.ID }

type ApertureShape struct {
	Type                        string             `json:"type"` // CIRCULAR, RECTANGULAR
	ApertureCenter              *Vector3D          `json:"apertureCenter,omitempty"`
	HalfApertureAngle           *HalfApertureAngle `json:"halfApertureAngle,omitempty"`
	FirstAxisVector             *Vector3D          `json:"firstAxisVector,omitempty"`
	FirstAxisHalfApertureAngle  *HalfApertureAngle `json:"firstAxisHalfApertureAngle,omitempty"`
	SecondAxisVector            *Vector3D          `json:"secondAxisVector,omitempty"`
	SecondAxisHalfApertureAngle *HalfApertureAngle `json:"secondAxisHalfApertureAngle,omitempty"`
}

type Vector3D struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
	Z float64 `json:"z"`
}

type HalfApertureAngle struct {
	Degrees float64 `json:"degrees"`
}
