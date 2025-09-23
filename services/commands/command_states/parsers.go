package command_states

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/leanspace/terraform-provider-leanspace/helper"
	"github.com/leanspace/terraform-provider-leanspace/helper/general_objects"
)

func (state *CommandState) ToMap() map[string]any {
	stateMap := make(map[string]any)
	stateMap["id"] = state.ID
	stateMap["name"] = state.Name
	stateMap["read_only"] = state.ReadOnly
	stateMap["created_at"] = state.CreatedAt
	stateMap["created_by"] = state.CreatedBy
	stateMap["last_modified_at"] = state.LastModifiedAt
	stateMap["last_modified_by"] = state.LastModifiedBy
	stateMap["tags"] = helper.ParseToMaps(state.Tags)

	return stateMap
}

func (state *CommandState) FromMap(stateMap map[string]any) error {
	state.ID = stateMap["id"].(string)
	state.Name = stateMap["name"].(string)
	state.ReadOnly = stateMap["read_only"].(bool)
	state.CreatedAt = stateMap["created_at"].(string)
	state.CreatedBy = stateMap["created_by"].(string)
	state.LastModifiedAt = stateMap["last_modified_at"].(string)
	state.LastModifiedBy = stateMap["last_modified_by"].(string)
	if tags, err := helper.ParseFromMaps[general_objects.KeyValue](stateMap["tags"].(*schema.Set).List()); err != nil {
		return err
	} else {
		state.Tags = tags
	}
	return nil
}
