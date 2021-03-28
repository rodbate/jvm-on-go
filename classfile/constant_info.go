package classfile

import "fmt"

const (
	ConstantUtf8               = 1
	ConstantInteger            = 3
	ConstantFloat              = 4
	ConstantLong               = 5
	ConstantDouble             = 6
	ConstantClass              = 7
	ConstantString             = 8
	ConstantFieldRef           = 9
	ConstantMethodRef          = 10
	ConstantInterfaceMethodRef = 11
	ConstantNameAndType        = 12
	ConstantMethodHandle       = 15
	ConstantMethodType         = 16
	ConstantInvokeDynamic      = 18
)

type ConstantInfo interface {
	readInfo(cr *ClassReader)
}

func newConstantInfo(tag uint8, cp ConstantPool) ConstantInfo {
	switch tag {
	case ConstantUtf8:
		return &ConstantUtf8Info{}
	case ConstantInteger:
		return &ConstantIntegerInfo{}
	case ConstantFloat:
		return &ConstantFloatInfo{}
	case ConstantLong:
		return &ConstantLongInfo{}
	case ConstantDouble:
		return &ConstantDoubleInfo{}
	case ConstantClass:
		return &ConstantClassInfo{cp: cp}
	case ConstantString:
		return &ConstantStringInfo{cp: cp}
	case ConstantFieldRef:
		return &ConstantFieldRefInfo{ConstantMemberRefInfo{cp: cp}}
	case ConstantMethodRef:
		return &ConstantMethodRefInfo{ConstantMemberRefInfo{cp: cp}}
	case ConstantInterfaceMethodRef:
		return &ConstantInterfaceMethodRefInfo{ConstantMemberRefInfo{cp: cp}}
	case ConstantNameAndType:
		return &ConstantNameAndTypeInfo{cp: cp}
	case ConstantMethodHandle:
		return &ConstantMethodHandleInfo{cp: cp}
	case ConstantMethodType:
		return &ConstantMethodTypeInfo{cp: cp}
	case ConstantInvokeDynamic:
		return &ConstantInvokeDynamicInfo{cp: cp}
	default:
		panic(fmt.Sprintf("unknown constant info tag=%v", tag))
	}
}

func readConstantInfo(cr *ClassReader, cp ConstantPool) ConstantInfo {
	tag := cr.readUint8()
	info := newConstantInfo(tag, cp)
	info.readInfo(cr)
	return info
}
