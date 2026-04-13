package units

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/leanspace/terraform-provider-leanspace/helper"
)

type UnitTF struct {
	ID          types.String `tfsdk:"id"`
	Symbol      types.String `tfsdk:"symbol"`
	DisplayName types.String `tfsdk:"display_name"`
}

func (s *Unit) ToTF() any {
	return &UnitTF{
		ID:          helper.TFStringValue(s.ID),
		Symbol:      helper.TFStringValue(s.Symbol),
		DisplayName: helper.TFStringValue(s.DisplayName),
	}
}

func (tf *UnitTF) ToAPI() any {
	return &Unit{
		ID:          helper.FromTFString(tf.ID),
		Symbol:      helper.FromTFString(tf.Symbol),
		DisplayName: helper.FromTFString(tf.DisplayName),
	}
}
