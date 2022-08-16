package units

type Unit struct {
	ID          string `json:"id"`
	Symbol      string `json:"symbol"`
	DisplayName string `json:"displayName"`
}

func (unit *Unit) GetID() string { return unit.ID }
