package command_queues

import (
	"github.com/hashicorp/terraform-plugin-framework-validators/setvalidator"
	datasourceschema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	resourceschema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/leanspace/terraform-provider-leanspace/helper"
)

var commandQueueSchema = map[string]resourceschema.Attribute{
	"id": resourceschema.StringAttribute{
		Computed: true,
	},
	"asset_id": resourceschema.StringAttribute{
		Required:      true,
		Validators:    helper.ValidUUID(),
		PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
	},
	"name": resourceschema.StringAttribute{
		Required: true,
	},
	"ground_station_ids": resourceschema.SetAttribute{
		ElementType: types.StringType,
		Optional:    true,
		Validators: []validator.Set{
			setvalidator.ValueStringsAre(helper.ValidUUID()...),
		},
	},
	"command_transformer_plugin_id": resourceschema.StringAttribute{
		Optional:    true,
		Description: "The Id of the Command Transformer's Plugin",
		Validators:  helper.ValidUUID(),
	},
	"protocol_transformer_plugin_id": resourceschema.StringAttribute{
		Optional:    true,
		Description: "The Id of the Protocol Transformer's Plugin",
		Validators:  helper.ValidUUID(),
	},
	"protocol_transformer_init_data": resourceschema.StringAttribute{
		Optional:    true,
		Description: "Initialization data used by the Protocol Transformer",
	},
	"created_at": resourceschema.StringAttribute{
		Computed:    true,
		Description: "When it was created",
	},
	"created_by": resourceschema.StringAttribute{
		Computed:    true,
		Description: "Who created it",
	},
	"last_modified_at": resourceschema.StringAttribute{
		Computed:    true,
		Description: "When it was last modified",
	},
	"last_modified_by": resourceschema.StringAttribute{
		Computed:    true,
		Description: "Who modified it the last",
	},
}

var dataSourceFilterSchema = map[string]datasourceschema.Attribute{
	"asset_ids": datasourceschema.ListAttribute{
		ElementType: types.StringType,
		Optional:    true,
	},
	"ground_station_ids": datasourceschema.ListAttribute{
		ElementType: types.StringType,
		Optional:    true,
	},
	"command_transformer_plugin_ids": datasourceschema.ListAttribute{
		ElementType: types.StringType,
		Optional:    true,
	},
	"protocol_transformer_plugin_ids": datasourceschema.ListAttribute{
		ElementType: types.StringType,
		Optional:    true,
	},
}
