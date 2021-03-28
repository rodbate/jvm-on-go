package comparison

import (
	"jvm-on-go/instructions/base"
	"jvm-on-go/rtda"
)

func IFACmpEq(reader *base.ByteCodeReader, frame *rtda.Frame) {
	offset := int(reader.ReadInt16())
	val2 := frame.OperandStack.PopRef()
	val1 := frame.OperandStack.PopRef()
	if val1 == val2 {
		base.Branch(frame, offset)
	}
}

func IFACmpNe(reader *base.ByteCodeReader, frame *rtda.Frame) {
	offset := int(reader.ReadInt16())
	val2 := frame.OperandStack.PopRef()
	val1 := frame.OperandStack.PopRef()
	if val1 != val2 {
		base.Branch(frame, offset)
	}
}
