package reference

import (
	"jvm-on-go/instructions/base"
	"jvm-on-go/rtda"
)

func MonitorEnter(reader *base.ByteCodeReader, frame *rtda.Frame) {
	ref := frame.OperandStack.PopRef()
	if ref == nil {
		panic("java.lang.NullPointer")
	}
}

func MonitorExit(reader *base.ByteCodeReader, frame *rtda.Frame) {
	ref := frame.OperandStack.PopRef()
	if ref == nil {
		panic("java.lang.NullPointer")
	}
}
