package leaf_space_contact_reservations

func (contactReservation *ContactReservation) ToMap() map[string]any {
	contactReservatrionStateMap := make(map[string]any)
	contactReservatrionStateMap["id"] = contactReservation.ID
	contactReservatrionStateMap["contact_state_id"] = contactReservation.ContactStateId
	contactReservatrionStateMap["created_at"] = contactReservation.CreatedAt
	contactReservatrionStateMap["status"] = contactReservation.Status
	contactReservatrionStateMap["created_by"] = contactReservation.CreatedBy
	contactReservatrionStateMap["last_modified_at"] = contactReservation.LastModifiedAt
	contactReservatrionStateMap["last_modified_by"] = contactReservation.LastModifiedBy

	return contactReservatrionStateMap
}

func (contactReservation *ContactReservation) FromMap(leafSpaceIntegrationMap map[string]any) error {
	contactReservation.ID = leafSpaceIntegrationMap["id"].(string)
	contactReservation.ContactStateId = leafSpaceIntegrationMap["contact_state_id"].(string)
	contactReservation.LeafspaceStatus = leafSpaceIntegrationMap["leafspace_status"].(string)
	contactReservation.Status = leafSpaceIntegrationMap["status"].(string)
	contactReservation.CreatedAt = leafSpaceIntegrationMap["created_at"].(string)
	contactReservation.CreatedBy = leafSpaceIntegrationMap["created_by"].(string)
	contactReservation.LastModifiedAt = leafSpaceIntegrationMap["last_modified_at"].(string)
	contactReservation.LastModifiedBy = leafSpaceIntegrationMap["last_modified_by"].(string)

	return nil
}
