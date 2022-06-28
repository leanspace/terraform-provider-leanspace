package asset

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type GenericResourceType[T any] struct {
	Client     *Client
	Path       string
	CreatePath func(T) string
}

type DataSourceType[T any] struct {
	// Will be used in the terraform file!
	ResourceIdentifier string
	// Will be used in the terraform file!
	Name string
	// The path to which API requests are sent (e.g. "asset-repository/nodes")
	// This path will be used for all requests (GET/POST/PUT/DELETE), except if `CreatePath` is specified.
	Path string
	// Optional. A function that returns the path to which API *creation* requests are sent.
	// This can be useful when the path to create a resource depends on the resource's owner (e.g. "nodes/NODE_ID/properties")
	CreatePath func(T) string

	// The schema to represent the data
	Schema map[string]*schema.Schema

	// A function that returns the ID of a data type instance.
	GetID func(*T) string
	// A function that converts the given map into a struct of this data type.
	MapToStruct func(map[string]any) (T, error)
	// A function that converts the given struct of this data type into a map.
	StructToMap func(T) map[string]any
}

func (dataSource DataSourceType[T]) convert(client *Client) GenericResourceType[T] {
	return GenericResourceType[T]{
		Client:     client,
		Path:       dataSource.Path,
		CreatePath: dataSource.CreatePath,
	}
}
