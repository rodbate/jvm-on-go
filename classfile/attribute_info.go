package classfile

/**
attribute_info {
    u2 attribute_name_index;
    u4 attribute_length;
    u1 info[attribute_length];
}
*/

type AttributeInfo interface {
	readInfo(cr *ClassReader)
}

func readAttributeInfos(cr *ClassReader, cp ConstantPool) []AttributeInfo {
	infos := make([]AttributeInfo, cr.readUint16())
	for i := range infos {
		infos[i] = readAttributeInfo(cr, cp)
	}
	return infos
}

func readAttributeInfo(cr *ClassReader, cp ConstantPool) AttributeInfo {
	attrIndex := cr.readUint16()
	attrName := cp.getUtf8(attrIndex)
	attrLen := cr.readUint32()
	attrInfo := newAttributeInfo(attrName, attrLen, cp)
	attrInfo.readInfo(cr)
	return attrInfo
}

func newAttributeInfo(attributeName string, attributeLen uint32, cp ConstantPool) AttributeInfo {
	switch attributeName {
	case "Deprecated":
		return &DeprecatedAttribute{}
	case "Synthetic":
		return &SyntheticAttribute{}
	case "SourceFile":
		return &SourceFileAttribute{cp: cp}
	case "LineNumberTable":
		return &LineNumberTableAttribute{}
	case "Exceptions":
		return &ExceptionsAttribute{cp: cp}
	case "ConstantValue":
		return &ConstantValueAttribute{}
	case "Code":
		return &CodeAttribute{cp: cp}
	case "LocalVariableTable":
		return &LocalVariableTableAttribute{cp: cp}
	case "Signature":
		return &SignatureAttribute{cp: cp}
	case "RuntimeVisibleAnnotations":
		return &RuntimeVisibleAnnotationsAttribute{attributeLen: attributeLen}
	case "RuntimeVisibleParameterAnnotations":
		return &RuntimeVisibleParameterAnnotationsAttribute{attributeLen: attributeLen}
	case "BootstrapMethods":
		return &BootstrapMethodsAttribute{cp: cp}
	default:
		return &UnparsedAttributeInfo{attributeName: attributeName, attributeLen: attributeLen}
	}
}
