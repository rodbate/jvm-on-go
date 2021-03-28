package conversion

import (
	"jvm-on-go/instructions/base"
	"jvm-on-go/rtda"
)

func I2B(reader *base.ByteCodeReader, frame *rtda.Frame) {
	val := frame.OperandStack.PopInt()
	frame.OperandStack.PushInt(val & 0xFF)
}

func I2C(reader *base.ByteCodeReader, frame *rtda.Frame) {
	val := frame.OperandStack.PopInt()
	frame.OperandStack.PushInt(val & 0xFFFF)
}

func I2D(reader *base.ByteCodeReader, frame *rtda.Frame) {
	val := frame.OperandStack.PopInt()
	frame.OperandStack.PushDouble(float64(val))
}

func I2F(reader *base.ByteCodeReader, frame *rtda.Frame) {
	val := frame.OperandStack.PopInt()
	frame.OperandStack.PushFloat(float32(val))
}

func I2L(reader *base.ByteCodeReader, frame *rtda.Frame) {
	val := frame.OperandStack.PopInt()
	frame.OperandStack.PushLong(int64(val))
}

func I2S(reader *base.ByteCodeReader, frame *rtda.Frame) {
	val := frame.OperandStack.PopInt()
	frame.OperandStack.PushInt(int32(int16(val)))
}
