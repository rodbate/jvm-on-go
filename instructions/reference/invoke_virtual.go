package reference

import (
	"github.com/rodbate/jvm-on-go/instructions/base"
	"github.com/rodbate/jvm-on-go/rtda"
)

func InvokeVirtual(reader *base.ByteCodeReader, frame *rtda.Frame) {
	index := reader.ReadUint16()
	currentClass := frame.Method().Class()
	cp := currentClass.ConstantPool()
	methodRef := cp.GetConstant(index).(*rtda.MethodRef)
	resolvedMethod := methodRef.ResolvedMethod()
	if resolvedMethod.IsStatic() {
		panic("java.lang.IncompatibleClassChangeError")
	}
	objectRef := frame.OperandStack.GetRefFromTop(uint16(resolvedMethod.ArgSlotCount() - 1))

	//check protected method
	if resolvedMethod.IsProtected() && resolvedMethod.Class().IsSuperClassOf(currentClass) &&
		currentClass.GetPackageName() != resolvedMethod.Class().GetPackageName() &&
		!currentClass.IsSuperClassOf(objectRef.Class()) &&
		!(objectRef.Class().IsArray() && resolvedMethod.Name() == "clone") {
		panic("java.lang.IllegalAccessError")
	}

	methodToInvoked := objectRef.Class().LookupMethod(methodRef.Name(), methodRef.Descriptor())
	if methodToInvoked == nil {
		panic("java.lang.NoSuchMethod, " + methodRef.Name() + "," + methodRef.Descriptor())
	}

	if methodToInvoked.IsAbstract() {
		panic("java.lang.AbstractMethod")
	}

	base.InvokeMethod(frame, methodToInvoked)
}
