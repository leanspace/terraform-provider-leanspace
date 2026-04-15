package pass_delay_configuration

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/leanspace/terraform-provider-leanspace/helper"
	"github.com/leanspace/terraform-provider-leanspace/helper/general_objects"
)

type PassDelayConfigurationTF struct {
	ID                    types.String  `tfsdk:"id"`
	Name                  types.String  `tfsdk:"name"`
	AosDelayInMillisecond types.Float64 `tfsdk:"aos_delay_in_millisecond"`
	LosDelayInMillisecond types.Float64 `tfsdk:"los_delay_in_millisecond"`
}

func (s *PassDelayConfiguration) ToTF() any {
	return general_objects.ReflectToTF[PassDelayConfigurationTF](s)
}

func (tf *PassDelayConfigurationTF) ToAPI() any {
	return general_objects.ReflectFromTF[PassDelayConfiguration](tf)
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
