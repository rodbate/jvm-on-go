package io

import (
	"github.com/rodbate/jvm-on-go/pkg/constants/classname"
	"github.com/rodbate/jvm-on-go/pkg/native"
	rtda2 "github.com/rodbate/jvm-on-go/pkg/rtda"
	"os"
	"path/filepath"
)

func init() {
	native.RegisterNative(classname.UnixFileSystem, "initIDs", "()V", initIDs)
	native.RegisterNative(classname.UnixFileSystem, "canonicalize0", "(Ljava/lang/String;)Ljava/lang/String;", canonicalize0)
	native.RegisterNative(classname.UnixFileSystem, "getBooleanAttributes0", "(Ljava/io/File;)I", getBooleanAttributes0)
}

// private native String canonicalize0(String path) throws IOException;
func canonicalize0(frame *rtda2.Frame) {
	vars := frame.LocalVars
	path := vars.GetRef(1)

	goPath := rtda2.GetGoString(path)
	goPath2 := filepath.Clean(goPath)
	if goPath2 != goPath {
		path = rtda2.GetJString(frame.Method().Class().ClassLoader(), goPath2)
	}

	frame.OperandStack.PushRef(path)
}

// public native int getBooleanAttributes0(File f);
func getBooleanAttributes0(frame *rtda2.Frame) {
	f := frame.LocalVars.GetRef(1)
	path := getPath(f)

	attributes0 := 0
	if exists(path) {
		attributes0 |= 0x01
	}
	if isDir(path) {
		attributes0 |= 0x04
	}

	frame.OperandStack.PushInt(int32(attributes0))
}

func getPath(fileObj *rtda2.Object) string {
	pathStr := fileObj.GetFieldValue("path", "Ljava/lang/String;")
	return rtda2.GetGoString(pathStr.(*rtda2.Object))
}

func exists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}

func isDir(path string) bool {
	fileInfo, err := os.Stat(path)
	if err == nil {
		return fileInfo.IsDir()
	}
	return false
}
