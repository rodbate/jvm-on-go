package io

import (
	"jvm-on-go/constants/classname"
	"jvm-on-go/native"
	"jvm-on-go/rtda"
)

func init() {
	native.RegisterNative(classname.FileInputStream, "initIDs", "()V", initIDs)
}

//private static native void initIDs()
func initIDs(frame *rtda.Frame) {
	//noop
}
