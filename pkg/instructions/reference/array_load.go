package reference

import (
	"github.com/rodbate/jvm-on-go/pkg/instructions/base"
	rtda2 "github.com/rodbate/jvm-on-go/pkg/rtda"
)

func AALoad(reader *base.ByteCodeReader, frame *rtda2.Frame) {
	doALoad(frame,
		func(obj *rtda2.Object) interface{} {
			return obj.Refs()
		}, func(i interface{}) int {
			return len(i.(rtda2.ArrayRef))
		},
		func(data interface{}, index int32) {
			frame.OperandStack.PushRef((data.(rtda2.ArrayRef))[index])
		},
	)
}

func BALoad(reader *base.ByteCodeReader, frame *rtda2.Frame) {
	doALoad(frame,
		func(obj *rtda2.Object) interface{} {
			return obj.Bytes()
		}, func(i interface{}) int {
			return len(i.(rtda2.ArrayByte))
		},
		func(data interface{}, index int32) {
			frame.OperandStack.PushInt(int32((data.(rtda2.ArrayByte))[index]))
		},
	)
}

func CALoad(reader *base.ByteCodeReader, frame *rtda2.Frame) {
	doALoad(frame,
		func(obj *rtda2.Object) interface{} {
			return obj.Chars()
		}, func(i interface{}) int {
			return len(i.(rtda2.ArrayChar))
		},
		func(data interface{}, index int32) {
			frame.OperandStack.PushInt(int32((data.(rtda2.ArrayChar))[index]))
		},
	)
}

func SALoad(reader *base.ByteCodeReader, frame *rtda2.Frame) {
	doALoad(frame,
		func(obj *rtda2.Object) interface{} {
			return obj.Shorts()
		}, func(i interface{}) int {
			return len(i.(rtda2.ArrayShort))
		},
		func(data interface{}, index int32) {
			frame.OperandStack.PushInt(int32((data.(rtda2.ArrayShort))[index]))
		},
	)
}

func IALoad(reader *base.ByteCodeReader, frame *rtda2.Frame) {
	doALoad(frame,
		func(obj *rtda2.Object) interface{} {
			return obj.Ints()
		}, func(i interface{}) int {
			return len(i.(rtda2.ArrayInt))
		},
		func(data interface{}, index int32) {
			frame.OperandStack.PushInt((data.(rtda2.ArrayInt))[index])
		},
	)
}

func LALoad(reader *base.ByteCodeReader, frame *rtda2.Frame) {
	doALoad(frame,
		func(obj *rtda2.Object) interface{} {
			return obj.Longs()
		}, func(i interface{}) int {
			return len(i.(rtda2.ArrayLong))
		},
		func(data interface{}, index int32) {
			frame.OperandStack.PushLong((data.(rtda2.ArrayLong))[index])
		},
	)
}

func FALoad(reader *base.ByteCodeReader, frame *rtda2.Frame) {
	doALoad(frame,
		func(obj *rtda2.Object) interface{} {
			return obj.Floats()
		}, func(i interface{}) int {
			return len(i.(rtda2.ArrayFloat))
		},
		func(data interface{}, index int32) {
			frame.OperandStack.PushFloat((data.(rtda2.ArrayFloat))[index])
		},
	)
}

func DALoad(reader *base.ByteCodeReader, frame *rtda2.Frame) {
	doALoad(frame,
		func(obj *rtda2.Object) interface{} {
			return obj.Doubles()
		}, func(i interface{}) int {
			return len(i.(rtda2.ArrayDouble))
		},
		func(data interface{}, index int32) {
			frame.OperandStack.PushDouble((data.(rtda2.ArrayDouble))[index])
		},
	)
}

func doALoad(frame *rtda2.Frame, dataGetter func(*rtda2.Object) interface{},
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
