package request_states

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/leanspace/terraform-provider-leanspace/helper/general_objects"
)

type RequestStateTF struct {
	general_objects.AuditModelTF
	Name types.String `tfsdk:"name"`
}

func (s *RequestState) ToTF() any {
	return general_objects.ReflectToTF[RequestStateTF](s)
}

func (tf *RequestStateTF) ToAPI() any {
	return general_objects.ReflectFromTF[RequestState](tf)
}
