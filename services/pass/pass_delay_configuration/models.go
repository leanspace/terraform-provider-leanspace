package pass_delay_configuration

type PassDelayConfiguration struct {
	ID                    string  `json:"id"`
	Name                  string  `json:"name"`
	AosDelayInMillisecond float64 `json:"aosDelayInMillisecond"`
	LosDelayInMillisecond float64 `json:"losDelayInMillisecond"`
}

func (passDelayConfiguration *PassDelayConfiguration) GetID() string {
	return passDelayConfiguration.ID
}
