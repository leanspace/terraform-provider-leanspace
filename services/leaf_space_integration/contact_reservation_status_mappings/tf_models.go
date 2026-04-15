package contact_reservation_status_mappings

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/leanspace/terraform-provider-leanspace/helper/general_objects"
)

type ContactReservationStatusMappingTF struct {
	general_objects.AuditModelTF
	ContactStateId  types.String `tfsdk:"contact_state_id"`
	LeafspaceStatus types.String `tfsdk:"leafspace_status"`
}

func (s *ContactReservationStatusMapping) ToTF() any {
	return general_objects.ReflectToTF[ContactReservationStatusMappingTF](s)
}

func (tf *ContactReservationStatusMappingTF) ToAPI() any {
	return general_objects.ReflectFromTF[ContactReservationStatusMapping](tf)
}
