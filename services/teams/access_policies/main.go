package access_policies

import "leanspace-terraform-provider/provider"

var AccessPolicyDataType = provider.DataSourceType[AccessPolicy, *AccessPolicy]{
	ResourceIdentifier: "leanspace_access_policies",
	Name:               "access_policy",
	Path:               "teams-repository/access-policies",
	Schema:             accessPolicySchema,
}
