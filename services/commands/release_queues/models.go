package release_queues

import (
	"github.com/leanspace/terraform-provider-leanspace/helper/general_objects"
)

type ReleaseQueue struct {
	general_objects.AuditModel
	AssetId                                   string                     `json:"assetId"`
	Name                                      string                     `json:"name"`
	Description                               string                     `json:"description"`
	CommandTransformerPluginId                string                     `json:"commandTransformerPluginId"`
	CommandTransformationStrategy             string                     `json:"commandTransformationStrategy"`
	CommandTransformerPluginConfigurationData string                     `json:"commandTransformerPluginConfigurationData"`
	GlobalTransmissionMetadata                []general_objects.KeyValue `json:"globalTransmissionMetadata"`
	LogicalLock                               bool                       `json:"logicalLock"`
	Tags                                      []general_objects.KeyValue `json:"tags,omitempty"`
}
