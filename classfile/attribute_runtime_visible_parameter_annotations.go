package classfile

type RuntimeVisibleParameterAnnotationsAttribute struct {
	attributeLen uint32
	Bytes        []byte
}

func (r *RuntimeVisibleParameterAnnotationsAttribute) readInfo(cr *ClassReader) {
	r.Bytes = cr.readBytes(r.attributeLen)
}
