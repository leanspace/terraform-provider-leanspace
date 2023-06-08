package plugins

import (
	"github.com/leanspace/terraform-provider-leanspace/helper"
	"io"
)

func (plugin *Plugin) CustomEncoding(data []byte) (io.Reader, string, error) {
	return helper.CustomEncoding(plugin.FilePath, data)
}
