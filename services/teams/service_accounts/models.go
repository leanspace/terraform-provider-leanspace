package service_accounts

type ServiceAccount struct {
	ID             string   `json:"id" terra:"id"`
	Name           string   `json:"name" terra:"name"`
	PolicyIds      []string `json:"policyIds" terra:"policy_ids"`
	CreatedAt      string   `json:"createdAt" terra:"created_at"`
	CreatedBy      string   `json:"createdBy" terra:"created_by"`
	LastModifiedAt string   `json:"lastModifiedAt" terra:"last_modified_at"`
	LastModifiedBy string   `json:"lastModifiedBy" terra:"last_modified_by"`
}

func (serviceAccount *ServiceAccount) GetID() string { return serviceAccount.ID }
