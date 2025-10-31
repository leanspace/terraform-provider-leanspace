package passive_resource_functions

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/leanspace/terraform-provider-leanspace/helper"
	"github.com/leanspace/terraform-provider-leanspace/helper/general_objects"
)

func (passiveResourceFunction *PassiveResourceFunction) ToMap() map[string]any {
	passiveResourceFunctionMap := make(map[string]any)
	passiveResourceFunctionMap["id"] = passiveResourceFunction.ID
	passiveResourceFunctionMap["resource_id"] = passiveResourceFunction.ResourceId
	passiveResourceFunctionMap["name"] = passiveResourceFunction.Name
	if passiveResourceFunction.ControlBound != nil {
		passiveResourceFunctionMap["control_bound"] = []any{float64(*passiveResourceFunction.ControlBound)}
	}
	passiveResourceFunctionMap["formula"] = []map[string]any{passiveResourceFunction.Formula.ToMap()}
	passiveResourceFunctionMap["tags"] = helper.ParseToMaps(passiveResourceFunction.Tags)
	passiveResourceFunctionMap["created_at"] = passiveResourceFunction.CreatedAt
	passiveResourceFunctionMap["created_by"] = passiveResourceFunction.CreatedBy
	passiveResourceFunctionMap["last_modified_at"] = passiveResourceFunction.LastModifiedAt
	passiveResourceFunctionMap["last_modified_by"] = passiveResourceFunction.LastModifiedBy

	return passiveResourceFunctionMap
}

func (formula *PassiveResourceFunctionFormula) ToMap() map[string]any {
	formulaMap := make(map[string]any)

	formulaMap["type"] = formula.Type

	if formula.Type == "LINEAR" {
		formulaMap["rate"] = formula.Rate
		formulaMap["time_unit"] = formula.TimeUnit
	}

	return formulaMap
}

func (passiveResourceFunction *PassiveResourceFunction) FromMap(passiveResourceFunctionMap map[string]any) error {
	passiveResourceFunction.ID = passiveResourceFunctionMap["id"].(string)
	passiveResourceFunction.ResourceId = passiveResourceFunctionMap["resource_id"].(string)
	passiveResourceFunction.Name = passiveResourceFunctionMap["name"].(string)

	if v, ok := passiveResourceFunctionMap["control_bound"]; ok && v != nil {
		if list, ok := v.([]interface{}); ok && len(list) > 0 {
			if floatVal, ok := list[0].(float64); ok {
				passiveResourceFunction.ControlBound = &floatVal
			}
		}
	}

	if len(passiveResourceFunctionMap["formula"].([]any)) > 0 && passiveResourceFunctionMap["formula"].([]any)[0] != nil {
		passiveResourceFunction.Formula = new(PassiveResourceFunctionFormula)
		if err := passiveResourceFunction.Formula.FromMap(passiveResourceFunctionMap["formula"].([]any)[0].(map[string]any)); err != nil {
			return err
		}
	}
	if tags, err := helper.ParseFromMaps[general_objects.KeyValue](passiveResourceFunctionMap["tags"].(*schema.Set).List()); err != nil {
		return err
	} else {
		passiveResourceFunction.Tags = tags
	}
	passiveResourceFunction.CreatedAt = passiveResourceFunctionMap["created_at"].(string)
	passiveResourceFunction.CreatedBy = passiveResourceFunctionMap["created_by"].(string)
	passiveResourceFunction.LastModifiedAt = passiveResourceFunctionMap["last_modified_at"].(string)
	passiveResourceFunction.LastModifiedBy = passiveResourceFunctionMap["last_modified_by"].(string)
	return nil
}

func (formula *PassiveResourceFunctionFormula) FromMap(formulaMap map[string]any) error {
	formula.Type = formulaMap["type"].(string)

	if formula.Type == "LINEAR" {
		formula.Rate = formulaMap["rate"].(float64)
		formula.TimeUnit = formulaMap["time_unit"].(string)
	}

	return nil
}
