package metrics

import (
	"github.com/leanspace/terraform-provider-leanspace/helper/general_objects"
	"github.com/leanspace/terraform-provider-leanspace/provider"
)

func (metric *Metric[T]) PostReadProcess(_ *provider.Client, newValue any) error {
	newMetric, ok := newValue.(*Metric[T])
	if !ok || newMetric == nil {
		return nil
	}
	newMetric.Tags = general_objects.ReorderKeyValues(metric.Tags, newMetric.Tags)
	return nil
}
