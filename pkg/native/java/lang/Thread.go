package lang

import (
	"github.com/rodbate/jvm-on-go/pkg/constants/classname"
	"github.com/rodbate/jvm-on-go/pkg/native"
	"github.com/rodbate/jvm-on-go/pkg/rtda"
)

func init() {
	native.RegisterNative(classname.Thread, "currentThread",
		"()Ljava/lang/Thread;", currentThread)
	native.RegisterNative(classname.Thread, "setPriority0",
		"(I)V", setPriority0)
	native.RegisterNative(classname.Thread, "isAlive",
		"()Z", isAlive)
	native.RegisterNative(classname.Thread, "start0",
		"()V", start0)
}

//public static native Thread currentThread();
func currentThread(frame *rtda.Frame) {
	cl := frame.Method().Class().ClassLoader()
	thread := cl.LoadClass(classname.Thread).NewInstance()
	threadGroup := cl.LoadClass(classname.ThreadGroup).NewInstance()
	thread.SetFieldValue("group", "Ljava/lang/ThreadGroup;", threadGroup)
	thread.SetFieldValue("priority", "I", int32(5))
	frame.OperandStack.PushRef(thread)
}

//private native void setPriority0(int newPriority);
func setPriority0(frame *rtda.Frame) {
	//noop
}

//public final native boolean isAlive();
func isAlive(frame *rtda.Frame) {
	frame.OperandStack.PushInt(0)
}

//private native void start0();
func start0(frame *rtda.Frame) {

}
