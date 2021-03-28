package classfile

type ConstantPool []ConstantInfo

func readConstantPool(cr *ClassReader) ConstantPool {
	cpCount := int(cr.readUint16())
	cp := make(ConstantPool, cpCount)
	for i := 1; i < cpCount; i++ {
		cp[i] = readConstantInfo(cr, cp)
		switch cp[i].(type) {
		case *ConstantLongInfo, *ConstantDoubleInfo:
			i++
		}
	}
	return cp
}

func (cp ConstantPool) getConstantInfo(index uint16) ConstantInfo {
	return cp[index]
}

func (cp ConstantPool) getUtf8(index uint16) string {
	utf8Info := cp[index].(*ConstantUtf8Info)
	return utf8Info.Val
}

func (cp ConstantPool) getConstantClassInfo(index uint16) *ConstantClassInfo {
	return cp[index].(*ConstantClassInfo)
}
