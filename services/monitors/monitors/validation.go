package monitors

import (
	. "github.com/leanspace/terraform-provider-leanspace/helper"
)

var monitorExpressionValidators = Validators{
	If(
		IsSet("tolerance"),
		Or(Equals("comparison_operator", "EQUAL_TO"), Equals("comparison_operator", "NOT_EQUAL_TO")),
	),
}

func (monitor *Monitor) Validate(obj map[string]any) error {
	return monitorExpressionValidators.Check(obj["expression"].([]any)[0].(map[string]any))
}
