package command_queues

import (
	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/setvalidator"
	datasourceschema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	resourceschema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/leanspace/terraform-provider-leanspace/helper"
	"github.com/leanspace/terraform-provider-leanspace/helper/general_objects"
)

var commandQueueSchema = general_objects.ResourceSchemaWith(map[string]resourceschema.Attribute{
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
})

var dataSourceFilterSchema = map[string]datasourceschema.Attribute{
	"asset_ids": datasourceschema.ListAttribute{
		ElementType: types.StringType,
		Optional:    true,
		Validators:  []validator.List{listvalidator.ValueStringsAre(helper.ValidUUID()...)},
	},
	"ground_station_ids": datasourceschema.ListAttribute{
		ElementType: types.StringType,
		Optional:    true,
		Validators:  []validator.List{listvalidator.ValueStringsAre(helper.ValidUUID()...)},
	},
	"command_transformer_plugin_ids": datasourceschema.ListAttribute{
		ElementType: types.StringType,
		Optional:    true,
		Validators:  []validator.List{listvalidator.ValueStringsAre(helper.ValidUUID()...)},
	},
	"protocol_transformer_plugin_ids": datasourceschema.ListAttribute{
		ElementType: types.StringType,
		Optional:    true,
		Validators:  []validator.List{listvalidator.ValueStringsAre(helper.ValidUUID()...)},
	},
}
