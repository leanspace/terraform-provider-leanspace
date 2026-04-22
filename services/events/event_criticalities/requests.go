package event_criticalities

import (
	"github.com/leanspace/terraform-provider-leanspace/helper/general_objects"
	"github.com/leanspace/terraform-provider-leanspace/provider"
)

func (ec *EventCriticalities) PostReadProcess(_ *provider.Client, newValue any) error {
	newEC, ok := newValue.(*EventCriticalities)
	if !ok || newEC == nil {
		return nil
	}
	newEC.Tags = general_objects.ReorderKeyValues(ec.Tags, newEC.Tags)
	return nil
}
