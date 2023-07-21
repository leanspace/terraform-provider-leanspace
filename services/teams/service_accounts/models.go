package service_accounts

type ServiceAccount struct {
	ID             string      `json:"id"`
	Name           string      `json:"name"`
	PolicyIds      []string    `json:"policyIds"`
	Credendials    Credentials `json:"credentials"`
	CreatedAt      string      `json:"createdAt"`
	CreatedBy      string      `json:"createdBy"`
	LastModifiedAt string      `json:"lastModifiedAt"`
	LastModifiedBy string      `json:"lastModifiedBy"`
}

func (serviceAccount *ServiceAccount) GetID() string { return serviceAccount.ID }

type Credentials struct {
	ClientId     string `json:"clientId"`
	ClientSecret string `json:"clientSecret"`
}
