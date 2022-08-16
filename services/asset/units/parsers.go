package units

func (unit *Unit) ToMap() map[string]any {
	unitMap := make(map[string]any)
	unitMap["id"] = unit.ID
	unitMap["symbol"] = unit.Symbol
	unitMap["display_name"] = unit.DisplayName
	return unitMap
}

func (unit *Unit) FromMap(unitMap map[string]any) error {
	unit.ID = unitMap["id"].(string)
	unit.DisplayName = unitMap["display_name"].(string)
	unit.Symbol = unitMap["symbol"].(string)
	return nil
}
