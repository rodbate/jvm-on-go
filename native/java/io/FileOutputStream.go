package io

import (
	"github.com/rodbate/jvm-on-go/constants/classname"
	"github.com/rodbate/jvm-on-go/native"
	"github.com/rodbate/jvm-on-go/rtda"
	"os"
	"unsafe"
)

func init() {
	native.RegisterNative(classname.FileOutputStream, "writeBytes",
		"([BIIZ)V", writeBytes)
	native.RegisterNative(classname.FileOutputStream, "initIDs",
		"()V", initIDsOutputStream)
}

//private native void writeBytes(byte b[], int off, int len, boolean append) throws IOException;
func writeBytes(frame *rtda.Frame) {
	bytes := frame.LocalVars.GetRef(1).Bytes()
	off := frame.LocalVars.GetInt(2)
	length := frame.LocalVars.GetInt(3)

	bytesPtr := unsafe.Pointer(&bytes)
	_, err := os.Stdout.Write((*(*[]byte)(bytesPtr))[off : off+length])
	if err != nil {
		panic(err)
	}
}

//private static native void initIDs()
func initIDsOutputStream(frame *rtda.Frame) {
	//noop
}
