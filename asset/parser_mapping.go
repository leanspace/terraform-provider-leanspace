package asset

import (
	"io"

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

type ExtraMarshallModel interface {
	// An optional extra function that is called before the object is marshalled.
	// This can be useful to encode specific fields to an API-compatible format, or to extrapolate optional data.
	// This function is guaranteed to only be called once for the instance of this model.
	PreMarshallProcess() error
}

type ExtraUnmarshallModel interface {
	// An optional extra function that is called after the object was unmarshalled.
	// This can be useful to decode specific fields from an API-compatible format, or to extrapolate optional data.
	// This function is guaranteed to only be called once for the instance of this model.
	PostUnmarshallProcess() error
}

type PostCreateModel interface {
	// An optional extra function that is called after this instance is created remotely by terraform.
	// Extra requests can be done here, as this method is exclusively called when the resource is created (unlike
	// PostUnmarshallProcess).
	// The parameter is the instance of the model that was created - it can be used to compare the desired result
	// with what was actually created.
	PostCreateProcess(*Client, any) error
}

type PostReadModel interface {
	// An optional extra function that is called after this instance was read remotely by terraform.
	// Extra requests (e.g. extra data fetching) can be done here, as this method is exclusively called after the resource
	// is read, and changes done to this instance will be persisted when saving to the state.
	// The parameter is the instance of the model after being updated. It can be used to compare the desired
	// state with what is currently present.
	PostReadProcess(*Client, any) error
}

type PostUpdateModel interface {
	// An optional extra function that is called after this instance was updated remotely by terraform.
	// Extra requests can be done here, as this method is exclusively called after the resource is updated (unlike
	// PostUnmarshallProcess).
	// The parameter is the instance of the model after being updated. It can be used to compare the desired
	// state with what is currently present.
	PostUpdateProcess(*Client, any) error
}

type PostDeleteModel interface {
	// An optional extra function that is called after this instance was delete remotely by terraform.
	// Extra requests (e.g. a cleanup) can be done here, as this method is exclusively called after the resource
	// is successfuly deleted.
	PostDeleteProcess(*Client) error
}

type CustomEncodingModel interface {
	// An optional extra function that is called when this instance needs to be encoded by terraform
	// for a request (when creating or updating).
	// If this is implemented, it will replace the default body and content types of the request.
	// This can be useful to properly encode multipart data, for instance.
	// The parameters are the JSON encoded representation of the model, and the client used.
	// It must return a reader to the body, the content type, and possibly an error.
	CustomEncoding([]byte) (io.Reader, string, error)
}

type GenericClient[T any, PT ParseableModel[T]] struct {
	Client     *Client
	Path       string
	CreatePath func(PT) string
	ReadPath   func(string) string
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
	// Optional. A function that returns the path to which API reading requests are sent.
	// This can be useful when the path to read from has extra subpaths (e.g. "plugins/PLUGIN_ID/metadata")
	ReadPath func(string) string
	// The schema to represent the data
	Schema map[string]*schema.Schema
}

func (dataSource DataSourceType[T, PT]) convert(client *Client) GenericClient[T, PT] {
	return GenericClient[T, PT]{
		Client:     client,
		Path:       dataSource.Path,
		CreatePath: dataSource.CreatePath,
		ReadPath:   dataSource.ReadPath,
	}
}
