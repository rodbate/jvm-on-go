package math

import (
	"jvm-on-go/instructions/base"
	"jvm-on-go/rtda"
)

func IDiv(reader *base.ByteCodeReader, frame *rtda.Frame) {
	i1 := frame.OperandStack.PopInt()
	i2 := frame.OperandStack.PopInt()
	if i1 == 0 {
		panic("ArithmeticException: by zero")
	}
	frame.OperandStack.PushInt(i2 / i1)
}

func FDiv(reader *base.ByteCodeReader, frame *rtda.Frame) {
	i1 := frame.OperandStack.PopFloat()
	i2 := frame.OperandStack.PopFloat()
	if i1 == 0 {
		panic("ArithmeticException: by zero")
	}
	frame.OperandStack.PushFloat(i2 / i1)
}

func LDiv(reader *base.ByteCodeReader, frame *rtda.Frame) {
	i1 := frame.OperandStack.PopLong()
	i2 := frame.OperandStack.PopLong()
	if i1 == 0 {
		panic("ArithmeticException: by zero")
	}
	frame.OperandStack.PushLong(i2 / i1)
}

func DDiv(reader *base.ByteCodeReader, frame *rtda.Frame) {
	i1 := frame.OperandStack.PopDouble()
	i2 := frame.OperandStack.PopDouble()
	if i1 == 0 {
		panic("ArithmeticException: by zero")
	}
	frame.OperandStack.PushDouble(i2 / i1)
}
