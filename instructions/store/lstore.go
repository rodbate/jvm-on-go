package store

import (
	"jvm-on-go/instructions/base"
	"jvm-on-go/rtda"
)

func LStore(reader *base.ByteCodeReader, frame *rtda.Frame) {
	index := uint16(reader.ReadUint8())
	DoStoreLong(index, frame)
}

func LStore0(reader *base.ByteCodeReader, frame *rtda.Frame) {
	DoStoreLong(0, frame)
}

func LStore1(reader *base.ByteCodeReader, frame *rtda.Frame) {
	DoStoreLong(1, frame)
}

func LStore2(reader *base.ByteCodeReader, frame *rtda.Frame) {
	DoStoreLong(2, frame)
}

func LStore3(reader *base.ByteCodeReader, frame *rtda.Frame) {
	DoStoreLong(3, frame)
}

func DoStoreLong(index uint16, frame *rtda.Frame) {
	frame.LocalVars.SetLong(uint(index), frame.OperandStack.PopLong())
}
