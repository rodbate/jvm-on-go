package load

import (
	"github.com/rodbate/jvm-on-go/pkg/instructions/base"
	"github.com/rodbate/jvm-on-go/pkg/rtda"
)

func DoLLoad(index uint16, frame *rtda.Frame) {
	val := frame.LocalVars.GetLong(uint(index))
	frame.OperandStack.PushLong(val)
}

func LLoad(reader *base.ByteCodeReader, frame *rtda.Frame) {
	DoLLoad(uint16(reader.ReadUint8()), frame)
}

func LLoad0(reader *base.ByteCodeReader, frame *rtda.Frame) {
	DoLLoad(uint16(0), frame)
}

func LLoad1(reader *base.ByteCodeReader, frame *rtda.Frame) {
	DoLLoad(uint16(1), frame)
}

func LLoad2(reader *base.ByteCodeReader, frame *rtda.Frame) {
	DoLLoad(uint16(2), frame)
}

func LLoad3(reader *base.ByteCodeReader, frame *rtda.Frame) {
	DoLLoad(uint16(3), frame)
}
