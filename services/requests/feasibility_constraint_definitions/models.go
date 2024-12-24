package feasibility_constraint_definitions

import "github.com/leanspace/terraform-provider-leanspace/helper/general_objects"

type FeasibilityConstraintDefinition struct {
	ID                  string                    `json:"id"`
	Name                string                    `json:"name"`
	Description         string                    `json:"description,omitempty"`
	Cloned              bool                      `json:"cloned"`
	ArgumentDefinitions []ArgumentDefinition[any] `json:"argumentDefinitions,omitempty"`
	CreatedAt           string                    `json:"createdAt"`
	CreatedBy           string                    `json:"createdBy"`
	LastModifiedAt      string                    `json:"lastModifiedAt"`
	LastModifiedBy      string                    `json:"lastModifiedBy"`
}

func (feasibilityConstraintDefinition *FeasibilityConstraintDefinition) GetID() string {
	return feasibilityConstraintDefinition.ID
}

type ArgumentDefinition[T any] struct {
	Name        string                                 `json:"name"`
	Description string                                 `json:"description,omitempty"`
	Attributes  general_objects.DefinitionAttribute[T] `json:"attributes"`
}
