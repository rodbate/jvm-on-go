package classfile

import "fmt"

/**
jvm8
doc: https://docs.oracle.com/javase/specs/jvms/se8/html/jvms-4.html
ClassFile {
    u4             magic;
    u2             minor_version;
    u2             major_version;
    u2             constant_pool_count;
    cp_info        constant_pool[constant_pool_count-1];
    u2             access_flags;
    u2             this_class;
    u2             super_class;
    u2             interfaces_count;
    u2             interfaces[interfaces_count];
    u2             fields_count;
    field_info     fields[fields_count];
    u2             methods_count;
    method_info    methods[methods_count];
    u2             attributes_count;
    attribute_info attributes[attributes_count];
}
*/
type ClassFile struct {
	magic        uint32
	minorVersion uint16
	majorVersion uint16
	constantPool ConstantPool
	accessFlags  uint16
	thisClass    uint16
	superClass   uint16
	interfaces   []uint16
	fields       []*MemberInfo
	methods      []*MemberInfo
	attributes   []AttributeInfo
}

func Parse(classData []byte) (cf *ClassFile, err error) {
	cr := &ClassReader{classData}
	cf = &ClassFile{}
	cf.read(cr)
	return
}

func (cf *ClassFile) read(cr *ClassReader) {
	cf.readAndCheckMagic(cr)
	cf.readAndCheckVersion(cr)

	cf.constantPool = readConstantPool(cr)

	cf.accessFlags = cr.readUint16()
	cf.thisClass = cr.readUint16()
	cf.superClass = cr.readUint16()
	cf.interfaces = cr.readUint16s()

	cf.fields = readMembers(cr, cf.constantPool)
	cf.methods = readMembers(cr, cf.constantPool)
	cf.attributes = readAttributeInfos(cr, cf.constantPool)
}

func (cf *ClassFile) readAndCheckMagic(cr *ClassReader) {
	magic := cr.readUint32()
	if magic != 0xCAFEBABE {
		panic("class file format error: invalid magic")
	}
	cf.magic = magic
}

func (cf *ClassFile) readAndCheckVersion(cr *ClassReader) {
	cf.minorVersion = cr.readUint16()
	cf.majorVersion = cr.readUint16()
	switch cf.majorVersion {
	case 52:
		if cf.minorVersion == 0 {
			return
		}
	}
	panic(fmt.Sprintf("class file format error: unsupported java version, [majorVersion=%d, minorVersion=%d]",
		cf.majorVersion, cf.minorVersion))
}

func (cf *ClassFile) MinorVersion() uint16 {
	return cf.minorVersion
}

func (cf *ClassFile) MajorVersion() uint16 {
	return cf.majorVersion
}

func (cf *ClassFile) ConstantPool() ConstantPool {
	return cf.constantPool
}

func (cf *ClassFile) AccessFlags() uint16 {
	return cf.accessFlags
}

func (cf *ClassFile) Fields() []*MemberInfo {
	return cf.fields
}

func (cf *ClassFile) Methods() []*MemberInfo {
	return cf.methods
}

func (cf *ClassFile) ClassName() string {
	return cf.constantPool.getConstantClassInfo(cf.thisClass).Name()
}

func (cf *ClassFile) SuperClassName() string {
	if cf.superClass == 0 {
		return ""
	}
	return cf.constantPool.getConstantClassInfo(cf.superClass).Name()
}

func (cf *ClassFile) InterfaceNames() []string {
	interfaces := make([]string, len(cf.interfaces))
	for i := range interfaces {
		interfaces[i] = cf.constantPool.getConstantClassInfo(cf.interfaces[i]).Name()
	}
	return interfaces
}

func (cf *ClassFile) SourceFile() string {
	for _, attr := range cf.attributes {
		switch attr.(type) {
		case *SourceFileAttribute:
			return attr.(*SourceFileAttribute).sourceFile
		}
	}
	return "UnknownSource"
}

func (cf *ClassFile) Signature() string {
	for _, attr := range cf.attributes {
		switch attr.(type) {
		case *SignatureAttribute:
			return attr.(*SignatureAttribute).Signature
		}
	}
	return ""
}

func (cf *ClassFile) RuntimeVisibleAnnotationData() []byte {
	for _, attrInfo := range cf.attributes {
		switch attrInfo.(type) {
		case *RuntimeVisibleAnnotationsAttribute:
			return attrInfo.(*RuntimeVisibleAnnotationsAttribute).Bytes
		}
	}
	return nil
}
