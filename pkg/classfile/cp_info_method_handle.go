package classfile

/**
CONSTANT_MethodHandle_info {
	u1 tag;
	u1 reference_kind;
	u2 reference_index;
}
*/

/**

Kind	Description	Interpretation
1	REF_getField	getfield C.f:T
2	REF_getStatic	getstatic C.f:T
3	REF_putField	putfield C.f:T
4	REF_putStatic	putstatic C.f:T
5	REF_invokeVirtual	invokevirtual C.m:(A*)T
6	REF_invokeStatic	invokestatic C.m:(A*)T
7	REF_invokeSpecial	invokespecial C.m:(A*)T
8	REF_newInvokeSpecial	new C; dup; invokespecial C.<init>:(A*)V
9	REF_invokeInterface	invokeinterface C.m:(A*)T
*/
const (
	Ref_getField uint8 = iota + 1
	Ref_getStatic
	Ref_putField
	Ref_putStatic
	Ref_invokeVirtual
	Ref_invokeStatic
	Ref_invokeSpecial
	Ref_newInvokeSpecial
	Ref_invokeInterface
)

type ConstantMethodHandleInfo struct {
	cp             ConstantPool
	referenceKind  uint8
	referenceIndex uint16
}

func (c *ConstantMethodHandleInfo) readInfo(cr *ClassReader) {
	c.referenceKind = cr.readUint8()
	c.referenceIndex = cr.readUint16()
}

func (c *ConstantMethodHandleInfo) GetRefInfo() *ConstantMemberRefInfo {
	return c.cp.getConstantInfo(c.referenceIndex).(*ConstantMemberRefInfo)
}

func (c *ConstantMethodHandleInfo) GetFieldRefInfo() *ConstantFieldRefInfo {
	switch c.referenceKind {
	case Ref_getField, Ref_getStatic, Ref_putField, Ref_putStatic:
		return c.cp.getConstantInfo(c.referenceIndex).(*ConstantFieldRefInfo)
	default:
		panic("Invalid referenceKind: ConstantMethodHandleInfo#GetFieldRefInfo")
	}
}

func (c *ConstantMethodHandleInfo) GetMethodRefInfo() *ConstantMethodRefInfo {
	switch c.referenceKind {
	case Ref_invokeVirtual, Ref_newInvokeSpecial:
		return c.cp.getConstantInfo(c.referenceIndex).(*ConstantMethodRefInfo)
	case Ref_invokeStatic, Ref_invokeSpecial:
		//only support jvm8
		constInfo := c.cp.getConstantInfo(c.referenceIndex)
		switch constInfo.(type) {
		case *ConstantMethodRefInfo:
			return constInfo.(*ConstantMethodRefInfo)
		default:
			return nil
		}
	default:
		panic("Invalid referenceKind: ConstantMethodHandleInfo#GetFieldRefInfo")
	}
}

func (c *ConstantMethodHandleInfo) GetInterfaceMethodRefInfo() *ConstantInterfaceMethodRefInfo {
	switch c.referenceKind {
	case Ref_invokeInterface:
		return c.cp.getConstantInfo(c.referenceIndex).(*ConstantInterfaceMethodRefInfo)
	default:
		panic("Invalid referenceKind: ConstantMethodHandleInfo#GetInterfaceMethodRefInfo")
	}
}
