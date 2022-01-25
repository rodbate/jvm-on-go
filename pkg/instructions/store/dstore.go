package store

import (
	"github.com/rodbate/jvm-on-go/pkg/instructions/base"
	"github.com/rodbate/jvm-on-go/pkg/rtda"
)

func DStore(reader *base.ByteCodeReader, frame *rtda.Frame) {
	index := uint16(reader.ReadUint8())
	DoStoreDouble(index, frame)
}

func DStore0(reader *base.ByteCodeReader, frame *rtda.Frame) {
	DoStoreDouble(0, frame)
}

func DStore1(reader *base.ByteCodeReader, frame *rtda.Frame) {
	DoStoreDouble(1, frame)
}

func DStore2(reader *base.ByteCodeReader, frame *rtda.Frame) {
	DoStoreDouble(2, frame)
}

func DStore3(reader *base.ByteCodeReader, frame *rtda.Frame) {
	DoStoreDouble(3, frame)
}

func DoStoreDouble(index uint16, frame *rtda.Frame) {
	frame.LocalVars.SetDouble(uint(index), frame.OperandStack.PopDouble())
}
