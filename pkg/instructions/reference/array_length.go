package reference

import (
	"github.com/rodbate/jvm-on-go/pkg/instructions/base"
	"github.com/rodbate/jvm-on-go/pkg/rtda"
)

func ArrayLength(reader *base.ByteCodeReader, frame *rtda.Frame) {
	arrayRef := frame.OperandStack.PopRef()
	if arrayRef == nil {
		panic("java.lang.NullPointerException")
	}
	frame.OperandStack.PushInt(arrayRef.ArrayLength())
}
