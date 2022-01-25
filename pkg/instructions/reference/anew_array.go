package reference

import (
	"github.com/rodbate/jvm-on-go/pkg/instructions/base"
	rtda2 "github.com/rodbate/jvm-on-go/pkg/rtda"
)

func ANewArray(reader *base.ByteCodeReader, frame *rtda2.Frame) {
	index := reader.ReadUint16()
	count := frame.OperandStack.PopInt()
	if count < 0 {
		panic("java.lang.NegativeArraySizeException")
	}
	cp := frame.Method().Class().ConstantPool()
	ref := cp.GetConstant(index).(*rtda2.ClassRef)
	componentClass := ref.ResolvedClass()
	arrayRef := componentClass.ArrayClass().NewArray(uint(count))
	frame.OperandStack.PushRef(arrayRef)
}
