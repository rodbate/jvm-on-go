package reference

import (
	"github.com/rodbate/jvm-on-go/instructions/base"
	"github.com/rodbate/jvm-on-go/rtda"
)

func New(reader *base.ByteCodeReader, frame *rtda.Frame) {
	index := reader.ReadUint16()
	cp := frame.Method().Class().ConstantPool()
	classRef := cp.GetConstant(index).(*rtda.ClassRef)
	class := classRef.ResolvedClass()
	if !class.InitStarted() {
		base.RevertNextPc(reader, frame)
		base.InitClass(frame.Thread(), class)
		return
	}
	if class.IsInterface() || class.IsAbstract() {
		panic("java.lang.InstantiationError")
	}
	frame.OperandStack.PushRef(class.NewInstance())
}
