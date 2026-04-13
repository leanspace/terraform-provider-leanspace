package generic_plugins

import (
	"github.com/leanspace/terraform-provider-leanspace/helper"
	"github.com/leanspace/terraform-provider-leanspace/provider"
	"github.com/leanspace/terraform-provider-leanspace/services/plugins"
)

func (genericPlugin *GenericPlugin) ToMap() map[string]any {
	genericPluginMap := genericPlugin.ToAuditMap()
	genericPluginMap["name"] = helper.NilIfEmpty(genericPlugin.Name)
	genericPluginMap["description"] = helper.NilIfEmpty(genericPlugin.Description)
	genericPluginMap["type"] = helper.NilIfEmpty(genericPlugin.Type)
	genericPluginMap["language"] = helper.NilIfEmpty(genericPlugin.Language)
	genericPluginMap["source_code_link"] = []any{genericPlugin.SourceCodeLink.ToMap()}
	genericPluginMap["status"] = helper.NilIfEmpty(genericPlugin.Status)
	genericPluginMap["source_code_path"] = helper.NilIfEmpty(genericPlugin.FilePath)
	genericPluginMap["source_code_sha"] = helper.NilIfEmpty(genericPlugin.FileSha)
	return genericPluginMap
}

func (sourceCodeLink *SourceCodeLink) ToMap() map[string]any {
	sourceCodeLinkMap := make(map[string]any)
	sourceCodeLinkMap["expiration_time"] = helper.NilIfEmpty(sourceCodeLink.ExpirationTime)
	sourceCodeLinkMap["source_code_id"] = helper.NilIfEmpty(sourceCodeLink.SourceCodeId)
	sourceCodeLinkMap["url"] = helper.NilIfEmpty(sourceCodeLink.Url)
	return sourceCodeLinkMap
}

func (genericPlugin *GenericPlugin) FromMap(genericPluginMap map[string]any) error {
	genericPlugin.FromAuditMap(genericPluginMap)
	genericPlugin.Name = helper.CastString(genericPluginMap, "name")
	genericPlugin.Description = helper.CastString(genericPluginMap, "description")
	genericPlugin.Type = helper.CastString(genericPluginMap, "type")
	genericPlugin.Language = helper.CastString(genericPluginMap, "language")
	if len(helper.CastSlice(genericPluginMap, "source_code_link")) > 0 {
		if err := genericPlugin.SourceCodeLink.FromMap(helper.CastSlice(genericPluginMap, "source_code_link")[0].(map[string]any)); err != nil {
			return err
		}
	}
	genericPlugin.Status = helper.CastString(genericPluginMap, "status")
	genericPlugin.FilePath = helper.CastString(genericPluginMap, "source_code_path")
	genericPlugin.FileSha = helper.CastString(genericPluginMap, "source_code_sha")
	return nil
}

func (sourceCodeLink *SourceCodeLink) FromMap(sourceCodeLinkMap map[string]any) error {
	sourceCodeLink.ExpirationTime = helper.CastString(sourceCodeLinkMap, "expiration_time")
	sourceCodeLink.SourceCodeId = helper.CastString(sourceCodeLinkMap, "source_code_id")
	sourceCodeLink.Url = helper.CastString(sourceCodeLinkMap, "url")
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
