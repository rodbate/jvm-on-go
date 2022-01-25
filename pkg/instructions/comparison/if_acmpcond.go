package comparison

import (
	base2 "github.com/rodbate/jvm-on-go/pkg/instructions/base"
	"github.com/rodbate/jvm-on-go/pkg/rtda"
)

func IFACmpEq(reader *base2.ByteCodeReader, frame *rtda.Frame) {
	offset := int(reader.ReadInt16())
	val2 := frame.OperandStack.PopRef()
	val1 := frame.OperandStack.PopRef()
	if val1 == val2 {
		base2.Branch(frame, offset)
	}
}

func IFACmpNe(reader *base2.ByteCodeReader, frame *rtda.Frame) {
	offset := int(reader.ReadInt16())
	val2 := frame.OperandStack.PopRef()
	val1 := frame.OperandStack.PopRef()
	if val1 != val2 {
		base2.Branch(frame, offset)
	}
}
