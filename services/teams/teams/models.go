package teams

import "github.com/leanspace/terraform-provider-leanspace/helper/general_objects"

type Team struct {
	ID             string                     `json:"id"`
	Name           string                     `json:"name"`
	PolicyIds      []string                   `json:"policyIds"`
	Members        []string                   `json:"members"`
	Tags           []general_objects.KeyValue `json:"tags,omitempty"`
	CreatedAt      string                     `json:"createdAt"`
	CreatedBy      string                     `json:"createdBy"`
	LastModifiedAt string                     `json:"lastModifiedAt"`
	LastModifiedBy string                     `json:"lastModifiedBy"`
}

func (team *Team) GetID() string { return team.ID }
