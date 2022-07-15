package asset

import (
	"io/ioutil"
	"log"
	"os"
	"strconv"
)

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
