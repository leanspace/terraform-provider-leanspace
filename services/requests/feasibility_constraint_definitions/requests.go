package feasibility_constraint_definitions

import (
	"github.com/leanspace/terraform-provider-leanspace/helper"
	"github.com/leanspace/terraform-provider-leanspace/provider"
	"github.com/leanspace/terraform-provider-leanspace/services/activities/activity_definitions"
)

func (def *FeasibilityConstraintDefinition) PostReadProcess(_ *provider.Client, newValue any) error {
	newDef, ok := newValue.(*FeasibilityConstraintDefinition)
	if !ok || newDef == nil {
		return nil
	}
	newDef.ArgumentDefinitions = helper.ReorderByKey(
		def.ArgumentDefinitions,
		newDef.ArgumentDefinitions,
		func(a activity_definitions.ArgumentDefinition[any]) string { return a.Name },
	)
	return nil
}
