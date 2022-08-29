package analysis_definitions

import "github.com/leanspace/terraform-provider-leanspace/provider"

var AnalysisDefinitionDataType = provider.DataSourceType[AnalysisDefinition, *AnalysisDefinition]{
	ResourceIdentifier: "leanspace_analysis_definitions",
	Path:               "analyses-repository/analysis-definitions",
	Schema:             analysisDefinitionSchema,
	FilterSchema:       dataSourceFilterSchema,
}
