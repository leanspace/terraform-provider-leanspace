package activity_states

type ActivityState struct {
	ID             string `json:"id"`
	Name           string `json:"name"`
	ReadOnly       bool   `json:"readOnly"`
	CreatedAt      string `json:"createdAt"`
	CreatedBy      string `json:"createdBy"`
	LastModifiedAt string `json:"lastModifiedAt"`
	LastModifiedBy string `json:"lastModifiedBy"`
}

func (state *ActivityState) GetID() string { return state.ID }
