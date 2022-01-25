package reference

import (
	"github.com/rodbate/jvm-on-go/pkg/constants/descriptors"
	"github.com/rodbate/jvm-on-go/pkg/instructions/base"
	"github.com/rodbate/jvm-on-go/pkg/native/java/lang"
	rtda2 "github.com/rodbate/jvm-on-go/pkg/rtda"
)

func AThrow(reader *base.ByteCodeReader, frame *rtda2.Frame) {
	ex := frame.OperandStack.PopRef()
	if ex == nil {
		panic("java.lang.NullPointerException")
	}
	thread := frame.Thread()
	if !canHandleException(thread, ex) {
		handleUncaughtException(thread, ex)
	}
}

func canHandleException(thread *rtda2.Thread, ex *rtda2.Object) bool {
	for {
		frame := thread.CurrentFrame()
		pc := frame.NextPc() - 1
		handlerPc := frame.Method().FindExceptionHandlerPc(ex.Class(), pc)
		if handlerPc > 0 {
			stack := frame.OperandStack
			stack.Clear()
			stack.PushRef(ex)
			frame.SetNextPc(handlerPc)
			return true
		}
		thread.PopFrame()
		if thread.IsStackEmpty() {
			break
		}
	}
	return false
}

func handleUncaughtException(thread *rtda2.Thread, ex *rtda2.Object) {
	thread.ClearStack()
	jMsg := ex.GetFieldValue("detailMessage", descriptors.String).(*rtda2.Object)
	goMsg := rtda2.GetGoString(jMsg)
	println(ex.Class().JavaName() + ": " + goMsg)
	stackTraceElements := ex.StackTraceElements.([]*lang.StackTraceElement)
	for i := 0; i < len(stackTraceElements); i++ {
		println("\tat " + stackTraceElements[i].String())
	}
	thread.ExitCode = 1
}
