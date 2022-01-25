package lang

import (
	"github.com/rodbate/jvm-on-go/pkg/constants/classname"
	"github.com/rodbate/jvm-on-go/pkg/native"
	rtda2 "github.com/rodbate/jvm-on-go/pkg/rtda"
)

func init() {
	native.RegisterNative(classname.String, "intern",
		"()Ljava/lang/String;", intern)
}

/**
public native String intern()
*/
func intern(frame *rtda2.Frame) {
	this := frame.LocalVars.GetThis()
	frame.OperandStack.PushRef(rtda2.InternString(this))
}
