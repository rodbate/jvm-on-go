package io

import (
	"jvm-on-go/constants/classname"
	"jvm-on-go/native"
)

func init() {
	native.RegisterNative(classname.FileDescriptor, "initIDs", "()V", initIDs)
}
