package pass_delay_configuration

import (
	"github.com/leanspace/terraform-provider-leanspace/helper"
)

func (passDelayConfiguration *PassDelayConfiguration) ToMap() map[string]any {
	passDelayConfigurationMap := make(map[string]any)
	passDelayConfigurationMap["id"] = helper.NilIfEmpty(passDelayConfiguration.ID)
	passDelayConfigurationMap["name"] = helper.NilIfEmpty(passDelayConfiguration.Name)
	passDelayConfigurationMap["aos_delay_in_millisecond"] = helper.NilIfEmpty(passDelayConfiguration.AosDelayInMillisecond)
	passDelayConfigurationMap["los_delay_in_millisecond"] = helper.NilIfEmpty(passDelayConfiguration.LosDelayInMillisecond)

	return passDelayConfigurationMap
}

func (passDelayConfiguration *PassDelayConfiguration) FromMap(passDelayConfigurationMap map[string]any) error {
	passDelayConfiguration.ID = helper.CastString(passDelayConfigurationMap, "id")
	passDelayConfiguration.Name = helper.CastString(passDelayConfigurationMap, "name")
	passDelayConfiguration.AosDelayInMillisecond = helper.CastFloat64(passDelayConfigurationMap, "aos_delay_in_millisecond")
	passDelayConfiguration.LosDelayInMillisecond = helper.CastFloat64(passDelayConfigurationMap, "los_delay_in_millisecond")

	return nil
}
