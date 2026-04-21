package release_queues

import (
	"github.com/leanspace/terraform-provider-leanspace/helper/general_objects"
	"github.com/leanspace/terraform-provider-leanspace/provider"
)

func (queue *ReleaseQueue) PostReadProcess(_ *provider.Client, newValue any) error {
	newQueue, ok := newValue.(*ReleaseQueue)
	if !ok || newQueue == nil {
		return nil
	}
	newQueue.Tags = general_objects.ReorderKeyValues(queue.Tags, newQueue.Tags)
	return nil
}
