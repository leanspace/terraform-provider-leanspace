package streams_queue

import (
	"github.com/leanspace/terraform-provider-leanspace/helper"
	"github.com/leanspace/terraform-provider-leanspace/services/streams/streams"
)

func (streamQueue *streamQueue) ToMap() map[string]any {
	helper.Logger.Printf("toMap: streamQueue: %v", streamQueue.streamQueueId)
	helper.Logger.Printf("toMap: streamQueue: %v", streamQueue.stream)
	streamQueueMap := make(map[string]any)
	streamQueueMap["stream_queue_id"] = streamQueue.streamQueueId

	streamQueueMap["stream"] = streamQueue.stream.ToMap()
	helper.Logger.Printf("toMap inthe end: streamQueueMap: %v", streamQueueMap)

	return streamQueueMap
}

func (streamQueue *streamQueue) FromMap(streamQueueMap map[string]any) error {
	helper.Logger.Printf("fromMap: streamQueueMap: %v", streamQueueMap)
	helper.Logger.Printf("fromMap: streamQueueId: %v", streamQueueMap["stream_queue_id"])
	if streamQueueMap != nil {
		streamQueue.streamQueueId = streamQueueMap["stream_queue_id"].(string)
	}
	streamQueue.stream = streams.Stream{}
	err := streamQueue.stream.FromMap(streamQueueMap)
	if err != nil {
		return err
	}
	helper.Logger.Printf("fromMap in the end: streamQueue: %v", streamQueue)
	return nil
}
