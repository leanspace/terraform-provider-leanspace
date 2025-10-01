package service_accounts

import "github.com/leanspace/terraform-provider-leanspace/helper/general_objects"

type ServiceAccount struct {
	ID             string                     `json:"id"`
	Name           string                     `json:"name"`
	PolicyIds      []string                   `json:"policyIds"`
	Credentials    Credentials                `json:"credentials"`
	Tags           []general_objects.KeyValue `json:"tags,omitempty"`
	CreatedAt      string                     `json:"createdAt"`
	CreatedBy      string                     `json:"createdBy"`
	LastModifiedAt string                     `json:"lastModifiedAt"`
	LastModifiedBy string                     `json:"lastModifiedBy"`
}

func (serviceAccount *ServiceAccount) GetID() string { return serviceAccount.ID }

type Credentials struct {
	ClientId     string `json:"clientId"`
	ClientSecret string `json:"clientSecret"`
}
