package load

import (
	"github.com/rodbate/jvm-on-go/pkg/instructions/base"
	"github.com/rodbate/jvm-on-go/pkg/rtda"
)

/**
iload
iload_0
iload_1
iload_2
iload_3
*/

func DoILoad(index uint16, frame *rtda.Frame) {
	val := frame.LocalVars.GetInt(uint(index))
	frame.OperandStack.PushInt(val)
}

func ILoad(reader *base.ByteCodeReader, frame *rtda.Frame) {
	DoILoad(uint16(reader.ReadUint8()), frame)
}

func ILoad0(reader *base.ByteCodeReader, frame *rtda.Frame) {
	DoILoad(uint16(0), frame)
}

func ILoad1(reader *base.ByteCodeReader, frame *rtda.Frame) {
	DoILoad(uint16(1), frame)
}

func ILoad2(reader *base.ByteCodeReader, frame *rtda.Frame) {
	DoILoad(uint16(2), frame)
}

func ILoad3(reader *base.ByteCodeReader, frame *rtda.Frame) {
	DoILoad(uint16(3), frame)
}
