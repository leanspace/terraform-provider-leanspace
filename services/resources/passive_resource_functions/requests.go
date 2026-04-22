package passive_resource_functions

import (
	"github.com/leanspace/terraform-provider-leanspace/helper/general_objects"
	"github.com/leanspace/terraform-provider-leanspace/provider"
)

func (prf *PassiveResourceFunction) PostReadProcess(_ *provider.Client, newValue any) error {
	newPRF, ok := newValue.(*PassiveResourceFunction)
	if !ok || newPRF == nil {
		return nil
	}
	newPRF.Tags = general_objects.ReorderKeyValues(prf.Tags, newPRF.Tags)
	return nil
}
