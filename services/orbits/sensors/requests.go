package sensors

import (
	"github.com/leanspace/terraform-provider-leanspace/helper/general_objects"
	"github.com/leanspace/terraform-provider-leanspace/provider"
)

func (sensor *Sensor) PostReadProcess(_ *provider.Client, newValue any) error {
	newSensor, ok := newValue.(*Sensor)
	if !ok || newSensor == nil {
		return nil
	}
	newSensor.Tags = general_objects.ReorderKeyValues(sensor.Tags, newSensor.Tags)
	return nil
}
