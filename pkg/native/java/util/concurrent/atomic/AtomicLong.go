package atomic

import (
	"github.com/rodbate/jvm-on-go/pkg/constants/classname"
	"github.com/rodbate/jvm-on-go/pkg/native"
	"github.com/rodbate/jvm-on-go/pkg/rtda"
)

func init() {
	native.RegisterNative(classname.AtomicLong, "VMSupportsCS8", ""+
		"()Z", VMSupportsCS8)
}

//private static native boolean VMSupportsCS8()
func VMSupportsCS8(frame *rtda.Frame) {
	frame.OperandStack.PushBoolean(false)
}
