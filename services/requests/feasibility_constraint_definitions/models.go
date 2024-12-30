package feasibility_constraint_definitions

import (
	"github.com/leanspace/terraform-provider-leanspace/services/activities/activity_definitions"
)

type FeasibilityConstraintDefinition struct {
	ID                  string                                         `json:"id"`
	Name                string                                         `json:"name"`
	Description         string                                         `json:"description,omitempty"`
	ArgumentDefinitions []activity_definitions.ArgumentDefinition[any] `json:"argumentDefinitions,omitempty"`
	CreatedAt           string                                         `json:"createdAt"`
	CreatedBy           string                                         `json:"createdBy"`
	LastModifiedAt      string                                         `json:"lastModifiedAt"`
	LastModifiedBy      string                                         `json:"lastModifiedBy"`
}

func (feasibilityConstraintDefinition *FeasibilityConstraintDefinition) GetID() string {
	return feasibilityConstraintDefinition.ID
}
