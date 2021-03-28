package io

import (
	"github.com/rodbate/jvm-on-go/constants/classname"
	"github.com/rodbate/jvm-on-go/native"
	"github.com/rodbate/jvm-on-go/rtda"
	"os"
	"path/filepath"
)

func init() {
	native.RegisterNative(classname.UnixFileSystem, "initIDs", "()V", initIDs)
	native.RegisterNative(classname.UnixFileSystem, "canonicalize0", "(Ljava/lang/String;)Ljava/lang/String;", canonicalize0)
	native.RegisterNative(classname.UnixFileSystem, "getBooleanAttributes0", "(Ljava/io/File;)I", getBooleanAttributes0)
}

// private native String canonicalize0(String path) throws IOException;
func canonicalize0(frame *rtda.Frame) {
	vars := frame.LocalVars
	path := vars.GetRef(1)

	goPath := rtda.GetGoString(path)
	goPath2 := filepath.Clean(goPath)
	if goPath2 != goPath {
		path = rtda.GetJString(frame.Method().Class().ClassLoader(), goPath2)
	}

	frame.OperandStack.PushRef(path)
}

// public native int getBooleanAttributes0(File f);
func getBooleanAttributes0(frame *rtda.Frame) {
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

func getPath(fileObj *rtda.Object) string {
	pathStr := fileObj.GetFieldValue("path", "Ljava/lang/String;")
	return rtda.GetGoString(pathStr.(*rtda.Object))
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
