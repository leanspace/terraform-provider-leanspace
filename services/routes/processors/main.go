package processors

import "github.com/leanspace/terraform-provider-leanspace/provider"

var ProcessorDataType = provider.DataSourceType[Processor, *Processor]{
	ResourceIdentifier: "leanspace_processors",
	Path:               "routes-repository/processors",
	Schema:             processorSchema,
	FilterSchema:       dataSourceFilterSchema,
	NewTFModel:         func() any { return &ProcessorTF{} },
}
