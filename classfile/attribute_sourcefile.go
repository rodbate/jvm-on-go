package classfile

/**
SourceFile_attribute {
    u2 attribute_name_index;
    u4 attribute_length;
    u2 sourcefile_index;
}
*/

type SourceFileAttribute struct {
	cp         ConstantPool
	sourceFile string
}

func (s *SourceFileAttribute) readInfo(cr *ClassReader) {
	s.sourceFile = s.cp.getUtf8(cr.readUint16())
}

func (s *SourceFileAttribute) SourceFile() string {
	return s.sourceFile
}
