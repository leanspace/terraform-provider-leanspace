package analysis_definitions

import "leanspace-terraform-provider/provider"

var AnalysisDefinitionDataType = provider.DataSourceType[AnalysisDefinition, *AnalysisDefinition]{
	ResourceIdentifier: "leanspace_analysis_definitions",
	Path:               "analyses-repository/analysis-definitions",
	Schema:             analysisDefinitionSchema,
	FilterSchema:       dataSourceFilterSchema,
}
