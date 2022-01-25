package lang

import (
	"github.com/rodbate/jvm-on-go/pkg/constants/classname"
	"github.com/rodbate/jvm-on-go/pkg/native"
	"github.com/rodbate/jvm-on-go/pkg/rtda"
	"unsafe"
)

func init() {
	native.RegisterNative(classname.Object, "getClass", "()Ljava/lang/Class;", getClass)
	native.RegisterNative(classname.Object, "hashCode", "()I", hashCode)
	native.RegisterNative(classname.Object, "clone", "()Ljava/lang/Object;", clone)
	native.RegisterNative(classname.Object, "notifyAll", "()V", notifyAll)
}

//public final native Class<?> getClass()
func getClass(frame *rtda.Frame) {
	this := frame.LocalVars.GetThis()
	frame.OperandStack.PushRef(this.Class().JClass())
}

//public native int hashCode()
func hashCode(frame *rtda.Frame) {
	this := frame.LocalVars.GetThis()
	frame.OperandStack.PushInt(int32(uintptr(unsafe.Pointer(this))))
}

//protected native Object clone() throws CloneNotSupportedException
func clone(frame *rtda.Frame) {
	this := frame.LocalVars.GetThis()
	cloneable := frame.Method().Class().ClassLoader().LoadClass(classname.Cloneable)
	if !cloneable.IsAssignableFrom(this.Class()) {
		panic("java.lang.CloneNotSupportedException")
	}
	frame.OperandStack.PushRef(this.Clone())
}

//public final native void notifyAll()
func notifyAll(frame *rtda.Frame) {
	//todo
}
