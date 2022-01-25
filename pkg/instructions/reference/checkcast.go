package reference

import (
	"github.com/rodbate/jvm-on-go/pkg/instructions/base"
	rtda2 "github.com/rodbate/jvm-on-go/pkg/rtda"
)

func CheckCast(reader *base.ByteCodeReader, frame *rtda2.Frame) {
	index := reader.ReadUint16()
	ref := frame.OperandStack.PopRef()
	frame.OperandStack.PushRef(ref)
	if ref == nil {
		return
	}
	classRef := frame.Method().Class().ConstantPool().GetConstant(index).(*rtda2.ClassRef)
	class := classRef.ResolvedClass()
	if !ref.IsInstanceOf(class) {
		panic("java.lang.ClassCastException")
	}
}
