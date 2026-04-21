package monitors

import (
	"github.com/leanspace/terraform-provider-leanspace/helper/general_objects"
	"github.com/leanspace/terraform-provider-leanspace/provider"
)

func (monitor *Monitor) PostReadProcess(_ *provider.Client, newValue any) error {
	newMonitor, ok := newValue.(*Monitor)
	if !ok || newMonitor == nil {
		return nil
	}
	newMonitor.Tags = general_objects.ReorderKeyValues(monitor.Tags, newMonitor.Tags)
	return nil
}

func (monitor *Monitor) PostUnmarshallProcess() error {
	monitor.ActionTemplateLinks = make([]ActionTemplateLink, len(monitor.ActionTemplates))
	for i, value := range monitor.ActionTemplates {
		monitor.ActionTemplateLinks[i].ID = value.ID
		monitor.ActionTemplateLinks[i].TriggeredOn = value.TriggeredOn
	}
	return nil
}
