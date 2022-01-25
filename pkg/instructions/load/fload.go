package load

import (
	"github.com/rodbate/jvm-on-go/pkg/instructions/base"
	"github.com/rodbate/jvm-on-go/pkg/rtda"
)

func DoFLoad(index uint16, frame *rtda.Frame) {
	val := frame.LocalVars.GetFloat(uint(index))
	frame.OperandStack.PushFloat(val)
}

func FLoad(reader *base.ByteCodeReader, frame *rtda.Frame) {
	DoFLoad(uint16(reader.ReadUint8()), frame)
}

func FLoad0(reader *base.ByteCodeReader, frame *rtda.Frame) {
	DoFLoad(uint16(0), frame)
}

func FLoad1(reader *base.ByteCodeReader, frame *rtda.Frame) {
	DoFLoad(uint16(1), frame)
}

func FLoad2(reader *base.ByteCodeReader, frame *rtda.Frame) {
	DoFLoad(uint16(2), frame)
}

func FLoad3(reader *base.ByteCodeReader, frame *rtda.Frame) {
	DoFLoad(uint16(3), frame)
}
