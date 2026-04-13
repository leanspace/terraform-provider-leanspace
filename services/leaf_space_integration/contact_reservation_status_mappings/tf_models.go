package contact_reservation_status_mappings

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/leanspace/terraform-provider-leanspace/helper"
	"github.com/leanspace/terraform-provider-leanspace/helper/general_objects"
)

type ContactReservationStatusMappingTF struct {
	general_objects.AuditModelTF
	ContactStateId  types.String `tfsdk:"contact_state_id"`
	LeafspaceStatus types.String `tfsdk:"leafspace_status"`
}

func (s *ContactReservationStatusMapping) ToTF() any {
	return &ContactReservationStatusMappingTF{
		AuditModelTF:    general_objects.AuditModelToTF(&s.AuditModel),
		ContactStateId:  helper.TFStringValue(s.ContactStateId),
		LeafspaceStatus: helper.TFStringValue(s.LeafspaceStatus),
	}
}

func (tf *ContactReservationStatusMappingTF) ToAPI() any {
	return &ContactReservationStatusMapping{
		AuditModel:      general_objects.AuditModelFromTF(tf.AuditModelTF),
		ContactStateId:  helper.FromTFString(tf.ContactStateId),
		LeafspaceStatus: helper.FromTFString(tf.LeafspaceStatus),
	}
}
