package units

//go:generate go run github.com/leanspace/terraform-provider-leanspace/tools/gen_tf_models -struct Unit

type Unit struct {
	ID          string `json:"id"`
	Symbol      string `json:"symbol"`
	DisplayName string `json:"displayName"`
}

func (unit *Unit) GetID() string { return unit.ID }
