package stack

import (
	"github.com/rodbate/jvm-on-go/instructions/base"
	"github.com/rodbate/jvm-on-go/rtda"
)

func Pop(reader *base.ByteCodeReader, frame *rtda.Frame) {
	frame.OperandStack.PopSlot()
}

func Pop2(reader *base.ByteCodeReader, frame *rtda.Frame) {
	frame.OperandStack.PopSlot()
	frame.OperandStack.PopSlot()
}
