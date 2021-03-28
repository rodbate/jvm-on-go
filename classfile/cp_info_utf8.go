package classfile

/**
CONSTANT_Utf8_info {
    u1 tag;
    u2 length;
    u1 bytes[length];
}
*/
type ConstantUtf8Info struct {
	Val string
}

func (c *ConstantUtf8Info) readInfo(cr *ClassReader) {
	length := uint32(cr.readUint16())
	bytes := cr.readBytes(length)
	c.Val = decodeMUtf8(bytes)
}

func decodeMUtf8(bytes []byte) string {
	return string(bytes)
}
