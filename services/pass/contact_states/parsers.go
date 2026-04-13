package contact_states

import "github.com/leanspace/terraform-provider-leanspace/helper"

func (state *ContactState) ToMap() map[string]any {
	stateMap := state.ToAuditMap()
	stateMap["name"] = helper.NilIfEmpty(state.Name)
	stateMap["read_only"] = helper.NilIfEmpty(state.ReadOnly)
	return stateMap
}

func (state *ContactState) FromMap(stateMap map[string]any) error {
	state.FromAuditMap(stateMap)
	state.Name = helper.CastString(stateMap, "name")
	state.ReadOnly = helper.CastBool(stateMap, "read_only")
	return nil
}
