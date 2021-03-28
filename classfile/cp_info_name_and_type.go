package classfile

/**
CONSTANT_NameAndType_info {
    u1 tag;
    u2 name_index;
    u2 descriptor_index;
}
*/

type ConstantNameAndTypeInfo struct {
	cp              ConstantPool
	nameIndex       uint16
	descriptorIndex uint16
}

func (c *ConstantNameAndTypeInfo) readInfo(cr *ClassReader) {
	c.nameIndex = cr.readUint16()
	c.descriptorIndex = cr.readUint16()
}

func (c *ConstantNameAndTypeInfo) Name() string {
	return c.cp.getUtf8(c.nameIndex)
}

func (c *ConstantNameAndTypeInfo) Descriptor() string {
	return c.cp.getUtf8(c.descriptorIndex)
}
