package lang

import (
	"jvm-on-go/constants/classname"
	"jvm-on-go/native"
	"jvm-on-go/rtda"
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
