package io

import (
	"github.com/rodbate/jvm-on-go/pkg/constants/classname"
	"github.com/rodbate/jvm-on-go/pkg/native"
	"github.com/rodbate/jvm-on-go/pkg/rtda"
)

func init() {
	native.RegisterNative(classname.FileInputStream, "initIDs", "()V", initIDs)
}

//private static native void initIDs()
func initIDs(frame *rtda.Frame) {
	//noop
}
