package load

import (
	"github.com/rodbate/jvm-on-go/instructions/base"
	"github.com/rodbate/jvm-on-go/rtda"
)

func DoALoad(index uint16, frame *rtda.Frame) {
	ref := frame.LocalVars.GetRef(uint(index))
	frame.OperandStack.PushRef(ref)
}

func ALoad(reader *base.ByteCodeReader, frame *rtda.Frame) {
	DoALoad(uint16(reader.ReadUint8()), frame)
}

func ALoad0(reader *base.ByteCodeReader, frame *rtda.Frame) {
	DoALoad(uint16(0), frame)
}

func ALoad1(reader *base.ByteCodeReader, frame *rtda.Frame) {
	DoALoad(uint16(1), frame)
}

func ALoad2(reader *base.ByteCodeReader, frame *rtda.Frame) {
	DoALoad(uint16(2), frame)
}

func ALoad3(reader *base.ByteCodeReader, frame *rtda.Frame) {
	DoALoad(uint16(3), frame)
}
