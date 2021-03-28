package classfile

/**
CONSTANT_InvokeDynamic_info {
	u1 tag;
	u2 bootstrap_method_attr_index;
	u2 name_and_type_index;
}
*/

type ConstantInvokeDynamicInfo struct {
	cp                       ConstantPool
	bootstrapMethodAttrIndex uint16
	nameAndTypeIndex         uint16
}

func (c *ConstantInvokeDynamicInfo) readInfo(cr *ClassReader) {
	c.bootstrapMethodAttrIndex = cr.readUint16()
	c.nameAndTypeIndex = cr.readUint16()
}
