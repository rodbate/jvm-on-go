package math

import (
	"github.com/rodbate/jvm-on-go/instructions/base"
	"github.com/rodbate/jvm-on-go/rtda"
)

func IAdd(reader *base.ByteCodeReader, frame *rtda.Frame) {
	val1 := frame.OperandStack.PopInt()
	val2 := frame.OperandStack.PopInt()
	frame.OperandStack.PushInt(val1 + val2)
}

func FAdd(reader *base.ByteCodeReader, frame *rtda.Frame) {
	val1 := frame.OperandStack.PopFloat()
	val2 := frame.OperandStack.PopFloat()
	frame.OperandStack.PushFloat(val1 + val2)
}

func LAdd(reader *base.ByteCodeReader, frame *rtda.Frame) {
	val1 := frame.OperandStack.PopLong()
	val2 := frame.OperandStack.PopLong()
	frame.OperandStack.PushLong(val1 + val2)
}

func DAdd(reader *base.ByteCodeReader, frame *rtda.Frame) {
	val1 := frame.OperandStack.PopDouble()
	val2 := frame.OperandStack.PopDouble()
	frame.OperandStack.PushDouble(val1 + val2)
}
