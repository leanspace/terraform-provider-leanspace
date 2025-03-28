package event_definitions

import (
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/leanspace/terraform-provider-leanspace/helper"
	"github.com/leanspace/terraform-provider-leanspace/helper/general_objects"
)

func (eventDefinition *EventsDefinition) ToMap() map[string]any {
	eventDefinitionMap := make(map[string]any)
	eventDefinitionMap["id"] = eventDefinition.ID
	eventDefinitionMap["name"] = eventDefinition.Name
	eventDefinitionMap["description"] = eventDefinition.Description
	eventDefinitionMap["source"] = eventDefinition.Source
	eventDefinitionMap["state"] = eventDefinition.State
	if eventDefinition.Rules != nil {
		eventDefinitionMap["rules"] = helper.ParseToMaps(eventDefinition.Rules)
	}

	eventDefinitionMap["created_at"] = eventDefinition.CreatedAt
	eventDefinitionMap["created_by"] = eventDefinition.CreatedBy
	eventDefinitionMap["last_modified_at"] = eventDefinition.LastModifiedAt
	eventDefinitionMap["last_modified_by"] = eventDefinition.LastModifiedBy
	eventDefinitionMap["tags"] = helper.ParseToMaps(eventDefinition.Tags)

	return eventDefinitionMap
}

func (rule *Rules[T]) ToMap() map[string]any {
	ruleMap := make(map[string]any)
	ruleMap["operator"] = rule.Operator
	ruleMap["path"] = rule.Path
	if rule.ComparisonValue != nil {
		ruleMap["comparison_value"] = []any{(rule.ComparisonValue).ToMap()}
	}

	return ruleMap
}

func (comparisonValue *ComparisonValue[T]) ToMap() map[string]any {
	comparisonValueMap := make(map[string]any)
	comparisonValueMap["type"] = comparisonValue.Type
	switch comparisonValue.Type {
	case "NUMERIC":
		comparisonValueMap["value"] = helper.ParseFloat(any(comparisonValue.Value).(float64))
	case "BOOLEAN":
		comparisonValueMap["value"] = strconv.FormatBool(any(comparisonValue.Value).(bool))
	case "TEXT":
		comparisonValueMap["value"] = comparisonValue.Value

	}
	return comparisonValueMap
}

func (eventDefinition *EventsDefinition) FromMap(eventDefinitionMap map[string]any) error {
	eventDefinition.ID = eventDefinitionMap["id"].(string)
	eventDefinition.Name = eventDefinitionMap["name"].(string)
	eventDefinition.Source = eventDefinitionMap["source"].(string)
	eventDefinition.State = eventDefinitionMap["state"].(string)
	eventDefinition.Description = eventDefinitionMap["description"].(string)
	eventDefinition.CreatedAt = eventDefinitionMap["created_at"].(string)
	eventDefinition.CreatedBy = eventDefinitionMap["created_by"].(string)
	eventDefinition.LastModifiedAt = eventDefinitionMap["last_modified_at"].(string)
	eventDefinition.LastModifiedBy = eventDefinitionMap["last_modified_by"].(string)
	if eventDefinitionMap["rules"] != nil {
		if rules, err := helper.ParseFromMaps[Rules[any]](
			eventDefinitionMap["rules"].(*schema.Set).List(),
		); err != nil {
			return err
		} else {
			eventDefinition.Rules = rules
		}
	}
	if tags, err := helper.ParseFromMaps[general_objects.KeyValue](eventDefinitionMap["tags"].(*schema.Set).List()); err != nil {
		return err
	} else {
		eventDefinition.Tags = tags
	}
	return nil
}

func (rules *Rules[T]) FromMap(ruleMap map[string]any) error {
	rules.Operator = ruleMap["operator"].(string)
	rules.Path = ruleMap["path"].(string)
	if ruleMap["comparison_value"] != nil && len(ruleMap["comparison_value"].([]any)) > 0 {
		rules.ComparisonValue = new(ComparisonValue[T])
		err := (*(rules.ComparisonValue)).FromMap(ruleMap["comparison_value"].([]any)[0].(map[string]any))
		if err != nil {
			return err
		}
	}

	return nil
}

func (comparisonValue *ComparisonValue[T]) FromMap(comparisonValueMap map[string]any) error {
	comparisonValue.Value = comparisonValueMap["value"].(T)
	comparisonValue.Type = comparisonValueMap["type"].(string)
	return nil
}
