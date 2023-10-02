package leaf_space_contact_reservation_status_mappings

func (contactReservationStatusMapping *ContactReservationStatusMapping) ToMap() map[string]any {
	contactReservatrionStatusMappingMap := make(map[string]any)
	contactReservatrionStatusMappingMap["id"] = contactReservationStatusMapping.ID
	contactReservatrionStatusMappingMap["contact_state_id"] = contactReservationStatusMapping.ContactStateId
	contactReservatrionStatusMappingMap["created_at"] = contactReservationStatusMapping.CreatedAt
	contactReservatrionStatusMappingMap["created_by"] = contactReservationStatusMapping.CreatedBy
	contactReservatrionStatusMappingMap["last_modified_at"] = contactReservationStatusMapping.LastModifiedAt
	contactReservatrionStatusMappingMap["last_modified_by"] = contactReservationStatusMapping.LastModifiedBy

	return contactReservatrionStatusMappingMap
}

func (contactReservationStatusMapping *ContactReservationStatusMapping) FromMap(leafSpaceIntegrationMap map[string]any) error {
	contactReservationStatusMapping.ID = leafSpaceIntegrationMap["id"].(string)
	contactReservationStatusMapping.ContactStateId = leafSpaceIntegrationMap["contact_state_id"].(string)
	contactReservationStatusMapping.LeafspaceStatus = leafSpaceIntegrationMap["leafspace_status"].(string)
	contactReservationStatusMapping.CreatedAt = leafSpaceIntegrationMap["created_at"].(string)
	contactReservationStatusMapping.CreatedBy = leafSpaceIntegrationMap["created_by"].(string)
	contactReservationStatusMapping.LastModifiedAt = leafSpaceIntegrationMap["last_modified_at"].(string)
	contactReservationStatusMapping.LastModifiedBy = leafSpaceIntegrationMap["last_modified_by"].(string)

	return nil
}
