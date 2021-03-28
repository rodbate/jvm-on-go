package extended

import (
	"github.com/rodbate/jvm-on-go/instructions/base"
	"github.com/rodbate/jvm-on-go/rtda"
)

func IfNull(reader *base.ByteCodeReader, frame *rtda.Frame) {
	offset := reader.ReadInt16()
	ref := frame.OperandStack.PopRef()
	if ref == nil {
		base.Branch(frame, int(offset))
	}
}

func IfNotNull(reader *base.ByteCodeReader, frame *rtda.Frame) {
	offset := reader.ReadInt16()
	ref := frame.OperandStack.PopRef()
	if ref != nil {
		base.Branch(frame, int(offset))
	}
}
