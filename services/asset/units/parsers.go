package units

import (
	"github.com/leanspace/terraform-provider-leanspace/helper"
)

func (unit *Unit) ToMap() map[string]any {
	unitMap := make(map[string]any)
	unitMap["id"] = helper.NilIfEmpty(unit.ID)
	unitMap["symbol"] = helper.NilIfEmpty(unit.Symbol)
	unitMap["display_name"] = helper.NilIfEmpty(unit.DisplayName)
	return unitMap
}

func (unit *Unit) FromMap(unitMap map[string]any) error {
	unit.ID = helper.CastString(unitMap, "id")
	unit.DisplayName = helper.CastString(unitMap, "display_name")
	unit.Symbol = helper.CastString(unitMap, "symbol")
	return nil
}
