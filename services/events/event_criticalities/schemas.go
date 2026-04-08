package event_criticalities

import (
	datasourceschema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	resourceschema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/leanspace/terraform-provider-leanspace/helper/general_objects"
)

var EventCriticalitiesSchema = general_objects.ResourceSchemaWith(map[string]resourceschema.Attribute{
	"name": resourceschema.StringAttribute{
		Required: true,
	},
	"read_only": resourceschema.BoolAttribute{
		Computed:    true,
		Description: "Indicates if the object is read-only",
	},
	"tags": general_objects.KeyValuesSchema,
})

var dataSourceFilterSchema = map[string]datasourceschema.Attribute{
	"tags": datasourceschema.ListAttribute{
		ElementType: types.StringType,
		Optional:    true,
	},
}
