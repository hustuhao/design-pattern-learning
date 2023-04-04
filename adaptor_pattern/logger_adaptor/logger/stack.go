// Package stack provides utilities to capture and pass around stack traces.
// 源自于facebookgo/stack，但由于[go1.12的改动](https://golang.org/doc/go1.12#compiler)，获取的堆栈信息不准确
package logger

import (
	"bytes"
	"fmt"
	"runtime"
	"strings"
)

const maxStackSize = 32

// Frame identifies a file, line & function name in the stack.
type Frame struct {
	File string
	Line int
	Name string
}

// String provides the standard file:line representation.
func (f Frame) String() string {
	return fmt.Sprintf("%s:%d %s", f.File, f.Line, f.Name)
}

// Stack represents an ordered set of Frames.
type Stack []Frame

// String provides the standard multi-line stack trace.
func (s Stack) String() string {
	var b bytes.Buffer
	writeStack(&b, s)
	return b.String()
}

// Multi represents a number of Stacks. This is useful to allow tracking a
// value as it travels thru code.
type Multi struct {
	stacks []Stack
}

// Stacks returns the tracked Stacks.
func (m *Multi) Stacks() []Stack {
	return m.stacks
}

// Add the given Stack to this Multi.
func (m *Multi) Add(s Stack) {
	m.stacks = append(m.stacks, s)
}

// AddCallers adds the Callers Stack to this Multi. The argument skip is
// the number of stack frames to ascend, with 0 identifying the caller of
// Callers.
func (m *Multi) AddCallers(skip int) {
	m.Add(Callers(skip+1, true))
}

// String provides a human readable multi-line stack trace.
func (m *Multi) String() string {
	var b bytes.Buffer
	for i, s := range m.stacks {
		if i != 0 {
			fmt.Fprintf(&b, "\n(Stack %d)\n", i+1)
		}
		writeStack(&b, s)
	}
	return b.String()
}

// Copy makes a copy of the stack which is safe to modify.
func (m *Multi) Copy() *Multi {
	m2 := &Multi{
		stacks: make([]Stack, len(m.stacks)),
	}
	copy(m2.stacks, m.stacks)
	return m2
}

// Caller returns a single Frame for the caller. The argument skip is the
// number of stack frames to ascend, with 0 identifying the caller of Callers.
func Caller(skip int) Frame {
	pc, file, line, _ := runtime.Caller(skip + 1)
	fun := runtime.FuncForPC(pc)
	return Frame{
		File: StripGOPATH(file),
		Line: line,
		Name: StripPackage(fun.Name()),
	}
}

// Callers returns a Stack of Frames for the callers. The argument skip is the
// number of stack frames to ascend, with 0 identifying the caller of Callers.
func Callers(skip int, deep bool) Stack {
	pcs := make([]uintptr, maxStackSize)
	num := runtime.Callers(skip+2, pcs)
	stack := make(Stack, 0, num)
	frames := runtime.CallersFrames(pcs[:num])
	for {
		frame, more := frames.Next()
		stack = append(stack, Frame{
			File: StripGOPATH(frame.File),
			Line: frame.Line,
			Name: StripPackage(frame.Function),
		})
		if !more || !deep {
			break
		}
	}
	return stack
}

// CallersMulti returns a Multi which includes one Stack for the
// current callers. The argument skip is the number of stack frames to ascend,
// with 0 identifying the caller of CallersMulti.
func CallersMulti(skip int) *Multi {
	m := new(Multi)
	m.AddCallers(skip + 1)
	return m
}

func writeStack(b *bytes.Buffer, s Stack) {
	var width int
	for _, f := range s {
		if l := len(f.File) + numDigits(f.Line) + 1; l > width {
			width = l
		}
	}
	last := len(s) - 1
	for i, f := range s {
		b.WriteString(f.File)
		b.WriteRune(rune(':'))
		n, _ := fmt.Fprintf(b, "%d", f.Line)
		for i := width - len(f.File) - n; i != 0; i-- {
			b.WriteRune(rune(' '))
		}
		b.WriteString(f.Name)
		if i != last {
			b.WriteRune(rune('\n'))
		}
	}
}

func numDigits(i int) int {
	var n int
	for {
		n++
		i = i / 10
		if i == 0 {
			return n
		}
	}
}

var (
	// This can be set by a build script. It will be the colon separated equivalent
	// of the environment variable.
	goPath  string
	pathLen int
)

func init() {
	_, file, _, _ := runtime.Caller(0)
	s := strings.Split(file, "src/")
	goPath = s[0] + "src/"
	pathLen = len(goPath)
}

// StripGOPATH strips the GOPATH prefix from the file path f.
// In development, this will be done using the GOPATH environment variable.
// For production builds, where the GOPATH environment will not be set, the
// GOPATH can be included in the binary by passing ldflags, for example:
//
//	GO_LDFLAGS="$GO_LDFLAGS -X github.com/facebookgo/stack.goPath $GOPATH"
//	go install "-ldflags=$GO_LDFLAGS" my/pkg
func StripGOPATH(f string) string {
	if strings.HasPrefix(f, goPath) {
		return f[pathLen:]
	}
	return f
}

// StripPackage strips the package name from the given Func.Name.
func StripPackage(n string) string {
	slashI := strings.LastIndex(n, "/")
	if slashI == -1 {
		slashI = 0 // for built-in packages
	}
	dotI := strings.Index(n[slashI:], ".")
	if dotI == -1 {
		return n
	}
	return n[slashI+dotI+1:]
}
