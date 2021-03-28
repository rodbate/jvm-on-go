package reflect

import (
	"github.com/rodbate/jvm-on-go/constants/classname"
	"github.com/rodbate/jvm-on-go/instructions/base"
	"github.com/rodbate/jvm-on-go/native"
	"github.com/rodbate/jvm-on-go/rtda"
)

func init() {
	native.RegisterNative(classname.NativeConstructorAccessorImpl, "newInstance0",
		"(Ljava/lang/reflect/Constructor;[Ljava/lang/Object;)Ljava/lang/Object;", newInstance0)
}

//private static native Object newInstance0(Constructor<?> c, Object[] args)
func newInstance0(frame *rtda.Frame) {
	ctr := frame.LocalVars.GetRef(0)
	args := frame.LocalVars.GetRef(1)
	ctrDescriptor := rtda.ToMethodDescriptor(args)
	classObj := ctr.GetFieldValue("clazz", "Ljava/lang/Class;").(*rtda.Object)
	javaCtr := classObj.Extra.(*rtda.Class).GetConstructor(ctrDescriptor)
	obj := classObj.Extra.(*rtda.Class).NewInstance()
	frame.OperandStack.PushRef(obj)

	thread := frame.Thread()
	stack := rtda.NewOperandStack(uint16(javaCtr.ArgSlotCount()))
	stack.PushRef(obj)

	newFrame := rtda.NewMockFrame(thread, stack)
	thread.PushFrame(newFrame)
	base.InvokeMethod(newFrame, javaCtr)
}
