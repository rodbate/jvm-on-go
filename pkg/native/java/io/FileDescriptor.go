package io

import (
	"github.com/rodbate/jvm-on-go/pkg/constants/classname"
	"github.com/rodbate/jvm-on-go/pkg/native"
)

func init() {
	native.RegisterNative(classname.FileDescriptor, "initIDs", "()V", initIDs)
}
