package io

import (
	"github.com/rodbate/jvm-on-go/constants/classname"
	"github.com/rodbate/jvm-on-go/native"
)

func init() {
	native.RegisterNative(classname.FileDescriptor, "initIDs", "()V", initIDs)
}
