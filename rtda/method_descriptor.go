package rtda

import (
	"github.com/rodbate/jvm-on-go/constants/descriptors"
	"strings"
)

type MethodDescriptor struct {
	ParameterTypes []string
	ReturnType     string
}

/**
(IDLjava/lang/String;[J)V
*/
func ParseMethodDescriptor(method *Method) *MethodDescriptor {
	md := &MethodDescriptor{}
	descriptor := method.descriptor
	index := strings.Index(descriptor, ")")
	md.ReturnType = descriptor[index+1:]
	descriptor = descriptor[1:index]

	var parameterTypes []string
	var prefix = ""
	for i := 0; i < len(descriptor); i++ {
		d := string(descriptor[i])
		switch d {
		case descriptors.Boolean, descriptors.Byte, descriptors.Short, descriptors.Char,
			descriptors.Int, descriptors.Float, descriptors.Double, descriptors.Long:
			parameterTypes = appendParamType(parameterTypes, d, prefix)
			prefix = ""
		case "L":
			index := strings.Index(descriptor[i:], ";")
			var des string
			if index == len(descriptor)-i-1 {
				des = descriptor[i:]
			} else {
				des = descriptor[i : index+i+1]
			}
			parameterTypes = appendParamType(parameterTypes, des, prefix)
			prefix = ""
			i = index + i
		case "[":
			prefix += d
		default:
			panic("Error: parse method descriptor: " + method.descriptor)
		}
	}

	md.ParameterTypes = parameterTypes
	return md
}

func appendParamType(types []string, paramType string, prefix string) []string {
	if prefix != "" {
		paramType = prefix + paramType
	}
	return append(types, paramType)
}
