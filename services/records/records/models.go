package records

import (
	"github.com/leanspace/terraform-provider-leanspace/helper/general_objects"
)

type Record struct {
	ID                   string                     `json:"id"`
	RecordTemplateId     string                     `json:"recordTemplateId"`
	Name                 string                     `json:"name"`
	RecordState          string                     `json:"state"`
	StartDateTime        string                     `json:"startDateTime"`
	StopDateTime         string                     `json:"stopDateTime"`
	StreamId             string                     `json:"streamId"`
	NodeIds              []string                   `json:"nodeIds"`
	MetricIds            []string                   `json:"metricIds"`
	Properties           []Property[any]            `json:"properties"`
	CommandDefinitionIds []string                   `json:"commandDefinitionIds"`
	Tags                 []general_objects.KeyValue `json:"tags,omitempty"`
	Comments             []string                   `json:"comments"`
	CreatedAt            string                     `json:"createdAt"`
	CreatedBy            string                     `json:"createdBy"`
	LastModifiedAt       string                     `json:"lastModifiedAt"`
	LastModifiedBy       string                     `json:"lastModifiedBy"`
}

func (record *Record) GetID() string { return record.ID }

type Property[T any] struct {
	Name       string                            `json:"name"`
	Attributes general_objects.ValueAttribute[T] `json:"attributes"`
}
