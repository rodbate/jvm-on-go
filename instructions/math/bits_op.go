package math

import (
	"jvm-on-go/instructions/base"
	"jvm-on-go/rtda"
)

func IAnd(reader *base.ByteCodeReader, frame *rtda.Frame) {
	v1 := frame.OperandStack.PopInt()
	v2 := frame.OperandStack.PopInt()
	frame.OperandStack.PushInt(v1 & v2)
}

func LAnd(reader *base.ByteCodeReader, frame *rtda.Frame) {
	v1 := frame.OperandStack.PopLong()
	v2 := frame.OperandStack.PopLong()
	frame.OperandStack.PushLong(v1 & v2)
}

func IOr(reader *base.ByteCodeReader, frame *rtda.Frame) {
	v1 := frame.OperandStack.PopInt()
	v2 := frame.OperandStack.PopInt()
	frame.OperandStack.PushInt(v1 | v2)
}

func LOr(reader *base.ByteCodeReader, frame *rtda.Frame) {
	v1 := frame.OperandStack.PopLong()
	v2 := frame.OperandStack.PopLong()
	frame.OperandStack.PushLong(v1 | v2)
}

func IXor(reader *base.ByteCodeReader, frame *rtda.Frame) {
	v1 := frame.OperandStack.PopInt()
	v2 := frame.OperandStack.PopInt()
	frame.OperandStack.PushInt(v1 ^ v2)
}

func LXor(reader *base.ByteCodeReader, frame *rtda.Frame) {
	v1 := frame.OperandStack.PopLong()
	v2 := frame.OperandStack.PopLong()
	frame.OperandStack.PushLong(v1 ^ v2)
}
