package access_policies

import (
	"github.com/leanspace/terraform-provider-leanspace/helper/general_objects"
	"github.com/leanspace/terraform-provider-leanspace/provider"
)

func (ap *AccessPolicy) PostReadProcess(_ *provider.Client, newValue any) error {
	newAP, ok := newValue.(*AccessPolicy)
	if !ok || newAP == nil {
		return nil
	}
	newAP.Tags = general_objects.ReorderKeyValues(ap.Tags, newAP.Tags)
	return nil
}
