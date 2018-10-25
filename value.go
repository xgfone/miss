// Copyright 2018 xgfone
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package miss

import (
	"runtime"
	"strconv"
	"strings"
)

// A Valuer generates a log value, which represents a dynamic value
// that is re-evaluated with each log event before firing it.
//
// For the caller depth, you maybe use the global variable DefaultDepth,
// which indicates the caller depth of the default implementation of the Logger.
type Valuer func() (interface{}, error)

// Caller returns a Valuer that returns a file and line from a specified depth
// in the callstack. Users will probably want to use DefaultCaller.
func Caller(depth int) Valuer {
	return func() (interface{}, error) {
		_, file, line, _ := runtime.Caller(depth + 1)
		idx := strings.LastIndexByte(file, '/')
		// using idx+1 below handles both of following cases:
		// idx == -1 because no "/" was found, or
		// idx >= 0 and we want to start at the character after the found "/".
		return file[idx+1:] + ":" + strconv.Itoa(line), nil
	}
}

// DefaultCaller is a Valuer that returns the file and line where the Log
// method was invoked. It can only be used with log.With.
var DefaultCaller = Caller(DefaultDepth)