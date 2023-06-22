package orbit_resources

type OrbitResource struct {
	ID                                        string                     `json:"id"`
	SatelliteId                               string                     `json:"satelliteId"`
	Name                                      string                     `json:"name"`
	DataSource                                string                     `json:"dataSource"`
	AutomaticTleUpdate                        bool                       `json:"automaticTleUpdate"`
	AutomaticPropagation                      bool                       `json:"automaticPropagation"`
	GpsMetricIds                              GpsMetricIds               `json:"gpsMetricIds,omitempty"`
	CreatedAt                                 string                     `json:"createdAt"`
	CreatedBy                                 string                     `json:"createdBy"`
	LastModifiedAt                            string                     `json:"lastModifiedAt"`
	LastModifiedBy                            string                     `json:"lastModifiedBy"`
}

func (orbitResource *OrbitResource) GetID() string { return orbitResource.ID }

type GpsMetricIds struct {
	MetricIdForPositionX                      string                     `json:"metricIdForPositionX,omitempty"`
	MetricIdForPositionY                      string                     `json:"metricIdForPositionY,omitempty"`
	MetricIdForPositionZ                      string                     `json:"metricIdForPositionZ,omitempty"`
	MetricIdForVelocityX                      string                     `json:"metricIdForVelocityX,omitempty"`
	MetricIdForVelocityY                      string                     `json:"metricIdForVelocityY,omitempty"`
	MetricIdForVelocityZ                      string                     `json:"metricIdForVelocityZ,omitempty"`
}

// ComplexOrbitResource && SimpleOrbitResource needed as workaround to "implement" the json marshalling of optional embedded object (GpsMetricIds)
// see parsers file for explanation
type ComplexOrbitResource struct {
	ID                                        string                     `json:"id"`
	SatelliteId                               string                     `json:"satelliteId"`
	Name                                      string                     `json:"name"`
	DataSource                                string                     `json:"dataSource"`
	AutomaticTleUpdate                        bool                       `json:"automaticTleUpdate"`
	AutomaticPropagation                      bool                       `json:"automaticPropagation"`
	GpsMetricIds                              GpsMetricIds               `json:"gpsMetricIds,omitempty"`
	CreatedAt                                 string                     `json:"createdAt"`
	CreatedBy                                 string                     `json:"createdBy"`
	LastModifiedAt                            string                     `json:"lastModifiedAt"`
	LastModifiedBy                            string                     `json:"lastModifiedBy"`
}

type SimpleOrbitResource struct {
	ID                                        string                     `json:"id"`
	SatelliteId                               string                     `json:"satelliteId"`
	Name                                      string                     `json:"name"`
	DataSource                                string                     `json:"dataSource"`
	AutomaticTleUpdate                        bool                       `json:"automaticTleUpdate"`
	AutomaticPropagation                      bool                       `json:"automaticPropagation"`
	CreatedAt                                 string                     `json:"createdAt"`
	CreatedBy                                 string                     `json:"createdBy"`
	LastModifiedAt                            string                     `json:"lastModifiedAt"`
	LastModifiedBy                            string                     `json:"lastModifiedBy"`
}