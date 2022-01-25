package misc

import (
	"github.com/rodbate/jvm-on-go/pkg/constants/classname"
	"github.com/rodbate/jvm-on-go/pkg/native"
	"github.com/rodbate/jvm-on-go/pkg/rtda"
)

func init() {
	native.RegisterNative(classname.Signal, "findSignal",
		"(Ljava/lang/String;)I", findSignal)
	native.RegisterNative(classname.Signal, "handle0",
		"(IJ)J", handle0)
}

// private static native int findSignal(String string);
func findSignal(frame *rtda.Frame) {
	//todo signal
	frame.OperandStack.PushInt(0)
}

// private static native long handle0(int i, long l);
func handle0(frame *rtda.Frame) {
	// todo
	frame.OperandStack.PushLong(0)
}
