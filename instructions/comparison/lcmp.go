package comparison

import (
	"github.com/rodbate/jvm-on-go/instructions/base"
	"github.com/rodbate/jvm-on-go/rtda"
)

func LCmp(reader *base.ByteCodeReader, frame *rtda.Frame) {
	val2 := frame.OperandStack.PopLong()
	val1 := frame.OperandStack.PopLong()
	if val1 > val2 {
		frame.OperandStack.PushInt(1)
	} else if val1 < val2 {
		frame.OperandStack.PushInt(-1)
	} else {
		frame.OperandStack.PushInt(0)
	}
}
