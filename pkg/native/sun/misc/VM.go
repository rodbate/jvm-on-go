package misc

import (
	"github.com/rodbate/jvm-on-go/pkg/constants/classname"
	"github.com/rodbate/jvm-on-go/pkg/instructions/base"
	"github.com/rodbate/jvm-on-go/pkg/native"
	"github.com/rodbate/jvm-on-go/pkg/rtda"
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
