package mlib

import (
	"os"
	"os/exec"
	"path"
	"strings"
	"sync"
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
