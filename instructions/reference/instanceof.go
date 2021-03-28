package reference

import (
	"github.com/rodbate/jvm-on-go/instructions/base"
	"github.com/rodbate/jvm-on-go/rtda"
)

func InstanceOf(reader *base.ByteCodeReader, frame *rtda.Frame) {
	index := reader.ReadUint16()
	instance := frame.OperandStack.PopRef()
	if instance == nil {
		frame.OperandStack.PushInt(0)
		return
	}
	classRef := frame.Method().Class().ConstantPool().GetConstant(index).(*rtda.ClassRef)
	class := classRef.ResolvedClass()
	if instance.IsInstanceOf(class) {
		frame.OperandStack.PushInt(1)
	} else {
		frame.OperandStack.PushInt(0)
	}
}
