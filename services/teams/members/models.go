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
	SecurityTags   []string `json:"-"`
}

func (member *Member) GetID() string { return member.ID }

func (member *Member) GetSecurityTags() []string { return member.SecurityTags }

func (member *Member) SetSecurityTags(tags []string) { member.SecurityTags = tags }