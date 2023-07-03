package orbit_resources

import (
	"github.com/leanspace/terraform-provider-leanspace/helper/general_objects"
)

type OrbitResource struct {
	ID                   string                     `json:"id"`
	SatelliteId          string                     `json:"satelliteId"`
	Name                 string                     `json:"name"`
	DataSource           string                     `json:"dataSource"`
	AutomaticPropagation bool                       `json:"automaticPropagation"`
	GpsMetricIds         *GpsMetricIds              `json:"gpsMetricIds,omitempty"`
	Tags                 []general_objects.KeyValue `json:"tags,omitempty"`
	CreatedAt            string                     `json:"createdAt"`
	CreatedBy            string                     `json:"createdBy"`
	LastModifiedAt       string                     `json:"lastModifiedAt"`
	LastModifiedBy       string                     `json:"lastModifiedBy"`
}

func (orbitResource *OrbitResource) GetID() string { return orbitResource.ID }

type GpsMetricIds struct {
	MetricIdForPositionX string `json:"metricIdForPositionX,omitempty"`
	MetricIdForPositionY string `json:"metricIdForPositionY,omitempty"`
	MetricIdForPositionZ string `json:"metricIdForPositionZ,omitempty"`
	MetricIdForVelocityX string `json:"metricIdForVelocityX,omitempty"`
	MetricIdForVelocityY string `json:"metricIdForVelocityY,omitempty"`
	MetricIdForVelocityZ string `json:"metricIdForVelocityZ,omitempty"`
}
