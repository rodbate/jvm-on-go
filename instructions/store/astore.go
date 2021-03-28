package store

import (
	"jvm-on-go/instructions/base"
	"jvm-on-go/rtda"
)

func AStore(reader *base.ByteCodeReader, frame *rtda.Frame) {
	index := uint16(reader.ReadUint8())
	DoStoreRef(index, frame)
}

func AStore0(reader *base.ByteCodeReader, frame *rtda.Frame) {
	DoStoreRef(0, frame)
}

func AStore1(reader *base.ByteCodeReader, frame *rtda.Frame) {
	DoStoreRef(1, frame)
}

func AStore2(reader *base.ByteCodeReader, frame *rtda.Frame) {
	DoStoreRef(2, frame)
}

func AStore3(reader *base.ByteCodeReader, frame *rtda.Frame) {
	DoStoreRef(3, frame)
}

func DoStoreRef(index uint16, frame *rtda.Frame) {
	frame.LocalVars.SetRef(uint(index), frame.OperandStack.PopRef())
}
