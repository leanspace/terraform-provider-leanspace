package record_templates

import (
	"github.com/leanspace/terraform-provider-leanspace/helper/general_objects"
)

type RecordTemplate struct {
	ID             string                     `json:"id"`
	Name           string                     `json:"name"`
	Description    string                     `json:"description"`
	RecordState    string                     `json:"state"`
	StartDateTime  string                     `json:"startDateTime"`
	StopDateTime   string                     `json:"stopDateTime"`
	DefaultParsers []DefaultParser            `json:"defaultParsers"`
	Nodes          []Node                     `json:"nodes"`
	Properties     []Property                 `json:"properties"`
	Tags           []general_objects.KeyValue `json:"tags,omitempty"`
	CreatedAt      string                     `json:"createdAt"`
	CreatedBy      string                     `json:"createdBy"`
	LastModifiedAt string                     `json:"lastModifiedAt"`
	LastModifiedBy string                     `json:"lastModifiedBy"`
}

func (recordTemplate *RecordTemplate) GetID() string { return recordTemplate.ID }

type DefaultParser struct {
	// TODO
}

type Node struct {
	// TODO
}

type Property struct {
	// TODO
}
