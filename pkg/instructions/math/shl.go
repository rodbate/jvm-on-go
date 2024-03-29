package math

import (
	"github.com/rodbate/jvm-on-go/pkg/instructions/base"
	"github.com/rodbate/jvm-on-go/pkg/rtda"
)

func IShl(reader *base.ByteCodeReader, frame *rtda.Frame) {
	s := frame.OperandStack.PopInt() & 0x1F
	val := frame.OperandStack.PopInt()
	frame.OperandStack.PushInt(val << s)
}

func LShl(reader *base.ByteCodeReader, frame *rtda.Frame) {
	s := frame.OperandStack.PopInt() & 0x3F
	val := frame.OperandStack.PopLong()
	frame.OperandStack.PushLong(val << s)
}
