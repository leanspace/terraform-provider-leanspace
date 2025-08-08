package areas_of_interest

import "github.com/leanspace/terraform-provider-leanspace/helper/general_objects"

type AreaOfInterest struct {
	ID             string                     `json:"id"`
	Name           string                     `json:"name"`
	Shape          *AreaOfInterestShape       `json:"shape"`
	Tags           []general_objects.KeyValue `json:"tags,omitempty"`
	CreatedAt      string                     `json:"createdAt"`
	CreatedBy      string                     `json:"createdBy"`
	LastModifiedAt string                     `json:"lastModifiedAt"`
	LastModifiedBy string                     `json:"lastModifiedBy"`
}

func (areaOfInterest *AreaOfInterest) GetID() string { return areaOfInterest.ID }

type AreaOfInterestShape struct {
	Type                string     `json:"type"` // POINT, CIRCLE, POLYGON
	Geolocation         *GeoPoint  `json:"geolocation,omitempty"`
	CenterGeolocation   *GeoPoint  `json:"centerGeolocation,omitempty"`
	RadiusInMeters      float64    `json:"radiusInMeters,omitempty"`
	VerticesGeolocation []GeoPoint `json:"verticesGeolocation,omitempty"`
}

type GeoPoint struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Altitude  float64 `json:"altitude"`
}
