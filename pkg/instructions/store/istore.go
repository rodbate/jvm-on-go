package store

import (
	"github.com/rodbate/jvm-on-go/pkg/instructions/base"
	"github.com/rodbate/jvm-on-go/pkg/rtda"
)

func IStore(reader *base.ByteCodeReader, frame *rtda.Frame) {
	index := uint16(reader.ReadUint8())
	DoStoreInt(index, frame)
}

func IStore0(reader *base.ByteCodeReader, frame *rtda.Frame) {
	DoStoreInt(0, frame)
}

func IStore1(reader *base.ByteCodeReader, frame *rtda.Frame) {
	DoStoreInt(1, frame)
}

func IStore2(reader *base.ByteCodeReader, frame *rtda.Frame) {
	DoStoreInt(2, frame)
}

func IStore3(reader *base.ByteCodeReader, frame *rtda.Frame) {
	DoStoreInt(3, frame)
}

func DoStoreInt(index uint16, frame *rtda.Frame) {
	frame.LocalVars.SetInt(uint(index), frame.OperandStack.PopInt())
}
