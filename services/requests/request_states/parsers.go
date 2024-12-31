package request_states

func (state *RequestState) ToMap() map[string]any {
	stateMap := make(map[string]any)
	stateMap["id"] = state.ID
	stateMap["name"] = state.Name
	stateMap["created_at"] = state.CreatedAt
	stateMap["created_by"] = state.CreatedBy
	stateMap["last_modified_at"] = state.LastModifiedAt
	stateMap["last_modified_by"] = state.LastModifiedBy

	return stateMap
}

func (state *RequestState) FromMap(stateMap map[string]any) error {
	state.ID = stateMap["id"].(string)
	state.Name = stateMap["name"].(string)
	state.CreatedAt = stateMap["created_at"].(string)
	state.CreatedBy = stateMap["created_by"].(string)
	state.LastModifiedAt = stateMap["last_modified_at"].(string)
	state.LastModifiedBy = stateMap["last_modified_by"].(string)

	return nil
}
