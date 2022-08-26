package access_policies

import "leanspace-terraform-provider/provider"

var AccessPolicyDataType = provider.DataSourceType[AccessPolicy, *AccessPolicy]{
	ResourceIdentifier: "leanspace_access_policies",
	Path:               "teams-repository/access-policies",
	Schema:             accessPolicySchema,
	FilterSchema:       dataSourceFilterSchema,
}
