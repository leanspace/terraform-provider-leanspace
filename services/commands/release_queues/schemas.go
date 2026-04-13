package release_queues

import (
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	datasourceschema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	resourceschema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/leanspace/terraform-provider-leanspace/helper"
	"github.com/leanspace/terraform-provider-leanspace/helper/general_objects"
)

var validCommandTransformationStrategies = []string{"TEST", "NO_TRANSFORMATION", "USE_PLUGIN"}

var releaseQueueSchema = general_objects.ResourceSchemaWith(map[string]resourceschema.Attribute{
	"asset_id": resourceschema.StringAttribute{
		Required:      true,
		Validators:    helper.ValidUUID(),
		PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
	},
	"name": resourceschema.StringAttribute{
		Required: true,
	},
	"description": resourceschema.StringAttribute{
		Optional: true,
	},
	"command_transformer_plugin_id": resourceschema.StringAttribute{
		Optional:    true,
		Description: "The Id of the Command Transformer Plugin",
		Validators:  helper.ValidUUID(),
	},
	"command_transformation_strategy": resourceschema.StringAttribute{
		Optional:    true,
		Description: "What transformation strategy shall be applied on created and updated Commands",
		Validators:  []validator.String{stringvalidator.OneOf(validCommandTransformationStrategies...)},
	},
	"command_transformer_plugin_configuration_data": resourceschema.StringAttribute{
		Optional:    true,
		Description: "Configuration data used by the Command Transformer Plugin (coming soon)",
	},
	"global_transmission_metadata": general_objects.KeyValuesSchema,
	"logical_lock": resourceschema.BoolAttribute{
		Computed: true,
	},
	"tags": general_objects.KeyValuesSchema,
})

var dataSourceFilterSchema = map[string]datasourceschema.Attribute{
	"asset_ids": datasourceschema.ListAttribute{
		ElementType: types.StringType,
		Optional:    true,
	},
	"command_transformer_plugin_ids": datasourceschema.ListAttribute{
		ElementType: types.StringType,
		Optional:    true,
	},
	"command_transformation_strategy": datasourceschema.StringAttribute{
		Optional:    true,
		Description: "What transformation strategy shall be applied on created and updated Commands",
		Validators:  []validator.String{stringvalidator.OneOf(validCommandTransformationStrategies...)},
	},
	"logical_lock": datasourceschema.BoolAttribute{
		Optional: true,
	},
	"tags": datasourceschema.ListAttribute{
		ElementType: types.StringType,
		Optional:    true,
	},
}
