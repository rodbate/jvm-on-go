package security

import (
	"jvm-on-go/constants/classname"
	"jvm-on-go/instructions/base"
	"jvm-on-go/native"
	"jvm-on-go/rtda"
)

func init() {
	native.RegisterNative(classname.AccessController, "doPrivileged",
		"(Ljava/security/PrivilegedExceptionAction;)Ljava/lang/Object;", doPrivileged)
	native.RegisterNative(classname.AccessController, "doPrivileged",
		"(Ljava/security/PrivilegedAction;)Ljava/lang/Object;", doPrivileged)
	native.RegisterNative(classname.AccessController, "doPrivileged",
		"(Ljava/security/PrivilegedAction;Ljava/security/AccessControlContext;)Ljava/lang/Object;", doPrivileged)
	native.RegisterNative(classname.AccessController, "doPrivileged",
		"(Ljava/security/PrivilegedExceptionAction;Ljava/security/AccessControlContext;)Ljava/lang/Object;", doPrivileged)
	native.RegisterNative(classname.AccessController, "getStackAccessControlContext",
		"()Ljava/security/AccessControlContext;", getStackAccessControlContext)
}

//public static native <T> T doPrivileged(PrivilegedAction<T> action);
//T run()
func doPrivileged(frame *rtda.Frame) {
	actionClass := frame.LocalVars.GetRef(0).Class()
	method := actionClass.GetInstanceMethod("run", "()Ljava/lang/Object;")
	frame.OperandStack.PushRef(frame.LocalVars.GetRef(0))
	base.InvokeMethod(frame, method)
}

func getStackAccessControlContext(frame *rtda.Frame) {
	frame.OperandStack.PushRef(nil)
}
