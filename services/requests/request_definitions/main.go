package request_definitions

import "github.com/leanspace/terraform-provider-leanspace/provider"

var RequestDefinitionDataType = provider.DataSourceType[RequestDefinition, *RequestDefinition]{
	ResourceIdentifier: "leanspace_request_definitions",
	Path:               "requests-repository/request-definitions",
	Schema:             requestDefinitionSchema,
	FilterSchema:       requestDefinitionFilterSchema,
}
