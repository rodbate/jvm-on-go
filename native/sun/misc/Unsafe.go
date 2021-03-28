package misc

import (
	"encoding/binary"
	"jvm-on-go/constants/classname"
	"jvm-on-go/native"
	"jvm-on-go/rtda"
	"unsafe"
)

func init() {
	native.RegisterNative(classname.Unsafe, "arrayBaseOffset",
		"(Ljava/lang/Class;)I", arrayBaseOffset)
	native.RegisterNative(classname.Unsafe, "arrayIndexScale",
		"(Ljava/lang/Class;)I", arrayIndexScale)
	native.RegisterNative(classname.Unsafe, "addressSize",
		"()I", addressSize)
	native.RegisterNative(classname.Unsafe, "objectFieldOffset",
		"(Ljava/lang/reflect/Field;)J", objectFieldOffset)
	native.RegisterNative(classname.Unsafe, "compareAndSwapObject",
		"(Ljava/lang/Object;JLjava/lang/Object;Ljava/lang/Object;)Z", compareAndSwapObject)
	native.RegisterNative(classname.Unsafe, "getIntVolatile",
		"(Ljava/lang/Object;J)I", getIntVolatile)
	native.RegisterNative(classname.Unsafe, "compareAndSwapInt",
		"(Ljava/lang/Object;JII)Z", compareAndSwapInt)
	native.RegisterNative(classname.Unsafe, "compareAndSwapLong",
		"(Ljava/lang/Object;JJJ)Z", compareAndSwapLong)
	native.RegisterNative(classname.Unsafe, "allocateMemory",
		"(J)J", allocateMemory)
	native.RegisterNative(classname.Unsafe, "putLong",
		"(JJ)V", putLong)
	native.RegisterNative(classname.Unsafe, "getByte",
		"(J)B", getByte)
	native.RegisterNative(classname.Unsafe, "freeMemory",
		"(J)V", freeMemory)
	native.RegisterNative(classname.Unsafe, "getObjectVolatile",
		"(Ljava/lang/Object;J)Ljava/lang/Object;", getObjectVolatile)
}

//public native int arrayBaseOffset(Class<?> arrayClass);
func arrayBaseOffset(frame *rtda.Frame) {
	frame.OperandStack.PushInt(0)
}

//public native int arrayIndexScale(Class<?> arrayClass);
func arrayIndexScale(frame *rtda.Frame) {
	frame.OperandStack.PushInt(1)
}

//public native int addressSize();
func addressSize(frame *rtda.Frame) {
	addressSize := unsafe.Sizeof(frame.LocalVars.GetThis())
	frame.OperandStack.PushInt(int32(addressSize))
}

//public native long objectFieldOffset(Field f);
func objectFieldOffset(frame *rtda.Frame) {
	field := frame.LocalVars.GetRef(1)
	slot := field.GetFieldValue("slot", "I").(int32)
	frame.OperandStack.PushLong(int64(slot))
}

/**
public final native boolean compareAndSwapObject(Object o, long offset,
												 Object expected,
												 Object x);
*/
func compareAndSwapObject(frame *rtda.Frame) {
	objData := frame.LocalVars.GetRef(1).Data
	offset := frame.LocalVars.GetLong(2)
	expectedVal := frame.LocalVars.GetRef(4)
	newVal := frame.LocalVars.GetRef(5)

	if data, ok := objData.(rtda.Slots); ok {
		//object
		val := data.GetRef(uint(offset))
		var cas bool
		if val == expectedVal {
			data.SetRef(uint(offset), newVal)
			cas = true
		} else {
			cas = false
		}
		frame.OperandStack.PushBoolean(cas)
	} else if data, ok := objData.(rtda.ArrayRef); ok {
		//array
		val := data[offset]
		var cas bool
		if val == expectedVal {
			data[offset] = newVal
			cas = true
		} else {
			cas = false
		}
		frame.OperandStack.PushBoolean(cas)
	} else {
		panic("Unsafe::compareAndSwapObject")
	}
}

//public native int getIntVolatile(Object o, long offset);
func getIntVolatile(frame *rtda.Frame) {
	obj := frame.LocalVars.GetRef(1)
	offset := frame.LocalVars.GetLong(2)
	if data, ok := obj.Data.(rtda.Slots); ok {
		val := data.GetInt(uint(offset))
		frame.OperandStack.PushInt(val)
	} else if data, ok := obj.Data.(rtda.ArrayInt); ok {
		frame.OperandStack.PushInt(data[offset])
	} else {
		panic("Unsafe::getIntVolatile")
	}
}

/**
public final native boolean compareAndSwapInt(Object o, long offset,
											  int expected,
											  int x);
*/
func compareAndSwapInt(frame *rtda.Frame) {
	objData := frame.LocalVars.GetRef(1).Data
	offset := frame.LocalVars.GetLong(2)
	expectedVal := frame.LocalVars.GetInt(4)
	newVal := frame.LocalVars.GetInt(5)

	if data, ok := objData.(rtda.Slots); ok {
		val := data.GetInt(uint(offset))
		var cas bool
		if val == expectedVal {
			data.SetInt(uint(offset), newVal)
			cas = true
		} else {
			cas = false
		}
		frame.OperandStack.PushBoolean(cas)
	} else if data, ok := objData.(rtda.ArrayInt); ok {
		val := data[offset]
		var cas bool
		if val == expectedVal {
			data[offset] = newVal
			cas = true
		} else {
			cas = false
		}
		frame.OperandStack.PushBoolean(cas)
	} else {
		panic("Unsafe::compareAndSwapInt")
	}
}

/**
public final native boolean compareAndSwapLong(Object o, long offset, long expected, long x);
*/
func compareAndSwapLong(frame *rtda.Frame) {
	objData := frame.LocalVars.GetRef(1).Data
	offset := frame.LocalVars.GetLong(2)
	expectedVal := frame.LocalVars.GetLong(4)
	newVal := frame.LocalVars.GetLong(6)

	if data, ok := objData.(rtda.Slots); ok {
		val := data.GetLong(uint(offset))
		var cas bool
		if val == expectedVal {
			data.SetLong(uint(offset), newVal)
			cas = true
		} else {
			cas = false
		}
		frame.OperandStack.PushBoolean(cas)
	} else if data, ok := objData.(rtda.ArrayLong); ok {
		val := data[offset]
		var cas bool
		if val == expectedVal {
			data[offset] = newVal
			cas = true
		} else {
			cas = false
		}
		frame.OperandStack.PushBoolean(cas)
	} else {
		panic("Unsafe::compareAndSwapLong")
	}
}

//public native long allocateMemory(long bytes)
func allocateMemory(frame *rtda.Frame) {
	bytesLen := frame.LocalVars.GetLong(1)
	frame.OperandStack.PushLong(allocate(bytesLen))
}

//public native void putLong(long address, long x)
func putLong(frame *rtda.Frame) {
	address := frame.LocalVars.GetLong(1)
	val := frame.LocalVars.GetLong(3)
	mem := memoryAt(address)
	binary.LittleEndian.PutUint64(mem, uint64(val))
}

//public native byte getByte(long address)
func getByte(frame *rtda.Frame) {
	address := frame.LocalVars.GetLong(1)
	mem := memoryAt(address)
	frame.OperandStack.PushInt(int32(mem[0]))
}

//public native void freeMemory(long address)
func freeMemory(frame *rtda.Frame) {
	address := frame.LocalVars.GetLong(1)
	free(address)
}

//public native Object getObjectVolatile(Object o, long offset);
func getObjectVolatile(frame *rtda.Frame) {
	obj := frame.LocalVars.GetRef(1)
	offset := frame.LocalVars.GetLong(2)
	if data, ok := obj.Data.(rtda.Slots); ok {
		val := data.GetRef(uint(offset))
		frame.OperandStack.PushRef(val)
	} else if data, ok := obj.Data.(rtda.ArrayRef); ok {
		frame.OperandStack.PushRef(data[offset])
	} else {
		panic("Unsafe::getObjectVolatile")
	}
}
