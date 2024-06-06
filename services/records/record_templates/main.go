package record_templates

import "github.com/leanspace/terraform-provider-leanspace/provider"

var RecordTemplateDataType = provider.DataSourceType[RecordTemplate, *RecordTemplate]{
	ResourceIdentifier: "leanspace_record_templates",
	Path:               "records/record-templates",
	Schema:             recordTemplateSchema,
	FilterSchema:       dataSourceFilterSchema,
}
