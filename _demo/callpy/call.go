package main

import (
	"github.com/goplus/cpython/py"
	"github.com/goplus/llgo/c"
)

func main() {
	py.Initialize()
	py.SetProgramName(*c.Argv)
	math := py.ImportModule(c.Str("math"))
	sqrt := math.GetAttrString(c.Str("sqrt"))
	_2 := py.FloatFromDouble(2)
	_sqrt2 := sqrt.CallOneArg(_2).FloatAsDouble()
	c.Printf(c.Str("sqrt(2) = %f\n"), _sqrt2)
	py.Finalize()
}
