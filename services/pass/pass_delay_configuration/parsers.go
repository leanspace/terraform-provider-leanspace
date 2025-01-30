package pass_delay_configuration

func (passDelayConfiguration *PassDelayConfiguration) ToMap() map[string]any {
	passDelayConfigurationMap := make(map[string]any)
	passDelayConfigurationMap["id"] = passDelayConfiguration.ID
	passDelayConfigurationMap["name"] = passDelayConfiguration.Name
	passDelayConfigurationMap["aos_delay_in_millisecond"] = passDelayConfiguration.AosDelayInMillisecond
	passDelayConfigurationMap["los_delay_in_millisecond"] = passDelayConfiguration.LosDelayInMillisecond

	return passDelayConfigurationMap
}

func (passDelayConfiguration *PassDelayConfiguration) FromMap(passDelayConfigurationMap map[string]any) error {
	passDelayConfiguration.ID = passDelayConfigurationMap["id"].(string)
	passDelayConfiguration.Name = passDelayConfigurationMap["name"].(string)
	passDelayConfiguration.AosDelayInMillisecond = passDelayConfigurationMap["aos_delay_in_millisecond"].(float64)
	passDelayConfiguration.LosDelayInMillisecond = passDelayConfigurationMap["los_delay_in_millisecond"].(float64)

	return nil
}
