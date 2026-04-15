package feasibility_constraint_definitions

import (
	"github.com/leanspace/terraform-provider-leanspace/helper/general_objects"
	"github.com/leanspace/terraform-provider-leanspace/services/activities/activity_definitions"
)

type FeasibilityConstraintDefinition struct {
	general_objects.AuditModel
	Name                string                                         `json:"name"`
	Description         *string                                        `json:"description,omitempty"`
	ArgumentDefinitions []activity_definitions.ArgumentDefinition[any] `json:"argumentDefinitions,omitempty"`
}
