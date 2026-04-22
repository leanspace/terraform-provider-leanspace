package orbits

import (
	"github.com/leanspace/terraform-provider-leanspace/helper/general_objects"
	"github.com/leanspace/terraform-provider-leanspace/provider"
)

func (orbit *Orbit) PostReadProcess(_ *provider.Client, newValue any) error {
	newOrbit, ok := newValue.(*Orbit)
	if !ok || newOrbit == nil {
		return nil
	}
	newOrbit.Tags = general_objects.ReorderKeyValues(orbit.Tags, newOrbit.Tags)
	return nil
}
