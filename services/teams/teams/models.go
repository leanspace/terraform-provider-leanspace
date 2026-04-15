package teams

//go:generate go run github.com/leanspace/terraform-provider-leanspace/tools/gen_tf_models -struct Team

import "github.com/leanspace/terraform-provider-leanspace/helper/general_objects"

type Team struct {
	general_objects.AuditModel
	Name      string                     `json:"name"`
	PolicyIds []string                   `json:"policyIds"`
	Members   []string                   `json:"members"`
	Tags      []general_objects.KeyValue `json:"tags,omitempty"`
}
