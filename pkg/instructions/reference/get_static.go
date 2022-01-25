package reference

import (
	"github.com/rodbate/jvm-on-go/pkg/constants/descriptors"
	base2 "github.com/rodbate/jvm-on-go/pkg/instructions/base"
	rtda2 "github.com/rodbate/jvm-on-go/pkg/rtda"
)

func GetStatic(reader *base2.ByteCodeReader, frame *rtda2.Frame) {
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
