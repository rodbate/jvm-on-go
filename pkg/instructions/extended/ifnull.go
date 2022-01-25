package extended

import (
	base2 "github.com/rodbate/jvm-on-go/pkg/instructions/base"
	"github.com/rodbate/jvm-on-go/pkg/rtda"
)

func IfNull(reader *base2.ByteCodeReader, frame *rtda.Frame) {
	offset := reader.ReadInt16()
	ref := frame.OperandStack.PopRef()
	if ref == nil {
		base2.Branch(frame, int(offset))
	}
}

func IfNotNull(reader *base2.ByteCodeReader, frame *rtda.Frame) {
	offset := reader.ReadInt16()
	ref := frame.OperandStack.PopRef()
	if ref != nil {
		base2.Branch(frame, int(offset))
	}
}
