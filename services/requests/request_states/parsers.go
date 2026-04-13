package request_states

import "github.com/leanspace/terraform-provider-leanspace/helper"

func (state *RequestState) ToMap() map[string]any {
	stateMap := state.ToAuditMap()
	stateMap["name"] = helper.NilIfEmpty(state.Name)
	return stateMap
}

func (state *RequestState) FromMap(stateMap map[string]any) error {
	state.FromAuditMap(stateMap)
	state.Name = helper.CastString(stateMap, "name")
	return nil
}
