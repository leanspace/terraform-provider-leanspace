package contact_reservation_status_mappings

import "github.com/leanspace/terraform-provider-leanspace/helper/general_objects"

type ContactReservationStatusMapping struct {
	general_objects.AuditModel
	ContactStateId  string `json:"contactStateId"`
	LeafspaceStatus string `json:"leafspaceStatus"`
}
