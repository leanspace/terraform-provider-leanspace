package resource_functions

type ResourceFunction struct {
	ID                   string                   `json:"id"`
	ActivityDefinitionId string                   `json:"activityDefinitionId"`
	ResourceId           string                   `json:"resourceId"`
	Name                 string                   `json:"name"`
	Formula              *ResourceFunctionFormula `json:"formula"`
	CreatedAt            string                   `json:"createdAt"`
	CreatedBy            string                   `json:"createdBy"`
	LastModifiedAt       string                   `json:"lastModifiedAt"`
	LastModifiedBy       string                   `json:"lastModifiedBy"`
}

func (resourceFunction *ResourceFunction) GetID() string { return resourceFunction.ID }

type ResourceFunctionFormula struct {
	Type      string  `json:"type"`
	Amplitude float64 `json:"amplitude"`
	Constant  float64 `json:"constant"`
	Rate      float64 `json:"rate"`
	TimeUnit  string  `json:"timeUnit,omitempty"`
}
