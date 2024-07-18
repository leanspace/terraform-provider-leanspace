package pass_plan_states

type PassPlanState struct {
	ID             string `json:"id"`
	Name           string `json:"name"`
	ReadOnly       bool   `json:"readOnly"`
	CreatedAt      string `json:"createdAt"`
	CreatedBy      string `json:"createdBy"`
	LastModifiedAt string `json:"lastModifiedAt"`
	LastModifiedBy string `json:"lastModifiedBy"`
}

func (state *PassPlanState) GetID() string { return state.ID }
