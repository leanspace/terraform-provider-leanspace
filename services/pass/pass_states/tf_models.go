package pass_states

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/leanspace/terraform-provider-leanspace/helper/general_objects"
)

type PassStateTF struct {
	general_objects.AuditModelTF
	Name     types.String `tfsdk:"name"`
	ReadOnly types.Bool   `tfsdk:"read_only"`
}

func (s *PassState) ToTF() any {
	return general_objects.ReflectToTF[PassStateTF](s)
}

func (tf *PassStateTF) ToAPI() any {
	return general_objects.ReflectFromTF[PassState](tf)
}
