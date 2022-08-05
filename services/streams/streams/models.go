package streams

import "leanspace-terraform-provider/helper"

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
	Valid        bool                                           `json:"valid,omitempty" terra:"valid,omitempty"`
	Errors       []Error                                        `json:"errors,omitempty" terra:"errors,omitempty"`
}

type StreamComponent struct {
	// Common
	Name   string  `json:"name"`
	Order  int     `json:"order"`
	Path   string  `json:"path"`
	Type   string  `json:"type"` // CONTAINER / FIELD / SWITCH
	Valid  bool    `json:"valid,omitempty" terra:"valid,omitempty"`
	Errors []Error `json:"errors,omitempty" terra:"errors,omitempty"`

	// Field only
	LengthInBits int    `json:"lengthInBits,omitempty" terra:"length_in_bits,omitempty"`
	Processor    string `json:"processor,omitempty" terra:"processor,omitempty"`
	DataType     string `json:"dataType,omitempty" terra:"data_type,omitempty"`
	Endianness   string `json:"endianness,omitempty" terra:"endianness,omitempty"`

	// Switch only
	Expression SwitchExpression `json:"expression,omitempty" terra:"expression,omitempty"`

	// Container and switch only
	Elements []StreamComponent `json:"elements,omitempty" terra:"elements,omitempty"`
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
	PacketID  ElementStatus       `json:"packetId,omitempty" terra:"packet_id,omitempty"`
	Timestamp TimestampDefinition `json:"timestamp"`
	Valid     bool                `json:"valid,omitempty" terra:"valid,omitempty"`
	Errors    []Error             `json:"errors,omitempty" terra:"errors,omitempty"`
}

type TimestampDefinition struct {
	Expression string  `json:"expression"`
	Valid      bool    `json:"valid,omitempty" terra:"valid,omitempty"`
	Errors     []Error `json:"errors,omitempty" terra:"errors,omitempty"`
}

type ElementList[T any, PT helper.ParseablePointer[T]] struct {
	Elements []T     `json:"elements"`
	Valid    bool    `json:"valid,omitempty" terra:"valid,omitempty"`
	Errors   []Error `json:"errors,omitempty" terra:"errors,omitempty"`
}

type Computation struct {
	Name       string  `json:"name"`
	Order      int     `json:"order"`
	Type       string  `json:"type"`
	DataType   string  `json:"dataType"`
	Expression string  `json:"expression"`
	Valid      bool    `json:"valid,omitempty" terra:"valid,omitempty"`
	Errors     []Error `json:"errors,omitempty" terra:"errors,omitempty"`
}

type Mapping struct {
	MetricId  string `json:"metricId"`
	Component string `json:"component"`
}

type ElementStatus struct {
	Valid  bool    `json:"valid,omitempty" terra:"valid,omitempty"`
	Errors []Error `json:"errors,omitempty" terra:"errors,omitempty"`
}

type Error struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}
