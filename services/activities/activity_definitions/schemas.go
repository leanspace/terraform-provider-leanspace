package activity_definitions

import (
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/setvalidator"
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

var allowedMappingStatuses = []string{"IN_SYNC", "OUT_OF_SYNC"}

var activityDefinitionSchema = general_objects.ResourceSchemaWith(map[string]resourceschema.Attribute{
	"node_id": resourceschema.StringAttribute{
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
	"mapping_status": resourceschema.StringAttribute{
		Computed:    true,
		Description: helper.AllowedValuesToDescription(allowedMappingStatuses),
	},
	"estimated_duration": resourceschema.Int64Attribute{
		Optional:   true,
		Validators: []validator.Int64{int64validator.AtLeast(0)},
	},
	"metadata": resourceschema.SetNestedAttribute{
		Optional: true,
		NestedObject: resourceschema.NestedAttributeObject{
			Attributes: metadataSchema,
		},
	},
	"argument_definitions": resourceschema.SetNestedAttribute{
		Optional: true,
		NestedObject: resourceschema.NestedAttributeObject{
			Attributes: argumentDefinitionSchema,
		},
	},
	"command_mappings": resourceschema.ListNestedAttribute{
		Optional: true,
		NestedObject: resourceschema.NestedAttributeObject{
			Attributes: commandMappingSchema,
		},
	},
	"tags": general_objects.KeyValuesSchema,
})

var metadataSchema = map[string]resourceschema.Attribute{
	"name": resourceschema.StringAttribute{
		Required: true,
	},
	"description": resourceschema.StringAttribute{
		Optional: true,
	},
	"attributes": resourceschema.SingleNestedAttribute{
		Required:   true,
		Attributes: general_objects.ValueAttributeSchema([]string{"ENUM", "STRUCTURE", "TLE"}),
	},
}

var argumentDefinitionSchema = map[string]resourceschema.Attribute{
	"name": resourceschema.StringAttribute{
		Required: true,
	},
	"description": resourceschema.StringAttribute{
		Optional: true,
	},
	"attributes": resourceschema.SingleNestedAttribute{
		Required: true,
		Attributes: general_objects.DefinitionAttributeSchema(
			[]string{"STRUCTURE", "TLE"}, // attribute types not allowed in command definition attributes
			nil,                          // All fields are used
			false,                        // Does not force recreation if the type changes
		),
	},
}

var commandMappingSchema = map[string]resourceschema.Attribute{
	"command_definition_id": resourceschema.StringAttribute{
		Required:   true,
		Validators: helper.ValidUUID(),
	},
	"position": resourceschema.Int64Attribute{
		Computed: true,
	},
	"delay_in_milliseconds": resourceschema.Int64Attribute{
		Required:    true,
		Description: "Delay to execute this command",
		Validators:  []validator.Int64{int64validator.AtLeast(0)},
	},
	"argument_mappings": resourceschema.SetNestedAttribute{
		Optional: true,
		NestedObject: resourceschema.NestedAttributeObject{
			Attributes: argumentMappingSchema,
		},
	},
	"metadata_mappings": resourceschema.SetNestedAttribute{
		Optional: true,
		NestedObject: resourceschema.NestedAttributeObject{
			Attributes: metadataMappingSchema,
		},
	},
}

var argumentMappingSchema = map[string]resourceschema.Attribute{
	"activity_definition_argument_name": resourceschema.StringAttribute{
		Required: true,
	},
	"command_definition_argument_name": resourceschema.StringAttribute{
		Required: true,
	},
	"mapping_status": resourceschema.StringAttribute{
		Computed:    true,
		Description: helper.AllowedValuesToDescription(allowedMappingStatuses),
	},
}

var metadataMappingSchema = map[string]resourceschema.Attribute{
	"activity_definition_metadata_name": resourceschema.StringAttribute{
		Required: true,
	},
	"command_definition_argument_name": resourceschema.StringAttribute{
		Required: true,
	},
	"mapping_status": resourceschema.StringAttribute{
		Computed:    true,
		Description: helper.AllowedValuesToDescription(allowedMappingStatuses),
	},
}

var dataSourceFilterSchema = map[string]datasourceschema.Attribute{
	"node_ids": datasourceschema.ListAttribute{
		ElementType: types.StringType,
		Optional:    true,
		Validators: []validator.List{
			listvalidator.ValueStringsAre(helper.ValidUUID()...),
		},
	},
	"with_arguments_metadata_and_command_mappings": datasourceschema.BoolAttribute{
		Optional:    true,
		Description: "Whether to include arguments, metadata and command mappings in the response. Setting this to true can significantly increase the response time.",
	},
	"mapping_statuses": datasourceschema.SetAttribute{
		ElementType: types.StringType,
		Optional:    true,
		Description: helper.AllowedValuesToDescription(allowedMappingStatuses),
		Validators: []validator.Set{
			setvalidator.ValueStringsAre(stringvalidator.OneOf(allowedMappingStatuses...)),
		},
	},
}
