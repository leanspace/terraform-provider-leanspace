package contact_reservation_status_mappings

import (

	resourceschema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	datasourceschema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)


var contactReservationStatusMappingSchema = map[string]resourceschema.Attribute{
	"id": resourceschema.StringAttribute{
		Computed: true,
	},
	"contact_state_id": resourceschema.StringAttribute{
		Required: true,
	},
	"leafspace_status": resourceschema.StringAttribute{
		Required: true,
	},
	"created_at": resourceschema.StringAttribute{
		Computed: true,
		Description: "When it was created",
	},
	"created_by": resourceschema.StringAttribute{
		Computed: true,
		Description: "Who created it",
	},
	"last_modified_at": resourceschema.StringAttribute{
		Computed: true,
		Description: "When it was last modified",
	},
	"last_modified_by": resourceschema.StringAttribute{
		Computed: true,
		Description: "Who modified it the last",
	},
}

var dataSourceFilterSchema = map[string]datasourceschema.Attribute{
	"leafspace_statuses": datasourceschema.ListAttribute{
		ElementType: types.StringType,
		Optional: true,
	},
}
