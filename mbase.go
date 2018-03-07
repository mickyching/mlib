package mlib

import (
	"bytes"
	"crypto/rand"
	"fmt"
	"os"
	"os/exec"
	"path"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"time"
)

// RunCmd exec cmd with args
func RunCmd(cmd string, args ...string) (string, error) {
	out, err := exec.Command(cmd, args...).Output()
	return strings.TrimSpace(string(out)), err
}

// PathExist check if path exist
func PathExist(p string) bool {
	_, err := os.Stat(p)
	return !os.IsNotExist(err)
}

// CreateDir create a directory
func CreateDir(dname string) error {
	return os.MkdirAll(dname, os.ModePerm)
}

// CreateFile create a file
func CreateFile(fname string) error {
	if PathExist(fname) {
		return nil
	}
	err := CreateDir(path.Dir(fname))
	if err != nil {
		return err
	}
	f, err := os.Create(fname)
	if err != nil {
		return err
	}
	f.Close()
	return nil
}

// GoFunc run func with num concurrency routines
func GoFunc(f func(), num int) *sync.WaitGroup {
	wg := sync.WaitGroup{}
	if num < 1 {
		return &wg
	}

	wg.Add(num)
	for i := 0; i < num; i++ {
		go func() {
			defer wg.Done()
			f()
		}()
	}
	return &wg
}

// GoId returns current Goroutine ID
func GoId() int64 {
	gid := func(s []byte) int64 {
		s = s[len("goroutine "):]
		s = s[:bytes.IndexByte(s, ' ')]
		gid, _ := strconv.ParseInt(string(s), 10, 64)
		return gid
	}
	var buf [64]byte
	return gid(buf[:runtime.Stack(buf[:], false)])
}

// Uuid returns uuid base on current time
func Uuid() string {
	unix32bits := uint32(time.Now().UTC().Unix())
	buff := make([]byte, 12)
	numRead, err := rand.Read(buff)
	if numRead != len(buff) || err != nil {
		Fatalf(err)
	}

	return fmt.Sprintf("%x-%x-%x-%x-%x-%x", unix32bits, buff[0:2], buff[2:4], buff[4:6], buff[6:8], buff[8:])
}