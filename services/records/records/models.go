package records

import (
	"github.com/leanspace/terraform-provider-leanspace/helper/general_objects"
)

type Record struct {
	ID               string                     `json:"id"`
	RecordTemplateId string                     `json:"recordTemplateId"`
	Name             string                     `json:"name"`
	RecordState      string                     `json:"state"`
	StartDateTime    string                     `json:"startDateTime"`
	StopDateTime     string                     `json:"stopDateTime"`
	Properties       []Property                 `json:"properties"`
	Tags             []general_objects.KeyValue `json:"tags,omitempty"`
	Comment          string                     `json:"comment"`
	CreatedAt        string                     `json:"createdAt"`
	CreatedBy        string                     `json:"createdBy"`
	LastModifiedAt   string                     `json:"lastModifiedAt"`
	LastModifiedBy   string                     `json:"lastModifiedBy"`
}

func (record *Record) GetID() string { return record.ID }

type Property struct {
	// TODO
}
