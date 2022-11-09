package states

import "github.com/leanspace/terraform-provider-leanspace/provider"

var StateDataType = provider.DataSourceType[State, *State]{
	ResourceIdentifier: "leanspace_states",
	Path:               "plans-repository/pass-plans/states",
	Schema:             stateSchema,
	FilterSchema:       nil,
}
