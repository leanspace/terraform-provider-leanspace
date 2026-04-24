package contact_reservation_status_mappings

import (
	datasourceschema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	resourceschema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/leanspace/terraform-provider-leanspace/helper/general_objects"
)

var contactReservationStatusMappingSchema = general_objects.ResourceSchemaWith(map[string]resourceschema.Attribute{
	"contact_state_id": resourceschema.StringAttribute{
		Required: true,
	},
	"leafspace_status": resourceschema.StringAttribute{
		Required: true,
	},
})

var dataSourceFilterSchema = map[string]datasourceschema.Attribute{
	"leafspace_statuses": datasourceschema.ListAttribute{
		ElementType: types.StringType,
		Optional:    true,
	},
}
