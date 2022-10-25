package logging

import (
	"fmt"
	"io"
	"log"
	"os"
)

type prefix string

const (
	prefixInfo  prefix = "[info] "
	prefixError prefix = "[error] "

	flags = log.Ldate | log.Ltime | log.Lmicroseconds | log.Llongfile | log.Lmsgprefix
)

// diff logger and output log file
var (
	infoLogger  *log.Logger
	errorLogger *log.Logger
)

// init for loggers
func init() {
	infoLogger = log.New(os.Stdout, string(prefixInfo), flags)
	errorLogger = log.New(os.Stderr, string(prefixError), flags)
}

func Infof(format string, v ...interface{}) {
	infoLogger.Output(2, fmt.Sprintf(format, v...))
}

func Errorf(format string, v ...interface{}) {
	errorLogger.Output(2, fmt.Sprintf(format, v...))
}

// SetPath sets the output path for the loggers.
func SetPath(path string) {
	// todo:这里的0666可以用标准库自带的啥常量代替嘛?
	file, err := os.OpenFile(path, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatalf("set log path error:%+v", err)
	}
	w := io.MultiWriter(os.Stderr, file)
	infoLogger.SetOutput(w)
	errorLogger.SetOutput(w)
}
