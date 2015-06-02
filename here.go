package here

import (
	"bytes"
	"runtime"

	"github.com/juju/loggo"
)

func Here() {
	here()
}

func here() {
	trace := make([]byte, 1024)
	runtime.Stack(trace, false)

	logger := loggo.GetLogger("here")
	b := bytes.NewBuffer(trace)
	for i := 0; i < 7; i++ {
		line, err := b.ReadString('\n')
		if err != nil {
			break
		}
		if i > 4 {
			logger.Infof(line)
		}
	}
}

func Is(v interface{}) {
	here()
	logger := loggo.GetLogger("here")
	logger.Infof("\t\t%#v\n", v)
}
