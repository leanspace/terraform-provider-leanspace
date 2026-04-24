package service_accounts

//go:generate go run github.com/leanspace/terraform-provider-leanspace/tools/gen_tf_models -struct ServiceAccount

import "github.com/leanspace/terraform-provider-leanspace/helper/general_objects"

type ServiceAccount struct {
	general_objects.AuditModel
	Name        string                     `json:"name"`
	PolicyIds   []string                   `json:"policyIds"`
	Credentials Credentials                `json:"credentials" tf:"object"` // The tf:"object" tag is needed to generate the correct Terraform schema, even though Credentials is not a pointer.
	Tags        []general_objects.KeyValue `json:"tags,omitempty"`
}

type Credentials struct {
	ClientId     string `json:"clientId"`
	ClientSecret string `json:"clientSecret"`
}
