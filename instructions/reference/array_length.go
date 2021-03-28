package reference

import (
	"jvm-on-go/instructions/base"
	"jvm-on-go/rtda"
)

func ArrayLength(reader *base.ByteCodeReader, frame *rtda.Frame) {
	arrayRef := frame.OperandStack.PopRef()
	if arrayRef == nil {
		panic("java.lang.NullPointerException")
	}
	frame.OperandStack.PushInt(arrayRef.ArrayLength())
}
