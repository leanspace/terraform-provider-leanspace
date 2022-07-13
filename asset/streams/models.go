package streams

import "terraform-provider-asset/asset"

type Stream struct {
	ID             string        `json:"id" terra:"id"`
	Version        int           `json:"version" terra:"version"`
	Name           string        `json:"name" terra:"name"`
	Description    string        `json:"description" terra:"description"`
	AssetId        string        `json:"assetId" terra:"asset_id"`
	Configuration  Configuration `json:"configuration" terra:"configuration"`
	Mappings       []Mapping     `json:"mappings" terra:"mappings"`
	CreatedAt      string        `json:"createdAt" terra:"created_at"`
	CreatedBy      string        `json:"createdBy" terra:"created_by"`
	LastModifiedAt string        `json:"lastModifiedAt" terra:"last_modified_at"`
	LastModifiedBy string        `json:"lastModifiedBy" terra:"last_modified_by"`
}

func (stream *Stream) GetID() string { return stream.ID }

type Configuration struct {
	Endianness   string                                         `json:"endianness" terra:"endianness"`
	Structure    ElementList[StreamComponent, *StreamComponent] `json:"structure" terra:"structure"`
	Metadata     Metadata                                       `json:"metadata" terra:"metadata"`
	Computations ElementList[Computation, *Computation]         `json:"computations" terra:"computations"`
	Valid        bool                                           `json:"valid,omitempty" terra:"valid,omitempty"`
	Errors       []Error                                        `json:"errors,omitempty" terra:"errors,omitempty"`
}

type StreamComponent struct {
	// Common
	Name   string  `json:"name" terra:"name"`
	Order  int     `json:"order" terra:"order"`
	Path   string  `json:"path" terra:"path"`
	Type   string  `json:"type" terra:"type"` // CONTAINER / FIELD / SWITCH
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
	SwitchOn string         `json:"switchOn" terra:"switch_on"`
	Options  []SwitchOption `json:"options" terra:"options"`
}

type SwitchOption struct {
	Value     SwitchValue[any] `json:"value" terra:"value"`
	Component string           `json:"component" terra:"component"`
}

type SwitchValue[T any] struct {
	DataType string `json:"dataType" terra:"data_type"`
	Data     T      `json:"data" terra:"data"`
}

type Metadata struct {
	PacketID  ElementStatus       `json:"packetId,omitempty" terra:"packet_id,omitempty"`
	Timestamp TimestampDefinition `json:"timestamp" terra:"timestamp"`
	Valid     bool                `json:"valid,omitempty" terra:"valid,omitempty"`
	Errors    []Error             `json:"errors,omitempty" terra:"errors,omitempty"`
}

type TimestampDefinition struct {
	Expression string  `json:"expression" terra:"expression"`
	Valid      bool    `json:"valid,omitempty" terra:"valid,omitempty"`
	Errors     []Error `json:"errors,omitempty" terra:"errors,omitempty"`
}

type ElementList[T any, PT asset.ParseablePointer[T]] struct {
	Elements []T     `json:"elements" terra:"elements"`
	Valid    bool    `json:"valid,omitempty" terra:"valid,omitempty"`
	Errors   []Error `json:"errors,omitempty" terra:"errors,omitempty"`
}

type Computation struct {
	Name       string  `json:"name" terra:"name"`
	Order      int     `json:"order" terra:"order"`
	Type       string  `json:"type" terra:"type"`
	DataType   string  `json:"dataType" terra:"data_type"`
	Expression string  `json:"expression" terra:"expression"`
	Valid      bool    `json:"valid,omitempty" terra:"valid,omitempty"`
	Errors     []Error `json:"errors,omitempty" terra:"errors,omitempty"`
}

type Mapping struct {
	MetricId  string `json:"metricId" terra:"metric_id"`
	Component string `json:"component" terra:"component"`
}

type ElementStatus struct {
	Valid  bool    `json:"valid,omitempty" terra:"valid,omitempty"`
	Errors []Error `json:"errors,omitempty" terra:"errors,omitempty"`
}

type Error struct {
	Code    string `json:"code" terra:"code"`
	Message string `json:"message" terra:"message"`
}
