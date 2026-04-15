package contact_states

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/leanspace/terraform-provider-leanspace/helper/general_objects"
)

type ContactStateTF struct {
	general_objects.AuditModelTF
	Name     types.String `tfsdk:"name"`
	ReadOnly types.Bool   `tfsdk:"read_only"`
}

func (s *ContactState) ToTF() any {
	return general_objects.ReflectToTF[ContactStateTF](s)
}

func (tf *ContactStateTF) ToAPI() any {
	return general_objects.ReflectFromTF[ContactState](tf)
}
