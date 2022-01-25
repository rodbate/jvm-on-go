package reference

import (
	base2 "github.com/rodbate/jvm-on-go/pkg/instructions/base"
	rtda2 "github.com/rodbate/jvm-on-go/pkg/rtda"
)

func InvokeStatic(reader *base2.ByteCodeReader, frame *rtda2.Frame) {
	index := reader.ReadUint16()
	cp := frame.Method().Class().ConstantPool()
	methodRef := cp.GetConstant(index).(*rtda2.MethodRef)
	method := methodRef.ResolvedMethod()
	if !method.Class().InitStarted() {
		base2.RevertNextPc(reader, frame)
		base2.InitClass(frame.Thread(), method.Class())
		return
	}
	if !method.IsStatic() {
		panic("java.lang.IncompatibleClassChangeError")
	}
	base2.InvokeMethod(frame, method)
}
