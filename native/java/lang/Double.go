package lang

import (
	"github.com/rodbate/jvm-on-go/constants/classname"
	"github.com/rodbate/jvm-on-go/native"
	"github.com/rodbate/jvm-on-go/rtda"
	"math"
)

func init() {
	native.RegisterNative(classname.Double, "doubleToRawLongBits",
		"(D)J", doubleToRawLongBits)
	native.RegisterNative(classname.Double, "longBitsToDouble",
		"(J)D", longBitsToDouble)
}

/**
public static native long doubleToRawLongBits(double value);
*/
func doubleToRawLongBits(frame *rtda.Frame) {
	double := frame.LocalVars.GetDouble(0)
	frame.OperandStack.PushLong(int64(math.Float64bits(double)))
}

/**
public static native double longBitsToDouble(long bits)
*/
func longBitsToDouble(frame *rtda.Frame) {
	long := frame.LocalVars.GetLong(0)
	frame.OperandStack.PushDouble(math.Float64frombits(uint64(long)))
}
