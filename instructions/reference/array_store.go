package reference

import (
	"jvm-on-go/constants/descriptors"
	"jvm-on-go/instructions/base"
	"jvm-on-go/rtda"
)

func AAStore(reader *base.ByteCodeReader, frame *rtda.Frame) {
	value := frame.OperandStack.PopRef()
	index := frame.OperandStack.PopInt()
	arrayRef := frame.OperandStack.PopRef()
	checkRefNotNull(arrayRef)
	data := arrayRef.Refs()
	checkArrayIndex(len(data), index)
	data[index] = value
}

func BAStore(reader *base.ByteCodeReader, frame *rtda.Frame) {
	value := frame.OperandStack.PopInt()
	index := frame.OperandStack.PopInt()
	arrayRef := frame.OperandStack.PopRef()
	checkRefNotNull(arrayRef)
	className := arrayRef.Class().Name()
	if className == descriptors.ArrayByte {
		data := arrayRef.Bytes()
		checkArrayIndex(len(data), index)
		data[index] = int8(value)
	} else if className == descriptors.ArrayBoolean {
		data := arrayRef.Booleans()
		checkArrayIndex(len(data), index)
		data[index] = int8(value)
	} else {
		panic("BAStore -> invalid array type: " + className)
	}
}

func CAStore(reader *base.ByteCodeReader, frame *rtda.Frame) {
	value := frame.OperandStack.PopInt()
	index := frame.OperandStack.PopInt()
	arrayRef := frame.OperandStack.PopRef()
	checkRefNotNull(arrayRef)
	data := arrayRef.Chars()
	checkArrayIndex(len(data), index)
	data[index] = uint16(value)
}

func SAStore(reader *base.ByteCodeReader, frame *rtda.Frame) {
	value := frame.OperandStack.PopInt()
	index := frame.OperandStack.PopInt()
	arrayRef := frame.OperandStack.PopRef()
	checkRefNotNull(arrayRef)
	data := arrayRef.Shorts()
	checkArrayIndex(len(data), index)
	data[index] = int16(value)
}

func IAStore(reader *base.ByteCodeReader, frame *rtda.Frame) {
	value := frame.OperandStack.PopInt()
	index := frame.OperandStack.PopInt()
	arrayRef := frame.OperandStack.PopRef()
	checkRefNotNull(arrayRef)
	data := arrayRef.Ints()
	checkArrayIndex(len(data), index)
	data[index] = value
}

func LAStore(reader *base.ByteCodeReader, frame *rtda.Frame) {
	value := frame.OperandStack.PopLong()
	index := frame.OperandStack.PopInt()
	arrayRef := frame.OperandStack.PopRef()
	checkRefNotNull(arrayRef)
	data := arrayRef.Longs()
	checkArrayIndex(len(data), index)
	data[index] = value
}

func FAStore(reader *base.ByteCodeReader, frame *rtda.Frame) {
	value := frame.OperandStack.PopFloat()
	index := frame.OperandStack.PopInt()
	arrayRef := frame.OperandStack.PopRef()
	checkRefNotNull(arrayRef)
	data := arrayRef.Floats()
	checkArrayIndex(len(data), index)
	data[index] = value
}

func DAStore(reader *base.ByteCodeReader, frame *rtda.Frame) {
	value := frame.OperandStack.PopDouble()
	index := frame.OperandStack.PopInt()
	arrayRef := frame.OperandStack.PopRef()
	checkRefNotNull(arrayRef)
	data := arrayRef.Doubles()
	checkArrayIndex(len(data), index)
	data[index] = value
}

func checkRefNotNull(ref *rtda.Object) {
	if ref == nil {
		panic("java.lang.NullPointerException")
	}
}

func checkArrayIndex(len int, index int32) {
	if index < 0 || index >= int32(len) {
		panic("ArrayIndexOutOfBoundsException")
	}
}
