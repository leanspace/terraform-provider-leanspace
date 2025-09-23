package command_states

import "github.com/leanspace/terraform-provider-leanspace/helper/general_objects"

type CommandState struct {
	ID             string                     `json:"id"`
	Name           string                     `json:"name"`
	ReadOnly       bool                       `json:"readOnly"`
	Tags           []general_objects.KeyValue `json:"tags,omitempty"`
	CreatedAt      string                     `json:"createdAt"`
	CreatedBy      string                     `json:"createdBy"`
	LastModifiedAt string                     `json:"lastModifiedAt"`
	LastModifiedBy string                     `json:"lastModifiedBy"`
}

func (state *CommandState) GetID() string { return state.ID }
