package reference

import (
	"jvm-on-go/instructions/base"
	"jvm-on-go/rtda"
)

func CheckCast(reader *base.ByteCodeReader, frame *rtda.Frame) {
	index := reader.ReadUint16()
	ref := frame.OperandStack.PopRef()
	frame.OperandStack.PushRef(ref)
	if ref == nil {
		return
	}
	classRef := frame.Method().Class().ConstantPool().GetConstant(index).(*rtda.ClassRef)
	class := classRef.ResolvedClass()
	if !ref.IsInstanceOf(class) {
		panic("java.lang.ClassCastException")
	}
}
