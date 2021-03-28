package lang

import (
	"github.com/rodbate/jvm-on-go/constants/classname"
	"github.com/rodbate/jvm-on-go/native"
	"github.com/rodbate/jvm-on-go/rtda"
	"math"
)

func init() {
	native.RegisterNative(classname.Float, "floatToRawIntBits",
		"(F)I", floatToRawIntBits)
}

/**
public static native int floatToRawIntBits(float value);
*/
func floatToRawIntBits(frame *rtda.Frame) {
	float := frame.LocalVars.GetFloat(0)
	frame.OperandStack.PushInt(int32(math.Float32bits(float)))
}
