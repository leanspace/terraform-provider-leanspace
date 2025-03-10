package contact_reservation_status_mappings

import (
	"github.com/leanspace/terraform-provider-leanspace/provider"
)

var LeafSpaceContactReservationStatusMappingDataType = provider.DataSourceType[ContactReservationStatusMapping, *ContactReservationStatusMapping]{
	ResourceIdentifier: "leanspace_leaf_space_contact_reservation_status_mappings",
	Path:               "integration-leafspace/contact-reservations/status/mappings",
	Schema:             contactReservationStatusMappingSchema,
	FilterSchema:       dataSourceFilterSchema,
}
