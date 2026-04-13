package command_definitions

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/leanspace/terraform-provider-leanspace/helper"
	"github.com/leanspace/terraform-provider-leanspace/helper/general_objects"
)

type CommandDefinitionTF struct {
	general_objects.AuditModelTF
	NodeId      types.String `tfsdk:"node_id"`
	Name        types.String `tfsdk:"name"`
	Description types.String `tfsdk:"description"`
	Identifier  types.String `tfsdk:"identifier"`
	Metadata    []MetadataTF `tfsdk:"metadata"`
	Arguments   []ArgumentTF `tfsdk:"arguments"`
}

type MetadataTF struct {
	ID          types.String                      `tfsdk:"id"`
	Name        types.String                      `tfsdk:"name"`
	Description types.String                      `tfsdk:"description"`
	Attributes  *general_objects.ValueAttributeTF `tfsdk:"attributes"`
}

type ArgumentTF struct {
	ID          types.String                           `tfsdk:"id"`
	Name        types.String                           `tfsdk:"name"`
	Identifier  types.String                           `tfsdk:"identifier"`
	Description types.String                           `tfsdk:"description"`
	Attributes  *general_objects.DefinitionAttributeTF `tfsdk:"attributes"`
}

func (x *CommandDefinition) ToTF() any {
	metadata := make([]MetadataTF, len(x.Metadata))
	for i, m := range x.Metadata {
		attr := general_objects.ValueAttributeToTF(&m.Attributes)
		metadata[i] = MetadataTF{
			ID:          helper.TFStringValue(m.ID),
			Name:        helper.TFStringValue(m.Name),
			Description: helper.TFStringValue(m.Description),
			Attributes:  &attr,
		}
	}

	arguments := make([]ArgumentTF, len(x.Arguments))
	for i, a := range x.Arguments {
		attr := general_objects.DefinitionAttributeToTF(&a.Attributes)
		arguments[i] = ArgumentTF{
			ID:          helper.TFStringValue(a.ID),
			Name:        helper.TFStringValue(a.Name),
			Identifier:  helper.TFStringValue(a.Identifier),
			Description: helper.TFStringValue(a.Description),
			Attributes:  &attr,
		}
	}

	return &CommandDefinitionTF{
		AuditModelTF: general_objects.AuditModelToTF(&x.AuditModel),
		NodeId:       helper.TFStringValue(x.NodeId),
		Name:         helper.TFStringValue(x.Name),
		Description:  helper.TFStringValue(x.Description),
		Identifier:   helper.TFStringValue(x.Identifier),
		Metadata:     metadata,
		Arguments:    arguments,
	}
}

func (tf *CommandDefinitionTF) ToAPI() any {
	metadata := make([]Metadata[any], len(tf.Metadata))
	for i, m := range tf.Metadata {
		metadata[i] = Metadata[any]{
			ID:          helper.FromTFString(m.ID),
			Name:        helper.FromTFString(m.Name),
			Description: helper.FromTFString(m.Description),
		}
		if m.Attributes != nil {
			metadata[i].Attributes = general_objects.ValueAttributeFromTF(*m.Attributes)
		}
	}

	arguments := make([]Argument[any], len(tf.Arguments))
	for i, a := range tf.Arguments {
		arguments[i] = Argument[any]{
			ID:          helper.FromTFString(a.ID),
			Name:        helper.FromTFString(a.Name),
			Identifier:  helper.FromTFString(a.Identifier),
			Description: helper.FromTFString(a.Description),
		}
		if a.Attributes != nil {
			arguments[i].Attributes = general_objects.DefinitionAttributeFromTF(*a.Attributes)
		}
	}

	return &CommandDefinition{
		AuditModel:  general_objects.AuditModelFromTF(tf.AuditModelTF),
		NodeId:      helper.FromTFString(tf.NodeId),
		Name:        helper.FromTFString(tf.Name),
		Description: helper.FromTFString(tf.Description),
		Identifier:  helper.FromTFString(tf.Identifier),
		Metadata:    metadata,
		Arguments:   arguments,
	}
}
