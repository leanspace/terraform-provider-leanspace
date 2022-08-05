package teams

type Team struct {
	ID             string   `json:"id"`
	Name           string   `json:"name"`
	PolicyIds      []string `json:"policyIds"`
	Members        []string `json:"members"`
	CreatedAt      string   `json:"createdAt"`
	CreatedBy      string   `json:"createdBy"`
	LastModifiedAt string   `json:"lastModifiedAt"`
	LastModifiedBy string   `json:"lastModifiedBy"`
}

func (team *Team) GetID() string { return team.ID }
