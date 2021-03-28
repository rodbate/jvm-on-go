package classfile

type UnparsedAttributeInfo struct {
	attributeName string
	attributeLen  uint32
	bytes         []byte
}

func (u *UnparsedAttributeInfo) readInfo(cr *ClassReader) {
	u.bytes = cr.readBytes(u.attributeLen)
}

func (u *UnparsedAttributeInfo) Name() string {
	return u.attributeName
}
