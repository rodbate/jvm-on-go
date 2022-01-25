package reference

import (
	base2 "github.com/rodbate/jvm-on-go/pkg/instructions/base"
	rtda2 "github.com/rodbate/jvm-on-go/pkg/rtda"
)

func InvokeSpecial(reader *base2.ByteCodeReader, frame *rtda2.Frame) {
	index := reader.ReadUint16()
	currentClass := frame.Method().Class()
	cp := currentClass.ConstantPool()
	methodRef := cp.GetConstant(index).(*rtda2.MethodRef)
	resolvedClass := methodRef.ResolvedClass()
	resolvedMethod := methodRef.ResolvedMethod()
	if resolvedMethod.Name() == "<init>" && resolvedClass != resolvedMethod.Class() {
		panic("java.lang.NoSuchMethodError")
	}
	if resolvedMethod.IsStatic() {
		panic("java.lang.IncompatibleClassChangeError")
	}
	objectRef := frame.OperandStack.GetRefFromTop(uint16(resolvedMethod.ArgSlotCount() - 1))
	if objectRef == nil {
		panic("java.lang.NullPointerException")
	}

	//check protected method
	if resolvedMethod.IsProtected() && resolvedMethod.Class().IsSuperClassOf(currentClass) &&
		currentClass.GetPackageName() != resolvedMethod.Class().GetPackageName() &&
		!currentClass.IsSuperClassOf(objectRef.Class()) && currentClass != objectRef.Class() {
		panic("java.lang.IllegalAccessError")
	}

	methodToInvoked := resolvedMethod
	//ACC_SUPER
	if resolvedMethod.Name() != "<init>" && !resolvedClass.IsInterface() &&
		resolvedClass.IsSuperClassOf(currentClass) && currentClass.IsSuper() {
		methodToInvoked = currentClass.SuperClass().LookupMethod(methodRef.Name(), methodRef.Descriptor())
	}

	if methodToInvoked == nil {
		panic("java.lang.NoSuchMethod")
	}

	if methodToInvoked.IsAbstract() {
		panic("java.lang.AbstractMethod")
	}

	base2.InvokeMethod(frame, methodToInvoked)
}
