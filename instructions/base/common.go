package base

import (
	"jvm-on-go/rtda"
	"reflect"
	"runtime"
)

type Instruction func(reader *ByteCodeReader, frame *rtda.Frame)

func (i Instruction) String() string {
	return runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
}

func Branch(frame *rtda.Frame, offset int) {
	frame.SetNextPc(frame.Thread().Pc() + uint32(offset))
}

func InvokeMethod(invokerFrame *rtda.Frame, method *rtda.Method) {
	thread := invokerFrame.Thread()
	newFrame := thread.NewFrame(method)
	thread.PushFrame(newFrame)
	argSlotCount := method.ArgSlotCount()
	for i := int(argSlotCount - 1); i >= 0; i-- {
		newFrame.LocalVars.SetSlot(uint(i), invokerFrame.OperandStack.PopSlot())
	}
}

func RevertNextPc(reader *ByteCodeReader, frame *rtda.Frame) {
	frame.RevertNextPc()
	reader.SetPc(frame.NextPc())
}
