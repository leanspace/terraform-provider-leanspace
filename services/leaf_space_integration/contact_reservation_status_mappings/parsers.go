package contact_reservation_status_mappings

import (
	"github.com/leanspace/terraform-provider-leanspace/helper"
)

func (contactReservationStatusMapping *ContactReservationStatusMapping) ToMap() map[string]any {
	contactReservatrionStatusMappingMap := contactReservationStatusMapping.ToAuditMap()
	contactReservatrionStatusMappingMap["contact_state_id"] = helper.NilIfEmpty(contactReservationStatusMapping.ContactStateId)
	contactReservatrionStatusMappingMap["leafspace_status"] = helper.NilIfEmpty(contactReservationStatusMapping.LeafspaceStatus)
	return contactReservatrionStatusMappingMap
}

func (contactReservationStatusMapping *ContactReservationStatusMapping) FromMap(leafSpaceIntegrationMap map[string]any) error {
	contactReservationStatusMapping.FromAuditMap(leafSpaceIntegrationMap)
	contactReservationStatusMapping.ContactStateId = helper.CastString(leafSpaceIntegrationMap, "contact_state_id")
	contactReservationStatusMapping.LeafspaceStatus = helper.CastString(leafSpaceIntegrationMap, "leafspace_status")
	return nil
}
