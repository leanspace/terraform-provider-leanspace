package service_accounts

import (
	"terraform-provider-asset/asset"
)

var ServiceAccountDataType = asset.DataSourceType[ServiceAccount, *ServiceAccount]{
	ResourceIdentifier: "leanspace_service_accounts",
	Name:               "service_account",
	Path:               "teams-repository/service-accounts",
	Schema:             serviceAccountSchema,
}
