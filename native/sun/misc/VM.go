package misc

import (
	"jvm-on-go/constants/classname"
	"jvm-on-go/instructions/base"
	"jvm-on-go/native"
	"jvm-on-go/rtda"
)

func init() {
	native.RegisterNative(classname.VM, "initialize", "()V", initialize)
}

/**
private native static void initialize()
*/
func initialize(frame *rtda.Frame) {
	classloader := frame.Method().Class().ClassLoader()
	systemClass := classloader.LoadClass(classname.System)
	initializeMethod := systemClass.GetStaticMethod("initializeSystemClass", "()V")
	base.InvokeMethod(frame, initializeMethod)
}
