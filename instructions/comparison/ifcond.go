package comparison

import (
	"jvm-on-go/instructions/base"
	"jvm-on-go/rtda"
)

func IFeq(reader *base.ByteCodeReader, frame *rtda.Frame) {
	offset := int(reader.ReadInt16())
	val := frame.OperandStack.PopInt()
	if val == 0 {
		base.Branch(frame, offset)
	}
}

func IFne(reader *base.ByteCodeReader, frame *rtda.Frame) {
	offset := int(reader.ReadInt16())
	val := frame.OperandStack.PopInt()
	if val != 0 {
		base.Branch(frame, offset)
	}
}

func IFlt(reader *base.ByteCodeReader, frame *rtda.Frame) {
	offset := int(reader.ReadInt16())
	val := frame.OperandStack.PopInt()
	if val < 0 {
		base.Branch(frame, offset)
	}
}

func IFle(reader *base.ByteCodeReader, frame *rtda.Frame) {
	offset := int(reader.ReadInt16())
	val := frame.OperandStack.PopInt()
	if val <= 0 {
		base.Branch(frame, offset)
	}
}

func IFgt(reader *base.ByteCodeReader, frame *rtda.Frame) {
	offset := int(reader.ReadInt16())
	val := frame.OperandStack.PopInt()
	if val > 0 {
		base.Branch(frame, offset)
	}
}

func IFge(reader *base.ByteCodeReader, frame *rtda.Frame) {
	offset := int(reader.ReadInt16())
	val := frame.OperandStack.PopInt()
	if val >= 0 {
		base.Branch(frame, offset)
	}
}
