package reference

import (
	"github.com/rodbate/jvm-on-go/pkg/constants/descriptors"
	base2 "github.com/rodbate/jvm-on-go/pkg/instructions/base"
	rtda2 "github.com/rodbate/jvm-on-go/pkg/rtda"
)

func PutStatic(reader *base2.ByteCodeReader, frame *rtda2.Frame) {
	index := reader.ReadUint16()
	fieldRef := frame.Method().Class().ConstantPool().GetConstant(index).(*rtda2.FieldRef)
	field := fieldRef.ResolvedField()
	if !field.Class().InitStarted() {
		base2.RevertNextPc(reader, frame)
		base2.InitClass(frame.Thread(), field.Class())
		return
	}
	if !field.IsStatic() {
		panic("java.lang.IncompatibleClassChangeError")
	}
	if field.IsFinal() && (frame.Method().Class() != field.Class() ||
		frame.Method().Name() != "<clinit>") {
		panic("java.lang.IllegalAccessError")
	}

	staticFields := field.Class().StaticFields()
	slotIndex := field.SlotIndex()
	switch string(field.Descriptor()[0]) {
	case descriptors.Boolean, descriptors.Byte, descriptors.Char, descriptors.Short, descriptors.Int:
		staticFields.SetInt(slotIndex, frame.OperandStack.PopInt())
	case descriptors.Float:
		staticFields.SetFloat(slotIndex, frame.OperandStack.PopFloat())
	case descriptors.Double:
		staticFields.SetDouble(slotIndex, frame.OperandStack.PopDouble())
	case descriptors.Long:
		staticFields.SetLong(slotIndex, frame.OperandStack.PopLong())
	case "L", "[":
		staticFields.SetRef(slotIndex, frame.OperandStack.PopRef())
	default:
		panic("unsupported field descriptor: " + field.Descriptor())
	}
}
