package action_templates

func (actionTemplate *ActionTemplate) ToMap() map[string]any {
	actionTemplateMap := make(map[string]any)
	actionTemplateMap["id"] = actionTemplate.ID
	actionTemplateMap["name"] = actionTemplate.Name
	actionTemplateMap["type"] = actionTemplate.Type
	actionTemplateMap["url"] = actionTemplate.URL
	actionTemplateMap["payload"] = actionTemplate.Payload
	actionTemplateMap["headers"] = actionTemplate.Headers
	actionTemplateMap["created_at"] = actionTemplate.CreatedAt
	actionTemplateMap["created_by"] = actionTemplate.CreatedBy
	actionTemplateMap["last_modified_at"] = actionTemplate.LastModifiedAt
	actionTemplateMap["last_modified_by"] = actionTemplate.LastModifiedBy
	return actionTemplateMap
}

func (actionTemplate *ActionTemplate) FromMap(actionTemplateMap map[string]any) error {
	actionTemplate.ID = actionTemplateMap["id"].(string)
	actionTemplate.Name = actionTemplateMap["name"].(string)
	actionTemplate.Type = actionTemplateMap["type"].(string)
	actionTemplate.URL = actionTemplateMap["url"].(string)
	actionTemplate.Payload = actionTemplateMap["payload"].(string)
	actionTemplate.Headers = make(map[string]string, len(actionTemplateMap["headers"].(map[string]any)))
	for key, value := range actionTemplateMap["headers"].(map[string]any) {
		actionTemplate.Headers[key] = value.(string)
	}
	actionTemplate.CreatedAt = actionTemplateMap["created_at"].(string)
	actionTemplate.CreatedBy = actionTemplateMap["created_by"].(string)
	actionTemplate.LastModifiedAt = actionTemplateMap["last_modified_at"].(string)
	actionTemplate.LastModifiedBy = actionTemplateMap["last_modified_by"].(string)
	return nil
}
