package contact_reservations

type ContactReservation struct {
	ID              string `json:"id"`
	ContactStateId  string `json:"contactStateId"`
	LeafspaceStatus string `json:"leafspaceStatus"`
	Status          string `json:"status"`
	CreatedAt       string `json:"createdAt"`
	CreatedBy       string `json:"createdBy"`
	LastModifiedAt  string `json:"lastModifiedAt"`
	LastModifiedBy  string `json:"lastModifiedBy"`
}

func (contactReservation *ContactReservation) GetID() string {
	return contactReservation.ID
}
