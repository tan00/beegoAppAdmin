package util

import (
	"io/ioutil"
	"os"
	"os/exec"
	"path"
	"strconv"
	"strings"
	"syscall"
	"unsafe"

	pool "./bytebufferpool"
)

const (
	ENOENT           int = 2
	DEFAULT_FILEMODE     = 0666
	PID_FILE             = "beegoAppAdmin.pid"
)

// does file exist
func DoesFileExist(name string) bool {
	// file path can not contains null
	if HasNullByte(name) {
		return false
	}

	return SyscallStat(name)
}

func HasNullByte(name string) bool {
	for i := 0; i < len(name); i++ {
		if 0 == name[i] {
			return true
		}
	}

	return false
}

func SyscallStat(name string) bool {
	var buf = pool.Get()
	defer pool.Put(buf)

	buf.WriteString(name)
	buf.AppendNull()
	return DoSyscallStat(&((buf.Bytes())[0]))
}

func DoSyscallStat(p *byte) bool {
	var stat = &syscall.Stat_t{}
	_, _, errNo := syscall.Syscall(syscall.SYS_STAT, uintptr(unsafe.Pointer(p)), uintptr(unsafe.Pointer(stat)), 0)
	return !(ENOENT == int(errNo))
}

// delete file
func DelFile(name string) error {
	cmd := exec.Command("rm -f", name)
	return cmd.Run()
}

// write file
func WriteFile(name string, data []byte) error {
	err := ioutil.WriteFile(name, data, DEFAULT_FILEMODE)
	return err
}

// write pid to dragon.pid
func WritePidFile() error {
	if DoesFileExist(PID_FILE) {
		DelFile(PID_FILE)
	}

	file, err := os.Create(PID_FILE)
	if nil != err {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(strconv.Itoa(os.Getpid()))
	return err
}

// dele pid file
func DelPidFile() error {
	err := DelFile(PID_FILE)
	println("delete pid file error : %v\n", err.Error())
	return err
}

// return file base name, remove suffix
func FileBaseName(file string) string {
	base := path.Base(file)
	results := strings.Split(base, ".")
	return results[0]
}
