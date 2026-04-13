package event_definitions

import (
	"strconv"

	"github.com/leanspace/terraform-provider-leanspace/helper"
	"github.com/leanspace/terraform-provider-leanspace/helper/general_objects"
)

func (eventDefinition *EventsDefinition) ToMap() map[string]any {
	eventDefinitionMap := eventDefinition.ToAuditMap()
	eventDefinitionMap["name"] = helper.NilIfEmpty(eventDefinition.Name)
	eventDefinitionMap["description"] = helper.NilIfEmpty(eventDefinition.Description)
	eventDefinitionMap["source"] = helper.NilIfEmpty(eventDefinition.Source)
	if eventDefinition.Criticality != "" {
		eventDefinitionMap["criticality"] = helper.NilIfEmpty(eventDefinition.Criticality)
	}
	eventDefinitionMap["state"] = helper.NilIfEmpty(eventDefinition.State)
	if eventDefinition.Rules != nil {
		eventDefinitionMap["rules"] = helper.ParseToMaps(eventDefinition.Rules)
	}
	eventDefinitionMap["tags"] = helper.ParseToMaps(eventDefinition.Tags)
	return eventDefinitionMap
}

func (rule *Rules[T]) ToMap() map[string]any {
	ruleMap := make(map[string]any)
	ruleMap["operator"] = helper.NilIfEmpty(rule.Operator)
	ruleMap["path"] = helper.NilIfEmpty(rule.Path)
	if rule.ComparisonValue != nil {
		ruleMap["comparison_value"] = (rule.ComparisonValue).ToMap()
	}

	return ruleMap
}

func (comparisonValue *ComparisonValue[T]) ToMap() map[string]any {
	comparisonValueMap := make(map[string]any)
	comparisonValueMap["type"] = helper.NilIfEmpty(comparisonValue.Type)
	switch comparisonValue.Type {
	case "NUMERIC":
		comparisonValueMap["value"] = helper.ParseFloat(any(comparisonValue.Value).(float64))
	case "BOOLEAN":
		comparisonValueMap["value"] = strconv.FormatBool(any(comparisonValue.Value).(bool))
	case "TEXT":
		comparisonValueMap["value"] = helper.NilIfEmpty(comparisonValue.Value)

	}
	return comparisonValueMap
}

func (eventDefinition *EventsDefinition) FromMap(eventDefinitionMap map[string]any) error {
	eventDefinition.FromAuditMap(eventDefinitionMap)
	eventDefinition.Name = helper.CastString(eventDefinitionMap, "name")
	eventDefinition.Source = helper.CastString(eventDefinitionMap, "source")
	eventDefinition.State = helper.CastString(eventDefinitionMap, "state")
	if criticality := helper.CastString(eventDefinitionMap, "criticality"); criticality != "" {
		eventDefinition.Criticality = criticality
	}
	eventDefinition.Description = helper.CastString(eventDefinitionMap, "description")
	if eventDefinitionMap["rules"] != nil {
		if rules, err := helper.ParseFromMaps[Rules[any]](
			helper.CastSlice(eventDefinitionMap, "rules"),
		); err != nil {
			return err
		} else {
			eventDefinition.Rules = rules
		}
	}
	if tags, err := helper.ParseFromMaps[general_objects.KeyValue](helper.CastSlice(eventDefinitionMap, "tags")); err != nil {
		return err
	} else {
		eventDefinition.Tags = tags
	}
	return nil
}

func (rules *Rules[T]) FromMap(ruleMap map[string]any) error {
	rules.Operator = helper.CastString(ruleMap, "operator")
	rules.Path = helper.CastString(ruleMap, "path")
	if ruleMap["comparison_value"] != nil {
		rules.ComparisonValue = new(ComparisonValue[T])
		err := (*(rules.ComparisonValue)).FromMap(helper.CastMapAny(ruleMap, "comparison_value"))
		if err != nil {
			return err
		}
	}

	return nil
}

func (comparisonValue *ComparisonValue[T]) FromMap(comparisonValueMap map[string]any) error {
	comparisonValue.Value = comparisonValueMap["value"].(T)
	comparisonValue.Type = helper.CastString(comparisonValueMap, "type")
	return nil
}
