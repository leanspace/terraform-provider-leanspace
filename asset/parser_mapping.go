package asset

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type Parseable interface {
	// A function that converts the given map into a struct of this data type.
	FromMap(map[string]any) error
	// A function that converts the given struct of this data type into a map.
	ToMap() map[string]any
}

type ParseableModel[T any] interface {
	*T
	Parseable
	// A function that returns the ID of a model instance instance.
	GetID() string
}

type GenericClient[T any, PT ParseableModel[T]] struct {
	Client     *Client
	Path       string
	CreatePath func(PT) string
}

type DataSourceType[T any, PT ParseableModel[T]] struct {
	// Will be used in the terraform file!
	ResourceIdentifier string
	// Will be used in the terraform file!
	Name string
	// The path to which API requests are sent (e.g. "asset-repository/nodes")
	// This path will be used for all requests (GET/POST/PUT/DELETE), except if `CreatePath` is specified.
	Path string
	// Optional. A function that returns the path to which API *creation* requests are sent.
	// This can be useful when the path to create a resource depends on the resource's owner (e.g. "nodes/NODE_ID/properties")
	CreatePath func(PT) string
	// The schema to represent the data
	Schema map[string]*schema.Schema
}

func (dataSource DataSourceType[T, PT]) convert(client *Client) GenericClient[T, PT] {
	return GenericClient[T, PT]{
		Client:     client,
		Path:       dataSource.Path,
		CreatePath: dataSource.CreatePath,
	}
}
