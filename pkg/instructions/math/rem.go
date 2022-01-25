package math

import (
	"github.com/rodbate/jvm-on-go/pkg/instructions/base"
	"github.com/rodbate/jvm-on-go/pkg/rtda"
	"math"
)

func IRem(reader *base.ByteCodeReader, frame *rtda.Frame) {
	i1 := frame.OperandStack.PopInt()
	i2 := frame.OperandStack.PopInt()
	if i1 == 0 {
		panic("ArithmeticException: % by zero")
	}
	frame.OperandStack.PushInt(i2 % i1)
}

func FRem(reader *base.ByteCodeReader, frame *rtda.Frame) {
	i1 := frame.OperandStack.PopFloat()
	i2 := frame.OperandStack.PopFloat()
	frame.OperandStack.PushFloat(float32(math.Mod(float64(i2), float64(i1))))
}

func LRem(reader *base.ByteCodeReader, frame *rtda.Frame) {
	i1 := frame.OperandStack.PopLong()
	i2 := frame.OperandStack.PopLong()
	if i1 == 0 {
		panic("ArithmeticException: % by zero")
	}
	frame.OperandStack.PushLong(i2 % i1)
}

func DRem(reader *base.ByteCodeReader, frame *rtda.Frame) {
	i1 := frame.OperandStack.PopDouble()
	i2 := frame.OperandStack.PopDouble()
	frame.OperandStack.PushDouble(math.Mod(i2, i1))
}
