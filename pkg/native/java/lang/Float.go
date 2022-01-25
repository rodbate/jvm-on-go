package lang

import (
	"github.com/rodbate/jvm-on-go/pkg/constants/classname"
	"github.com/rodbate/jvm-on-go/pkg/native"
	"github.com/rodbate/jvm-on-go/pkg/rtda"
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
