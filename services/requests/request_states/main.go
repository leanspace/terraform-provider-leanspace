package request_states

import "github.com/leanspace/terraform-provider-leanspace/provider"

var RequestStateDataType = provider.DataSourceType[RequestState, *RequestState]{
	ResourceIdentifier: "leanspace_request_states",
	Path:               "requests-repository/requests/states",
	Schema:             requestStateSchema,
	FilterSchema:       requestStateFilterSchema,
}
