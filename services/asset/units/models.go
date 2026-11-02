package units

type Unit struct {
	SecurityTags []string `json:"-"`
	ID           string   `json:"id"`
	Symbol       string   `json:"symbol"`
	DisplayName  string   `json:"displayName"`
}

func (unit *Unit) GetID() string { return unit.ID }

func (unit *Unit) GetSecurityTags() []string { return unit.SecurityTags }

func (unit *Unit) SetSecurityTags(tags []string) { unit.SecurityTags = tags }
