package request_definitions

import (
	"github.com/leanspace/terraform-provider-leanspace/services/activities/activity_definitions"
)

type RequestDefinition struct {
	ID                               string                                         `json:"id"`
	Name                             string                                         `json:"name"`
	Description                      string                                         `json:"description,omitempty"`
	PlanTemplateIds                  []string                                       `json:"planTemplateIds"`
	FeasibilityConstraintDefinitions []FeasibilityConstraintDefinition              `json:"feasibilityConstraintDefinitions"`
	ConfigurationArgumentDefinitions []activity_definitions.ArgumentDefinition[any] `json:"configurationArgumentDefinitions,omitempty"`
	ConfigurationArgumentMappings    []ArgumentMapping                              `json:"configurationArgumentMappings,omitempty"`
	CreatedAt                        string                                         `json:"createdAt"`
	CreatedBy                        string                                         `json:"createdBy"`
	LastModifiedAt                   string                                         `json:"lastModifiedAt"`
	LastModifiedBy                   string                                         `json:"lastModifiedBy"`
}

func (requestDefinition *RequestDefinition) GetID() string {
	return requestDefinition.ID
}

type FeasibilityConstraintDefinition struct {
	ID                  string                                         `json:"id"`
	Name                string                                         `json:"name"`
	Description         string                                         `json:"description,omitempty"`
	Required            bool                                           `json:"required"`
	ArgumentDefinitions []activity_definitions.ArgumentDefinition[any] `json:"argumentDefinitions,omitempty"`
	CreatedAt           string                                         `json:"createdAt"`
	CreatedBy           string                                         `json:"createdBy"`
	LastModifiedAt      string                                         `json:"lastModifiedAt"`
	LastModifiedBy      string                                         `json:"lastModifiedBy"`
}

type ArgumentMapping struct {
	PlanTemplateId                           string `json:"planTemplateId"`
	ActivityDefinitionPosition               int    `json:"activityDefinitionPosition"`
	ConfigurationArgumentDefinitionName      string `json:"configurationArgumentDefinitionName"`
	ActivityDefinitionArgumentDefinitionName string `json:"activityDefinitionArgumentDefinitionName"`
}
