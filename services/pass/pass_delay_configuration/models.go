package pass_delay_configuration

//go:generate go run github.com/leanspace/terraform-provider-leanspace/tools/gen_tf_models -struct PassDelayConfiguration

type PassDelayConfiguration struct {
	ID                    string  `json:"id"`
	Name                  string  `json:"name"`
	AosDelayInMillisecond float64 `json:"aosDelayInMillisecond"`
	LosDelayInMillisecond float64 `json:"losDelayInMillisecond"`
}

func (passDelayConfiguration *PassDelayConfiguration) GetID() string {
	return passDelayConfiguration.ID
}
