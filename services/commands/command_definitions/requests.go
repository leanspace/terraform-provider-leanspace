package command_definitions

import (
	"github.com/leanspace/terraform-provider-leanspace/helper"
	"github.com/leanspace/terraform-provider-leanspace/provider"
)

func (commandDefinition *CommandDefinition) PostReadProcess(_ *provider.Client, newValue any) error {
	newCmdDef, ok := newValue.(*CommandDefinition)
	if !ok || newCmdDef == nil {
		return nil
	}
	newCmdDef.Metadata = helper.ReorderByKey(commandDefinition.Metadata, newCmdDef.Metadata, func(m Metadata[any]) string { return m.Name })
	newCmdDef.Arguments = helper.ReorderByKey(commandDefinition.Arguments, newCmdDef.Arguments, func(a Argument[any]) string { return a.Name })
	return nil
}
