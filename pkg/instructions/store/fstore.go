package store

import (
	"github.com/rodbate/jvm-on-go/pkg/instructions/base"
	"github.com/rodbate/jvm-on-go/pkg/rtda"
)

func FStore(reader *base.ByteCodeReader, frame *rtda.Frame) {
	index := uint16(reader.ReadUint8())
	DoStoreFloat(index, frame)
}

func FStore0(reader *base.ByteCodeReader, frame *rtda.Frame) {
	DoStoreFloat(0, frame)
}

func FStore1(reader *base.ByteCodeReader, frame *rtda.Frame) {
	DoStoreFloat(1, frame)
}

func FStore2(reader *base.ByteCodeReader, frame *rtda.Frame) {
	DoStoreFloat(2, frame)
}

func FStore3(reader *base.ByteCodeReader, frame *rtda.Frame) {
	DoStoreFloat(3, frame)
}

func DoStoreFloat(index uint16, frame *rtda.Frame) {
	frame.LocalVars.SetFloat(uint(index), frame.OperandStack.PopFloat())
}
