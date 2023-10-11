package contact_states

import "github.com/leanspace/terraform-provider-leanspace/provider"

var ContactStateDataType = provider.DataSourceType[ContactState, *ContactState]{
	ResourceIdentifier: "leanspace_contact_states",
	Path:               "passes-repository/contacts/states",
	Schema:             contactStateSchema,
	FilterSchema:       nil,
}
