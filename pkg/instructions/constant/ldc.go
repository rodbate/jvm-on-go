package constant

import (
	"fmt"
	"github.com/rodbate/jvm-on-go/pkg/instructions/base"
	rtda2 "github.com/rodbate/jvm-on-go/pkg/rtda"
)

/**
ldc
ldc_w
ldc2_w
*/

func Ldc(reader *base.ByteCodeReader, frame *rtda2.Frame) {
	index := reader.ReadUint8()
	_ldc(frame, uint16(index))
}

func LdcW(reader *base.ByteCodeReader, frame *rtda2.Frame) {
	index := reader.ReadUint16()
	_ldc(frame, index)
}

func Ldc2W(reader *base.ByteCodeReader, frame *rtda2.Frame) {
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

func _ldc(frame *rtda2.Frame, index uint16) {
	constant := frame.Method().Class().ConstantPool().GetConstant(index)
	switch constant.(type) {
	case int32:
		frame.OperandStack.PushInt(constant.(int32))
	case float32:
		frame.OperandStack.PushFloat(constant.(float32))
	case string:
		jStr := rtda2.GetJString(frame.Method().Class().ClassLoader(), constant.(string))
		frame.OperandStack.PushRef(jStr)
	case *rtda2.ClassRef:
		classRef := constant.(*rtda2.ClassRef)
		frame.OperandStack.PushRef(classRef.ResolvedClass().JClass())
	default:
		panic(fmt.Sprintf("not supported ldc const type: %v", constant))
	}
}
