package pass_states

type PassState struct {
	ID             string `json:"id"`
	Name           string `json:"name"`
	ReadOnly       bool   `json:"readOnly"`
	CreatedAt      string `json:"createdAt"`
	CreatedBy      string `json:"createdBy"`
	LastModifiedAt string `json:"lastModifiedAt"`
	LastModifiedBy string `json:"lastModifiedBy"`
}

func (state *PassState) GetID() string { return state.ID }
