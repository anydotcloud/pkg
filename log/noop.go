// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

package log

var (
	// NoopLogger is useful for testing/mocking
	NoopLogger Logger = &voidLogger{}
)

// voidLogger implements Logger but does nothing. Useful for
// testing and mocking...
type voidLogger struct{}

func (l *voidLogger) IsDebugEnabled() bool         { return false }
func (l *voidLogger) WithValues(...interface{})    {}
func (l *voidLogger) Info(string, ...interface{})  {}
func (l *voidLogger) Debug(string, ...interface{}) {}
func (l *voidLogger) Trace(name string, additionalValues ...interface{}) TraceExiter {
	f := func(err error, args ...interface{}) {
		l.Exit(name, err, args...)
	}
	return f
}
func (l *voidLogger) Enter(string, ...interface{})       {}
func (l *voidLogger) Exit(string, error, ...interface{}) {}
