package mlib

import (
	"errors"
	"fmt"
	"os"
	"path"
	"runtime"
	"strings"
	"sync"
	"time"
)

var (
	LogFile = ""
	LogUuid = true

	logUuidCache = make(map[int64]string)
	logUuidMutex = sync.Mutex{}
)

// InitLogPrefix used to set log file prefix
// empty prefix will print to stdout
func InitLogPrefix(s string) {
	LogFile = s
}

func LogUuidCacheSize() int {
	logUuidMutex.Lock()
	num := len(logUuidCache)
	logUuidMutex.Unlock()
	return num
}

func LogSetUuid(uuid string) {
	logUuidMutex.Lock()
	logUuidCache[GoId()] = uuid
	logUuidMutex.Unlock()
}

func LogGetUuid() string {
	logUuidMutex.Lock()
	uuid := logUuidCache[GoId()]
	logUuidMutex.Unlock()
	if uuid == "" {
		uuid = Uuid()
		LogSetUuid(uuid)
	}
	return uuid
}

func LogDelUuid() {
	logUuidMutex.Lock()
	defer logUuidMutex.Unlock()
	delete(logUuidCache, GoId())
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
func logf(f interface{}, args ...interface{}) string {
	fs, ok := f.(string)
	if !ok {
		fs = fmt.Sprintf("%v", f) + strings.Repeat(" [%v]", len(args))
	}

	pc, file, line, _ := runtime.Caller(2)
	fun := runtime.FuncForPC(pc)
	key := fmt.Sprintf("[%s] %s:%s():%d: ",
		time.Now().Format(STD_TIME_FORMAT),
		path.Base(file), path.Base(fun.Name()), line)
	if LogUuid {
		key += fmt.Sprintf("%s:%d ", LogGetUuid(), LogUuidCacheSize())
	}
	s := strings.TrimSpace(fmt.Sprintf(fs, args...))
	logwrite(key + s)
	return s
}

// Debugf log with debug level
func Debugf(f interface{}, args ...interface{}) {
	logf(f, args...)
}

// Infof log with info level
func Infof(f interface{}, args ...interface{}) {
	logf(f, args...)
}

// Errorf log with error level
func Errorf(f interface{}, args ...interface{}) error {
	return errors.New(logf(f, args...))
}

// Fatalf log with fatal level
func Fatalf(f interface{}, args ...interface{}) {
	panic(logf(f, args...))
}
