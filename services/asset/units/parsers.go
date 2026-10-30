package units

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

func (unit *Unit) ToMap() map[string]any {
	unitMap := make(map[string]any)
	unitMap["id"] = unit.ID
	unitMap["symbol"] = unit.Symbol
	unitMap["display_name"] = unit.DisplayName
	if unit.SecurityTags != nil {
		unitMap["security_tags"] = unit.SecurityTags
	}
	return unitMap
}

func (unit *Unit) FromMap(unitMap map[string]any) error {
	if tagSet, ok := unitMap["security_tags"].(*schema.Set); ok {
		unit.SecurityTags = make([]string, tagSet.Len())
		for index, tag := range tagSet.List() {
			unit.SecurityTags[index] = tag.(string)
		}
	}
	unit.ID = unitMap["id"].(string)
	unit.DisplayName = unitMap["display_name"].(string)
	unit.Symbol = unitMap["symbol"].(string)
	return nil
}
