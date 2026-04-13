package action_templates

import (
	. "github.com/leanspace/terraform-provider-leanspace/helper"
)

var actionTemplateValidator = Validators{
	If(
		Equals("type", "WEBHOOK"),
		And(IsSet("url"), IsSet("payload")),
	),
}

func (actionTemplate *ActionTemplate) Validate() error {
	obj := map[string]any{
		"type":    actionTemplate.Type,
		"url":     actionTemplate.URL,
		"payload": actionTemplate.Payload,
	}
	return actionTemplateValidator.Check(obj)
}
