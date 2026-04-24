package resources

import (
	"fmt"

	"github.com/leanspace/terraform-provider-leanspace/helper"
	"github.com/leanspace/terraform-provider-leanspace/helper/general_objects"
	"github.com/leanspace/terraform-provider-leanspace/provider"
)

func (resource *Resource) PostReadProcess(_ *provider.Client, newValue any) error {
	newResource, ok := newValue.(*Resource)
	if !ok || newResource == nil {
		return nil
	}
	newResource.Tags = general_objects.ReorderKeyValues(resource.Tags, newResource.Tags)
	newResource.Thresholds = helper.ReorderByKey(resource.Thresholds, newResource.Thresholds,
		func(t ResourceThreshold) string { return fmt.Sprintf("%s:%v", t.Kind, t.Value) })
	return nil
}
