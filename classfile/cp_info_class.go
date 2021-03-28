package classfile

/**
CONSTANT_Class_info {
    u1 tag;
    u2 name_index;
}
*/

type ConstantClassInfo struct {
	cp        ConstantPool
	nameIndex uint16
}

func (c *ConstantClassInfo) readInfo(cr *ClassReader) {
	c.nameIndex = cr.readUint16()
}

func (c *ConstantClassInfo) Name() string {
	return c.cp.getUtf8(c.nameIndex)
}
