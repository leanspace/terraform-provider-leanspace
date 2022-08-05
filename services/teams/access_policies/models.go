package access_policies

type AccessPolicy struct {
	ID             string      `json:"id"`
	Name           string      `json:"name"`
	Description    string      `json:"description"`
	ReadOnly       bool        `json:"readOnly"`
	Statements     []Statement `json:"statements"`
	CreatedAt      string      `json:"createdAt"`
	CreatedBy      string      `json:"createdBy"`
	LastModifiedAt string      `json:"lastModifiedAt"`
	LastModifiedBy string      `json:"lastModifiedBy"`
}

func (policy *AccessPolicy) GetID() string { return policy.ID }

type Statement struct {
	Name    string   `json:"name"`
	Actions []string `json:"actions"`
}
