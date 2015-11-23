package here

import (
	"bytes"
	"fmt"
	"os"
	"runtime"
	"time"

	"github.com/juju/loggo"
	"github.com/kr/pretty"
)

var haveSetLogLevel bool
var logger = loggo.GetLogger("here")

// Here prints just enough stack trace to find where it was called from
func Here() {
	here()
}

type formatter struct{}

// Format returns the parameters separated by spaces except for filename and
// line which are separated by a colon.  The timestamp is shown to second
// resolution in UTC.
func (*formatter) Format(level loggo.Level, module, filename string, line int, timestamp time.Time, message string) string {
	return fmt.Sprintf("%s", message)
}

func write(line string) {
	if haveSetLogLevel == false {

		logger.SetLogLevel(loggo.INFO)
		haveSetLogLevel = true
	}
	logger.Infof(line)
}

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
	write(pretty.Sprintf("\t---> %# v\n", v))
}

// V prints the value of v
func V(name string, v interface{}) {
	write(pretty.Sprintf("%s: %# v\n", name, v))
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
