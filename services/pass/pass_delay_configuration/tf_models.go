package pass_delay_configuration

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/leanspace/terraform-provider-leanspace/helper"
)

type PassDelayConfigurationTF struct {
	ID                    types.String  `tfsdk:"id"`
	Name                  types.String  `tfsdk:"name"`
	AosDelayInMillisecond types.Float64 `tfsdk:"aos_delay_in_millisecond"`
	LosDelayInMillisecond types.Float64 `tfsdk:"los_delay_in_millisecond"`
}

func (s *PassDelayConfiguration) ToTF() any {
	return &PassDelayConfigurationTF{
		ID:                    helper.TFStringValue(s.ID),
		Name:                  helper.TFStringValue(s.Name),
		AosDelayInMillisecond: helper.TFFloat64Value(s.AosDelayInMillisecond),
		LosDelayInMillisecond: helper.TFFloat64Value(s.LosDelayInMillisecond),
	}
}

func (tf *PassDelayConfigurationTF) ToAPI() any {
	return &PassDelayConfiguration{
		ID:                    helper.FromTFString(tf.ID),
		Name:                  helper.FromTFString(tf.Name),
		AosDelayInMillisecond: helper.FromTFFloat64(tf.AosDelayInMillisecond),
		LosDelayInMillisecond: helper.FromTFFloat64(tf.LosDelayInMillisecond),
	}
}

// PassDelayConfigurationDSTF is the data-source TF model for unique reads.
// The FilterSchema is nil, so the data source only exposes "id".
type PassDelayConfigurationDSTF struct {
	ID types.String `tfsdk:"id"`
}

func (s *PassDelayConfiguration) ToDSTF() any {
	return &PassDelayConfigurationDSTF{
		ID: helper.TFStringValue(s.ID),
	}
}
