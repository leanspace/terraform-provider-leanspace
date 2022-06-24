package asset

type GenericResourceType[T any] struct {
	Client *Client
	Path   string
}
