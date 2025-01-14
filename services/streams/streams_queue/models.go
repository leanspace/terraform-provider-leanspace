package streams_queue

import (
	"github.com/leanspace/terraform-provider-leanspace/services/streams/streams"
)

type streamQueue struct {
	stream        streams.Stream `json:"stream"`
	streamQueueId string         `json:"stream_queue_id"`
}

func (streamQueue *streamQueue) GetID() string { return streamQueue.stream.ID }
