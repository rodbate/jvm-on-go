package reflect

import (
	"jvm-on-go/constants/classname"
	"jvm-on-go/native"
	"jvm-on-go/rtda"
)

func init() {
	native.RegisterNative(classname.Reflection, "getCallerClass",
		"()Ljava/lang/Class;", getCallerClass)
	native.RegisterNative(classname.Reflection, "getClassAccessFlags",
		"(Ljava/lang/Class;)I", getClassAccessFlags)
}

//public static native Class<?> getCallerClass();
func getCallerClass(frame *rtda.Frame) {
	callerFrame := frame.Thread().GetFrames()[2]
	callerClass := callerFrame.Method().Class().JClass()
	frame.OperandStack.PushRef(callerClass)
}

//public static native int getClassAccessFlags(Class<?> c);
func getClassAccessFlags(frame *rtda.Frame) {
	classObj := frame.LocalVars.GetRef(0)
	accessFlags := classObj.Extra.(*rtda.Class).AccessFlags
	frame.OperandStack.PushInt(int32(accessFlags))
}
