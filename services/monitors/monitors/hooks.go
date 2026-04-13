package monitors

func (monitor *Monitor) PostUnmarshallProcess() error {
	monitor.ActionTemplateLinks = make([]ActionTemplateLink, len(monitor.ActionTemplates))
	for i, value := range monitor.ActionTemplates {
		monitor.ActionTemplateLinks[i].ID = value.ID
		monitor.ActionTemplateLinks[i].TriggeredOn = value.TriggeredOn
	}
	return nil
}
