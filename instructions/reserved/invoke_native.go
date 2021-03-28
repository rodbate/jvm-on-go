package reserved

import (
	"fmt"
	"github.com/rodbate/jvm-on-go/instructions/base"
	"github.com/rodbate/jvm-on-go/native"
	_ "github.com/rodbate/jvm-on-go/native/java/io"
	_ "github.com/rodbate/jvm-on-go/native/java/lang"
	_ "github.com/rodbate/jvm-on-go/native/java/security"
	_ "github.com/rodbate/jvm-on-go/native/java/util/concurrent/atomic"
	_ "github.com/rodbate/jvm-on-go/native/sun/misc"
	_ "github.com/rodbate/jvm-on-go/native/sun/reflect"
	"github.com/rodbate/jvm-on-go/rtda"
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
