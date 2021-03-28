package lang

import (
	"jvm-on-go/constants/classname"
	"jvm-on-go/native"
	"jvm-on-go/rtda"
)

func init() {
	native.RegisterNative(classname.ClassLoader, "findBuiltinLib",
		"(Ljava/lang/String;)Ljava/lang/String;", findBuiltinLib)
	native.RegisterNative(classname.ClassLoader, "findLoadedClass0",
		"(Ljava/lang/String;)Ljava/lang/Class;", findLoadedClass0)
	native.RegisterNative(classname.ClassLoader, "findBootstrapClass",
		"(Ljava/lang/String;)Ljava/lang/Class;", findBootstrapClass)
}

//todo
var builtinLibs = map[string]bool{
	"zip": true,
}

//private static native String findBuiltinLib(String name);
func findBuiltinLib(frame *rtda.Frame) {
	lib := frame.LocalVars.GetRef(0)
	libName := rtda.GetGoString(lib)
	if ok := builtinLibs[libName]; ok {
		frame.OperandStack.PushRef(lib)
	} else {
		frame.OperandStack.PushRef(nil)
	}
}

//private native final Class<?> findLoadedClass0(String name)
func findLoadedClass0(frame *rtda.Frame) {
	cl := frame.Method().Class().ClassLoader()
	className := frame.LocalVars.GetRef(1)
	class := cl.FindLoadedClass(rtda.GetGoString(className))
	if class == nil {
		frame.OperandStack.PushRef(nil)
		return
	}
	frame.OperandStack.PushRef(class.JClass())
}

//private native Class<?> findBootstrapClass(String name)
func findBootstrapClass(frame *rtda.Frame) {
	frame.OperandStack.PushRef(nil)
}
