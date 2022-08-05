package command_definitions

import "leanspace-terraform-provider/helper/general_objects"

type CommandDefinition struct {
	ID             string          `json:"id"`
	NodeId         string          `json:"nodeId"`
	Name           string          `json:"name"`
	Description    string          `json:"description,omitempty"`
	Identifier     string          `json:"identifier,omitempty"`
	Metadata       []Metadata[any] `json:"metadata,omitempty"`
	Arguments      []Argument[any] `json:"arguments,omitempty"`
	CreatedAt      string          `json:"createdAt"`
	CreatedBy      string          `json:"createdBy"`
	LastModifiedAt string          `json:"lastModifiedAt"`
	LastModifiedBy string          `json:"lastModifiedBy"`
}

func (cmdDefinition *CommandDefinition) GetID() string { return cmdDefinition.ID }

type Metadata[T any] struct {
	ID          string                            `json:"id"`
	Name        string                            `json:"name"`
	Description string                            `json:"description,omitempty"`
	Attributes  general_objects.ValueAttribute[T] `json:"attributes"`
}

type Argument[T any] struct {
	ID          string                                 `json:"id"`
	Name        string                                 `json:"name"`
	Identifier  string                                 `json:"identifier"`
	Description string                                 `json:"description,omitempty"`
	Attributes  general_objects.DefinitionAttribute[T] `json:"attributes"`
}
