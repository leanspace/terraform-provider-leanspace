package plan_templates

import "github.com/leanspace/terraform-provider-leanspace/provider"

var PlanTemplateDataType = provider.DataSourceType[PlanTemplate, *PlanTemplate]{
	ResourceIdentifier: "leanspace_plan_templates",
	Path:               "plans-repository/plan-templates",
	Schema:             planTemplateSchema,
	FilterSchema:       nil,
}
