package areas_of_interest

import "github.com/leanspace/terraform-provider-leanspace/helper/general_objects"

type AreaOfInterest struct {
	general_objects.AuditModel
	Name  string                     `json:"name"`
	Shape *AreaOfInterestShape       `json:"shape"`
	Tags  []general_objects.KeyValue `json:"tags,omitempty"`
}

type AreaOfInterestShape struct {
	Type                string     `json:"type"` // POINT, CIRCLE, POLYGON
	Geolocation         *GeoPoint  `json:"geolocation,omitempty"`
	CenterGeolocation   *GeoPoint  `json:"centerGeolocation,omitempty"`
	RadiusInMeters      *float64   `json:"radiusInMeters,omitempty"`
	VerticesGeolocation []GeoPoint `json:"verticesGeolocation,omitempty"`
}

type GeoPoint struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Altitude  float64 `json:"altitude"`
}
