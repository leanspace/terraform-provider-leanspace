package members

import "github.com/leanspace/terraform-provider-leanspace/helper/general_objects"

type Member struct {
	general_objects.AuditModel
	Name      string   `json:"name"`
	Email     string   `json:"email"`
	Status    string   `json:"status"`
	PolicyIds []string `json:"policyIds"`
}
