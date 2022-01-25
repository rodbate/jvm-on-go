package descriptors

import "fmt"

const (
	Boolean   = "Z"
	Byte      = "B"
	Short     = "S"
	Char      = "C"
	Int       = "I"
	Float     = "F"
	Double    = "D"
	Long      = "J"
	Void      = "V"
	RefPrefix = "L"
	String    = "Ljava/lang/String;"

	ArrayPrefix  = "["
	ArrayBoolean = "[Z"
	ArrayByte    = "[B"
	ArrayShort   = "[S"
	ArrayChar    = "[C"
	ArrayInt     = "[I"
	ArrayFloat   = "[F"
	ArrayDouble  = "[D"
	ArrayLong    = "[J"
	ArrayString  = "[Ljava/lang/String;"
)

var PrimitiveTypeToArrayDescriptor = map[string]string{
	"void":    Void,
	"boolean": Boolean,
	"byte":    Byte,
	"short":   Short,
	"int":     Int,
	"long":    Long,
	"char":    Char,
	"float":   Float,
	"double":  Double,
}

func ConvertToDescriptorFrom(className string) string {
	if className[0] == '[' {
		return className
	}
	if d, ok := PrimitiveTypeToArrayDescriptor[className]; ok {
		return d
	}
	return RefPrefix + className + ";"
}

func ConvertToClassNameFrom(descriptor string) string {
	if descriptor[0] == '[' {
		return descriptor
	}
	if descriptor[0] == 'L' {
		return descriptor[1 : len(descriptor)-1]
	}
	for k, v := range PrimitiveTypeToArrayDescriptor {
		if v == descriptor {
			return k
		}
	}
	panic(fmt.Sprintf("ConvertToClassNameFrom -> invalid descriptor: %v", descriptor))
}
