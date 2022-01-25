package comparison

import (
	base2 "github.com/rodbate/jvm-on-go/pkg/instructions/base"
	"github.com/rodbate/jvm-on-go/pkg/rtda"
)

func IFICmpEq(reader *base2.ByteCodeReader, frame *rtda.Frame) {
	_ifICmp(reader, frame, func(i1 int32, i2 int32) bool {
		return i1 == i2
	})
}

func IFICmpNe(reader *base2.ByteCodeReader, frame *rtda.Frame) {
	_ifICmp(reader, frame, func(i1 int32, i2 int32) bool {
		return i1 != i2
	})
}

func IFICmpLt(reader *base2.ByteCodeReader, frame *rtda.Frame) {
	_ifICmp(reader, frame, func(i1 int32, i2 int32) bool {
		return i1 < i2
	})
}

func IFICmpLe(reader *base2.ByteCodeReader, frame *rtda.Frame) {
	_ifICmp(reader, frame, func(i1 int32, i2 int32) bool {
		return i1 <= i2
	})
}

func IFICmpGt(reader *base2.ByteCodeReader, frame *rtda.Frame) {
	_ifICmp(reader, frame, func(i1 int32, i2 int32) bool {
		return i1 > i2
	})
}

func IFICmpGe(reader *base2.ByteCodeReader, frame *rtda.Frame) {
	_ifICmp(reader, frame, func(i1 int32, i2 int32) bool {
		return i1 >= i2
	})
}

func _ifICmp(reader *base2.ByteCodeReader, frame *rtda.Frame,
	compFunc func(int32, int32) bool) {
	offset := int(reader.ReadInt16())
	val2 := frame.OperandStack.PopInt()
	val1 := frame.OperandStack.PopInt()
	if compFunc(val1, val2) {
		base2.Branch(frame, offset)
	}
}
