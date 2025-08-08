package sensors

import "github.com/leanspace/terraform-provider-leanspace/provider"

var SensorDataType = provider.DataSourceType[Sensor, *Sensor]{
	ResourceIdentifier: "leanspace_sensors",
	Path:               "orbits-repository/sensors",
	Schema:             sensorSchema,
	FilterSchema:       dataSourceFilterSchema,
}
