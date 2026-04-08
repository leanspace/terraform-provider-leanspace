package request_definitions

import (
	"github.com/leanspace/terraform-provider-leanspace/helper/general_objects"
	"github.com/leanspace/terraform-provider-leanspace/services/activities/activity_definitions"
)

type RequestDefinition struct {
	general_objects.AuditModel
	Name                             string                                         `json:"name"`
	Description                      string                                         `json:"description,omitempty"`
	PlanTemplateIds                  []string                                       `json:"planTemplateIds"`
	FeasibilityConstraintDefinitions []FeasibilityConstraintDefinition              `json:"feasibilityConstraintDefinitions"`
	ConfigurationArgumentDefinitions []activity_definitions.ArgumentDefinition[any] `json:"configurationArgumentDefinitions,omitempty"`
	ConfigurationArgumentMappings    []ArgumentMapping                              `json:"configurationArgumentMappings,omitempty"`
}

type FeasibilityConstraintDefinition struct {
	general_objects.AuditModel
	Name                string                                         `json:"name"`
	Description         string                                         `json:"description,omitempty"`
	Required            bool                                           `json:"required"`
	ArgumentDefinitions []activity_definitions.ArgumentDefinition[any] `json:"argumentDefinitions,omitempty"`
}

type ArgumentMapping struct {
	PlanTemplateId                           string `json:"planTemplateId"`
	ActivityDefinitionPosition               int    `json:"activityDefinitionPosition"`
	ConfigurationArgumentDefinitionName      string `json:"configurationArgumentDefinitionName"`
	ActivityDefinitionArgumentDefinitionName string `json:"activityDefinitionArgumentDefinitionName"`
}
