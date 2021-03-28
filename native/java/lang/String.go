package lang

import (
	"github.com/rodbate/jvm-on-go/constants/classname"
	"github.com/rodbate/jvm-on-go/native"
	"github.com/rodbate/jvm-on-go/rtda"
)

func init() {
	native.RegisterNative(classname.String, "intern",
		"()Ljava/lang/String;", intern)
}

/**
public native String intern()
*/
func intern(frame *rtda.Frame) {
	this := frame.LocalVars.GetThis()
	frame.OperandStack.PushRef(rtda.InternString(this))
}
