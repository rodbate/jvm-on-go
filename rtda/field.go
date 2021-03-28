package rtda

import (
	"jvm-on-go/classfile"
	"jvm-on-go/constants/descriptors"
)

type Field struct {
	Member
	constantValueIndex uint16
	slotIndex          uint
}

func newField(class *Class, fieldInfo *classfile.MemberInfo) *Field {
	field := &Field{}
	field.class = class
	field.initInfo(fieldInfo)
	field.initAttribute(fieldInfo)
	return field
}

func newFields(class *Class, fieldInfos []*classfile.MemberInfo) []*Field {
	fields := make([]*Field, len(fieldInfos))
	for i := range fieldInfos {
		fields[i] = newField(class, fieldInfos[i])
	}
	return fields
}

func (f *Field) initAttribute(fieldInfo *classfile.MemberInfo) {
	if attr := fieldInfo.ConstantValueAttribute(); attr != nil {
		f.constantValueIndex = attr.ConstantValueIndex()
	}
}

func (f *Field) IsVolatile() bool {
	return f.AccessFlags&AccVolatile != 0
}

func (f *Field) IsTransient() bool {
	return f.AccessFlags&AccTransient != 0
}

func (f *Field) IsEnum() bool {
	return f.AccessFlags&AccEnum != 0
}

func (f *Field) IsLongOrDouble() bool {
	return f.descriptor == descriptors.Long || f.descriptor == descriptors.Double
}

func (f *Field) Match(name string, descriptor string) bool {
	return f.name == name && f.descriptor == descriptor
}

func (f *Field) SlotIndex() uint {
	return f.slotIndex
}

func (f *Field) SetStaticValue(value interface{}) {
	switch string(f.descriptor[0]) {
	case descriptors.Boolean, descriptors.Byte, descriptors.Char, descriptors.Short, descriptors.Int:
		f.class.staticFields.SetInt(f.slotIndex, value.(int32))
	case descriptors.Float:
		f.class.staticFields.SetFloat(f.slotIndex, value.(float32))
	case descriptors.Double:
		f.class.staticFields.SetDouble(f.slotIndex, value.(float64))
	case descriptors.Long:
		f.class.staticFields.SetLong(f.slotIndex, value.(int64))
	case "L", "[":
		f.class.staticFields.SetRef(f.slotIndex, value.(*Object))
	default:
		panic("unsupported field descriptor: " + f.descriptor)
	}
}

func (f *Field) Type() *Class {
	className := descriptors.ConvertToClassNameFrom(f.descriptor)
	return f.class.classloader.LoadClass(className)
}
