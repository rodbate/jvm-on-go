package classfile

/**
field_info {
    u2             access_flags;
    u2             name_index;
    u2             descriptor_index;
    u2             attributes_count;
    attribute_info attributes[attributes_count];
}

method_info {
    u2             access_flags;
    u2             name_index;
    u2             descriptor_index;
    u2             attributes_count;
    attribute_info attributes[attributes_count];
}
*/

type MemberInfo struct {
	constantPool    ConstantPool
	accessFlags     uint16
	nameIndex       uint16
	descriptorIndex uint16
	attributes      []AttributeInfo
}

func readMembers(cr *ClassReader, cp ConstantPool) []*MemberInfo {
	memberCount := cr.readUint16()
	members := make([]*MemberInfo, memberCount)
	for i := range members {
		members[i] = readMember(cr, cp)
	}
	return members
}

func readMember(cr *ClassReader, cp ConstantPool) *MemberInfo {
	return &MemberInfo{
		constantPool:    cp,
		accessFlags:     cr.readUint16(),
		nameIndex:       cr.readUint16(),
		descriptorIndex: cr.readUint16(),
		attributes:      readAttributeInfos(cr, cp),
	}
}

func (info *MemberInfo) AccessFlags() uint16 {
	return info.accessFlags
}

func (info *MemberInfo) Name() string {
	return info.constantPool.getUtf8(info.nameIndex)
}

func (info *MemberInfo) Descriptor() string {
	return info.constantPool.getUtf8(info.descriptorIndex)
}

func (info *MemberInfo) Attributes() []AttributeInfo {
	return info.attributes
}

func (info *MemberInfo) CodeAttribute() *CodeAttribute {
	for _, attrInfo := range info.attributes {
		switch attrInfo.(type) {
		case *CodeAttribute:
			return attrInfo.(*CodeAttribute)
		}
	}
	return nil
}

func (info *MemberInfo) ConstantValueAttribute() *ConstantValueAttribute {
	for _, attrInfo := range info.attributes {
		switch attrInfo.(type) {
		case *ConstantValueAttribute:
			return attrInfo.(*ConstantValueAttribute)
		}
	}
	return nil
}

func (info *MemberInfo) Signature() string {
	for _, attrInfo := range info.attributes {
		switch attrInfo.(type) {
		case *SignatureAttribute:
			return attrInfo.(*SignatureAttribute).Signature
		}
	}
	return ""
}

func (info *MemberInfo) Exceptions() []string {
	for _, attrInfo := range info.attributes {
		switch attrInfo.(type) {
		case *ExceptionsAttribute:
			return attrInfo.(*ExceptionsAttribute).exceptionClasses
		}
	}
	return nil
}

func (info *MemberInfo) RuntimeVisibleAnnotationData() []byte {
	for _, attrInfo := range info.attributes {
		switch attrInfo.(type) {
		case *RuntimeVisibleAnnotationsAttribute:
			return attrInfo.(*RuntimeVisibleAnnotationsAttribute).Bytes
		}
	}
	return nil
}

func (info *MemberInfo) RuntimeVisibleParameterAnnotationData() []byte {
	for _, attrInfo := range info.attributes {
		switch attrInfo.(type) {
		case *RuntimeVisibleParameterAnnotationsAttribute:
			return attrInfo.(*RuntimeVisibleParameterAnnotationsAttribute).Bytes
		}
	}
	return nil
}
