package access_policies

type AccessPolicy struct {
	ID             string      `json:"id" terra:"id"`
	Name           string      `json:"name" terra:"name"`
	Description    string      `json:"description" terra:"description"`
	ReadOnly       bool        `json:"readOnly" terra:"read_only"`
	Statements     []Statement `json:"statements" terra:"statements"`
	CreatedAt      string      `json:"createdAt" terra:"created_at"`
	CreatedBy      string      `json:"createdBy" terra:"created_by"`
	LastModifiedAt string      `json:"lastModifiedAt" terra:"last_modified_at"`
	LastModifiedBy string      `json:"lastModifiedBy" terra:"last_modified_by"`
}

func (policy *AccessPolicy) GetID() string { return policy.ID }

type Statement struct {
	Name    string   `json:"name" terra:"name"`
	Actions []string `json:"actions" terra:"actions"`
}
