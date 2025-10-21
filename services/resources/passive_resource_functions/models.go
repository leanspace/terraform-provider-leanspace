package passive_resource_functions

import "github.com/leanspace/terraform-provider-leanspace/helper/general_objects"

type PassiveResourceFunction struct {
	ID             string                          `json:"id"`
	ResourceId     string                          `json:"resourceId"`
	Name           string                          `json:"name"`
	Formula        *PassiveResourceFunctionFormula `json:"formula"`
	ControlBound   float64                         `json:"controlBound"`
	Tags           []general_objects.KeyValue      `json:"tags,omitempty"`
	CreatedAt      string                          `json:"createdAt"`
	CreatedBy      string                          `json:"createdBy"`
	LastModifiedAt string                          `json:"lastModifiedAt"`
	LastModifiedBy string                          `json:"lastModifiedBy"`
}

func (passiveResourceFunction *PassiveResourceFunction) GetID() string {
	return passiveResourceFunction.ID
}

type PassiveResourceFunctionFormula struct {
	Type     string  `json:"type"`
	Constant float64 `json:"constant"`
	Rate     float64 `json:"rate"`
	TimeUnit string  `json:"timeUnit"`
}
