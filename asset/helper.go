package asset

import (
	"io/ioutil"
	"log"
	"os"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Returns a function that casts its input to a map, and makes a hash of the value
// for the given key.
// This requires three assumptions that must hold:
//
// - the input is a map[string]any
//
// - the map contains the given key
//
// - the value mapped to the given key is a string
func HashFromMapValue(key string) func(any) int {
	return func(m any) int {
		return schema.HashString(m.(map[string]any)[key].(string))
	}
}

func GenericSliceToAny[T any](slice []T) []any {
	anySlice := make([]any, len(slice))
	for i, value := range slice {
		anySlice[i] = value
	}
	return anySlice
}

// Maps a list of type S to type T
func Map[S any, T any](source []S, converter func(S) T) []T {
	target := make([]T, len(source))
	for index, value := range source {
		target[index] = converter(value)
	}
	return target
}

// Maps a list of type S to type T, casting with intermediate type I
func CastMap[S any, I any, T any](source []S, converter func(I) T) []T {
	target := make([]T, len(source))
	for index, value := range source {
		target[index] = converter(any(value).(I))
	}
	return target
}

type ParseablePointer[T any] interface {
	*T
	Parseable
}

func ParseToMaps[T any, PT ParseablePointer[T]](parseables []T) []map[string]any {
	maps := make([]map[string]any, len(parseables))
	for index, value := range parseables {
		maps[index] = any(&value).(PT).ToMap()
	}
	return maps
}

func ParseFromMaps[T any, PT ParseablePointer[T]](maps []any) []T {
	parseables := make([]T, len(maps))
	for index, m := range maps {
		any(&parseables[index]).(PT).FromMap(m.(map[string]any))
	}
	return parseables
}

// Parses a float to a string. Use this method to ensure consistency.
func ParseFloat(num float64) string {
	return strconv.FormatFloat(num, 'g', -1, 64)
}

const Debug = false

var Logger = func() *log.Logger {
	if Debug {
		var logFile, _ = os.OpenFile("log_plug.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
		var logger = log.New(logFile, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
		return logger
	}
	log.SetOutput(ioutil.Discard)
	return log.Default()
}()
