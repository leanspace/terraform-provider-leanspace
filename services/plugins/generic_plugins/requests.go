package generic_plugins

import (
	"io"

	"github.com/leanspace/terraform-provider-leanspace/helper"
	"github.com/leanspace/terraform-provider-leanspace/provider"
	"github.com/leanspace/terraform-provider-leanspace/services/plugins"
)

func (genericPlugin *GenericPlugin) CustomEncoding(data []byte, isUpdating bool) (io.Reader, string, error) {
	multipartMap := map[string]any{
		"name":     genericPlugin.Name,
		"type":     genericPlugin.Type,
		"language": genericPlugin.Language,
	}
	if genericPlugin.Description != nil {
		multipartMap["description"] = *genericPlugin.Description
	}
	return helper.FileAndDatasToMultipart(genericPlugin.FilePath, "sourceCode", multipartMap)
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
