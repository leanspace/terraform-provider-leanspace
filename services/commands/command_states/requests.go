package command_states

import (
	"github.com/leanspace/terraform-provider-leanspace/helper/general_objects"
	"github.com/leanspace/terraform-provider-leanspace/provider"
)

func (cs *CommandState) PostReadProcess(_ *provider.Client, newValue any) error {
	newCS, ok := newValue.(*CommandState)
	if !ok || newCS == nil {
		return nil
	}
	newCS.Tags = general_objects.ReorderKeyValues(cs.Tags, newCS.Tags)
	return nil
}
