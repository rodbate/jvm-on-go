package stack

import (
	"jvm-on-go/instructions/base"
	"jvm-on-go/rtda"
)

func Dup(reader *base.ByteCodeReader, frame *rtda.Frame) {
	frame.OperandStack.PushSlot(frame.OperandStack.PeekSlot())
}

func DupX1(reader *base.ByteCodeReader, frame *rtda.Frame) {
	slot1 := frame.OperandStack.PopSlot()
	slot2 := frame.OperandStack.PopSlot()
	frame.OperandStack.PushSlot(slot1)
	frame.OperandStack.PushSlot(slot2)
	frame.OperandStack.PushSlot(slot1)
}

func Dup2(reader *base.ByteCodeReader, frame *rtda.Frame) {
	s1 := frame.OperandStack.PopSlot()
	s2 := frame.OperandStack.PopSlot()
	frame.OperandStack.PushSlot(s2)
	frame.OperandStack.PushSlot(s1)
	frame.OperandStack.PushSlot(s2)
	frame.OperandStack.PushSlot(s1)
}

func Dup2X1(reader *base.ByteCodeReader, frame *rtda.Frame) {
	s1 := frame.OperandStack.PopSlot()
	s2 := frame.OperandStack.PopSlot()
	s3 := frame.OperandStack.PopSlot()
	frame.OperandStack.PushSlot(s2)
	frame.OperandStack.PushSlot(s1)
	frame.OperandStack.PushSlot(s3)
	frame.OperandStack.PushSlot(s2)
	frame.OperandStack.PushSlot(s1)
}
