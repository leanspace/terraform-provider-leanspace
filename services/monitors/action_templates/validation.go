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

func (actionTemplate *ActionTemplate) Validate(obj map[string]any) error {
	if err := actionTemplateValidator.Check(obj); err != nil {
		return err
	}

	return nil
}
