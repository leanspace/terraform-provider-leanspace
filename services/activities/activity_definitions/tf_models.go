package activity_definitions

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/leanspace/terraform-provider-leanspace/helper"
	"github.com/leanspace/terraform-provider-leanspace/helper/general_objects"
)

type ActivityDefinitionTF struct {
	general_objects.AuditModelTF
	NodeId              types.String                 `tfsdk:"node_id"`
	Name                types.String                 `tfsdk:"name"`
	Description         types.String                 `tfsdk:"description"`
	EstimatedDuration   types.Int64                  `tfsdk:"estimated_duration"`
	MappingStatus       types.String                 `tfsdk:"mapping_status"`
	Metadata            []MetadataTF                 `tfsdk:"metadata"`
	ArgumentDefinitions []ArgumentDefinitionTF       `tfsdk:"argument_definitions"`
	CommandMappings     []CommandMappingTF           `tfsdk:"command_mappings"`
	Tags                []general_objects.KeyValueTF `tfsdk:"tags"`
}

type MetadataTF struct {
	Name        types.String                      `tfsdk:"name"`
	Description types.String                      `tfsdk:"description"`
	Attributes  *general_objects.ValueAttributeTF `tfsdk:"attributes"`
}

type ArgumentDefinitionTF struct {
	Name        types.String                           `tfsdk:"name"`
	Description types.String                           `tfsdk:"description"`
	Attributes  *general_objects.DefinitionAttributeTF `tfsdk:"attributes"`
}

type CommandMappingTF struct {
	CommandDefinitionId types.String        `tfsdk:"command_definition_id"`
	Position            types.Int64         `tfsdk:"position"`
	DelayInMilliseconds types.Int64         `tfsdk:"delay_in_milliseconds"`
	ArgumentMappings    []ArgumentMappingTF `tfsdk:"argument_mappings"`
	MetadataMappings    []MetadataMappingTF `tfsdk:"metadata_mappings"`
}

type ArgumentMappingTF struct {
	ActivityDefinitionArgumentName types.String `tfsdk:"activity_definition_argument_name"`
	CommandDefinitionArgumentName  types.String `tfsdk:"command_definition_argument_name"`
	MappingStatus                  types.String `tfsdk:"mapping_status"`
}

type MetadataMappingTF struct {
	ActivityDefinitionMetadataName types.String `tfsdk:"activity_definition_metadata_name"`
	CommandDefinitionArgumentName  types.String `tfsdk:"command_definition_argument_name"`
	MappingStatus                  types.String `tfsdk:"mapping_status"`
}

func (x *ActivityDefinition) ToTF() any {
	metadata := make([]MetadataTF, len(x.Metadata))
	for i, m := range x.Metadata {
		attr := general_objects.ValueAttributeToTF(&m.Attributes)
		metadata[i] = MetadataTF{
			Name:        types.StringValue(m.Name),
			Description: helper.TFStringPtrValue(m.Description),
			Attributes:  &attr,
		}
	}

	argDefs := make([]ArgumentDefinitionTF, len(x.ArgumentDefinitions))
	for i, a := range x.ArgumentDefinitions {
		attr := general_objects.DefinitionAttributeToTF(&a.Attributes)
		argDefs[i] = ArgumentDefinitionTF{
			Name:        types.StringValue(a.Name),
			Description: helper.TFStringPtrValue(a.Description),
			Attributes:  &attr,
		}
	}

	cmdMappings := make([]CommandMappingTF, len(x.CommandMappings))
	for i, cm := range x.CommandMappings {
		argMappings := make([]ArgumentMappingTF, len(cm.ArgumentMappings))
		for j, am := range cm.ArgumentMappings {
			argMappings[j] = ArgumentMappingTF{
				ActivityDefinitionArgumentName: types.StringValue(am.ActivityDefinitionArgumentName),
				CommandDefinitionArgumentName:  types.StringValue(am.CommandDefinitionArgumentName),
				MappingStatus:                  helper.TFStringPtrValue(am.MappingStatus),
			}
		}
		metaMappings := make([]MetadataMappingTF, len(cm.MetadataMappings))
		for j, mm := range cm.MetadataMappings {
			metaMappings[j] = MetadataMappingTF{
				ActivityDefinitionMetadataName: types.StringValue(mm.ActivityDefinitionMetadataName),
				CommandDefinitionArgumentName:  types.StringValue(mm.CommandDefinitionArgumentName),
				MappingStatus:                  helper.TFStringPtrValue(mm.MappingStatus),
			}
		}
		cmdMappings[i] = CommandMappingTF{
			CommandDefinitionId: types.StringValue(cm.CommandDefinitionId),
			Position:            helper.TFInt64Value(cm.Position),
			DelayInMilliseconds: helper.TFInt64Value(cm.DelayInMilliseconds),
			ArgumentMappings:    argMappings,
			MetadataMappings:    metaMappings,
		}
	}

	return &ActivityDefinitionTF{
		AuditModelTF:        general_objects.AuditModelToTF(&x.AuditModel),
		NodeId:              types.StringValue(x.NodeId),
		Name:                types.StringValue(x.Name),
		Description:         helper.TFStringPtrValue(x.Description),
		EstimatedDuration:   helper.TFInt64Value(x.EstimatedDuration),
		MappingStatus:       helper.TFStringPtrValue(x.MappingStatus),
		Metadata:            metadata,
		ArgumentDefinitions: argDefs,
		CommandMappings:     cmdMappings,
		Tags:                general_objects.KeyValuesToTF(x.Tags),
	}
}

func (tf *ActivityDefinitionTF) ToAPI() any {
	metadata := make([]Metadata[any], len(tf.Metadata))
	for i, m := range tf.Metadata {
		metadata[i] = Metadata[any]{
			Name:        helper.FromTFString(m.Name),
			Description: helper.FromTFStringPtr(m.Description),
		}
		if m.Attributes != nil {
			metadata[i].Attributes = general_objects.ValueAttributeFromTF(*m.Attributes)
		}
	}

	argDefs := make([]ArgumentDefinition[any], len(tf.ArgumentDefinitions))
	for i, a := range tf.ArgumentDefinitions {
		argDefs[i] = ArgumentDefinition[any]{
			Name:        helper.FromTFString(a.Name),
			Description: helper.FromTFStringPtr(a.Description),
		}
		if a.Attributes != nil {
			argDefs[i].Attributes = general_objects.DefinitionAttributeFromTF(*a.Attributes)
		}
	}

	cmdMappings := make([]CommandMapping, len(tf.CommandMappings))
	for i, cm := range tf.CommandMappings {
		argMappings := make([]ArgumentMapping, len(cm.ArgumentMappings))
		for j, am := range cm.ArgumentMappings {
			argMappings[j] = ArgumentMapping{
				ActivityDefinitionArgumentName: helper.FromTFString(am.ActivityDefinitionArgumentName),
				CommandDefinitionArgumentName:  helper.FromTFString(am.CommandDefinitionArgumentName),
				MappingStatus:                  helper.FromTFStringPtr(am.MappingStatus),
			}
		}
		metaMappings := make([]MetadataMapping, len(cm.MetadataMappings))
		for j, mm := range cm.MetadataMappings {
			metaMappings[j] = MetadataMapping{
				ActivityDefinitionMetadataName: helper.FromTFString(mm.ActivityDefinitionMetadataName),
				CommandDefinitionArgumentName:  helper.FromTFString(mm.CommandDefinitionArgumentName),
				MappingStatus:                  helper.FromTFStringPtr(mm.MappingStatus),
			}
		}
		cmdMappings[i] = CommandMapping{
			CommandDefinitionId: helper.FromTFString(cm.CommandDefinitionId),
			Position:            helper.FromTFInt64(cm.Position),
			DelayInMilliseconds: helper.FromTFInt64(cm.DelayInMilliseconds),
			ArgumentMappings:    argMappings,
			MetadataMappings:    metaMappings,
		}
	}

	return &ActivityDefinition{
		AuditModel:          general_objects.AuditModelFromTF(tf.AuditModelTF),
		NodeId:              helper.FromTFString(tf.NodeId),
		Name:                helper.FromTFString(tf.Name),
		Description:         helper.FromTFStringPtr(tf.Description),
		EstimatedDuration:   helper.FromTFInt64(tf.EstimatedDuration),
		MappingStatus:       helper.FromTFStringPtr(tf.MappingStatus),
		Metadata:            metadata,
		ArgumentDefinitions: argDefs,
		CommandMappings:     cmdMappings,
		Tags:                general_objects.KeyValuesFromTF(tf.Tags),
	}
}
