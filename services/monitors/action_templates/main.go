package action_templates

import "github.com/leanspace/terraform-provider-leanspace/provider"

var ActionTemplateDataType = provider.DataSourceType[ActionTemplate, *ActionTemplate]{
	ResourceIdentifier: "leanspace_action_templates",
	Path:               "monitors-repository/action-templates",
	Schema:             ActionTemplateSchema,
	FilterSchema:       dataSourceFilterSchema,
}
