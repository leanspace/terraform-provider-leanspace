package release_queues

import (
	"github.com/leanspace/terraform-provider-leanspace/helper/general_objects"
)

type ReleaseQueue struct {
	ID                                        string                     `json:"id"`
	AssetId                                   string                     `json:"assetId"`
	Name                                      string                     `json:"name"`
	Description                               string                     `json:"description"`
	CommandTransformerPluginId                string                     `json:"commandTransformerPluginId"`
	CommandTransformationStrategy             string                     `json:"commandTransformationStrategy"`
	CommandTransformerPluginConfigurationData string                     `json:"commandTransformerPluginConfigurationData"`
	GlobalTransmissionMetadata                []general_objects.KeyValue `json:"globalTransmissionMetadata"`
	LogicalLock                               bool                       `json:"logicalLock"`
	CreatedAt                                 string                     `json:"createdAt"`
	CreatedBy                                 string                     `json:"createdBy"`
	LastModifiedAt                            string                     `json:"lastModifiedAt"`
	LastModifiedBy                            string                     `json:"lastModifiedBy"`
}

func (queue *ReleaseQueue) GetID() string { return queue.ID }
