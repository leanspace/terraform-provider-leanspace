package connections

type LeafSpaceConnection struct {
	ID                  string `json:"id"`
	Name                string `json:"name"`
	DomainUrl           string `json:"domainUrl"`
	AuthenticationToken string `json:"authenticationToken"`
	Password            string `json:"password"`
	Username            string `json:"username"`
	Status              string `json:"status"`
	CreatedAt           string `json:"createdAt"`
	CreatedBy           string `json:"createdBy"`
	LastModifiedAt      string `json:"lastModifiedAt"`
	LastModifiedBy      string `json:"lastModifiedBy"`
}

func (leafSpaceConnection *LeafSpaceConnection) GetID() string {
	return leafSpaceConnection.ID
}
