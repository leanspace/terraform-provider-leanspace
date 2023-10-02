package leaf_space_contact_reservation_status_mappings

type ContactReservationStatusMapping struct {
	ID              string `json:"id"`
	ContactStateId  string `json:"contactStateId"`
	LeafspaceStatus string `json:"leafspaceStatus"`
	CreatedAt       string `json:"createdAt"`
	CreatedBy       string `json:"createdBy"`
	LastModifiedAt  string `json:"lastModifiedAt"`
	LastModifiedBy  string `json:"lastModifiedBy"`
}

func (contactReservationStatusMapping *ContactReservationStatusMapping) GetID() string {
	return contactReservationStatusMapping.ID
}
