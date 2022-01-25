package conversion

import (
	"github.com/rodbate/jvm-on-go/pkg/instructions/base"
	"github.com/rodbate/jvm-on-go/pkg/rtda"
)

func D2I(reader *base.ByteCodeReader, frame *rtda.Frame) {
	val := frame.OperandStack.PopDouble()
	frame.OperandStack.PushInt(int32(val))
}

func D2F(reader *base.ByteCodeReader, frame *rtda.Frame) {
	val := frame.OperandStack.PopDouble()
	frame.OperandStack.PushFloat(float32(val))
}

func D2L(reader *base.ByteCodeReader, frame *rtda.Frame) {
	val := frame.OperandStack.PopDouble()
	frame.OperandStack.PushLong(int64(val))
}
