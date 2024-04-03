package records

import "github.com/leanspace/terraform-provider-leanspace/provider"

var RecordDataType = provider.DataSourceType[Record, *Record]{
	ResourceIdentifier: "leanspace_records",
	Path:               "records/records",
	Schema:             recordSchema,
	FilterSchema:       dataSourceFilterSchema,
}
