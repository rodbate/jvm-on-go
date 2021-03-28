package reference

import (
	"github.com/rodbate/jvm-on-go/instructions/base"
	"github.com/rodbate/jvm-on-go/rtda"
)

func ANewArray(reader *base.ByteCodeReader, frame *rtda.Frame) {
	index := reader.ReadUint16()
	count := frame.OperandStack.PopInt()
	if count < 0 {
		panic("java.lang.NegativeArraySizeException")
	}
	cp := frame.Method().Class().ConstantPool()
	ref := cp.GetConstant(index).(*rtda.ClassRef)
	componentClass := ref.ResolvedClass()
	arrayRef := componentClass.ArrayClass().NewArray(uint(count))
	frame.OperandStack.PushRef(arrayRef)
}
