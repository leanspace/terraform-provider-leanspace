package processors

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"strings"
)

func (processor *Processor) CustomEncoding(data []byte) (io.Reader, string, error) {
	processorFile, err := os.Open(processor.FilePath)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, "", fmt.Errorf("file '%s' was not found", processor.FilePath)
		}
		return nil, "", err
	}

	var b bytes.Buffer
	formWriter := multipart.NewWriter(&b)

	// Add file field
	fileWriter, err := formWriter.CreateFormFile("file", processorFile.Name())
	if err != nil {
		return nil, "", err
	}
	_, err = io.Copy(fileWriter, processorFile)
	if err != nil {
		return nil, "", err
	}

	// Add data field
	dataWriter, err := formWriter.CreateFormField("command")
	if err != nil {
		return nil, "", err
	}
	_, err = io.Copy(dataWriter, strings.NewReader(string(data)))
	if err != nil {
		return nil, "", err
	}

	// Close the form and return
	formWriter.Close()
	return &b, formWriter.FormDataContentType(), nil
}