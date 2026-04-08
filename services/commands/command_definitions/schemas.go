package command_definitions

import (
	datasourceschema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	resourceschema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/leanspace/terraform-provider-leanspace/helper"
	"github.com/leanspace/terraform-provider-leanspace/helper/general_objects"
)

var commandDefinitionSchema = general_objects.ResourceSchemaWith(map[string]resourceschema.Attribute{
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
	"identifier": resourceschema.StringAttribute{
		Optional: true,
	},
	"metadata": resourceschema.SetNestedAttribute{
		Optional: true,
		NestedObject: resourceschema.NestedAttributeObject{
			Attributes: metadataSchema,
		},
	},
	"arguments": resourceschema.SetNestedAttribute{
		Optional: true,
		NestedObject: resourceschema.NestedAttributeObject{
			Attributes: argumentSchema,
		},
	},
})

var metadataSchema = map[string]resourceschema.Attribute{
	"id": resourceschema.StringAttribute{
		Computed: true,
	},
	"name": resourceschema.StringAttribute{
		Required: true,
	},
	"description": resourceschema.StringAttribute{
		Optional: true,
	},
	"attributes": resourceschema.ListNestedAttribute{
		Required: true,
		NestedObject: resourceschema.NestedAttributeObject{
			Attributes: general_objects.ValueAttributeSchema([]string{"ENUM", "STRUCTURE", "GEOPOINT", "TLE", "BINARY", "ARRAY"}),
		},
	},
}

var argumentSchema = map[string]resourceschema.Attribute{
	"id": resourceschema.StringAttribute{
		Computed: true,
	},
	"name": resourceschema.StringAttribute{
		Required: true,
	},
	"identifier": resourceschema.StringAttribute{
		Optional: true,
	},
	"description": resourceschema.StringAttribute{
		Optional: true,
	},
	"attributes": resourceschema.SingleNestedAttribute{
		Required: true,
		Attributes: general_objects.DefinitionAttributeSchema(
			[]string{"STRUCTURE", "GEOPOINT", "TLE"}, // attribute types not allowed in command definition attributes
			nil,                                      // All fields are used
			false,                                    // Does not force recreation if the type changes
		),
	},
}

var dataSourceFilterSchema = map[string]datasourceschema.Attribute{
	"node_ids": datasourceschema.ListAttribute{
		ElementType: types.StringType,
		Optional:    true,
	},
	"node_types": datasourceschema.ListAttribute{
		ElementType: types.StringType,
		Optional:    true,
		Description: "Filter on the Node type. Allowed values : GROUP, ASSET, COMPONENT",
	},
	"node_kinds": datasourceschema.ListAttribute{
		ElementType: types.StringType,
		Optional:    true,
		Description: "Filter on the Node kind. Allowed values : GENERIC, SATELLITE, GROUND_STATION",
	},
	"with_arguments_and_metadata": datasourceschema.BoolAttribute{
		Optional: true,
	},
	"created_bys": datasourceschema.ListAttribute{
		ElementType: types.StringType,
		Optional:    true,
		Description: "Filter on the user who created the Node. If you have no wish to use this field as a filter, either provide a null value or remove the field.",
	},
}
