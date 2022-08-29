package access_policies

import "github.com/leanspace/terraform-provider-leanspace/provider"

var AccessPolicyDataType = provider.DataSourceType[AccessPolicy, *AccessPolicy]{
	ResourceIdentifier: "leanspace_access_policies",
	Path:               "teams-repository/access-policies",
	Schema:             accessPolicySchema,
	FilterSchema:       dataSourceFilterSchema,
}
