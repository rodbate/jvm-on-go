package classfile

/**
CONSTANT_Fieldref_info {
    u1 tag;
    u2 class_index;
    u2 name_and_type_index;
}

CONSTANT_Methodref_info {
    u1 tag;
    u2 class_index;
    u2 name_and_type_index;
}

CONSTANT_InterfaceMethodref_info {
    u1 tag;
    u2 class_index;
    u2 name_and_type_index;
}
*/

type ConstantMemberRefInfo struct {
	cp               ConstantPool
	classIndex       uint16
	nameAndTypeIndex uint16
}

func (c *ConstantMemberRefInfo) readInfo(cr *ClassReader) {
	c.classIndex = cr.readUint16()
	c.nameAndTypeIndex = cr.readUint16()
}

type ConstantFieldRefInfo struct {
	ConstantMemberRefInfo
}

type ConstantMethodRefInfo struct {
	ConstantMemberRefInfo
}

type ConstantInterfaceMethodRefInfo struct {
	ConstantMemberRefInfo
}

func (c *ConstantMemberRefInfo) ClassName() string {
	return c.cp.getConstantClassInfo(c.classIndex).Name()
}

func (c *ConstantMemberRefInfo) NameAndType() (string, string) {
	nameAndType := c.cp.getConstantInfo(c.nameAndTypeIndex).(*ConstantNameAndTypeInfo)
	return nameAndType.Name(), nameAndType.Descriptor()
}
