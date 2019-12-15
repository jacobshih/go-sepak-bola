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

package dbg

import (
	"fmt"
	"log"
	"runtime"
)

func funcName() string {
	pc := make([]uintptr, 10) // at least 1 entry needed
	runtime.Callers(3, pc)
	f := runtime.FuncForPC(pc[0])
	return f.Name()
}

// Trace prefixes function name to debug message.
func Trace(format string, a ...interface{}) {
	info := fmt.Sprintf(format, a...)
	log.Printf("%s() - %v", funcName(), info)
}

// Debug prefixes function name to debug message.
func Debug(format string, a ...interface{}) {
	info := fmt.Sprintf(format, a...)
	fmt.Printf("%s() - %v", funcName(), info)
}
