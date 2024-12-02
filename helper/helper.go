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

type Parseable interface {
	// A function that converts the given map into a struct of this data type.
	FromMap(map[string]any) error
	// A function that converts the given struct of this data type into a map.
	ToMap() map[string]any
}

type ParseablePointer[T any] interface {
	*T
	Parseable
}

func ParseToMaps[T any, PT ParseablePointer[T]](parseables []T) []map[string]any {
	maps := make([]map[string]any, len(parseables))
	for index, value := range parseables {
		var pointer PT = &value
		maps[index] = pointer.ToMap()
	}
	return maps
}

func ParseFromMaps[T any, PT ParseablePointer[T]](maps []any) ([]T, error) {
	parseables := make([]T, len(maps))
	for index, m := range maps {
		var pointer PT = &parseables[index]
		if err := pointer.FromMap(m.(map[string]any)); err != nil {
			return parseables, err
		}
	}
	return parseables, nil
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

func AllowedIntValuesToDescription(allowedValues []int) string {
	return "it must be one of these values: " + strings.Trim(strings.Join(strings.Fields(fmt.Sprint(allowedValues)), ", "), "[]")
}

func Ptr[T any](value T) *T {
	return &value
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

func Implements[T any, I any]() bool {
	var ptr *T
	_, isInstance := any(ptr).(I)
	return isInstance
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
