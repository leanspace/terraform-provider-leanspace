package request_states

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/leanspace/terraform-provider-leanspace/helper"
	"github.com/leanspace/terraform-provider-leanspace/helper/general_objects"
)

type RequestStateTF struct {
	general_objects.AuditModelTF
	Name types.String `tfsdk:"name"`
}

func (s *RequestState) ToTF() any {
	return &RequestStateTF{
		AuditModelTF: general_objects.AuditModelToTF(&s.AuditModel),
		Name:         helper.TFStringValue(s.Name),
	}
}

func (tf *RequestStateTF) ToAPI() any {
	return &RequestState{
		AuditModel: general_objects.AuditModelFromTF(tf.AuditModelTF),
		Name:       helper.FromTFString(tf.Name),
	}
}
