package new_plan_states

import "github.com/leanspace/terraform-provider-leanspace/provider"

var NewPlanStateDataType = provider.DataSourceType[PlanState, *PlanState]{
	ResourceIdentifier: "leanspace_new_plan_states",
	Path:               "plans-repository/plans/states",
	Schema:             planStateSchema,
	FilterSchema:       nil,
}
