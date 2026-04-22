package metrics

import (
	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	datasourceschema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	resourceschema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/leanspace/terraform-provider-leanspace/helper"
	"github.com/leanspace/terraform-provider-leanspace/helper/general_objects"
)

var metricSchema = general_objects.ResourceSchemaWith(map[string]resourceschema.Attribute{
	"node_id": resourceschema.StringAttribute{
		Required:      true,
		Validators:    helper.ValidUUID(),
		PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
	},
	"ancestor_asset_id": resourceschema.StringAttribute{
		Computed: true,
	},
	"name": resourceschema.StringAttribute{
		Required: true,
	},
	"description": resourceschema.StringAttribute{
		Optional: true,
	},
	"attributes": resourceschema.SingleNestedAttribute{
		Required: true,
		Attributes: general_objects.DefinitionAttributeSchema(
			[]string{"TIME", "ARRAY"}, // TIME and ARRAY type not allowed
			[]string{"required", "default_value", "min_size", "max_size", "unique", "constraint"}, // Fields unused
			true, // Force recreation if the type changes
		),
	},
	"tags": general_objects.KeyValuesSchema,
})

var dataSourceFilterSchema = map[string]datasourceschema.Attribute{
	"node_ids": datasourceschema.ListAttribute{
		ElementType: types.StringType,
		Optional:    true,
		Validators:  []validator.List{listvalidator.ValueStringsAre(helper.ValidUUID()...)},
	},
	"ancestor_node_ids": datasourceschema.ListAttribute{
		ElementType: types.StringType,
		Optional:    true,
		Validators:  []validator.List{listvalidator.ValueStringsAre(helper.ValidUUID()...)},
	},
	"attribute_types": datasourceschema.ListAttribute{
		ElementType: types.StringType,
		Optional:    true,
	},
	"tags": datasourceschema.ListAttribute{
		ElementType: types.StringType,
		Optional:    true,
	},
}
