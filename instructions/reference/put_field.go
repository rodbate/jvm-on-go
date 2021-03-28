package reference

import (
	"jvm-on-go/constants/descriptors"
	"jvm-on-go/instructions/base"
	"jvm-on-go/rtda"
)

func PutField(reader *base.ByteCodeReader, frame *rtda.Frame) {
	index := reader.ReadUint16()
	fieldRef := frame.Method().Class().ConstantPool().GetConstant(index).(*rtda.FieldRef)
	field := fieldRef.ResolvedField()
	if field.IsStatic() {
		panic("java.lang.IncompatibleClassChangeError: " + field.Name())
	}
	if field.IsFinal() && (field.Class() != frame.Method().Class() ||
		frame.Method().Name() != "<init>") {
		panic("java.lang.IllegalAccessError: " + field.Name())
	}

	switch string(field.Descriptor()[0]) {
	case descriptors.Boolean, descriptors.Byte, descriptors.Char, descriptors.Short, descriptors.Int:
		val := frame.OperandStack.PopInt()
		popInstanceRefAndCheckNotNull(frame).Fields().SetInt(field.SlotIndex(), val)
	case descriptors.Float:
		val := frame.OperandStack.PopFloat()
		popInstanceRefAndCheckNotNull(frame).Fields().SetFloat(field.SlotIndex(), val)
	case descriptors.Double:
		val := frame.OperandStack.PopDouble()
		popInstanceRefAndCheckNotNull(frame).Fields().SetDouble(field.SlotIndex(), val)
	case descriptors.Long:
		val := frame.OperandStack.PopLong()
		popInstanceRefAndCheckNotNull(frame).Fields().SetLong(field.SlotIndex(), val)
	case "L", "[":
		ref := frame.OperandStack.PopRef()
		popInstanceRefAndCheckNotNull(frame).Fields().SetRef(field.SlotIndex(), ref)
	default:
		panic("unsupported field descriptor: " + field.Descriptor())
	}
}

func popInstanceRefAndCheckNotNull(frame *rtda.Frame) *rtda.Object {
	instanceRef := frame.OperandStack.PopRef()
	if instanceRef == nil {
		panic("java.lang.NullPointerException")
	}
	return instanceRef
}
