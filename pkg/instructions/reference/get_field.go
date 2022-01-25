package reference

import (
	"github.com/rodbate/jvm-on-go/pkg/constants/descriptors"
	"github.com/rodbate/jvm-on-go/pkg/instructions/base"
	rtda2 "github.com/rodbate/jvm-on-go/pkg/rtda"
)

func GetField(reader *base.ByteCodeReader, frame *rtda2.Frame) {
	index := reader.ReadUint16()
	fieldRef := frame.Method().Class().ConstantPool().GetConstant(index).(*rtda2.FieldRef)
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
