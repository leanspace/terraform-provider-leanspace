package event_criticalities

import "github.com/leanspace/terraform-provider-leanspace/helper/general_objects"

type EventCriticalities struct {
	general_objects.AuditModel
	Name     string                     `json:"name"`
	ReadOnly bool                       `json:"readOnly"`
	Tags     []general_objects.KeyValue `json:"tags,omitempty"`
}
