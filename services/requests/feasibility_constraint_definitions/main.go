package feasibility_constraint_definitions

import "github.com/leanspace/terraform-provider-leanspace/provider"

var FeasibilityConstraintDefinitionDataType = provider.DataSourceType[FeasibilityConstraintDefinition, *FeasibilityConstraintDefinition]{
	ResourceIdentifier: "leanspace_feasibility_constraint_definitions",
	Path:               "requests-repository/feasibility-constraint-definitions",
	Schema:             feasibilityConstraintDefinitionSchema,
	FilterSchema:       feasibilityConstraintDefinitionFilterSchema,
}
