package helper

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"os"
	"regexp"
	"strconv"
	"strings"
)

// ParseablePointer constrains a type parameter to be a pointer to T.
// Used by PaginatedList and GenericClient for JSON unmarshalling.
type ParseablePointer[T any] interface {
	*T
}

// ReorderByKey reorders apiItems to match the order of stateOrder using keyFn to
// derive a stable string key per element. Elements in apiItems not in stateOrder
// are appended at the end. Use this in PostReadProcess to eliminate perpetual diffs
// caused by the API returning list elements in a different order.
func ReorderByKey[T any](stateOrder []T, apiItems []T, keyFn func(T) string) []T {
	if len(stateOrder) == 0 {
		return apiItems
	}
	used := make([]bool, len(apiItems))
	reordered := make([]T, 0, len(apiItems))
	for _, state := range stateOrder {
		key := keyFn(state)
		for i, api := range apiItems {
			if !used[i] && keyFn(api) == key {
				reordered = append(reordered, api)
				used[i] = true
				break
			}
		}
	}
	for i, api := range apiItems {
		if !used[i] {
			reordered = append(reordered, api)
		}
	}
	return reordered
}

// Parses a float to a string. Use this method to ensure consistency.
func ParseFloat(num float64) string {
	return strconv.FormatFloat(num, 'g', -1, 64)
}

func Contains[T comparable](slice []T, value T) bool {
	for _, v := range slice {
		if v == value {
			return true
		}
	}
	return false
}

func AllowedValuesToDescription(allowedValues []string) string {
	return "it must be one of these values: " + strings.Join(allowedValues, ", ")
}

func FileAndDataToMultipart(filePath string, data []byte) (io.Reader, string, error) {
	return FileAndDatasToMultipart(filePath, "file", map[string]any{"command": string(data)})
}

func FileAndDatasToMultipart(filePath string, fileFieldName string, datas map[string]any) (io.Reader, string, error) {
	processorFile, err := os.Open(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, "", fmt.Errorf("file '%s' was not found", filePath)
		}
		return nil, "", err
	}

	var b bytes.Buffer
	formWriter := multipart.NewWriter(&b)

	// Add file field
	fileWriter, err := formWriter.CreateFormFile(fileFieldName, processorFile.Name())
	if err != nil {
		return nil, "", err
	}
	_, err = io.Copy(fileWriter, processorFile)
	if err != nil {
		return nil, "", err
	}

	// Add data fields
	for key, value := range datas {
		_, _, err := AddFormField(formWriter, key, fmt.Sprint(value))
		if err != nil {
			return nil, "", err
		}
	}

	// Close the form and return
	formWriter.Close()
	return &b, formWriter.FormDataContentType(), nil
}

func AddFormField(formWriter *multipart.Writer, fieldName string, field string) (io.Reader, string, error) {
	dataWriter, err := formWriter.CreateFormField(fieldName)
	if err != nil {
		return nil, "", err
	}
	_, err = io.Copy(dataWriter, strings.NewReader(field))
	if err != nil {
		return nil, "", err
	}
	return nil, "", nil
}

func SnakeToCamelCase(str string) string {
	parts := strings.Split(str, "_")
	base := strings.ToLower(parts[0])
	for i := 1; i < len(parts); i++ {
		base += strings.Title(parts[i])
	}
	return base
}

const Debug = false

var Logger = func() *log.Logger {
	if Debug {
		var logFile, _ = os.OpenFile("log_plug.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
		var logger = log.New(logFile, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
		return logger
	}
	log.SetOutput(io.Discard)
	return log.Default()
}()

var PathToJarFileRegex *regexp.Regexp = regexp.MustCompile(`^.*\.jar$`)
