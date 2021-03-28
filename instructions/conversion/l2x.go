package conversion

import (
	"github.com/rodbate/jvm-on-go/instructions/base"
	"github.com/rodbate/jvm-on-go/rtda"
)

func L2I(reader *base.ByteCodeReader, frame *rtda.Frame) {
	val := frame.OperandStack.PopLong()
	frame.OperandStack.PushInt(int32(val & 0xFFFFFFFF))
}

func L2F(reader *base.ByteCodeReader, frame *rtda.Frame) {
	val := frame.OperandStack.PopLong()
	frame.OperandStack.PushFloat(float32(val))
}

func L2D(reader *base.ByteCodeReader, frame *rtda.Frame) {
	val := frame.OperandStack.PopLong()
	frame.OperandStack.PushDouble(float64(val))
}
