package command_definitions

import "terraform-provider-asset/asset/general_objects"

type CommandDefinition struct {
	ID             string          `json:"id" terra:"id"`
	NodeId         string          `json:"nodeId" terra:"node_id"`
	Name           string          `json:"name" terra:"name"`
	Description    string          `json:"description,omitempty" terra:"description"`
	Identifier     string          `json:"identifier,omitempty" terra:"identifier"`
	Metadata       []Metadata[any] `json:"metadata,omitempty" terra:"metadata"`
	Arguments      []Argument[any] `json:"arguments,omitempty" terra:"arguments"`
	CreatedAt      string          `json:"createdAt" terra:"created_at"`
	CreatedBy      string          `json:"createdBy" terra:"created_by"`
	LastModifiedAt string          `json:"lastModifiedAt" terra:"last_modified_at"`
	LastModifiedBy string          `json:"lastModifiedBy" terra:"last_modified_by"`
}

func (cmdDefinition *CommandDefinition) GetID() string { return cmdDefinition.ID }

type Metadata[T any] struct {
	ID          string            `json:"id" terra:"id"`
	Name        string            `json:"name" terra:"name"`
	Description string            `json:"description,omitempty" terra:"description"`
	Attributes  ValueAttribute[T] `json:"attributes" terra:"attributes"`
}

type ValueAttribute[T any] struct {
	Value  T      `json:"value,omitempty" terra:"value"`
	Type   string `json:"type" terra:"type"`
	UnitId string `json:"unit_id,omitempty" terra:"unit_id,omitempty"`
}

type Argument[T any] struct {
	ID          string                                 `json:"id" terra:"id"`
	Name        string                                 `json:"name" terra:"name"`
	Identifier  string                                 `json:"identifier" terra:"identifier"`
	Description string                                 `json:"description,omitempty" terra:"description"`
	Attributes  general_objects.DefinitionAttribute[T] `json:"attributes" terra:"attributes"`
}
