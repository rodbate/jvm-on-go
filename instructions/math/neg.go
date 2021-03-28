package math

import (
	"jvm-on-go/instructions/base"
	"jvm-on-go/rtda"
)

func INeg(reader *base.ByteCodeReader, frame *rtda.Frame) {
	val := frame.OperandStack.PopInt()
	frame.OperandStack.PushInt(-val)
}

func FNeg(reader *base.ByteCodeReader, frame *rtda.Frame) {
	val := frame.OperandStack.PopFloat()
	frame.OperandStack.PushFloat(-val)
}

func LNeg(reader *base.ByteCodeReader, frame *rtda.Frame) {
	val := frame.OperandStack.PopLong()
	frame.OperandStack.PushLong(-val)
}

func DNeg(reader *base.ByteCodeReader, frame *rtda.Frame) {
	val := frame.OperandStack.PopDouble()
	frame.OperandStack.PushDouble(-val)
}
