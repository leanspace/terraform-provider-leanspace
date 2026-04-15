package contact_reservation_status_mappings

//go:generate go run github.com/leanspace/terraform-provider-leanspace/tools/gen_tf_models -struct ContactReservationStatusMapping

import "github.com/leanspace/terraform-provider-leanspace/helper/general_objects"

type ContactReservationStatusMapping struct {
	general_objects.AuditModel
	ContactStateId  string `json:"contactStateId"`
	LeafspaceStatus string `json:"leafspaceStatus"`
}
