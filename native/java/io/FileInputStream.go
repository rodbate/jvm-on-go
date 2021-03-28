package io

import (
	"github.com/rodbate/jvm-on-go/constants/classname"
	"github.com/rodbate/jvm-on-go/native"
	"github.com/rodbate/jvm-on-go/rtda"
)

func init() {
	native.RegisterNative(classname.FileInputStream, "initIDs", "()V", initIDs)
}

//private static native void initIDs()
func initIDs(frame *rtda.Frame) {
	//noop
}
