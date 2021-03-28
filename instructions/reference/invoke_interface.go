package reference

import (
	"jvm-on-go/instructions/base"
	"jvm-on-go/rtda"
)

func InvokeInterface(reader *base.ByteCodeReader, frame *rtda.Frame) {
	index := reader.ReadUint16()
	reader.ReadUint8()
	reader.ReadUint8()

	cp := frame.Method().Class().ConstantPool()
	methodRef := cp.GetConstant(index).(*rtda.InterfaceMethodRef)
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

	base.InvokeMethod(frame, methodToInvoked)
}
