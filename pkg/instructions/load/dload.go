package load

import (
	"github.com/rodbate/jvm-on-go/pkg/instructions/base"
	"github.com/rodbate/jvm-on-go/pkg/rtda"
)

func DoDLoad(index uint16, frame *rtda.Frame) {
	val := frame.LocalVars.GetDouble(uint(index))
	frame.OperandStack.PushDouble(val)
}

func DLoad(reader *base.ByteCodeReader, frame *rtda.Frame) {
	DoDLoad(uint16(reader.ReadUint8()), frame)
}

func DLoad0(reader *base.ByteCodeReader, frame *rtda.Frame) {
	DoDLoad(uint16(0), frame)
}

func DLoad1(reader *base.ByteCodeReader, frame *rtda.Frame) {
	DoDLoad(uint16(0), frame)
}

func DLoad2(reader *base.ByteCodeReader, frame *rtda.Frame) {
	DoDLoad(uint16(0), frame)
}

func DLoad3(reader *base.ByteCodeReader, frame *rtda.Frame) {
	DoDLoad(uint16(0), frame)
}
