package reference

import (
	"github.com/rodbate/jvm-on-go/pkg/instructions/base"
	rtda2 "github.com/rodbate/jvm-on-go/pkg/rtda"
)

func MultiANewArray(reader *base.ByteCodeReader, frame *rtda2.Frame) {
	index := reader.ReadUint16()
	dimensions := reader.ReadUint8()

	cp := frame.Method().Class().ConstantPool()
	classRef := cp.GetConstant(index).(*rtda2.ClassRef)
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

func newMultiArray(counts []uint, arrayClass *rtda2.Class) *rtda2.Object {
	array := arrayClass.NewArray(counts[0])
	if len(counts) > 1 {
		refs := array.Refs()
		for i := range refs {
			refs[i] = newMultiArray(counts[1:], arrayClass.ComponentClass())
		}
	}
	return array
}
