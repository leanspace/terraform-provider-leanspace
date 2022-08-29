package analysis_definitions

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/leanspace/terraform-provider-leanspace/helper"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func (analysisDefinition *AnalysisDefinition) ToMap() map[string]any {
	analysisDefMap := make(map[string]any)
	analysisDefMap["id"] = analysisDefinition.ID
	analysisDefMap["name"] = analysisDefinition.Name
	analysisDefMap["description"] = analysisDefinition.Description
	analysisDefMap["framework"] = analysisDefinition.Framework
	analysisDefMap["model_id"] = analysisDefinition.ModelId
	analysisDefMap["node_id"] = analysisDefinition.NodeId
	analysisDefMap["statistics"] = []map[string]any{analysisDefinition.Statistics.ToMap()}
	analysisDefMap["inputs"] = []map[string]any{analysisDefinition.Inputs.ToMap()}
	analysisDefMap["created_at"] = analysisDefinition.CreatedAt
	analysisDefMap["created_by"] = analysisDefinition.CreatedBy
	analysisDefMap["last_modified_at"] = analysisDefinition.LastModifiedAt
	analysisDefMap["last_modified_by"] = analysisDefinition.LastModifiedBy
	return analysisDefMap
}

func (statistics *Statistics) ToMap() map[string]any {
	statisticsMap := make(map[string]any)
	statisticsMap["number_of_executions"] = statistics.NumberOfExecutions
	statisticsMap["last_executed_at"] = statistics.LastExecutedAt
	return statisticsMap
}

func (field *Field) ToMap() map[string]any {
	fieldMap := make(map[string]any)
	fieldMap["type"] = field.Type
	fieldMap["items"] = []any{}
	fieldMap["fields"] = []any{}
	if field.Type == "STRUCTURE" {
		subfields := []map[string]any{}
		for subfieldName, subfield := range field.Fields {
			// We can't have a map of objects in the current terraform sdk,
			// so instead we extract the map key and set it inside the object.
			// The reverse operation is done when parsing in .FromMap()
			subfieldMap := subfield.ToMap()
			subfieldMap["name"] = subfieldName
			subfields = append(subfields, subfieldMap)
		}
		fieldMap["fields"] = subfields
	} else if field.Type == "ARRAY" {
		fieldMap["items"] = helper.ParseToMaps(field.Items)
	} else {
		fieldMap["source"] = field.Source
		fieldMap["ref"] = field.Ref
		if field.Value != nil {
			if jsonData, err := json.Marshal(field.Value); err == nil && !strings.HasPrefix(string(jsonData), "\"") {
				fieldMap["value"] = string(jsonData)
			} else {
				fieldMap["value"] = field.Value
			}
		}
	}
	return fieldMap
}

func (analysisDefinition *AnalysisDefinition) FromMap(analysisDefMap map[string]any) error {
	analysisDefinition.ID = analysisDefMap["id"].(string)
	analysisDefinition.Name = analysisDefMap["name"].(string)
	analysisDefinition.Description = analysisDefMap["description"].(string)
	analysisDefinition.Framework = analysisDefMap["framework"].(string)
	analysisDefinition.ModelId = analysisDefMap["model_id"].(string)
	analysisDefinition.NodeId = analysisDefMap["node_id"].(string)
	if statisticsMap, ok := analysisDefMap["statistics"].([]any); ok && len(statisticsMap) > 0 {
		if err := analysisDefinition.Statistics.FromMap(statisticsMap[0].(map[string]any)); err != nil {
			return err
		}
	}
	if err := analysisDefinition.Inputs.FromMap(
		analysisDefMap["inputs"].([]any)[0].(map[string]any),
	); err != nil {
		return err
	}
	analysisDefinition.CreatedAt = analysisDefMap["created_at"].(string)
	analysisDefinition.CreatedBy = analysisDefMap["created_by"].(string)
	analysisDefinition.LastModifiedAt = analysisDefMap["last_modified_at"].(string)
	analysisDefinition.LastModifiedBy = analysisDefMap["last_modified_by"].(string)
	return nil
}

func (statistics *Statistics) FromMap(statisticsMap map[string]any) error {
	statistics.NumberOfExecutions = statisticsMap["number_of_executions"].(int)
	statistics.LastExecutedAt = statisticsMap["last_executed_at"].(string)
	return nil
}

func (field *Field) FromMap(fieldMap map[string]any) error {
	field.Type = fieldMap["type"].(string)
	if field.Type == "STRUCTURE" {
		field.Fields = make(map[string]Field)
		for _, subfieldMapRaw := range fieldMap["fields"].(*schema.Set).List() {
			// We can't have maps of objects in the current terraform sdk
			// So instead we extract the name of the field and use it
			// when making the dictionary. The reverse is done in .ToMap()
			subfieldMap, isMap := subfieldMapRaw.(map[string]any)
			if !isMap {
				continue
			}
			subfieldName, ok := subfieldMap["name"].(string)
			if !ok {
				return fmt.Errorf("expected name to be set for field, got %p", &subfieldName)
			}
			subfield := Field{}
			subfield.FromMap(subfieldMap)
			field.Fields[subfieldName] = subfield
		}
	} else if field.Type == "ARRAY" {
		if items, err := helper.ParseFromMaps[Field](fieldMap["items"].([]any)); err != nil {
			return err
		} else {
			field.Items = items
		}
	} else {
		field.Source = fieldMap["source"].(string)
		field.Ref = fieldMap["ref"].(string)
		if field.Source == "STATIC" {
			valueString := fieldMap["value"].(string)
			var jsonValue any
			if err := json.Unmarshal([]byte(valueString), jsonValue); err == nil {
				field.Value = jsonValue
			} else {
				field.Value = valueString
			}
		}
	}
	return nil
}
