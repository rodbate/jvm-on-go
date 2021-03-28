package math

import (
	"jvm-on-go/instructions/base"
	"jvm-on-go/rtda"
)

func IMul(reader *base.ByteCodeReader, frame *rtda.Frame) {
	i1 := frame.OperandStack.PopInt()
	i2 := frame.OperandStack.PopInt()
	frame.OperandStack.PushInt(i1 * i2)
}

func FMul(reader *base.ByteCodeReader, frame *rtda.Frame) {
	i1 := frame.OperandStack.PopFloat()
	i2 := frame.OperandStack.PopFloat()
	frame.OperandStack.PushFloat(i1 * i2)
}

func LMul(reader *base.ByteCodeReader, frame *rtda.Frame) {
	i1 := frame.OperandStack.PopLong()
	i2 := frame.OperandStack.PopLong()
	frame.OperandStack.PushLong(i1 * i2)
}

func DMul(reader *base.ByteCodeReader, frame *rtda.Frame) {
	i1 := frame.OperandStack.PopDouble()
	i2 := frame.OperandStack.PopDouble()
	frame.OperandStack.PushDouble(i1 * i2)
}
