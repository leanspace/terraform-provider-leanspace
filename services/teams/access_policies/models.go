package access_policies

import "github.com/leanspace/terraform-provider-leanspace/helper/general_objects"

type AccessPolicy struct {
	ID             string                     `json:"id"`
	Name           string                     `json:"name"`
	Description    string                     `json:"description"`
	ReadOnly       bool                       `json:"readOnly"`
	Statements     []Statement                `json:"statements"`
	Tags           []general_objects.KeyValue `json:"tags,omitempty"`
	CreatedAt      string                     `json:"createdAt"`
	CreatedBy      string                     `json:"createdBy"`
	LastModifiedAt string                     `json:"lastModifiedAt"`
	LastModifiedBy string                     `json:"lastModifiedBy"`
}

func (policy *AccessPolicy) GetID() string { return policy.ID }

type Statement struct {
	Name    string   `json:"name"`
	Actions []string `json:"actions"`
}
