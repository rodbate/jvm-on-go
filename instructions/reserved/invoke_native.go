package reserved

import (
	"fmt"
	"jvm-on-go/instructions/base"
	"jvm-on-go/native"
	_ "jvm-on-go/native/java/io"
	_ "jvm-on-go/native/java/lang"
	_ "jvm-on-go/native/java/security"
	_ "jvm-on-go/native/java/util/concurrent/atomic"
	_ "jvm-on-go/native/sun/misc"
	_ "jvm-on-go/native/sun/reflect"
	"jvm-on-go/rtda"
)

func InvokeNative(reader *base.ByteCodeReader, frame *rtda.Frame) {
	method := frame.Method()
	className := method.Class().Name()
	methodName := method.Name()
	methodDescriptor := method.Descriptor()

	nativeMethod := native.GetNativeMethod(className, methodName, methodDescriptor)
	if nativeMethod == nil {
		panic(fmt.Sprintf("java.lang.UnsatisifiedLinkException: %v.%v%v",
			className, methodName, methodDescriptor))
	}
	nativeMethod(frame)
}
