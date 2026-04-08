package service_accounts

import "github.com/leanspace/terraform-provider-leanspace/helper/general_objects"

type ServiceAccount struct {
	general_objects.AuditModel
	Name        string                     `json:"name"`
	PolicyIds   []string                   `json:"policyIds"`
	Credentials Credentials                `json:"credentials"`
	Tags        []general_objects.KeyValue `json:"tags,omitempty"`
}

type Credentials struct {
	ClientId     string `json:"clientId"`
	ClientSecret string `json:"clientSecret"`
}
