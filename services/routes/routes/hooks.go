package routes

import (
	"github.com/leanspace/terraform-provider-leanspace/helper/general_objects"
	"github.com/leanspace/terraform-provider-leanspace/provider"
)

func (route *Route) PostReadProcess(_ *provider.Client, newValue any) error {
	newRoute, ok := newValue.(*Route)
	if !ok || newRoute == nil {
		return nil
	}
	newRoute.Tags = general_objects.ReorderKeyValues(route.Tags, newRoute.Tags)
	return nil
}
