package classfile

/**
CONSTANT_MethodType_info {
	u1 tag;
	u2 descriptor_index;
}
*/

type ConstantMethodTypeInfo struct {
	cp              ConstantPool
	descriptorIndex uint16
}

func (c *ConstantMethodTypeInfo) readInfo(cr *ClassReader) {
	c.descriptorIndex = cr.readUint16()
}
