package generic_plugins

import (
	"io"

	"github.com/leanspace/terraform-provider-leanspace/helper"
)

func (genericPlugin *GenericPlugin) CustomEncoding(data []byte, isUpdating bool) (io.Reader, string, error) {
	multipartMap := map[string]any{
		"name":        genericPlugin.Name,
		"description": genericPlugin.Description,
		"type":        genericPlugin.Type,
		"language":    genericPlugin.Language,
	}
	return helper.FileAndDatasToMultipart(genericPlugin.FilePath, "sourceCode", multipartMap)
}
