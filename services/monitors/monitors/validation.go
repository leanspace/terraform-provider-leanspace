package monitors

import (
	. "github.com/leanspace/terraform-provider-leanspace/helper"
)

var monitorRuleThresholdConditionValidators = Validators{
	If(
		IsSet("tolerance"),
		Or(Equals("comparison_operator", "EQUAL_TO"), Equals("comparison_operator", "NOT_EQUAL_TO")),
	),
}

func (monitor *Monitor) Validate(obj map[string]any) error {
	rule := obj["rule"].([]any)[0].(map[string]any)

	if _, ok := rule["trigger_condition"].([]any); !ok {
		if err := monitorRuleThresholdConditionValidators.Check(rule); err != nil {
			return err
		}
	}

	for _, conditionName := range []string{"trigger_condition", "clear_condition"} {
		conditions, ok := rule[conditionName].([]any)
		if !ok || len(conditions) == 0 || conditions[0] == nil {
			continue
		}

		if err := monitorRuleThresholdConditionValidators.Check(conditions[0].(map[string]any)); err != nil {
			return err
		}
	}

	return nil
}
