package pass_states

import "github.com/leanspace/terraform-provider-leanspace/provider"

var PassStateDataType = provider.DataSourceType[PassState, *PassState]{
	ResourceIdentifier: "leanspace_pass_states",
	Path:               "passes-repository/passes/states",
	Schema:             passStateSchema,
	FilterSchema:       nil,
}
