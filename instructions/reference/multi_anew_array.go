package reference

import (
	"jvm-on-go/instructions/base"
	"jvm-on-go/rtda"
)

func MultiANewArray(reader *base.ByteCodeReader, frame *rtda.Frame) {
	index := reader.ReadUint16()
	dimensions := reader.ReadUint8()

	cp := frame.Method().Class().ConstantPool()
	classRef := cp.GetConstant(index).(*rtda.ClassRef)
	class := classRef.ResolvedClass() //[[[J

	counts := make([]uint, dimensions)
	for i := int(dimensions - 1); i >= 0; i-- {
		count := frame.OperandStack.PopInt()
		if count < 0 {
			panic("java.lang.NegativeArraySizeException")
		}
		counts[i] = uint(count)
	}

	frame.OperandStack.PushRef(newMultiArray(counts, class))
}

func newMultiArray(counts []uint, arrayClass *rtda.Class) *rtda.Object {
	array := arrayClass.NewArray(counts[0])
	if len(counts) > 1 {
		refs := array.Refs()
		for i := range refs {
			refs[i] = newMultiArray(counts[1:], arrayClass.ComponentClass())
		}
	}
	return array
}
