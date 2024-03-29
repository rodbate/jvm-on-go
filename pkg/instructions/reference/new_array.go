package reference

import (
	"fmt"
	"github.com/rodbate/jvm-on-go/pkg/constants/descriptors"
	"github.com/rodbate/jvm-on-go/pkg/instructions/base"
	rtda2 "github.com/rodbate/jvm-on-go/pkg/rtda"
)

const (
	TBoolean = 4
	TChar    = 5
	TFloat   = 6
	TDouble  = 7
	TByte    = 8
	TShort   = 9
	TInt     = 10
	TLong    = 11
)

func NewArray(reader *base.ByteCodeReader, frame *rtda2.Frame) {
	aType := reader.ReadUint8()
	count := frame.OperandStack.PopInt()
	if count < 0 {
		panic("java.lang.NegativeArraySizeException")
	}
	classloader := frame.Method().Class().ClassLoader()
	class := loadPrimitiveArrayClass(classloader, aType)
	frame.OperandStack.PushRef(class.NewArray(uint(count)))
}

func loadPrimitiveArrayClass(loader *rtda2.ClassLoader, aType uint8) *rtda2.Class {
	switch aType {
	case TBoolean:
		return loader.LoadClass(descriptors.ArrayBoolean)
	case TChar:
		return loader.LoadClass(descriptors.ArrayChar)
	case TFloat:
		return loader.LoadClass(descriptors.ArrayFloat)
	case TDouble:
		return loader.LoadClass(descriptors.ArrayDouble)
	case TByte:
		return loader.LoadClass(descriptors.ArrayByte)
	case TShort:
		return loader.LoadClass(descriptors.ArrayShort)
	case TInt:
		return loader.LoadClass(descriptors.ArrayInt)
	case TLong:
		return loader.LoadClass(descriptors.ArrayLong)
	default:
		panic(fmt.Sprintf("NewArray -> invalid aType: %v", aType))
	}
}
