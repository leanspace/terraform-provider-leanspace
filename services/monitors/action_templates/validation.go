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
	return actionTemplateValidator.CheckValue(actionTemplate)
}
