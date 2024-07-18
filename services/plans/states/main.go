package pass_plan_states

import "github.com/leanspace/terraform-provider-leanspace/provider"

var PassPlanStateDataType = provider.DataSourceType[PassPlanState, *PassPlanState]{
	ResourceIdentifier: "leanspace_pass_plan_states",
	Path:               "plans-repository/pass-plans/states",
	Schema:             planStateSchema,
	FilterSchema:       nil,
}
