package comparison

import (
	base2 "github.com/rodbate/jvm-on-go/pkg/instructions/base"
	"github.com/rodbate/jvm-on-go/pkg/rtda"
)

func IFeq(reader *base2.ByteCodeReader, frame *rtda.Frame) {
	offset := int(reader.ReadInt16())
	val := frame.OperandStack.PopInt()
	if val == 0 {
		base2.Branch(frame, offset)
	}
}

func IFne(reader *base2.ByteCodeReader, frame *rtda.Frame) {
	offset := int(reader.ReadInt16())
	val := frame.OperandStack.PopInt()
	if val != 0 {
		base2.Branch(frame, offset)
	}
}

func IFlt(reader *base2.ByteCodeReader, frame *rtda.Frame) {
	offset := int(reader.ReadInt16())
	val := frame.OperandStack.PopInt()
	if val < 0 {
		base2.Branch(frame, offset)
	}
}

func IFle(reader *base2.ByteCodeReader, frame *rtda.Frame) {
	offset := int(reader.ReadInt16())
	val := frame.OperandStack.PopInt()
	if val <= 0 {
		base2.Branch(frame, offset)
	}
}

func IFgt(reader *base2.ByteCodeReader, frame *rtda.Frame) {
	offset := int(reader.ReadInt16())
	val := frame.OperandStack.PopInt()
	if val > 0 {
		base2.Branch(frame, offset)
	}
}

func IFge(reader *base2.ByteCodeReader, frame *rtda.Frame) {
	offset := int(reader.ReadInt16())
	val := frame.OperandStack.PopInt()
	if val >= 0 {
		base2.Branch(frame, offset)
	}
}
