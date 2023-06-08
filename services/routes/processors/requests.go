package processors

import (
	"github.com/leanspace/terraform-provider-leanspace/helper"
	"io"
)

func (processor *Processor) CustomEncoding(data []byte) (io.Reader, string, error) {
	return helper.CustomEncoding(processor.FilePath, data)
}
