package action_templates

import (
	"github.com/leanspace/terraform-provider-leanspace/helper"
)

func (actionTemplate *ActionTemplate) ToMap() map[string]any {
	actionTemplateMap := actionTemplate.ToAuditMap()
	actionTemplateMap["name"] = helper.NilIfEmpty(actionTemplate.Name)
	actionTemplateMap["type"] = helper.NilIfEmpty(actionTemplate.Type)
	actionTemplateMap["url"] = helper.NilIfEmpty(actionTemplate.URL)
	actionTemplateMap["payload"] = helper.NilIfEmpty(actionTemplate.Payload)
	actionTemplateMap["content"] = helper.NilIfEmpty(actionTemplate.Content)
	actionTemplateMap["headers"] = helper.NilIfEmpty(actionTemplate.Headers)
	return actionTemplateMap
}

func (actionTemplate *ActionTemplate) FromMap(actionTemplateMap map[string]any) error {
	actionTemplate.FromAuditMap(actionTemplateMap)
	actionTemplate.Name = helper.CastString(actionTemplateMap, "name")
	actionTemplate.Type = helper.CastString(actionTemplateMap, "type")
	actionTemplate.URL = helper.CastString(actionTemplateMap, "url")
	actionTemplate.Payload = helper.CastString(actionTemplateMap, "payload")
	actionTemplate.Content = helper.CastString(actionTemplateMap, "content")
	actionTemplate.Headers = make(map[string]string, len(helper.CastMapAny(actionTemplateMap, "headers")))
	for key, value := range helper.CastMapAny(actionTemplateMap, "headers") {
		actionTemplate.Headers[key] = value.(string)
	}
	return nil
}
