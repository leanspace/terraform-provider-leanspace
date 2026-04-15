package action_templates

import "github.com/leanspace/terraform-provider-leanspace/helper/general_objects"

type ActionTemplate struct {
	general_objects.AuditModel
	Name    string            `json:"name"`
	Type    string            `json:"type"`
	URL     *string           `json:"url,omitempty"`
	Payload *string           `json:"payload,omitempty"`
	Content *string           `json:"content,omitempty"`
	Headers map[string]string `json:"headers,omitempty"`
}
