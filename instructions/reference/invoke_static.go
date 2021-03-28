package reference

import (
	"github.com/rodbate/jvm-on-go/instructions/base"
	"github.com/rodbate/jvm-on-go/rtda"
)

func InvokeStatic(reader *base.ByteCodeReader, frame *rtda.Frame) {
	index := reader.ReadUint16()
	cp := frame.Method().Class().ConstantPool()
	methodRef := cp.GetConstant(index).(*rtda.MethodRef)
	method := methodRef.ResolvedMethod()
	if !method.Class().InitStarted() {
		base.RevertNextPc(reader, frame)
		base.InitClass(frame.Thread(), method.Class())
		return
	}
	if !method.IsStatic() {
		panic("java.lang.IncompatibleClassChangeError")
	}
	base.InvokeMethod(frame, method)
}
