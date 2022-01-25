package native

import (
	"github.com/rodbate/jvm-on-go/pkg/rtda"
)

type JNativeMethod func(frame *rtda.Frame)

var EmptyNativeMethod = func(frame *rtda.Frame) {}

var registry = map[string]JNativeMethod{}

func RegisterNative(className, methodName, methodDescriptor string, nativeMethod JNativeMethod) {
	registry[buildKey(className, methodName, methodDescriptor)] = nativeMethod
}

func GetNativeMethod(className, methodName, methodDescriptor string) JNativeMethod {
	if m, ok := registry[buildKey(className, methodName, methodDescriptor)]; ok {
		return m
	}
	if methodName == "registerNatives" && methodDescriptor == "()V" {
		return EmptyNativeMethod
	}
	return nil
}

func buildKey(className, methodName, methodDescriptor string) string {
	return className + "-" + methodName + "-" + methodDescriptor
}
