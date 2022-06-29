package asset

import (
	"hash/fnv"
	"io/ioutil"
	"log"
	"os"
)

func Hash(str string) int {
	h := fnv.New32a()
	h.Write([]byte(str))
	return int(h.Sum32())
}

func GenericSliceToAny[T any](slice []T) []any {
	anySlice := make([]any, len(slice))
	for i, value := range slice {
		anySlice[i] = value
	}
	return anySlice
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
