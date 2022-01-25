package rtda

import (
	"fmt"
	"github.com/rodbate/jvm-on-go/pkg/classfile"
)

type Constant interface{}

type ConstantPool struct {
	class     *Class
	constants []Constant
}

func newConstantPool(class *Class, pool classfile.ConstantPool) *ConstantPool {
	cp := &ConstantPool{}
	cp.class = class
	cpCount := len(pool)
	constants := make([]Constant, cpCount)
	cp.constants = constants

	for i := 1; i < cpCount; i++ {
		cpInfo := pool[i]
		switch cpInfo.(type) {
		case *classfile.ConstantIntegerInfo:
			constants[i] = (cpInfo.(*classfile.ConstantIntegerInfo)).Value()
		case *classfile.ConstantFloatInfo:
			constants[i] = (cpInfo.(*classfile.ConstantFloatInfo)).Value()
		case *classfile.ConstantLongInfo:
			constants[i] = (cpInfo.(*classfile.ConstantLongInfo)).Value()
			i++
		case *classfile.ConstantDoubleInfo:
			constants[i] = (cpInfo.(*classfile.ConstantDoubleInfo)).Value()
			i++
		case *classfile.ConstantStringInfo:
			constants[i] = (cpInfo.(*classfile.ConstantStringInfo)).String()
		case *classfile.ConstantClassInfo:
			constants[i] = newClassRef(cp, cpInfo.(*classfile.ConstantClassInfo))
		case *classfile.ConstantMethodRefInfo:
			constants[i] = newMethodRef(cp, &cpInfo.(*classfile.ConstantMethodRefInfo).ConstantMemberRefInfo)
		case *classfile.ConstantInterfaceMethodRefInfo:
			constants[i] = newInterfaceMethodRef(cp, cpInfo.(*classfile.ConstantInterfaceMethodRefInfo))
		case *classfile.ConstantFieldRefInfo:
			constants[i] = newFieldRef(cp, cpInfo.(*classfile.ConstantFieldRefInfo))
		case *classfile.ConstantNameAndTypeInfo:
			constants[i] = newNameAndTypeInfo(cpInfo.(*classfile.ConstantNameAndTypeInfo))
		case *classfile.ConstantUtf8Info:
			constants[i] = newUtf8Info(cpInfo.(*classfile.ConstantUtf8Info))
		case *classfile.ConstantMethodHandleInfo:
			constants[i] = newMethodHandleInfo(cpInfo.(*classfile.ConstantMethodHandleInfo))
		case *classfile.ConstantMethodTypeInfo:
			constants[i] = nil
		case *classfile.ConstantInvokeDynamicInfo:
			constants[i] = nil
		default:
			panic(fmt.Sprintf("not supported constant info types: %v\n", cpInfo))
		}
	}
	return cp
}

func (cp *ConstantPool) GetConstant(index uint16) Constant {
	if v := cp.constants[index]; v != nil {
		return v
	}
	panic(fmt.Sprintf("not found constant at index: %v\n", index))
}
