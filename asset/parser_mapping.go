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
	ResourceIdentifier string // Will be used in the terraform file!
	Name               string // Will be used in the terraform file!
	Path               string
	CreatePath         func(T) string

	Schema     map[string]*schema.Schema
	RootSchema map[string]*schema.Schema

	GetID       func(*T) string
	MapToStruct func(map[string]any) (T, error)
	StructToMap func(T) map[string]any
}

func (dataSource DataSourceType[T]) convert(client *Client) GenericResourceType[T] {
	return GenericResourceType[T]{
		Client:     client,
		Path:       dataSource.Path,
		CreatePath: dataSource.CreatePath,
	}
}
