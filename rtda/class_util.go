package rtda

import (
	"fmt"
	"github.com/rodbate/jvm-on-go/constants/descriptors"
)

func getArrayClassName(componentClassName string) string {
	return descriptors.ArrayPrefix + descriptors.ConvertToDescriptorFrom(componentClassName)
}

func getComponentClassName(arrayClassName string) string {
	if arrayClassName[0] != '[' {
		panic(fmt.Sprintf("getComponentClassName -> not array class: %v", arrayClassName))
	}
	return descriptors.ConvertToClassNameFrom(arrayClassName[1:])
}

func ToJavaClasses(classloader *ClassLoader, classNames []string, isDescriptor bool) *Object {
	classLen := uint(len(classNames))
	javaClasses := classloader.LoadClass("[Ljava/lang/Class;").NewArray(classLen)
	if classLen > 0 {
		for i, name := range classNames {
			if isDescriptor {
				name = descriptors.ConvertToClassNameFrom(name)
			}
			javaClasses.Refs()[i] = classloader.LoadClass(name).JClass()
		}
	}
	return javaClasses
}

func ToMethodDescriptor(args *Object) string {
	if args == nil || len(args.Refs()) == 0 {
		return "()V"
	}
	var desc = ""
	for _, arg := range args.Refs() {
		desc = desc + descriptors.ConvertToDescriptorFrom(arg.class.name)
	}
	return "(" + desc + ")V"
}
