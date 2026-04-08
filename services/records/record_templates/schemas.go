package record_templates

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

var validFileTypes = []string{"CSV"}

var recordTemplateSchema = general_objects.ResourceSchemaWith(map[string]resourceschema.Attribute{
	"name": resourceschema.StringAttribute{
		Required:      true,
		PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
	},
	"description": resourceschema.StringAttribute{
		Optional:      true,
		PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
	},
	"stream_id": resourceschema.StringAttribute{
		Optional:      true,
		PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
	},
	"default_parsers": resourceschema.SetNestedAttribute{
		Computed: true,
		NestedObject: resourceschema.NestedAttributeObject{
			Attributes: recordTemplateDefaultParserSchema,
		},
	},
	"node_ids": resourceschema.SetAttribute{
		ElementType: types.StringType,
		Optional:    true,
		Validators:  []validator.Set{setvalidator.ValueStringsAre(helper.ValidUUID()...)},
	},
	"metric_ids": resourceschema.SetAttribute{
		ElementType: types.StringType,
		Optional:    true,
		Validators:  []validator.Set{setvalidator.ValueStringsAre(helper.ValidUUID()...)},
	},
	"command_definition_ids": resourceschema.SetAttribute{
		ElementType: types.StringType,
		Optional:    true,
		Validators:  []validator.Set{setvalidator.ValueStringsAre(helper.ValidUUID()...)},
	},
	"properties": resourceschema.SetNestedAttribute{
		Optional: true,
		NestedObject: resourceschema.NestedAttributeObject{
			Attributes: recordTemplatePropertySchema,
		},
	},
	"tags": general_objects.KeyValuesSchema,
})

var recordTemplateDefaultParserSchema = map[string]resourceschema.Attribute{
	"id": resourceschema.StringAttribute{
		Computed: true,
	},
	"file_type": resourceschema.StringAttribute{
		Computed:    true,
		Description: helper.AllowedValuesToDescription(validFileTypes),
	},
}

var recordTemplatePropertySchema = map[string]resourceschema.Attribute{
	"name": resourceschema.StringAttribute{
		Required:      true,
		PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
	},
	"attributes": resourceschema.SingleNestedAttribute{
		Required: true,
		Attributes: general_objects.DefinitionAttributeSchema(
			[]string{"BINARY", "GEOPOINT", "TLE"}, // Attribute types not allowed in attributes
			nil,                                   // All fields are used
			false,                                 // Does not force recreation if the type changes
		),
	},
}

var dataSourceFilterSchema = map[string]datasourceschema.Attribute{
	"names": datasourceschema.ListAttribute{
		ElementType: types.StringType,
		Optional:    true,
		Description: "Only returns Record Templates who's name matches one of the provided values.",
	},
	"node_ids": datasourceschema.ListAttribute{
		ElementType: types.StringType,
		Optional:    true,
		Description: "Only returns Record Templates with at least one nodeId that matches one of the provided values.",
		Validators:  []validator.List{listvalidator.ValueStringsAre(helper.ValidUUID()...)},
	},
	"metric_ids": datasourceschema.ListAttribute{
		ElementType: types.StringType,
		Optional:    true,
		Description: "Only returns Record Templates with at least one metricId that matches one of the provided values.",
		Validators:  []validator.List{listvalidator.ValueStringsAre(helper.ValidUUID()...)},
	},
	"created_by": datasourceschema.ListAttribute{
		ElementType: types.StringType,
		Optional:    true,
		Description: "Filter on the user who created the RecordTemplate. If you have no wish to use this field as a filter, either provide a null value or remove the field.",
	},
	"last_modified_by": datasourceschema.ListAttribute{
		ElementType: types.StringType,
		Optional:    true,
		Description: "Filter on the user who last modified the RecordTemplate. If you have no wish to use this field as a filter, either provide a null value or remove the field.",
	},
	"from_created_at": datasourceschema.StringAttribute{
		Optional:    true,
		Description: "Filter on the RecordTemplate creation date. RecordTemplates with a creation date greater or equals than the filter value will be selected (if they are not excluded by other filters). If you have no wish to use this field as a filter, either provide a null value or remove the field.",
	},
	"to_created_at": datasourceschema.StringAttribute{
		Optional:    true,
		Description: "Filter on the RecordTemplate creation date. RecordTemplates with a creation date lower or equals than the filter value will be selected (if they are not excluded by other filters). If you have no wish to use this field as a filter, either provide a null value or remove the field.",
	},
	"from_last_modified_at": datasourceschema.StringAttribute{
		Optional:    true,
		Description: "Filter on the RecordTemplate last modification date. RecordTemplates with a last modification date greater or equals than the filter value will be selected (if they are not excluded by other filters). If you have no wish to use this field as a filter, either provide a null value or remove the field.",
	},
	"to_last_modified_at": datasourceschema.StringAttribute{
		Optional:    true,
		Description: "Filter on the RecordTemplate last modification date. RecordTemplates with a last modification date lower or equals than the filter value will be selected (if they are not excluded by other filters). If you have no wish to use this field as a filter, either provide a null value or remove the field.",
	},
	"tags": datasourceschema.ListAttribute{
		ElementType: types.StringType,
		Optional:    true,
	},
}
