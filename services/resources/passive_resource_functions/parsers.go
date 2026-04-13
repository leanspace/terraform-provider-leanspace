package passive_resource_functions

import (
	"github.com/leanspace/terraform-provider-leanspace/helper"
	"github.com/leanspace/terraform-provider-leanspace/helper/general_objects"
)

func (passiveResourceFunction *PassiveResourceFunction) ToMap() map[string]any {
	passiveResourceFunctionMap := passiveResourceFunction.ToAuditMap()
	passiveResourceFunctionMap["resource_id"] = helper.NilIfEmpty(passiveResourceFunction.ResourceId)
	passiveResourceFunctionMap["name"] = helper.NilIfEmpty(passiveResourceFunction.Name)
	passiveResourceFunctionMap["control_bound"] = helper.Float64PtrToAny(passiveResourceFunction.ControlBound)
	passiveResourceFunctionMap["formula"] = passiveResourceFunction.Formula.ToMap()
	passiveResourceFunctionMap["tags"] = helper.ParseToMaps(passiveResourceFunction.Tags)

	return passiveResourceFunctionMap
}

func (formula *PassiveResourceFunctionFormula) ToMap() map[string]any {
	formulaMap := make(map[string]any)

	formulaMap["type"] = helper.NilIfEmpty(formula.Type)

	if formula.Type == "LINEAR" {
		formulaMap["rate"] = helper.NilIfEmpty(formula.Rate)
		formulaMap["time_unit"] = helper.NilIfEmpty(formula.TimeUnit)
	}

	return formulaMap
}

func (passiveResourceFunction *PassiveResourceFunction) FromMap(passiveResourceFunctionMap map[string]any) error {
	passiveResourceFunction.FromAuditMap(passiveResourceFunctionMap)
	passiveResourceFunction.ResourceId = helper.CastString(passiveResourceFunctionMap, "resource_id")
	passiveResourceFunction.Name = helper.CastString(passiveResourceFunctionMap, "name")
	passiveResourceFunction.ControlBound = helper.CastFloat64Ptr(passiveResourceFunctionMap, "control_bound")

	passiveResourceFunction.Formula = new(PassiveResourceFunctionFormula)
	if err := passiveResourceFunction.Formula.FromMap(helper.CastMapAny(passiveResourceFunctionMap, "formula")); err != nil {
		return err
	}
	if tags, err := helper.ParseFromMaps[general_objects.KeyValue](helper.CastSlice(passiveResourceFunctionMap, "tags")); err != nil {
		return err
	} else {
		passiveResourceFunction.Tags = tags
	}
	return nil
}

func (formula *PassiveResourceFunctionFormula) FromMap(formulaMap map[string]any) error {
	formula.Type = helper.CastString(formulaMap, "type")

	if formula.Type == "LINEAR" {
		formula.Rate = helper.CastFloat64(formulaMap, "rate")
		formula.TimeUnit = helper.CastString(formulaMap, "time_unit")
	}

	return nil
}
