package reference

import (
	"github.com/rodbate/jvm-on-go/instructions/base"
	"github.com/rodbate/jvm-on-go/rtda"
)

func AALoad(reader *base.ByteCodeReader, frame *rtda.Frame) {
	doALoad(frame,
		func(obj *rtda.Object) interface{} {
			return obj.Refs()
		}, func(i interface{}) int {
			return len(i.(rtda.ArrayRef))
		},
		func(data interface{}, index int32) {
			frame.OperandStack.PushRef((data.(rtda.ArrayRef))[index])
		},
	)
}

func BALoad(reader *base.ByteCodeReader, frame *rtda.Frame) {
	doALoad(frame,
		func(obj *rtda.Object) interface{} {
			return obj.Bytes()
		}, func(i interface{}) int {
			return len(i.(rtda.ArrayByte))
		},
		func(data interface{}, index int32) {
			frame.OperandStack.PushInt(int32((data.(rtda.ArrayByte))[index]))
		},
	)
}

func CALoad(reader *base.ByteCodeReader, frame *rtda.Frame) {
	doALoad(frame,
		func(obj *rtda.Object) interface{} {
			return obj.Chars()
		}, func(i interface{}) int {
			return len(i.(rtda.ArrayChar))
		},
		func(data interface{}, index int32) {
			frame.OperandStack.PushInt(int32((data.(rtda.ArrayChar))[index]))
		},
	)
}

func SALoad(reader *base.ByteCodeReader, frame *rtda.Frame) {
	doALoad(frame,
		func(obj *rtda.Object) interface{} {
			return obj.Shorts()
		}, func(i interface{}) int {
			return len(i.(rtda.ArrayShort))
		},
		func(data interface{}, index int32) {
			frame.OperandStack.PushInt(int32((data.(rtda.ArrayShort))[index]))
		},
	)
}

func IALoad(reader *base.ByteCodeReader, frame *rtda.Frame) {
	doALoad(frame,
		func(obj *rtda.Object) interface{} {
			return obj.Ints()
		}, func(i interface{}) int {
			return len(i.(rtda.ArrayInt))
		},
		func(data interface{}, index int32) {
			frame.OperandStack.PushInt((data.(rtda.ArrayInt))[index])
		},
	)
}

func LALoad(reader *base.ByteCodeReader, frame *rtda.Frame) {
	doALoad(frame,
		func(obj *rtda.Object) interface{} {
			return obj.Longs()
		}, func(i interface{}) int {
			return len(i.(rtda.ArrayLong))
		},
		func(data interface{}, index int32) {
			frame.OperandStack.PushLong((data.(rtda.ArrayLong))[index])
		},
	)
}

func FALoad(reader *base.ByteCodeReader, frame *rtda.Frame) {
	doALoad(frame,
		func(obj *rtda.Object) interface{} {
			return obj.Floats()
		}, func(i interface{}) int {
			return len(i.(rtda.ArrayFloat))
		},
		func(data interface{}, index int32) {
			frame.OperandStack.PushFloat((data.(rtda.ArrayFloat))[index])
		},
	)
}

func DALoad(reader *base.ByteCodeReader, frame *rtda.Frame) {
	doALoad(frame,
		func(obj *rtda.Object) interface{} {
			return obj.Doubles()
		}, func(i interface{}) int {
			return len(i.(rtda.ArrayDouble))
		},
		func(data interface{}, index int32) {
			frame.OperandStack.PushDouble((data.(rtda.ArrayDouble))[index])
		},
	)
}

func doALoad(frame *rtda.Frame, dataGetter func(*rtda.Object) interface{},
	arrayLenGetter func(interface{}) int, framePusher func(interface{}, int32)) {
	index := frame.OperandStack.PopInt()
	arrayRef := frame.OperandStack.PopRef()
	if arrayRef == nil {
		panic("java.lang.NullPointerException")
	}
	data := dataGetter(arrayRef)
	if index < 0 || index >= int32(arrayLenGetter(data)) {
		panic("ArrayIndexOutOfBoundsException")
	}
	framePusher(data, index)
}
