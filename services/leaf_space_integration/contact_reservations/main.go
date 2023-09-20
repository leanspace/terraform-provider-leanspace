package leaf_space_contact_reservations

import (
	"github.com/leanspace/terraform-provider-leanspace/provider"
)

var LeafSpaceContactReservationDataType = provider.DataSourceType[ContactReservation, *ContactReservation]{
	ResourceIdentifier: "leanspace_leaf_space_contact_reservations",
	Path:               "integration-leafspace/contact-reservations/status/mappings",
	Schema:             contactReservationSchema,
	FilterSchema:       dataSourceFilterSchema,
}
