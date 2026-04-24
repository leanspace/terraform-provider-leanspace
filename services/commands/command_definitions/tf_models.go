package command_definitions

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
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
			ID:          types.StringValue(m.ID),
			Name:        types.StringValue(m.Name),
			Description: types.StringPointerValue(m.Description),
			Attributes:  &attr,
		}
	}

	arguments := make([]ArgumentTF, len(x.Arguments))
	for i, a := range x.Arguments {
		attr := general_objects.DefinitionAttributeToTF(&a.Attributes)
		arguments[i] = ArgumentTF{
			ID:          types.StringValue(a.ID),
			Name:        types.StringValue(a.Name),
			Identifier:  types.StringValue(a.Identifier),
			Description: types.StringPointerValue(a.Description),
			Attributes:  &attr,
		}
	}

	return &CommandDefinitionTF{
		AuditModelTF: general_objects.AuditModelToTF(&x.AuditModel),
		NodeId:       types.StringValue(x.NodeId),
		Name:         types.StringValue(x.Name),
		Description:  types.StringPointerValue(x.Description),
		Identifier:   types.StringPointerValue(x.Identifier),
		Metadata:     metadata,
		Arguments:    arguments,
	}
}

func (tf *CommandDefinitionTF) ToAPI() any {
	metadata := make([]Metadata[any], len(tf.Metadata))
	for i, m := range tf.Metadata {
		metadata[i] = Metadata[any]{
			ID:          m.ID.ValueString(),
			Name:        m.Name.ValueString(),
			Description: m.Description.ValueStringPointer(),
		}
		if m.Attributes != nil {
			metadata[i].Attributes = general_objects.ValueAttributeFromTF(*m.Attributes)
		}
	}

	arguments := make([]Argument[any], len(tf.Arguments))
	for i, a := range tf.Arguments {
		arguments[i] = Argument[any]{
			ID:          a.ID.ValueString(),
			Name:        a.Name.ValueString(),
			Identifier:  a.Identifier.ValueString(),
			Description: a.Description.ValueStringPointer(),
		}
		if a.Attributes != nil {
			arguments[i].Attributes = general_objects.DefinitionAttributeFromTF(*a.Attributes)
		}
	}

	return &CommandDefinition{
		AuditModel:  general_objects.AuditModelFromTF(tf.AuditModelTF),
		NodeId:      tf.NodeId.ValueString(),
		Name:        tf.Name.ValueString(),
		Description: tf.Description.ValueStringPointer(),
		Identifier:  tf.Identifier.ValueStringPointer(),
		Metadata:    metadata,
		Arguments:   arguments,
	}
}
