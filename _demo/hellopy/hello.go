package main

import (
	"github.com/goplus/cpython/py"
	"github.com/goplus/llgo/c"
)

func main() {
	py.Initialize()
	py.SetProgramName(*c.Argv)
	py.RunSimpleString(c.Str(`print('Hello, World!')`))
	py.Finalize()
}
