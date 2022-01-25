package lang

import (
	"github.com/rodbate/jvm-on-go/pkg/constants/classname"
	"github.com/rodbate/jvm-on-go/pkg/native"
	"github.com/rodbate/jvm-on-go/pkg/rtda"
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
