package command_states

import (
	"github.com/leanspace/terraform-provider-leanspace/helper"
	"github.com/leanspace/terraform-provider-leanspace/helper/general_objects"
)

func (state *CommandState) ToMap() map[string]any {
	stateMap := state.ToAuditMap()
	stateMap["name"] = helper.NilIfEmpty(state.Name)
	stateMap["read_only"] = helper.NilIfEmpty(state.ReadOnly)
	stateMap["tags"] = helper.ParseToMaps(state.Tags)
	return stateMap
}

func (state *CommandState) FromMap(stateMap map[string]any) error {
	state.FromAuditMap(stateMap)
	state.Name = helper.CastString(stateMap, "name")
	state.ReadOnly = helper.CastBool(stateMap, "read_only")
	if tags, err := helper.ParseFromMaps[general_objects.KeyValue](helper.CastSlice(stateMap, "tags")); err != nil {
		return err
	} else {
		state.Tags = tags
	}
	return nil
}
