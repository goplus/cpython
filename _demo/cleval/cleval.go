package main

import (
	"github.com/goplus/cpython/py"
	"github.com/goplus/llgo/c"
)

func main() {
	py.Initialize()
	py.SetProgramName(*c.Argv)
	code := py.CompileString(c.Str(`print('Hello, World!')`), c.Str(`hello.py`), py.EvalInput)
	if code != nil {
		mod := py.ImportModule(c.Str("__main__"))
		gbl := mod.GetDict()
		py.EvalCode(code, gbl, nil)
	}
	py.Finalize()
}
