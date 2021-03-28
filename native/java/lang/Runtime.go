package lang

import (
	"jvm-on-go/constants/classname"
	"jvm-on-go/native"
	"jvm-on-go/rtda"
	"runtime"
)

func init() {
	native.RegisterNative(classname.Runtime, "availableProcessors",
		"()I", availableProcessors)
}

//public native int availableProcessors()
func availableProcessors(frame *rtda.Frame) {
	num := runtime.NumCPU()
	frame.OperandStack.PushInt(int32(num))
}
