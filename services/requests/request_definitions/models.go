package request_definitions

import "github.com/leanspace/terraform-provider-leanspace/helper/general_objects"

type RequestDefinition struct {
	ID                               string                            `json:"id"`
	Name                             string                            `json:"name"`
	Description                      string                            `json:"description,omitempty"`
	PlanTemplateIds                  []string                          `json:"planTemplateIds"`
	FeasibilityConstraintDefinitions []FeasibilityConstraintDefinition `json:"feasibilityConstraintDefinitions"`
	ConfigurationArgumentDefinitions []ArgumentDefinition[any]         `json:"configurationArgumentDefinitions"`
	ConfigurationArgumentMappings    []ArgumentMapping                 `json:"configurationArgumentMappings"`
	CreatedAt                        string                            `json:"createdAt"`
	CreatedBy                        string                            `json:"createdBy"`
	LastModifiedAt                   string                            `json:"lastModifiedAt"`
	LastModifiedBy                   string                            `json:"lastModifiedBy"`
}

func (requestDefinition *RequestDefinition) GetID() string {
	return requestDefinition.ID
}

type FeasibilityConstraintDefinition struct {
	ID                  string                    `json:"id"`
	Name                string                    `json:"name"`
	Description         string                    `json:"description,omitempty"`
	Required            bool                      `json:"required,omitempty"`
	Cloned              bool                      `json:"cloned"`
	ArgumentDefinitions []ArgumentDefinition[any] `json:"argumentDefinitions,omitempty"`
	CreatedAt           string                    `json:"createdAt"`
	CreatedBy           string                    `json:"createdBy"`
	LastModifiedAt      string                    `json:"lastModifiedAt"`
	LastModifiedBy      string                    `json:"lastModifiedBy"`
}

type ArgumentDefinition[T any] struct {
	Name        string                                 `json:"name"`
	Description string                                 `json:"description,omitempty"`
	Attributes  general_objects.DefinitionAttribute[T] `json:"attributes"`
}

type ArgumentMapping struct {
	PlanTemplateId                           string `json:"planTemplateId"`
	ActivityDefinitionPosition               int    `json:"activityDefinitionPosition"`
	ConfigurationArgumentDefinitionName      string `json:"configurationArgumentDefinitionName"`
	ActivityDefinitionArgumentDefinitionName string `json:"activityDefinitionArgumentDefinitionName"`
}
