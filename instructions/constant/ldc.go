package constant

import (
	"fmt"
	"github.com/rodbate/jvm-on-go/instructions/base"
	"github.com/rodbate/jvm-on-go/rtda"
)

/**
ldc
ldc_w
ldc2_w
*/

func Ldc(reader *base.ByteCodeReader, frame *rtda.Frame) {
	index := reader.ReadUint8()
	_ldc(frame, uint16(index))
}

func LdcW(reader *base.ByteCodeReader, frame *rtda.Frame) {
	index := reader.ReadUint16()
	_ldc(frame, index)
}

func Ldc2W(reader *base.ByteCodeReader, frame *rtda.Frame) {
	index := reader.ReadUint16()
	constant := frame.Method().Class().ConstantPool().GetConstant(index)
	switch constant.(type) {
	case int64:
		frame.OperandStack.PushLong(constant.(int64))
	case float64:
		frame.OperandStack.PushDouble(constant.(float64))
	default:
		panic("ClassFormat Error")
	}
}

func _ldc(frame *rtda.Frame, index uint16) {
	constant := frame.Method().Class().ConstantPool().GetConstant(index)
	switch constant.(type) {
	case int32:
		frame.OperandStack.PushInt(constant.(int32))
	case float32:
		frame.OperandStack.PushFloat(constant.(float32))
	case string:
		jStr := rtda.GetJString(frame.Method().Class().ClassLoader(), constant.(string))
		frame.OperandStack.PushRef(jStr)
	case *rtda.ClassRef:
		classRef := constant.(*rtda.ClassRef)
		frame.OperandStack.PushRef(classRef.ResolvedClass().JClass())
	default:
		panic(fmt.Sprintf("not supported ldc const type: %v", constant))
	}
}
