package event_definitions

import (
	"github.com/leanspace/terraform-provider-leanspace/helper/general_objects"
	"github.com/leanspace/terraform-provider-leanspace/provider"
)

func (def *EventsDefinition) PostReadProcess(_ *provider.Client, newValue any) error {
	newDef, ok := newValue.(*EventsDefinition)
	if !ok || newDef == nil {
		return nil
	}
	newDef.Tags = general_objects.ReorderKeyValues(def.Tags, newDef.Tags)
	return nil
}
