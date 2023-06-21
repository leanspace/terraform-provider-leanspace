package orbit_resources

type OrbitResource struct {
	ID                                        string                     `json:"id"`
	SatelliteId                               string                     `json:"satelliteId"`
	Name                                      string                     `json:"name"`
	DataSource                                string                     `json:"dataSource"`
	AutomaticTleUpdate                        string                     `json:"automaticTleUpdate"`
	AutomaticPropagation                      string                     `json:"automaticPropagation"`
	GpsMetricIds                              GpsMetricIds               `json:"gpsMetricIds,omitempty"`
	CreatedAt                                 string                     `json:"createdAt"`
	CreatedBy                                 string                     `json:"createdBy"`
	LastModifiedAt                            string                     `json:"lastModifiedAt"`
	LastModifiedBy                            string                     `json:"lastModifiedBy"`
}

func (queue *OrbitResource) GetID() string { return queue.ID }

type GpsMetricIds struct {
	MetricIdForPositionX                      string                     `json:"metricIdForPositionX"`
	MetricIdForPositionY                      string                     `json:"metricIdForPositionY"`
	MetricIdForPositionZ                      string                     `json:"metricIdForPositionZ"`
	MetricIdForVelocityX                      string                     `json:"metricIdForVelocityX"`
	MetricIdForVelocityY                      string                     `json:"metricIdForVelocityY"`
	MetricIdForVelocityZ                      string                     `json:"metricIdForVelocityZ"`
}