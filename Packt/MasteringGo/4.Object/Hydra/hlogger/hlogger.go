package hlogger

import (
	"log"
	"sync"
	"os"
)

type hydraLogger struct {
	*log.Logger
	filename string
}

var hlogger *hydraLogger
var once sync.Once

func GetInstance() *hydraLogger {
	once.Do(func() {
		hlogger = createLogger("hydralogger.log")
	})
	return hlogger
}

func createLogger(fname string) *hydraLogger {
	file, _ := os.OpenFile(fname, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)

	return &hydraLogger{
		filename:fname,
		Logger: log.New(file, "Hydra", log.Lshortfile),
	}
}

