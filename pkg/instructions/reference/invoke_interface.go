package reference

import (
	base2 "github.com/rodbate/jvm-on-go/pkg/instructions/base"
	rtda2 "github.com/rodbate/jvm-on-go/pkg/rtda"
)

func InvokeInterface(reader *base2.ByteCodeReader, frame *rtda2.Frame) {
	index := reader.ReadUint16()
	reader.ReadUint8()
	reader.ReadUint8()

	cp := frame.Method().Class().ConstantPool()
	methodRef := cp.GetConstant(index).(*rtda2.InterfaceMethodRef)
	method := methodRef.ResolvedInterfaceMethod()
	objectRef := frame.OperandStack.GetRefFromTop(uint16(method.ArgSlotCount() - 1))
	if objectRef == nil {
		panic("java.lang.NullPointerException")
	}

	if !method.Class().IsAssignableFrom(objectRef.Class()) {
		panic("java.lang.IncompatibleClassChangeError")
	}

	methodToInvoked := objectRef.Class().LookupMethod(methodRef.Name(), methodRef.Descriptor())
	if methodToInvoked == nil {
		panic("java.lang.NoSuchMethodError")
	}
	if methodToInvoked.IsAbstract() {
		panic("java.lang.AbstractMethodError")
	}
	if !methodToInvoked.IsPublic() {
		panic("java.lang.AccessError")
	}

	base2.InvokeMethod(frame, methodToInvoked)
}
