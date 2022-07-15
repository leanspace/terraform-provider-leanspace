package access_policies

import (
	"terraform-provider-asset/asset"
)

var AccessPolicyDataType = asset.DataSourceType[AccessPolicy, *AccessPolicy]{
	ResourceIdentifier: "leanspace_access_policies",
	Name:               "access_policy",
	Path:               "teams-repository/access-policies",
	Schema:             accessPolicySchema,
}
