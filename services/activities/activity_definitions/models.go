package activity_definitions

import "leanspace-terraform-provider/helper/general_objects"

type ActivityDefinition struct {
	ID                  string                    `json:"id" terra:"id"`
	NodeId              string                    `json:"nodeId" terra:"node_id"`
	Name                string                    `json:"name" terra:"name"`
	Description         string                    `json:"description,omitempty" terra:"description"`
	EstimatedDuration   int                       `json:"estimatedDuration" terra:"estimated_duration"`
	Metadata            []Metadata[any]           `json:"metadata,omitempty" terra:"metadata"`
	ArgumentDefinitions []ArgumentDefinition[any] `json:"argumentDefinitions,omitempty" terra:"argument_definitions"`
	CommandMappings     []CommandMapping          `json:"commandMappings" terra:"command_mappings"`
	CreatedAt           string                    `json:"createdAt" terra:"created_at"`
	CreatedBy           string                    `json:"createdBy" terra:"created_by"`
	LastModifiedAt      string                    `json:"lastModifiedAt" terra:"last_modified_at"`
	LastModifiedBy      string                    `json:"lastModifiedBy" terra:"last_modified_by"`
}

func (actDefinition *ActivityDefinition) GetID() string { return actDefinition.ID }

type Metadata[T any] struct {
	ID          string                            `json:"id" terra:"id"`
	Name        string                            `json:"name" terra:"name"`
	Description string                            `json:"description,omitempty" terra:"description"`
	Attributes  general_objects.ValueAttribute[T] `json:"attributes" terra:"attributes"`
}

type ArgumentDefinition[T any] struct {
	ID          string                                 `json:"id" terra:"id"`
	Name        string                                 `json:"name" terra:"name"`
	Description string                                 `json:"description,omitempty" terra:"description"`
	Attributes  general_objects.DefinitionAttribute[T] `json:"attributes" terra:"attributes"`
}

type CommandMapping struct {
	CommandDefinitionId string            `json:"commandDefinitionId" terra:"command_definition_id"`
	Position            int               `json:"position" terra:"position"`
	DelayInMilliseconds int               `json:"delayInMilliseconds" terra:"delay_in_milliseconds"`
	ArgumentMappings    []ArgumentMapping `json:"argumentMappings" terra:"argument_mappings"`
	MetadataMappings    []MetadataMapping `json:"metadataMappings" terra:"metadata_mappings"`
}

type ArgumentMapping struct {
	ActivityDefinitionArgumentName string `json:"activityDefinitionArgumentName" terra:"activity_definition_argument_name"`
	CommandDefinitionArgumentName  string `json:"commandDefinitionArgumentName" terra:"command_definition_argument_name"`
}

type MetadataMapping struct {
	ActivityDefinitionMetadataName string `json:"activityDefinitionMetadataName" terra:"activity_definition_metadata_name"`
	CommandDefinitionArgumentName  string `json:"commandDefinitionArgumentName" terra:"command_definition_argument_name"`
}
