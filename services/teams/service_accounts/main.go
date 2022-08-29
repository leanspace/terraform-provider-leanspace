package service_accounts

import "github.com/leanspace/terraform-provider-leanspace/provider"

var ServiceAccountDataType = provider.DataSourceType[ServiceAccount, *ServiceAccount]{
	ResourceIdentifier: "leanspace_service_accounts",
	Path:               "teams-repository/service-accounts",
	Schema:             serviceAccountSchema,
}
