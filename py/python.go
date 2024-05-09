/*
 * Copyright (c) 2024 The GoPlus Authors (goplus.org). All rights reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package py

import (
	_ "unsafe"

	"github.com/goplus/llgo/c"
)

const (
	LLGoPackage = "link: $LLGO_PYTHON_ROOT/python3.12"
)

type (
	Char = c.Char
	Int  = c.Int
)

// -----------------------------------------------------------------------------

//go:linkname SetProgramName C.Py_SetProgramName
func SetProgramName(name *Char)

//go:linkname Initialize C.Py_Initialize
func Initialize()

// This function works like Initialize() if initsigs is 1.
// If initsigs is 0, it skips initialization registration of signal handlers,
// which might be useful when Python is embedded.
//
//go:linkname InitializeEx C.Py_InitializeEx
func InitializeEx(initsigs Int)

//go:linkname Finalize C.Py_Finalize
func Finalize()

// -----------------------------------------------------------------------------

// llgo:type C
type CompilerFlags struct {
	CfFlags Int
}

//go:linkname RunSimpleString C.PyRun_SimpleString
func RunSimpleString(command *Char) Int

//go:linkname RunSimpleStringFlags C.PyRun_SimpleStringFlags
func RunSimpleStringFlags(command *Char, flags *CompilerFlags) Int

//go:linkname RunSimpleFile C.PyRun_SimpleFile
func RunSimpleFile(fp c.FilePtr, filename *Char) Int

//go:linkname RunSimpleFileFlags C.PyRun_SimpleFileFlags
func RunSimpleFileFlags(fp c.FilePtr, filename *Char, flags *CompilerFlags) Int

// -----------------------------------------------------------------------------
