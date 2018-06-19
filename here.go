package here

import (
	"bytes"
	"fmt"
	"os"
	"runtime"
	"strings"

	"github.com/juju/loggo"
	"github.com/kr/pretty"
)

var haveSetLogLevel bool
var logger = loggo.GetLogger("here")
var indentLevel = 0

// Here prints just enough stack trace to find where it was called from
func Here() {
	here()
}

// Uncluttered formatter
func Formatter(entry loggo.Entry) string {
	return entry.Message
}

// prefixSize is used internally to trim the user specific path from the
// front of the returned filenames from the runtime call stack.
var prefixSize int

// goPath is the deduced path based on the location of this file as compiled.
var goPath string

func init() {
	_, file, _, ok := runtime.Caller(0)
	if ok {
		// We know that the end of the file should be:
		// github.com/juju/errors/path.go
		size := len(file)
		suffix := len("github.com/juju/errors/path.go")
		goPath = file[:size-suffix]
		prefixSize = len(goPath)
	}
}

func trimGoPath(filename string) string {
	if strings.HasPrefix(filename, goPath) {
		return filename[prefixSize:]
	}
	return filename
}

func write(message string) {
	//if haveSetLogLevel == false {
	logger.SetLogLevel(loggo.INFO)
	//	haveSetLogLevel = true
	//}
	prefix := "| "
	f := fmt.Sprintf("### %%%ds%%s", indentLevel+len(prefix))
	m := fmt.Sprintf(f, prefix, message)
	logger.Infof(m)
}

// OverwriteWriter replaces the default writer with the above here.formatter
func OverwriteWriter() {
	loggo.RemoveWriter("default")
	w := loggo.NewSimpleWriter(os.Stdout, Formatter)
	loggo.RegisterWriter("default", w)
}

func here() {
	trace := make([]byte, 1024)
	runtime.Stack(trace, false)

	b := bytes.NewBuffer(trace)
	for i := 0; i < 7; i++ {
		line, err := b.ReadString('\n')
		if err != nil {
			break
		}
		if i > 4 {
			write(line)
		}
	}
}

func Loc() string {
	_, file, line, _ := runtime.Caller(1)
	return fmt.Sprintf("%s:%d", trimGoPath(file), line)
}

// Is prints the value of v with just enough stack trace to find where it was called from
func Is(vars ...interface{}) {
	here()
	for _, v := range(vars) {
		lines := strings.Split(pretty.Sprintf("---> %# v\n", v), "\n")
		for _, line := range lines {
			write(line)
		}
	}
}

// V prints the value of v, prefixed by a name
func V(name string, v interface{}) {
	lines := strings.Split(pretty.Sprintf("%s: %# v\n", name, v), "\n")
	for _, line := range lines {
		write(line)
	}
}

// M prints the string s
func M(s ...string) {
	for _, m := range(s) {
		write(m)
	}
}

// HR prints s in the middle of a horizontal line
func HR(s string) {
	write("---------------- " + s + " ----------------")
}

// Stack prints a full stack trace
func Stack() {
	trace := make([]byte, 10240)
	runtime.Stack(trace, false)

	b := bytes.NewBuffer(trace)
	write(">>>>>>>>>>>>")
	for i := 0; true; i++ {
		line, err := b.ReadString('\n')
		if err != nil {
			break
		}
		if i > 2 {
			write(line)
		}
	}
	write("<<<<<<<<<<<<<")
}

func Indent() {
	indentLevel += 2
}

func Dedent() {
	indentLevel -= 2
	if indentLevel < 0 {
		indentLevel = 0
	}
}
