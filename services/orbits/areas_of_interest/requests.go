package areas_of_interest

import (
	"github.com/leanspace/terraform-provider-leanspace/helper/general_objects"
	"github.com/leanspace/terraform-provider-leanspace/provider"
)

func (aoi *AreaOfInterest) PostReadProcess(_ *provider.Client, newValue any) error {
	newAOI, ok := newValue.(*AreaOfInterest)
	if !ok || newAOI == nil {
		return nil
	}
	newAOI.Tags = general_objects.ReorderKeyValues(aoi.Tags, newAOI.Tags)
	return nil
}
