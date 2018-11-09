package xrecovery

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"runtime"
)

var (
	unknow    = []byte("???")
	centerDot = []byte("·")
	dot       = []byte(".")
	slash     = []byte("/")
)

// Stack 调用栈
func Stack(skip int) []byte {
	buf := new(bytes.Buffer)
	var lines [][]byte
	var lastFile string
	for i := skip; ; i++ {
		pc, file, line, ok := runtime.Caller(i)
		if !ok {
			break
		}
		fmt.Fprintf(buf, "%s:%d (0x%x)\n", file, line, pc)
		if file != lastFile {
			data, err := ioutil.ReadFile(file)
			if err != nil {
				continue
			}
			lines = bytes.Split(data, []byte{'\n'})
			lastFile = file
		}
		fmt.Fprintf(buf, "\t%s: %s\n", function(pc), source(lines, line))
	}
	return buf.Bytes()
}

// 取调用函数
func source(lines [][]byte, n int) []byte {
	n--
	if n < 0 || n >= len(lines) {
		return unknow
	}
	return bytes.TrimSpace(lines[n])
}

// 取当前函数名
func function(pc uintptr) []byte {
	fn := runtime.FuncForPC(pc)
	if fn == nil {
		return unknow
	}

	name := []byte(fn.Name())
	if lastslash := bytes.LastIndex(name, slash); lastslash >= 0 {
		name = name[lastslash+1:]
	}

	if period := bytes.Index(name, dot); period >= 0 {
		name = name[period+1:]
	}

	name = bytes.Replace(name, centerDot, dot, -1)
	return name
}
