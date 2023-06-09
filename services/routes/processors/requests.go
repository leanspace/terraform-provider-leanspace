package processors

import (
	"io"
	"strings"

	"github.com/leanspace/terraform-provider-leanspace/helper"
)

func (processor *Processor) CustomEncoding(data []byte, isUpdating bool) (io.Reader, string, error) {
	if isUpdating {
		return strings.NewReader(string(data)), "application/json", nil
	}
	return helper.FileAndDataToMultipart(processor.FilePath, data)
}
