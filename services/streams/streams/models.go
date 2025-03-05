package streams

import "github.com/leanspace/terraform-provider-leanspace/helper"

type Stream struct {
	ID             string        `json:"id"`
	Version        int           `json:"version"`
	Name           string        `json:"name"`
	Description    string        `json:"description"`
	AssetId        string        `json:"assetId"`
	Configuration  Configuration `json:"configuration"`
	Mappings       []Mapping     `json:"mappings"`
	CreatedAt      string        `json:"createdAt"`
	CreatedBy      string        `json:"createdBy"`
	LastModifiedAt string        `json:"lastModifiedAt"`
	LastModifiedBy string        `json:"lastModifiedBy"`
}

func (stream *Stream) GetID() string { return stream.ID }

type Configuration struct {
	Endianness   string                                         `json:"endianness"`
	Structure    ElementList[StreamComponent, *StreamComponent] `json:"structure"`
	Metadata     Metadata                                       `json:"metadata"`
	Computations ElementList[Computation, *Computation]         `json:"computations"`
}

type StreamComponent struct {
	// Common
	Name       string      `json:"name"`
	Order      int         `json:"order"`
	Path       string      `json:"path"`
	Type       string      `json:"type"` // CONTAINER / FIELD / SWITCH
	Repetitive *Repetitive `json:"repetitive,omitempty"`

	// Field only
	Length     *Length `json:"length,omitempty"`
	Processor  string  `json:"processor,omitempty"`
	DataType   string  `json:"dataType,omitempty"`
	Endianness string  `json:"endianness,omitempty"`

	// Switch only
	Expression SwitchExpression `json:"expression,omitempty"`

	// Container and switch only
	Elements []StreamComponent `json:"elements,omitempty"`
}

type SwitchExpression struct {
	SwitchOn string         `json:"switchOn"`
	Options  []SwitchOption `json:"options"`
}

type SwitchOption struct {
	Value     SwitchValue[any] `json:"value"`
	Component string           `json:"component"`
}

type SwitchValue[T any] struct {
	DataType string `json:"dataType"`
	Data     T      `json:"data"`
}

type Metadata struct {
	Timestamp TimestampDefinition `json:"timestamp"`
}

type TimestampDefinition struct {
	Expression string `json:"expression"`
}

type ElementList[T any, PT helper.ParseablePointer[T]] struct {
	Elements []T `json:"elements"`
}

type Computation struct {
	Name       string `json:"name"`
	Order      int    `json:"order"`
	Type       string `json:"type"`
	DataType   string `json:"dataType"`
	Expression string `json:"expression"`
}

type Mapping struct {
	MetricId   string `json:"metricId"`
	Expression string `json:"expression,omitempty"`
}

type Repetitive struct {
	// Fixed
	Value int `json:"value,omitempty"`

	// Dynamic
	Path string `json:"path,omitempty"`
}

type Length struct {
	Type string `json:"type"` // [FIXED | DYNAMIC]
	Unit string `json:"unit"` // [BITS | BYTES]

	// Fixed
	Value int `json:"value,omitempty"`

	// Dynamic
	Path string `json:"path,omitempty"`
}
