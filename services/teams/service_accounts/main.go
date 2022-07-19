package service_accounts

import "leanspace-terraform-provider/provider"

var ServiceAccountDataType = provider.DataSourceType[ServiceAccount, *ServiceAccount]{
	ResourceIdentifier: "leanspace_service_accounts",
	Name:               "service_account",
	Path:               "teams-repository/service-accounts",
	Schema:             serviceAccountSchema,
}
