package command_definitions

import "github.com/leanspace/terraform-provider-leanspace/helper/general_objects"

type CommandDefinition struct {
	general_objects.AuditModel
	NodeId      string          `json:"nodeId"`
	Name        string          `json:"name"`
	Description string          `json:"description,omitempty"`
	Identifier  string          `json:"identifier,omitempty"`
	Metadata    []Metadata[any] `json:"metadata,omitempty"`
	Arguments   []Argument[any] `json:"arguments,omitempty"`
}

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
