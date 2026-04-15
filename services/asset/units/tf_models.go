package units

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/leanspace/terraform-provider-leanspace/helper/general_objects"
)

type UnitTF struct {
	ID          types.String `tfsdk:"id"`
	Symbol      types.String `tfsdk:"symbol"`
	DisplayName types.String `tfsdk:"display_name"`
}

func (s *Unit) ToTF() any {
	return general_objects.ReflectToTF[UnitTF](s)
}

func (tf *UnitTF) ToAPI() any {
	return general_objects.ReflectFromTF[Unit](tf)
}
