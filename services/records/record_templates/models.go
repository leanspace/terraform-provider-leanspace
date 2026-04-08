package record_templates

import (
	"github.com/leanspace/terraform-provider-leanspace/helper/general_objects"
)

type RecordTemplate struct {
	general_objects.AuditModel
	Name                 string                     `json:"name"`
	Description          string                     `json:"description"`
	StreamId             string                     `json:"streamId"`
	DefaultParsers       []DefaultParser            `json:"defaultParsers"`
	NodeIds              []string                   `json:"nodeIds"`
	MetricIds            []string                   `json:"metricIds"`
	CommandDefinitionIds []string                   `json:"commandDefinitionIds"`
	Properties           []Property[any]            `json:"properties"`
	Tags                 []general_objects.KeyValue `json:"tags,omitempty"`
}

type DefaultParser struct {
	ID       string `json:"id"`
	FileType string `json:"fileType"`
}

type Property[T any] struct {
	Name       string                                 `json:"name"`
	Attributes general_objects.DefinitionAttribute[T] `json:"attributes"`
}
