package record_templates

import (
	"github.com/leanspace/terraform-provider-leanspace/helper"
	"github.com/leanspace/terraform-provider-leanspace/helper/general_objects"
	"github.com/leanspace/terraform-provider-leanspace/provider"
)

func (recordTemplate *RecordTemplate) PostReadProcess(_ *provider.Client, newValue any) error {
	newRecordTemplate, ok := newValue.(*RecordTemplate)
	if !ok || newRecordTemplate == nil {
		return nil
	}
	newRecordTemplate.Tags = general_objects.ReorderKeyValues(recordTemplate.Tags, newRecordTemplate.Tags)
	newRecordTemplate.Properties = helper.ReorderByKey(recordTemplate.Properties, newRecordTemplate.Properties, func(p Property[any]) string { return p.Name })
	return nil
}
