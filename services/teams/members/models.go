package members

type Member struct {
	ID             string   `json:"id"`
	Name           string   `json:"name"`
	Email          string   `json:"email"`
	Status         string   `json:"status"`
	PolicyIds      []string `json:"policyIds"`
	CreatedAt      string   `json:"createdAt"`
	CreatedBy      string   `json:"createdBy"`
	LastModifiedAt string   `json:"lastModifiedAt"`
	LastModifiedBy string   `json:"lastModifiedBy"`
}

func (member *Member) GetID() string { return member.ID }
