package plan_states

import "github.com/leanspace/terraform-provider-leanspace/provider"

var PlanStateDataType = provider.DataSourceType[PlanState, *PlanState]{
	ResourceIdentifier: "leanspace_plan_states",
	Path:               "plans-repository/plans/states",
	Schema:             planStateSchema,
	FilterSchema:       nil,
}
