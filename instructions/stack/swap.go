package stack

import (
	"jvm-on-go/instructions/base"
	"jvm-on-go/rtda"
)

func Swap(reader *base.ByteCodeReader, frame *rtda.Frame) {
	slot1 := frame.OperandStack.PopSlot()
	slot2 := frame.OperandStack.PopSlot()
	frame.OperandStack.PushSlot(slot1)
	frame.OperandStack.PushSlot(slot2)
}
