package conversion

import (
	"github.com/rodbate/jvm-on-go/instructions/base"
	"github.com/rodbate/jvm-on-go/rtda"
)

func F2L(reader *base.ByteCodeReader, frame *rtda.Frame) {
	val := frame.OperandStack.PopFloat()
	frame.OperandStack.PushLong(int64(val))
}

func F2I(reader *base.ByteCodeReader, frame *rtda.Frame) {
	val := frame.OperandStack.PopFloat()
	frame.OperandStack.PushInt(int32(val))
}

func F2D(reader *base.ByteCodeReader, frame *rtda.Frame) {
	val := frame.OperandStack.PopFloat()
	frame.OperandStack.PushDouble(float64(val))
}
