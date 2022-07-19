package action_templates

import "leanspace-terraform-provider/provider"

var ActionTemplateDataType = provider.DataSourceType[ActionTemplate, *ActionTemplate]{
	ResourceIdentifier: "leanspace_action_templates",
	Name:               "action_template",
	Path:               "monitors-repository/action-templates",
	Schema:             actionTemplateSchema,
}
