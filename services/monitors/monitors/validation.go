package monitors

import (
	. "github.com/leanspace/terraform-provider-leanspace/helper"
)

var monitorRuleValidators = Validators{
	If(
		IsSet("tolerance"),
		Or(Equals("comparison_operator", "EQUAL_TO"), Equals("comparison_operator", "NOT_EQUAL_TO")),
	),
}

func (monitor *Monitor) Validate() error {
	return monitorRuleValidators.CheckValue(monitor)
}
