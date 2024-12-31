package request_states

type RequestState struct {
	ID             string `json:"id"`
	Name           string `json:"name"`
	CreatedAt      string `json:"createdAt"`
	CreatedBy      string `json:"createdBy"`
	LastModifiedAt string `json:"lastModifiedAt"`
	LastModifiedBy string `json:"lastModifiedBy"`
}

func (state *RequestState) GetID() string { return state.ID }
