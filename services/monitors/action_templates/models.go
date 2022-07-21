package action_templates

type ActionTemplate struct {
	ID             string            `json:"id" terra:"id"`
	Name           string            `json:"name" terra:"name"`
	Type           string            `json:"type" terra:"type"`
	URL            string            `json:"url" terra:"url"`
	Payload        string            `json:"payload" terra:"payload"`
	Headers        map[string]string `json:"headers" terra:"headers"`
	CreatedAt      string            `json:"createdAt" terra:"created_at"`
	CreatedBy      string            `json:"createdBy" terra:"created_by"`
	LastModifiedAt string            `json:"lastModifiedAt" terra:"last_modified_at"`
	LastModifiedBy string            `json:"lastModifiedBy" terra:"last_modified_by"`
}

func (actionTemplate *ActionTemplate) GetID() string { return actionTemplate.ID }
