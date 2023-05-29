package processors

import "github.com/leanspace/terraform-provider-leanspace/provider"

func (processor *Processor) ToMap() map[string]any {
	processorMap := make(map[string]any)
	processorMap["id"] = processor.ID
	processorMap["name"] = processor.Name
	processorMap["description"] = processor.Description
	processorMap["version"] = processor.Version
	processorMap["type"] = processor.Type
	processorMap["file_path"] = processor.FilePath

	processorMap["created_at"] = processor.CreatedAt
	processorMap["created_by"] = processor.CreatedBy
	processorMap["last_modified_at"] = processor.LastModifiedAt
	processorMap["last_modified_by"] = processor.LastModifiedBy

	return processorMap
}

func (processor *Processor) FromMap(processorMap map[string]any) error {
	processor.ID = processorMap["id"].(string)
	processor.Name = processorMap["name"].(string)
	processor.Description = processorMap["description"].(string)
	processor.Version = processorMap["version"].(string)
	processor.Type = processorMap["type"].(string)
	processor.FilePath = processorMap["file_path"].(string)

	processor.CreatedAt = processorMap["created_at"].(string)
	processor.CreatedBy = processorMap["created_by"].(string)
	processor.LastModifiedAt = processorMap["last_modified_at"].(string)
	processor.LastModifiedBy = processorMap["last_modified_by"].(string)

	return nil
}

// Persist the file path - this data is not returned from the backend, so when the resource
// is loaded (from create/read/update) the path is empty, and so terraform thinks the field was
// changed. This workaround prevents the value from changing - it's loaded by terraform
// when reading the config and never changes again (except if the config changes).
func (processor *Processor) persistFilePath(destProcessor *Processor) error {
	destProcessor.FilePath = processor.FilePath
	return nil
}

func (processor *Processor) PostCreateProcess(_ *provider.Client, destProcessorRaw any) error {
	return processor.persistFilePath(destProcessorRaw.(*Processor))
}
func (processor *Processor) PostUpdateProcess(_ *provider.Client, destProcessorRaw any) error {
	return processor.persistFilePath(destProcessorRaw.(*Processor))
}
func (processor *Processor) PostReadProcess(_ *provider.Client, destProcessorRaw any) error {
	return processor.persistFilePath(destProcessorRaw.(*Processor))
}
