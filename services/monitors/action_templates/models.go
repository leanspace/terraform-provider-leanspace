package action_templates

type ActionTemplate struct {
	ID             string            `json:"id"`
	Name           string            `json:"name"`
	Type           string            `json:"type"`
	URL            string            `json:"url,omitempty"`
	Payload        string            `json:"payload,omitempty"`
	Content        string            `json:"content,omitempty"`
	Headers        map[string]string `json:"headers,omitempty"`
	CreatedAt      string            `json:"createdAt"`
	CreatedBy      string            `json:"createdBy"`
	LastModifiedAt string            `json:"lastModifiedAt"`
	LastModifiedBy string            `json:"lastModifiedBy"`
}

func (actionTemplate *ActionTemplate) GetID() string { return actionTemplate.ID }
