package analysis_definitions

type AnalysisDefinition struct {
	ID             string     `json:"id"`
	Name           string     `json:"name"`
	Description    string     `json:"description"`
	Framework      string     `json:"framework"`
	ModelId        string     `json:"modelId"`
	NodeId         string     `json:"nodeId"`
	Statistics     Statistics `json:"statistics"`
	Inputs         Field      `json:"inputs"`
	CreatedAt      string     `json:"createdAt"`
	CreatedBy      string     `json:"createdBy"`
	LastModifiedAt string     `json:"lastModifiedAt"`
	LastModifiedBy string     `json:"lastModifiedBy"`
}

func (analysisDefinition *AnalysisDefinition) GetID() string { return analysisDefinition.ID }

type Statistics struct {
	NumberOfExecutions int    `json:"numberOfExecutions"`
	LastExecutedAt     string `json:"lastExecutedAt"`
}

type Field struct {
	Type   string           `json:"type"`
	Source *string          `json:"source,omitempty"`
	Ref    *string          `json:"ref,omitempty"`
	Value  any              `json:"value,omitempty"`
	Fields map[string]Field `json:"fields,omitempty"`
	Items  []Field          `json:"items,omitempty"`
}
