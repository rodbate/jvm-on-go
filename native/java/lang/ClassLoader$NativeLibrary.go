package lang

import (
	"jvm-on-go/native"
	"jvm-on-go/rtda"
)

func init() {
	native.RegisterNative("java/lang/ClassLoader$NativeLibrary", "load",
		"(Ljava/lang/String;Z)V", load)
}

//native void load(String name, boolean isBuiltin);
func load(frame *rtda.Frame) {
	this := frame.LocalVars.GetThis()
	this.SetFieldValue("loaded", "Z", int32(1))
}
