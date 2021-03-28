package comparison

import (
	"github.com/rodbate/jvm-on-go/instructions/base"
	"github.com/rodbate/jvm-on-go/rtda"
)

func IFICmpEq(reader *base.ByteCodeReader, frame *rtda.Frame) {
	_ifICmp(reader, frame, func(i1 int32, i2 int32) bool {
		return i1 == i2
	})
}

func IFICmpNe(reader *base.ByteCodeReader, frame *rtda.Frame) {
	_ifICmp(reader, frame, func(i1 int32, i2 int32) bool {
		return i1 != i2
	})
}

func IFICmpLt(reader *base.ByteCodeReader, frame *rtda.Frame) {
	_ifICmp(reader, frame, func(i1 int32, i2 int32) bool {
		return i1 < i2
	})
}

func IFICmpLe(reader *base.ByteCodeReader, frame *rtda.Frame) {
	_ifICmp(reader, frame, func(i1 int32, i2 int32) bool {
		return i1 <= i2
	})
}

func IFICmpGt(reader *base.ByteCodeReader, frame *rtda.Frame) {
	_ifICmp(reader, frame, func(i1 int32, i2 int32) bool {
		return i1 > i2
	})
}

func IFICmpGe(reader *base.ByteCodeReader, frame *rtda.Frame) {
	_ifICmp(reader, frame, func(i1 int32, i2 int32) bool {
		return i1 >= i2
	})
}

func _ifICmp(reader *base.ByteCodeReader, frame *rtda.Frame,
	compFunc func(int32, int32) bool) {
	offset := int(reader.ReadInt16())
	val2 := frame.OperandStack.PopInt()
	val1 := frame.OperandStack.PopInt()
	if compFunc(val1, val2) {
		base.Branch(frame, offset)
	}
}
