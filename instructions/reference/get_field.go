package reference

import (
	"jvm-on-go/constants/descriptors"
	"jvm-on-go/instructions/base"
	"jvm-on-go/rtda"
)

func GetField(reader *base.ByteCodeReader, frame *rtda.Frame) {
	index := reader.ReadUint16()
	fieldRef := frame.Method().Class().ConstantPool().GetConstant(index).(*rtda.FieldRef)
	field := fieldRef.ResolvedField()

	if field.IsStatic() {
		panic("java.lang.IncompatibleClassChangeError: " + field.Name())
	}

	switch string(field.Descriptor()[0]) {
	case descriptors.Boolean, descriptors.Byte, descriptors.Char, descriptors.Short, descriptors.Int:
		frame.OperandStack.PushInt(popInstanceRefAndCheckNotNull(frame).Fields().GetInt(field.SlotIndex()))
	case descriptors.Float:
		frame.OperandStack.PushFloat(popInstanceRefAndCheckNotNull(frame).Fields().GetFloat(field.SlotIndex()))
	case descriptors.Double:
		frame.OperandStack.PushDouble(popInstanceRefAndCheckNotNull(frame).Fields().GetDouble(field.SlotIndex()))
	case descriptors.Long:
		frame.OperandStack.PushLong(popInstanceRefAndCheckNotNull(frame).Fields().GetLong(field.SlotIndex()))
	case "L", "[":
		frame.OperandStack.PushRef(popInstanceRefAndCheckNotNull(frame).Fields().GetRef(field.SlotIndex()))
	default:
		panic("unsupported field descriptor: " + field.Descriptor())
	}
}
