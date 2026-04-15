package access_policies

import "github.com/leanspace/terraform-provider-leanspace/helper/general_objects"

type AccessPolicy struct {
	general_objects.AuditModel
	Name        string                     `json:"name"`
	Description *string                    `json:"description,omitempty"`
	ReadOnly    bool                       `json:"readOnly"`
	Statements  []Statement                `json:"statements"`
	Tags        []general_objects.KeyValue `json:"tags,omitempty"`
}

type Statement struct {
	Name    string   `json:"name"`
	Actions []string `json:"actions"`
}
