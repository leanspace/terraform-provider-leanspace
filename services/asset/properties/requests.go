package properties

import (
	"github.com/leanspace/terraform-provider-leanspace/helper/general_objects"
	"github.com/leanspace/terraform-provider-leanspace/provider"
)

func (property *Property[T]) PostReadProcess(_ *provider.Client, newValue any) error {
	newProperty, ok := newValue.(*Property[T])
	if !ok || newProperty == nil {
		return nil
	}
	newProperty.Tags = general_objects.ReorderKeyValues(property.Tags, newProperty.Tags)
	return nil
}
