package lang

import (
	"fmt"
	"github.com/rodbate/jvm-on-go/constants/classname"
	"github.com/rodbate/jvm-on-go/native"
	"github.com/rodbate/jvm-on-go/rtda"
)

//java.lang.StackTraceElements
type StackTraceElement struct {
	DeclaringClass string
	MethodName     string
	FileName       string
	LineNumber     int
}

func init() {
	native.RegisterNative(classname.Throwable, "fillInStackTrace",
		"(I)Ljava/lang/Throwable;", fillInStackTrace)
}

//private native Throwable fillInStackTrace(int dummy)
func fillInStackTrace(frame *rtda.Frame) {
	this := frame.LocalVars.GetThis()
	frame.OperandStack.PushRef(this)
	stackTraceElements := buildStackTraceElements(this, frame.Thread())
	this.StackTraceElements = stackTraceElements
}

func buildStackTraceElements(throwable *rtda.Object, thread *rtda.Thread) []*StackTraceElement {
	var skipFrames = 2 //fillInStackTrace() and fillInStackTrace(int dummy)
	for c := throwable.Class().SuperClass(); c != nil; c = c.SuperClass() {
		skipFrames++
	}
	frames := thread.GetFrames()[skipFrames:]
	stackTraceElements := make([]*StackTraceElement, len(frames))
	for i, frame := range frames {
		stackTraceElements[i] = createStackTraceElement(frame)
	}
	return stackTraceElements
}

func createStackTraceElement(frame *rtda.Frame) *StackTraceElement {
	method := frame.Method()
	class := frame.Method().Class()
	return &StackTraceElement{
		DeclaringClass: class.JavaName(),
		MethodName:     method.Name(),
		FileName:       class.SourceFile,
		LineNumber:     method.GetLineNumber(frame.NextPc() - 1),
	}
}

func (ste *StackTraceElement) String() string {
	return fmt.Sprintf("%v.%v(%v:%v)", ste.DeclaringClass, ste.MethodName, ste.FileName,
		ste.LineNumber)
}
