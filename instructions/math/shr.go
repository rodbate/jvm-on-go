package math

import (
	"github.com/rodbate/jvm-on-go/instructions/base"
	"github.com/rodbate/jvm-on-go/rtda"
)

func IShr(reader *base.ByteCodeReader, frame *rtda.Frame) {
	s := frame.OperandStack.PopInt() & 0x1F
	val := frame.OperandStack.PopInt()
	frame.OperandStack.PushInt(val >> s)
}

func LShr(reader *base.ByteCodeReader, frame *rtda.Frame) {
	s := frame.OperandStack.PopInt() & 0x3F
	val := frame.OperandStack.PopLong()
	frame.OperandStack.PushLong(val >> s)
}
