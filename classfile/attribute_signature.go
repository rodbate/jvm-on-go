package classfile

/**
Signature_attribute {
	u2 attribute_name_index;
	u4 attribute_length;
	u2 signature_index;
}
*/
type SignatureAttribute struct {
	cp        ConstantPool
	Signature string
}

func (s *SignatureAttribute) readInfo(cr *ClassReader) {
	s.Signature = s.cp.getUtf8(cr.readUint16())
}
