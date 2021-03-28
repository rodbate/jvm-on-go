package reference

import (
	"jvm-on-go/constants/descriptors"
	"jvm-on-go/instructions/base"
	"jvm-on-go/rtda"
)

func PutStatic(reader *base.ByteCodeReader, frame *rtda.Frame) {
	index := reader.ReadUint16()
	fieldRef := frame.Method().Class().ConstantPool().GetConstant(index).(*rtda.FieldRef)
	field := fieldRef.ResolvedField()
	if !field.Class().InitStarted() {
		base.RevertNextPc(reader, frame)
		base.InitClass(frame.Thread(), field.Class())
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
