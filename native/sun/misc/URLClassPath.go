package misc

import (
	"github.com/rodbate/jvm-on-go/constants/classname"
	"github.com/rodbate/jvm-on-go/native"
	"github.com/rodbate/jvm-on-go/rtda"
)

func init() {
	native.RegisterNative(classname.URLClassPath, "getLookupCacheURLs",
		"(Ljava/lang/ClassLoader;)[Ljava/net/URL;", getLookupCacheURLs)
}

//private static native URL[] getLookupCacheURLs(ClassLoader loader);
func getLookupCacheURLs(frame *rtda.Frame) {
	//todo: disable jvm look up cache
	frame.OperandStack.PushRef(nil)
}
