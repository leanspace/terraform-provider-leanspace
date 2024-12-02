package generic_plugins

import (
	"github.com/leanspace/terraform-provider-leanspace/provider"
	"github.com/leanspace/terraform-provider-leanspace/services/plugins"
)

func (genericPlugin *GenericPlugin) ToMap() map[string]any {
	genericPluginMap := make(map[string]any)
	genericPluginMap["id"] = genericPlugin.ID
	genericPluginMap["name"] = genericPlugin.Name
	genericPluginMap["description"] = genericPlugin.Description
	genericPluginMap["type"] = genericPlugin.Type
	genericPluginMap["language"] = genericPlugin.Language
	genericPluginMap["source_code_link"] = []any{genericPlugin.SourceCodeLink.ToMap()}
	genericPluginMap["created_at"] = genericPlugin.CreatedAt
	genericPluginMap["created_by"] = genericPlugin.CreatedBy
	genericPluginMap["last_modified_at"] = genericPlugin.LastModifiedAt
	genericPluginMap["last_modified_by"] = genericPlugin.LastModifiedBy
	genericPluginMap["status"] = genericPlugin.Status
	genericPluginMap["source_code_path"] = genericPlugin.FilePath
	genericPluginMap["source_code_sha"] = genericPlugin.FileSha
	return genericPluginMap
}

func (sourceCodeLink *SourceCodeLink) ToMap() map[string]any {
	sourceCodeLinkMap := make(map[string]any)
	sourceCodeLinkMap["expiration_time"] = sourceCodeLink.ExpirationTime
	sourceCodeLinkMap["source_code_id"] = sourceCodeLink.SourceCodeId
	sourceCodeLinkMap["url"] = sourceCodeLink.Url
	return sourceCodeLinkMap
}

func (genericPlugin *GenericPlugin) FromMap(genericPluginMap map[string]any) error {
	genericPlugin.ID = genericPluginMap["id"].(string)
	genericPlugin.Name = genericPluginMap["name"].(string)
	genericPlugin.Description = genericPluginMap["description"].(string)
	genericPlugin.Type = genericPluginMap["type"].(string)
	genericPlugin.Language = genericPluginMap["language"].(string)
	if len(genericPluginMap["source_code_link"].([]any)) > 0 {
		if err := genericPlugin.SourceCodeLink.FromMap(genericPluginMap["source_code_link"].([]any)[0].(map[string]any)); err != nil {
			return err
		}
	}
	genericPlugin.CreatedAt = genericPluginMap["created_at"].(string)
	genericPlugin.CreatedBy = genericPluginMap["created_by"].(string)
	genericPlugin.LastModifiedAt = genericPluginMap["last_modified_at"].(string)
	genericPlugin.LastModifiedBy = genericPluginMap["last_modified_by"].(string)
	genericPlugin.Status = genericPluginMap["status"].(string)
	genericPlugin.FilePath = genericPluginMap["source_code_path"].(string)
	genericPlugin.FileSha = genericPluginMap["source_code_sha"].(string)
	return nil
}

func (sourceCodeLink *SourceCodeLink) FromMap(sourceCodeLinkMap map[string]any) error {
	sourceCodeLink.ExpirationTime = sourceCodeLinkMap["expiration_time"].(string)
	sourceCodeLink.SourceCodeId = sourceCodeLinkMap["source_code_id"].(string)
	sourceCodeLink.Url = sourceCodeLinkMap["url"].(string)
	return nil
}

func (genericPlugin *GenericPlugin) PostCreateProcess(client *provider.Client, destGenericPluginRaw any) error {
	return plugins.DoPostCreateProcess[*GenericPlugin](client, genericPlugin, GenericPluginDataType.Path, destGenericPluginRaw)
}

func (plugin *GenericPlugin) PostUpdateProcess(client *provider.Client, destPluginRaw any) error {
	return plugin.PostCreateProcess(client, destPluginRaw)
}

func (genericPlugin *GenericPlugin) PostReadProcess(client *provider.Client, destPluginRaw any) error {
	createdGenericPlugin := destPluginRaw.(*GenericPlugin)
	genericPlugin.SourceCodeLink = createdGenericPlugin.SourceCodeLink
	return plugins.DoPostReadProcess[*GenericPlugin](client, genericPlugin, destPluginRaw)
}
