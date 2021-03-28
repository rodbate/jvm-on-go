package rtda

import (
	"fmt"
	"github.com/rodbate/jvm-on-go/constants/descriptors"
)

type Object struct {
	class              *Class
	Data               interface{}
	Extra              interface{}
	StackTraceElements interface{}
}

func newObject(class *Class) *Object {
	return &Object{
		class: class,
		Data:  make(Slots, class.instanceSlotCount),
	}
}

func (obj *Object) Fields() Slots {
	return obj.Data.(Slots)
}

func (obj *Object) IsInstanceOf(class *Class) bool {
	if class == nil {
		return false
	}
	return class.IsAssignableFrom(obj.class)
}

func (obj *Object) Class() *Class {
	return obj.class
}

func (obj *Object) SetFieldValue(name, descriptor string, value interface{}) {
	f := obj.GetField(name, descriptor)
	if f == nil {
		panic(fmt.Sprintf("SetFieldValue -> not find field: %v %v of class: %v", name, descriptor,
			obj.class.name))
	}
	switch string(f.descriptor[0]) {
	case descriptors.Boolean, descriptors.Byte, descriptors.Char, descriptors.Short, descriptors.Int:
		obj.Data.(Slots).SetInt(f.slotIndex, value.(int32))
	case descriptors.Float:
		obj.Data.(Slots).SetFloat(f.slotIndex, value.(float32))
	case descriptors.Double:
		obj.Data.(Slots).SetDouble(f.slotIndex, value.(float64))
	case descriptors.Long:
		obj.Data.(Slots).SetLong(f.slotIndex, value.(int64))
	case "L", "[":
		obj.Data.(Slots).SetRef(f.slotIndex, value.(*Object))
	default:
		panic("unsupported field descriptor: " + f.descriptor)
	}
}

func (obj *Object) GetFieldValue(name, descriptor string) interface{} {
	f := obj.GetField(name, descriptor)
	if f == nil {
		panic(fmt.Sprintf("GetFieldValue -> not find field: %v %v of class: %v", name, descriptor,
			obj.class.name))
	}
	switch string(f.descriptor[0]) {
	case descriptors.Boolean, descriptors.Byte, descriptors.Char, descriptors.Short, descriptors.Int:
		return obj.Data.(Slots).GetInt(f.slotIndex)
	case descriptors.Float:
		return obj.Data.(Slots).GetFloat(f.slotIndex)
	case descriptors.Double:
		return obj.Data.(Slots).GetDouble(f.slotIndex)
	case descriptors.Long:
		return obj.Data.(Slots).GetLong(f.slotIndex)
	case "L", "[":
		return obj.Data.(Slots).GetRef(f.slotIndex)
	default:
		panic("unsupported field descriptor: " + f.descriptor)
	}
}

func (obj *Object) GetField(name, descriptor string) *Field {
	return obj.class.GetField(func(field *Field) bool {
		return field.Name() == name && field.descriptor == descriptor
	})
}

func (obj *Object) Clone() *Object {
	return &Object{
		class: obj.class,
		Data:  obj.cloneData(),
	}
}

func (obj *Object) cloneData() interface{} {
	switch obj.Data.(type) {
	case ArrayBoolean:
		srcData := obj.Data.(ArrayBoolean)
		clonedData := make(ArrayBoolean, len(srcData))
		copy(clonedData, srcData)
		return clonedData
	case ArrayByte:
		srcData := obj.Data.(ArrayByte)
		clonedData := make(ArrayByte, len(srcData))
		copy(clonedData, srcData)
		return clonedData
	case ArrayChar:
		srcData := obj.Data.(ArrayChar)
		clonedData := make(ArrayChar, len(srcData))
		copy(clonedData, srcData)
		return clonedData
	case ArrayShort:
		srcData := obj.Data.(ArrayShort)
		clonedData := make(ArrayShort, len(srcData))
		copy(clonedData, srcData)
		return clonedData
	case ArrayInt:
		srcData := obj.Data.(ArrayInt)
		clonedData := make(ArrayInt, len(srcData))
		copy(clonedData, srcData)
		return clonedData
	case ArrayLong:
		srcData := obj.Data.(ArrayLong)
		clonedData := make(ArrayLong, len(srcData))
		copy(clonedData, srcData)
		return clonedData
	case ArrayFloat:
		srcData := obj.Data.(ArrayFloat)
		clonedData := make(ArrayFloat, len(srcData))
		copy(clonedData, srcData)
		return clonedData
	case ArrayDouble:
		srcData := obj.Data.(ArrayDouble)
		clonedData := make(ArrayDouble, len(srcData))
		copy(clonedData, srcData)
		return clonedData
	case ArrayRef:
		srcData := obj.Data.(ArrayRef)
		clonedData := make(ArrayRef, len(srcData))
		copy(clonedData, srcData)
		return clonedData
	case Slots:
		srcData := obj.Data.(Slots)
		clonedData := make(Slots, len(srcData))
		copy(clonedData, srcData)
		return clonedData
	default:
		panic(fmt.Sprintf("invalid array type: %T", obj))
	}
}

//================== Array Object =====================//

func (obj *Object) Booleans() ArrayBoolean {
	return obj.Data.(ArrayBoolean)
}

func (obj *Object) Bytes() ArrayByte {
	return obj.Data.(ArrayByte)
}

func (obj *Object) Shorts() ArrayShort {
	return obj.Data.(ArrayShort)
}

func (obj *Object) Chars() ArrayChar {
	return obj.Data.(ArrayChar)
}

func (obj *Object) Ints() ArrayInt {
	return obj.Data.(ArrayInt)
}

func (obj *Object) Longs() ArrayLong {
	return obj.Data.(ArrayLong)
}

func (obj *Object) Floats() ArrayFloat {
	return obj.Data.(ArrayFloat)
}

func (obj *Object) Doubles() ArrayDouble {
	return obj.Data.(ArrayDouble)
}

func (obj *Object) Refs() ArrayRef {
	return obj.Data.(ArrayRef)
}

func (obj *Object) ArrayLength() int32 {
	switch obj.Data.(type) {
	case ArrayBoolean:
		return int32(len(obj.Data.(ArrayBoolean)))
	case ArrayByte:
		return int32(len(obj.Data.(ArrayByte)))
	case ArrayChar:
		return int32(len(obj.Data.(ArrayChar)))
	case ArrayInt:
		return int32(len(obj.Data.(ArrayInt)))
	case ArrayLong:
		return int32(len(obj.Data.(ArrayLong)))
	case ArrayFloat:
		return int32(len(obj.Data.(ArrayFloat)))
	case ArrayDouble:
		return int32(len(obj.Data.(ArrayDouble)))
	case ArrayRef:
		return int32(len(obj.Data.(ArrayRef)))
	default:
		panic(fmt.Sprintf("ArrayLength not support types: %T", obj))
	}
}

func ArrayCopy(src, dest *Object, srcPos, destPos, length int32) {
	switch src.Data.(type) {
	case ArrayBoolean:
		srcSlice := src.Data.(ArrayBoolean)[srcPos : srcPos+length]
		destSlice := dest.Data.(ArrayBoolean)[destPos : destPos+length]
		copy(destSlice, srcSlice)
	case ArrayByte:
		srcSlice := src.Data.(ArrayByte)[srcPos : srcPos+length]
		destSlice := dest.Data.(ArrayByte)[destPos : destPos+length]
		copy(destSlice, srcSlice)
	case ArrayChar:
		srcSlice := src.Data.(ArrayChar)[srcPos : srcPos+length]
		destSlice := dest.Data.(ArrayChar)[destPos : destPos+length]
		copy(destSlice, srcSlice)
	case ArrayShort:
		srcSlice := src.Data.(ArrayShort)[srcPos : srcPos+length]
		destSlice := dest.Data.(ArrayShort)[destPos : destPos+length]
		copy(destSlice, srcSlice)
	case ArrayInt:
		srcSlice := src.Data.(ArrayInt)[srcPos : srcPos+length]
		destSlice := dest.Data.(ArrayInt)[destPos : destPos+length]
		copy(destSlice, srcSlice)
	case ArrayLong:
		srcSlice := src.Data.(ArrayLong)[srcPos : srcPos+length]
		destSlice := dest.Data.(ArrayLong)[destPos : destPos+length]
		copy(destSlice, srcSlice)
	case ArrayFloat:
		srcSlice := src.Data.(ArrayFloat)[srcPos : srcPos+length]
		destSlice := dest.Data.(ArrayFloat)[destPos : destPos+length]
		copy(destSlice, srcSlice)
	case ArrayDouble:
		srcSlice := src.Data.(ArrayDouble)[srcPos : srcPos+length]
		destSlice := dest.Data.(ArrayDouble)[destPos : destPos+length]
		copy(destSlice, srcSlice)
	case ArrayRef:
		srcSlice := src.Data.(ArrayRef)[srcPos : srcPos+length]
		destSlice := dest.Data.(ArrayRef)[destPos : destPos+length]
		copy(destSlice, srcSlice)
	default:
		panic(fmt.Sprintf("invalid array type: %T", src))
	}
}
