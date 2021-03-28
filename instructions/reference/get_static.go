package reference

import (
	"github.com/rodbate/jvm-on-go/constants/descriptors"
	"github.com/rodbate/jvm-on-go/instructions/base"
	"github.com/rodbate/jvm-on-go/rtda"
)

func GetStatic(reader *base.ByteCodeReader, frame *rtda.Frame) {
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
	staticFields := field.Class().StaticFields()
	slotIndex := field.SlotIndex()
	switch string(field.Descriptor()[0]) {
	case descriptors.Boolean, descriptors.Byte, descriptors.Char, descriptors.Short, descriptors.Int:
		frame.OperandStack.PushInt(staticFields.GetInt(slotIndex))
	case descriptors.Float:
		frame.OperandStack.PushFloat(staticFields.GetFloat(slotIndex))
	case descriptors.Double:
		frame.OperandStack.PushDouble(staticFields.GetDouble(slotIndex))
	case descriptors.Long:
		frame.OperandStack.PushLong(staticFields.GetLong(slotIndex))
	case "L", "[":
		frame.OperandStack.PushRef(staticFields.GetRef(slotIndex))
	default:
		panic("unsupported field descriptor: " + field.Descriptor())
	}
}
