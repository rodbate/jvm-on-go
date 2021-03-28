package math

import (
	"jvm-on-go/instructions/base"
	"jvm-on-go/rtda"
)

func IUShr(reader *base.ByteCodeReader, frame *rtda.Frame) {
	s := frame.OperandStack.PopInt() & 0x1F
	val := frame.OperandStack.PopInt()
	frame.OperandStack.PushInt(int32(uint32(val) >> s))
}

func LUShr(reader *base.ByteCodeReader, frame *rtda.Frame) {
	s := frame.OperandStack.PopInt() & 0x3F
	val := frame.OperandStack.PopLong()
	frame.OperandStack.PushLong(int64(uint64(val) >> s))
}
