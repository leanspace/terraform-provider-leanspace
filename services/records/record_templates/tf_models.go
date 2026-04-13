package record_templates

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/leanspace/terraform-provider-leanspace/helper"
	"github.com/leanspace/terraform-provider-leanspace/helper/general_objects"
)

type RecordTemplateTF struct {
	general_objects.AuditModelTF
	Name                 types.String                 `tfsdk:"name"`
	Description          types.String                 `tfsdk:"description"`
	StreamId             types.String                 `tfsdk:"stream_id"`
	DefaultParsers       []DefaultParserTF            `tfsdk:"default_parsers"`
	NodeIds              []types.String               `tfsdk:"node_ids"`
	MetricIds            []types.String               `tfsdk:"metric_ids"`
	CommandDefinitionIds []types.String               `tfsdk:"command_definition_ids"`
	Properties           []PropertyTF                 `tfsdk:"properties"`
	Tags                 []general_objects.KeyValueTF `tfsdk:"tags"`
}

type DefaultParserTF struct {
	ID       types.String `tfsdk:"id"`
	FileType types.String `tfsdk:"file_type"`
}

type PropertyTF struct {
	Name       types.String                           `tfsdk:"name"`
	Attributes *general_objects.DefinitionAttributeTF `tfsdk:"attributes"`
}

func (x *RecordTemplate) ToTF() any {
	defaultParsers := make([]DefaultParserTF, len(x.DefaultParsers))
	for i, dp := range x.DefaultParsers {
		defaultParsers[i] = DefaultParserTF{
			ID:       helper.TFStringValue(dp.ID),
			FileType: helper.TFStringValue(dp.FileType),
		}
	}

	properties := make([]PropertyTF, len(x.Properties))
	for i, p := range x.Properties {
		attr := general_objects.DefinitionAttributeToTF(&p.Attributes)
		properties[i] = PropertyTF{
			Name:       helper.TFStringValue(p.Name),
			Attributes: &attr,
		}
	}

	return &RecordTemplateTF{
		AuditModelTF:         general_objects.AuditModelToTF(&x.AuditModel),
		Name:                 helper.TFStringValue(x.Name),
		Description:          helper.TFStringValue(x.Description),
		StreamId:             helper.TFStringValue(x.StreamId),
		DefaultParsers:       defaultParsers,
		NodeIds:              helper.TFStringsValue(x.NodeIds),
		MetricIds:            helper.TFStringsValue(x.MetricIds),
		CommandDefinitionIds: helper.TFStringsValue(x.CommandDefinitionIds),
		Properties:           properties,
		Tags:                 general_objects.KeyValuesToTF(x.Tags),
	}
}

func (tf *RecordTemplateTF) ToAPI() any {
	defaultParsers := make([]DefaultParser, len(tf.DefaultParsers))
	for i, dp := range tf.DefaultParsers {
		defaultParsers[i] = DefaultParser{
			ID:       helper.FromTFString(dp.ID),
			FileType: helper.FromTFString(dp.FileType),
		}
	}

	properties := make([]Property[any], len(tf.Properties))
	for i, p := range tf.Properties {
		properties[i] = Property[any]{
			Name: helper.FromTFString(p.Name),
		}
		if p.Attributes != nil {
			properties[i].Attributes = general_objects.DefinitionAttributeFromTF(*p.Attributes)
		}
	}

	return &RecordTemplate{
		AuditModel:           general_objects.AuditModelFromTF(tf.AuditModelTF),
		Name:                 helper.FromTFString(tf.Name),
		Description:          helper.FromTFString(tf.Description),
		StreamId:             helper.FromTFString(tf.StreamId),
		DefaultParsers:       defaultParsers,
		NodeIds:              helper.FromTFStrings(tf.NodeIds),
		MetricIds:            helper.FromTFStrings(tf.MetricIds),
		CommandDefinitionIds: helper.FromTFStrings(tf.CommandDefinitionIds),
		Properties:           properties,
		Tags:                 general_objects.KeyValuesFromTF(tf.Tags),
	}
}
