package here

import (
	"bytes"
	"fmt"
	"os"
	"runtime"
	"strings"
	"time"

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

type formatter struct{}

// Format returns the parameters separated by spaces except for filename and
// line which are separated by a colon.  The timestamp is shown to second
// resolution in UTC.
func (*formatter) Format(level loggo.Level, module, filename string, line int, timestamp time.Time, message string) string {
	return message
}

func write(message string) {
	if haveSetLogLevel == false {
		logger.SetLogLevel(loggo.INFO)
		haveSetLogLevel = true
	}
	prefix := "| "
	f := fmt.Sprintf("### %%%ds%%s", indentLevel+len(prefix))
	m := fmt.Sprintf(f, prefix, message)
	logger.Infof(m)
}

// OverwriteWriter replaces the default writer with the above here.formatter
func OverwriteWriter() {
	loggo.RemoveWriter("default")
	w := loggo.NewSimpleWriter(os.Stdout, &formatter{})
	loggo.RegisterWriter("default", w, loggo.INFO)
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

// Is prints the value of v with just enough stack trace to find where it was called from
func Is(v interface{}) {
	here()
	lines := strings.Split(pretty.Sprintf("---> %# v\n", v), "\n")
	for _, line := range lines {
		write(line)
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
func M(s string) {
	write(s)
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
