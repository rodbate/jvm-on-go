package reference

import (
	"github.com/rodbate/jvm-on-go/constants/descriptors"
	"github.com/rodbate/jvm-on-go/instructions/base"
	"github.com/rodbate/jvm-on-go/native/java/lang"
	"github.com/rodbate/jvm-on-go/rtda"
)

func AThrow(reader *base.ByteCodeReader, frame *rtda.Frame) {
	ex := frame.OperandStack.PopRef()
	if ex == nil {
		panic("java.lang.NullPointerException")
	}
	thread := frame.Thread()
	if !canHandleException(thread, ex) {
		handleUncaughtException(thread, ex)
	}
}

func canHandleException(thread *rtda.Thread, ex *rtda.Object) bool {
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

func handleUncaughtException(thread *rtda.Thread, ex *rtda.Object) {
	thread.ClearStack()
	jMsg := ex.GetFieldValue("detailMessage", descriptors.String).(*rtda.Object)
	goMsg := rtda.GetGoString(jMsg)
	println(ex.Class().JavaName() + ": " + goMsg)
	stackTraceElements := ex.StackTraceElements.([]*lang.StackTraceElement)
	for i := 0; i < len(stackTraceElements); i++ {
		println("\tat " + stackTraceElements[i].String())
	}
	thread.ExitCode = 1
}
