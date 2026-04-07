package event_criticalities

import (

	resourceschema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	datasourceschema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/leanspace/terraform-provider-leanspace/helper/general_objects"
)


var EventCriticalitiesSchema = map[string]resourceschema.Attribute{
	"id": resourceschema.StringAttribute{
		Computed: true,
	},
	"name": resourceschema.StringAttribute{
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
	"read_only": resourceschema.BoolAttribute{
		Computed: true,
		Description: "Indicates if the object is read-only",
	},
	"tags": general_objects.KeyValuesSchema,
}

var dataSourceFilterSchema = map[string]datasourceschema.Attribute{
	"tags": datasourceschema.ListAttribute{
		ElementType: types.StringType,
		Optional: true,
	},
}
