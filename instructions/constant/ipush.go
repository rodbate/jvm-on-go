package constant

import (
	"github.com/rodbate/jvm-on-go/instructions/base"
	"github.com/rodbate/jvm-on-go/rtda"
)

/**
bipush
sipush
*/

func BIPush(reader *base.ByteCodeReader, frame *rtda.Frame) {
	val := reader.ReadInt8()
	frame.OperandStack.PushInt(int32(val))
}

func SIPush(reader *base.ByteCodeReader, frame *rtda.Frame) {
	val := reader.ReadInt16()
	frame.OperandStack.PushInt(int32(val))
}
