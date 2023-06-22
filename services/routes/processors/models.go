package processors

type Processor struct {
	ID             string `json:"id"`
	Name           string `json:"name"`
	Description    string `json:"description,omitempty"`
	Version        string `json:"version"`
	Type           string `json:"type"`
	FilePath       string `json:"filePath"`
	CreatedAt      string `json:"createdAt"`
	CreatedBy      string `json:"createdBy"`
	LastModifiedAt string `json:"lastModifiedAt"`
	LastModifiedBy string `json:"lastModifiedBy"`
}

func (processor *Processor) GetID() string { return processor.ID }