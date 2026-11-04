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
	SecurityTags   []string                   `json:"-"`
}

func (policy *AccessPolicy) GetID() string { return policy.ID }

func (policy *AccessPolicy) GetSecurityTags() []string { return policy.SecurityTags }

func (policy *AccessPolicy) SetSecurityTags(tags []string) { policy.SecurityTags = tags }

type Statement struct {
	SecurityTags []string `json:"securityTags,omitempty"`
	Name         string   `json:"name"`
	Actions      []string `json:"actions"`
}