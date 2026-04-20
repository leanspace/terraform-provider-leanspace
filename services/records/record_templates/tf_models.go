package record_templates

import (
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/leanspace/terraform-provider-leanspace/helper"
	"github.com/leanspace/terraform-provider-leanspace/helper/general_objects"
)

var defaultParserAttrTypes = map[string]attr.Type{
	"id":        types.StringType,
	"file_type": types.StringType,
}

type RecordTemplateTF struct {
	general_objects.AuditModelTF
	Name                 types.String                 `tfsdk:"name"`
	Description          types.String                 `tfsdk:"description"`
	StreamId             types.String                 `tfsdk:"stream_id"`
	DefaultParsers       types.Set                    `tfsdk:"default_parsers"`
	NodeIds              []types.String               `tfsdk:"node_ids"`
	MetricIds            []types.String               `tfsdk:"metric_ids"`
	CommandDefinitionIds []types.String               `tfsdk:"command_definition_ids"`
	Properties           []PropertyTF                 `tfsdk:"properties"`
	Tags                 []general_objects.KeyValueTF `tfsdk:"tags"`
}

type PropertyTF struct {
	Name       types.String                           `tfsdk:"name"`
	Attributes *general_objects.DefinitionAttributeTF `tfsdk:"attributes"`
}

func (x *RecordTemplate) ToTF() any {
	dpElems := make([]attr.Value, len(x.DefaultParsers))
	for i, dp := range x.DefaultParsers {
		dpObj, _ := types.ObjectValue(defaultParserAttrTypes, map[string]attr.Value{
			"id":        types.StringValue(dp.ID),
			"file_type": types.StringValue(dp.FileType),
		})
		dpElems[i] = dpObj
	}
	defaultParsers := types.SetValueMust(types.ObjectType{AttrTypes: defaultParserAttrTypes}, dpElems)

	properties := make([]PropertyTF, len(x.Properties))
	for i, p := range x.Properties {
		attrVal := general_objects.DefinitionAttributeToTF(&p.Attributes)
		properties[i] = PropertyTF{
			Name:       types.StringValue(p.Name),
			Attributes: &attrVal,
		}
	}

	return &RecordTemplateTF{
		AuditModelTF:         general_objects.AuditModelToTF(&x.AuditModel),
		Name:                 types.StringValue(x.Name),
		Description:          helper.TFStringPtrValue(x.Description),
		StreamId:             types.StringValue(x.StreamId),
		DefaultParsers:       defaultParsers,
		NodeIds:              helper.TFStringsValue(x.NodeIds),
		MetricIds:            helper.TFStringsValue(x.MetricIds),
		CommandDefinitionIds: helper.TFStringsValue(x.CommandDefinitionIds),
		Properties:           properties,
		Tags:                 general_objects.KeyValuesToTF(x.Tags),
	}
}

func (tf *RecordTemplateTF) ToAPI() any {
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
		Description:          helper.FromTFStringPtr(tf.Description),
		StreamId:             helper.FromTFString(tf.StreamId),
		NodeIds:              helper.FromTFStrings(tf.NodeIds),
		MetricIds:            helper.FromTFStrings(tf.MetricIds),
		CommandDefinitionIds: helper.FromTFStrings(tf.CommandDefinitionIds),
		Properties:           properties,
		Tags:                 general_objects.KeyValuesFromTF(tf.Tags),
	}
}
