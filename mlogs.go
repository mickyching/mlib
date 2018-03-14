package mlib

import (
	"errors"
	"fmt"
	"os"
	"path"
	"runtime"
	"strings"
	"time"
)

var (
	LogFile = ""
	LogUuid = true
)

// InitLogPrefix used to set log file prefix
// empty prefix will print to stdout
func InitLogPrefix(s string) {
	LogFile = s
}

func logwrite(s string) {
	if LogFile == "" {
		fmt.Println(s)
		return
	}

	day := time.Now().Format("20060102")
	fname := LogFile + "-" + day
	err := CreateFile(fname)
	if err != nil {
		panic(err)
	}
	f, err := os.OpenFile(fname, os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	f.WriteString(s + "\n")
}

// support format
// 1. logf(s)
// 2. logf(fmt, v...)
// 3. logf(any, v...)
func logf(level string, f interface{}, args ...interface{}) string {
	fs, ok := f.(string)
	if !ok {
		fs = fmt.Sprintf("%v", f) + strings.Repeat(" [%v]", len(args))
	}

	pc, file, line, _ := runtime.Caller(2)
	fun := runtime.FuncForPC(pc).Name()
	fn := len(fun)
	if fn > 20 {
		fun = "..." + fun[fn-16:fn-1]
	}
	key := fmt.Sprintf("[%s] %s %s:%s():%d: ",
		time.Now().Format(STD_TIME_FORMAT), level,
		path.Base(file), fun, line)
	if LogUuid {
		if uuid := GetUuid(); uuid != "" {
			key += fmt.Sprintf("%s:%d ", uuid, UuidCacheSize())
		}
	}
	s := strings.TrimSpace(fmt.Sprintf(fs, args...))
	logwrite(key + s)
	return s
}

// Debugf log with debug level
func Debugf(f interface{}, args ...interface{}) {
	logf("DEBUG", f, args...)
}

// Infof log with info level
func Infof(f interface{}, args ...interface{}) {
	logf("INFOF", f, args...)
}

// Errorf log with error level
func Errorf(f interface{}, args ...interface{}) error {
	return errors.New(logf("ERROR", f, args...))
}

// Fatalf log with fatal level
func Fatalf(f interface{}, args ...interface{}) {
	panic(logf("FATAL", f, args...))
}
