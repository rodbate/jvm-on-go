package math

import (
	"github.com/rodbate/jvm-on-go/pkg/instructions/base"
	"github.com/rodbate/jvm-on-go/pkg/rtda"
)

func ISub(reader *base.ByteCodeReader, frame *rtda.Frame) {
	i1 := frame.OperandStack.PopInt()
	i2 := frame.OperandStack.PopInt()
	frame.OperandStack.PushInt(i2 - i1)
}

func FSub(reader *base.ByteCodeReader, frame *rtda.Frame) {
	i1 := frame.OperandStack.PopFloat()
	i2 := frame.OperandStack.PopFloat()
	frame.OperandStack.PushFloat(i2 - i1)
}

func LSub(reader *base.ByteCodeReader, frame *rtda.Frame) {
	i1 := frame.OperandStack.PopLong()
	i2 := frame.OperandStack.PopLong()
	frame.OperandStack.PushLong(i2 - i1)
}

func DSub(reader *base.ByteCodeReader, frame *rtda.Frame) {
	i1 := frame.OperandStack.PopDouble()
	i2 := frame.OperandStack.PopDouble()
	frame.OperandStack.PushDouble(i2 - i1)
}
