package reserved

import (
	"fmt"
	"github.com/rodbate/jvm-on-go/pkg/instructions/base"
	"github.com/rodbate/jvm-on-go/pkg/native"
	"github.com/rodbate/jvm-on-go/pkg/rtda"
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
