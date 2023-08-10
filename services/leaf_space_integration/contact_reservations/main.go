package contact_reservations

import (
	"github.com/leanspace/terraform-provider-leanspace/provider"
)

var LeafSpaceContactReservation = provider.DataSourceType[contactReservation, *contactReservation]{
	ResourceIdentifier: "leanspace_leaf_space_contact_reservations",
	Path:               "integration-leafspace/contact-reservations/status/mappings",
	Schema:             contactReservationSchema,
	FilterSchema:       dataSourceFilterSchema,
}
