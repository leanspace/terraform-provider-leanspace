package record_templates

import (
	"github.com/leanspace/terraform-provider-leanspace/helper/general_objects"
)

type RecordTemplate struct {
	ID                   string                     `json:"id"`
	Name                 string                     `json:"name"`
	Description          string                     `json:"description"`
	State                string                     `json:"state"`
	StartDateTime        string                     `json:"startDateTime"`
	StopDateTime         string                     `json:"stopDateTime"`
	StreamId             string                     `json:"streamId"`
	DefaultParsers       []DefaultParser            `json:"defaultParsers"`
	NodeIds              []string                   `json:"nodeIds"`
	MetricIds            []string                   `json:"metricIds"`
	CommandDefinitionIds []string                   `json:"commandDefinitionIds"`
	Properties           []Property[any]            `json:"properties"`
	Tags                 []general_objects.KeyValue `json:"tags,omitempty"`
	CreatedAt            string                     `json:"createdAt"`
	CreatedBy            string                     `json:"createdBy"`
	LastModifiedAt       string                     `json:"lastModifiedAt"`
	LastModifiedBy       string                     `json:"lastModifiedBy"`
}

func (recordTemplate *RecordTemplate) GetID() string { return recordTemplate.ID }

type DefaultParser struct {
	ID       string `json:"id"`
	FileType string `json:"fileType"`
}

type Property[T any] struct {
	Name       string                                 `json:"name"`
	Attributes general_objects.DefinitionAttribute[T] `json:"attributes"`
}
