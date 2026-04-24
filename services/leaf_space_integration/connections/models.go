package connections

//go:generate go run github.com/leanspace/terraform-provider-leanspace/tools/gen_tf_models -struct LeafSpaceConnection

import "github.com/leanspace/terraform-provider-leanspace/helper/general_objects"

type LeafSpaceConnection struct {
	general_objects.AuditModel
	Name                string  `json:"name"`
	DomainUrl           string  `json:"domainUrl"`
	AuthenticationToken *string `json:"authenticationToken,omitempty"`
	Password            string  `json:"password"`
	Username            string  `json:"username"`
	Status              string  `json:"status"`
}
