package event_criticalities

//go:generate go run github.com/leanspace/terraform-provider-leanspace/tools/gen_tf_models -struct EventCriticalities

import "github.com/leanspace/terraform-provider-leanspace/helper/general_objects"

type EventCriticalities struct {
	general_objects.AuditModel
	Name     string                     `json:"name"`
	ReadOnly bool                       `json:"readOnly"`
	Tags     []general_objects.KeyValue `json:"tags,omitempty"`
}
