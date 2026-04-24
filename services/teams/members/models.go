package members

//go:generate go run github.com/leanspace/terraform-provider-leanspace/tools/gen_tf_models -struct Member

import "github.com/leanspace/terraform-provider-leanspace/helper/general_objects"

type Member struct {
	general_objects.AuditModel
	Name      string   `json:"name"`
	Email     string   `json:"email"`
	State     string   `json:"state"`
	PolicyIds []string `json:"policyIds"`
}
