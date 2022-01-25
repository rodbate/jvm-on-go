package reference

import (
	base2 "github.com/rodbate/jvm-on-go/pkg/instructions/base"
	rtda2 "github.com/rodbate/jvm-on-go/pkg/rtda"
)

func New(reader *base2.ByteCodeReader, frame *rtda2.Frame) {
	index := reader.ReadUint16()
	cp := frame.Method().Class().ConstantPool()
	classRef := cp.GetConstant(index).(*rtda2.ClassRef)
	class := classRef.ResolvedClass()
	if !class.InitStarted() {
		base2.RevertNextPc(reader, frame)
		base2.InitClass(frame.Thread(), class)
		return
	}
	if class.IsInterface() || class.IsAbstract() {
		panic("java.lang.InstantiationError")
	}
	frame.OperandStack.PushRef(class.NewInstance())
}
