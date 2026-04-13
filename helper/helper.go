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
	parseables := make([]T, 0, len(maps))
	for _, m := range maps {
		if m == nil {
			continue
		}
		var t T
		var pointer PT = &t
		if err := pointer.FromMap(m.(map[string]any)); err != nil {
			return parseables, err
		}
		parseables = append(parseables, t)
	}
	return parseables, nil
}

// CastString safely extracts a string from a map. Returns "" if the key is missing or nil.
func CastString(m map[string]any, key string) string {
	if v, ok := m[key]; ok && v != nil {
		if s, ok := v.(string); ok {
			return s
		}
		return fmt.Sprint(v)
	}
	return ""
}

// CastBool safely extracts a bool from a map. Returns false if the key is missing or nil.
func CastBool(m map[string]any, key string) bool {
	if v, ok := m[key]; ok && v != nil {
		if b, ok := v.(bool); ok {
			return b
		}
	}
	return false
}

// CastInt safely extracts an int from a map. Returns 0 if the key is missing or nil.
// Handles int, int64, and float64 (JSON numbers).
func CastInt(m map[string]any, key string) int {
	if v, ok := m[key]; ok && v != nil {
		switch n := v.(type) {
		case int:
			return n
		case int64:
			return int(n)
		case float64:
			return int(n)
		}
	}
	return 0
}

// CastFloat64 safely extracts a float64 from a map. Returns 0.0 if the key is missing or nil.
func CastFloat64(m map[string]any, key string) float64 {
	if v, ok := m[key]; ok && v != nil {
		switch n := v.(type) {
		case float64:
			return n
		case int:
			return float64(n)
		case int64:
			return float64(n)
		}
	}
	return 0.0
}

// CastSlice safely extracts a []any from a map. Returns []any{} if the key is missing, nil, or wrong type.
func CastSlice(m map[string]any, key string) []any {
	if v, ok := m[key]; ok && v != nil {
		switch s := v.(type) {
		case []any:
			return s
		case []map[string]any:
			result := make([]any, len(s))
			for i, item := range s {
				result[i] = item
			}
			return result
		case []string:
			result := make([]any, len(s))
			for i, item := range s {
				result[i] = item
			}
			return result
		}
	}
	return []any{}
}

// CastMapAny safely extracts a map[string]any from a map. Returns map[string]any{} if missing or nil.
func CastMapAny(m map[string]any, key string) map[string]any {
	if v, ok := m[key]; ok && v != nil {
		if m2, ok := v.(map[string]any); ok {
			return m2
		}
	}
	return map[string]any{}
}

// NilIfEmpty converts an empty string to nil; all other values pass through unchanged.
// Using any as the parameter type makes it safe to apply to any struct field in ToMap()
// regardless of its type — non-string values are returned as-is.
func NilIfEmpty(v any) any {
	if s, ok := v.(string); ok && s == "" {
		return nil
	}
	return v
}

// CastIntPtr safely extracts an int from a map, returning nil if the key is absent or null.
// Use this for Optional model fields typed *int so that "not set" is preserved as nil
// rather than collapsed to the zero value.
func CastIntPtr(m map[string]any, key string) *int {
	v, ok := m[key]
	if !ok || v == nil {
		return nil
	}
	switch n := v.(type) {
	case int:
		return &n
	case int64:
		i := int(n)
		return &i
	case float64:
		i := int(n)
		return &i
	}
	return nil
}

// CastFloat64Ptr safely extracts a float64 from a map, returning nil if the key is absent or null.
func CastFloat64Ptr(m map[string]any, key string) *float64 {
	v, ok := m[key]
	if !ok || v == nil {
		return nil
	}
	switch n := v.(type) {
	case float64:
		return &n
	case int:
		f := float64(n)
		return &f
	case int64:
		f := float64(n)
		return &f
	}
	return nil
}

// CastBoolPtr safely extracts a bool from a map, returning nil if the key is absent or null.
func CastBoolPtr(m map[string]any, key string) *bool {
	v, ok := m[key]
	if !ok || v == nil {
		return nil
	}
	switch n := v.(type) {
	case bool:
		return &n
	}
	return nil
}

// IntPtrToAny converts a *int to any: nil pointer → nil (map will carry null), non-nil → int value.
func IntPtrToAny(v *int) any {
	if v == nil {
		return nil
	}
	return *v
}

// Float64PtrToAny converts a *float64 to any: nil pointer → nil, non-nil → float64 value.
func Float64PtrToAny(v *float64) any {
	if v == nil {
		return nil
	}
	return *v
}

// BoolPtrToAny converts a *bool to any: nil pointer → nil (map will carry null), non-nil → bool value.
func BoolPtrToAny(v *bool) any {
	if v == nil {
		return nil
	}
	return *v
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
