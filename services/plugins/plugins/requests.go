package plugins

import (
	"io"

	"github.com/leanspace/terraform-provider-leanspace/helper"
)

func (plugin *Plugin) CustomEncoding(data []byte, isUpdating bool) (io.Reader, string, error) {
	return helper.FileAndDataToMultipart(plugin.FilePath, data)
}
