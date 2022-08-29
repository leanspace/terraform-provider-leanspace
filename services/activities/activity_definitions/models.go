package activity_definitions

import "github.com/leanspace/terraform-provider-leanspace/helper/general_objects"

type ActivityDefinition struct {
	ID                  string                    `json:"id"`
	NodeId              string                    `json:"nodeId"`
	Name                string                    `json:"name"`
	Description         string                    `json:"description,omitempty"`
	EstimatedDuration   int                       `json:"estimatedDuration"`
	Metadata            []Metadata[any]           `json:"metadata,omitempty"`
	ArgumentDefinitions []ArgumentDefinition[any] `json:"argumentDefinitions,omitempty"`
	CommandMappings     []CommandMapping          `json:"commandMappings"`
	CreatedAt           string                    `json:"createdAt"`
	CreatedBy           string                    `json:"createdBy"`
	LastModifiedAt      string                    `json:"lastModifiedAt"`
	LastModifiedBy      string                    `json:"lastModifiedBy"`
}

func (actDefinition *ActivityDefinition) GetID() string { return actDefinition.ID }

type Metadata[T any] struct {
	ID          string                            `json:"id"`
	Name        string                            `json:"name"`
	Description string                            `json:"description,omitempty"`
	Attributes  general_objects.ValueAttribute[T] `json:"attributes"`
}

type ArgumentDefinition[T any] struct {
	ID          string                                 `json:"id"`
	Name        string                                 `json:"name"`
	Description string                                 `json:"description,omitempty"`
	Attributes  general_objects.DefinitionAttribute[T] `json:"attributes"`
}

type CommandMapping struct {
	CommandDefinitionId string            `json:"commandDefinitionId"`
	Position            int               `json:"position"`
	DelayInMilliseconds int               `json:"delayInMilliseconds"`
	ArgumentMappings    []ArgumentMapping `json:"argumentMappings"`
	MetadataMappings    []MetadataMapping `json:"metadataMappings"`
}

type ArgumentMapping struct {
	ActivityDefinitionArgumentName string `json:"activityDefinitionArgumentName"`
	CommandDefinitionArgumentName  string `json:"commandDefinitionArgumentName"`
}

type MetadataMapping struct {
	ActivityDefinitionMetadataName string `json:"activityDefinitionMetadataName"`
	CommandDefinitionArgumentName  string `json:"commandDefinitionArgumentName"`
}
