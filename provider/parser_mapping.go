package provider

import (
	"io"

	datasourceschema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	resourceschema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

type ParseableModel[T any] interface {
	*T
	// A function that returns the ID of a model instance.
	GetID() string
}

// into JSON, XML,
type ExtraMarshallModel interface {
	// An optional extra function that is called before the object is marshalled.
	// This can be useful to encode specific fields to an API-compatible format, or to extrapolate optional data.
	// This function is guaranteed to only be called once for the instance of this model.
	PreMarshallProcess() error
}

// into Go objects
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

type PreDeleteModel interface {
	// An optional extra function that is called before this instance was delete remotely by terraform.
	// Extra requests (e.g. detaching linked entity) can be done here to allow for the deletion of the resource,
	// as this method is exclusively called before the resource is deleted..
	PreDeleteProcess(*Client, any) error
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
	CustomEncoding([]byte, bool) (io.Reader, string, error)
}

type ValidationModel interface {
	// An optional extra function that is called before creating or updating.
	// This can be used to ensure all values are valid and coherent.
	// If an error is thrown the action is stopped and the error is displayed to the user.
	Validate() error
}

// APIToTF is implemented by API models that can produce a Terraform model directly,
// bypassing the intermediate map[string]any layer.
type APIToTF interface {
	ToTF() any // returns a pointer to the TF model struct
}

// TFToAPI is implemented by Terraform models that can produce their API model,
// bypassing the intermediate map[string]any layer.
type TFToAPI interface {
	ToAPI() any // returns the API model (value, not pointer)
}

// APIToDSTF is implemented by API models that can produce a data-source-specific
// Terraform model. Used by unique (non-paginated) data sources.
type APIToDSTF interface {
	ToDSTF() any // returns a pointer to the DS TF model struct
}

type GenericClient[T any, PT ParseableModel[T]] struct {
	Client         *Client
	Path           string
	CreatePath     func(PT) string
	ReadPath       func(string) string
	DeletePath     func(string) string
	UpdatePath     func(string) string
	UpdateFunction func(*Client, string, PT) (PT, error)
	CreateFunction func(*Client, PT) (PT, error)
	IsUnique       bool `default:"false"`
}

type DataSourceType[T any, PT ParseableModel[T]] struct {
	// Will be used in the terraform file!
	ResourceIdentifier string
	// The path to which API requests are sent (e.g. "asset-repository/nodes")
	// This path will be used for all requests (GET/POST/PUT/DELETE), except if `CreatePath` is specified.
	Path string
	// Optional. A function that returns the path to which API *creation* requests are sent.
	// This can be useful when the path to create a resource depends on the resource's owner (e.g. "nodes/NODE_ID/properties")
	CreatePath func(PT) string
	// Optional. A function that returns the path to which API reading requests are sent.
	// This can be useful when the path to read from has extra subpaths (e.g. "plugins/PLUGIN_ID/metadata")
	ReadPath func(string) string
	// Optional. A function that returns the path to which API delete requests are sent.
	// This can be useful when the path to delete from has extra subpaths (e.g. "LEAFSPACE/delete")
	DeletePath func(string) string
	// Optional. A function that returns the path to which API update requests are sent.
	// This can be useful when the path to update from has extra subpaths (e.g. "LEAFSPACE/update")
	UpdatePath func(string) string
	// Optional. A function that is called when an update is requested. This can be useful when the update
	// request is different from the default Update request.
	UpdateFunction func(*Client, string, PT) (PT, error)
	// Optional. A function that is called when an update is requested. This can be useful when the update
	// request is different from the default Create request.
	CreateFunction func(*Client, PT) (PT, error)
	// The schema to represent the data as a managed resource
	Schema map[string]resourceschema.Attribute
	// The filters used for this resource's data source. The only allowed fields are primitives and lists of
	// strings. Note that some fields are already declared and don't need to be redefined: ids, query, page, size, sort.
	// A value of nil is treated as an empty map, and only the fields specified previously will be usable.
	FilterSchema map[string]datasourceschema.Attribute
	// If the filet endpoint is paginated or not. Defaults to true.
	IsUnique bool `default:"false"`
	// Factory that returns a pointer to a new empty TF model struct (e.g. &NodeTF{}).
	// GenericResource uses the direct TF conversion path (Plan.Get/State.Set)
	TFModelFactory func() any
}

func (dataSource DataSourceType[T, PT]) convert(client *Client) GenericClient[T, PT] {
	return GenericClient[T, PT]{
		Client:         client,
		Path:           dataSource.Path,
		CreatePath:     dataSource.CreatePath,
		ReadPath:       dataSource.ReadPath,
		DeletePath:     dataSource.DeletePath,
		UpdatePath:     dataSource.UpdatePath,
		UpdateFunction: dataSource.UpdateFunction,
		CreateFunction: dataSource.CreateFunction,
		IsUnique:       dataSource.IsUnique,
	}
}
