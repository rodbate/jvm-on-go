package constant

import (
	"jvm-on-go/instructions/base"
	"jvm-on-go/rtda"
)

/**
aconst_null
iconst_m1
iconst_0
iconst_1
iconst_2
iconst_3
iconst_4
iconst_5
lconst_0
lconst_1
fconst_0
fconst_1
fconst_2
dconst_0
dconst_1
*/

func AConstNull(reader *base.ByteCodeReader, frame *rtda.Frame) {
	frame.OperandStack.PushRef(nil)
}

func IConstM1(reader *base.ByteCodeReader, frame *rtda.Frame) {
	frame.OperandStack.PushInt(-1)
}

func IConst0(reader *base.ByteCodeReader, frame *rtda.Frame) {
	frame.OperandStack.PushInt(0)
}

func IConst1(reader *base.ByteCodeReader, frame *rtda.Frame) {
	frame.OperandStack.PushInt(1)
}

func IConst2(reader *base.ByteCodeReader, frame *rtda.Frame) {
	frame.OperandStack.PushInt(2)
}

func IConst3(reader *base.ByteCodeReader, frame *rtda.Frame) {
	frame.OperandStack.PushInt(3)
}

func IConst4(reader *base.ByteCodeReader, frame *rtda.Frame) {
	frame.OperandStack.PushInt(4)
}

func IConst5(reader *base.ByteCodeReader, frame *rtda.Frame) {
	frame.OperandStack.PushInt(5)
}

func LConst0(reader *base.ByteCodeReader, frame *rtda.Frame) {
	frame.OperandStack.PushLong(0)
}

func LConst1(reader *base.ByteCodeReader, frame *rtda.Frame) {
	frame.OperandStack.PushLong(1)
}

func FConst0(reader *base.ByteCodeReader, frame *rtda.Frame) {
	frame.OperandStack.PushFloat(0)
}

func FConst1(reader *base.ByteCodeReader, frame *rtda.Frame) {
	frame.OperandStack.PushFloat(1)
}

func FConst2(reader *base.ByteCodeReader, frame *rtda.Frame) {
	frame.OperandStack.PushFloat(2)
}

func DConst0(reader *base.ByteCodeReader, frame *rtda.Frame) {
	frame.OperandStack.PushDouble(0)
}

func DConst1(reader *base.ByteCodeReader, frame *rtda.Frame) {
	frame.OperandStack.PushDouble(1)
}
