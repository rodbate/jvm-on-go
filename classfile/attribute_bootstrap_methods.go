package classfile

/**
Belong to ClassFile Attributes Table

BootstrapMethods_attribute {
    u2 attribute_name_index;
    u4 attribute_length;
    u2 num_bootstrap_methods;
    {   u2 bootstrap_method_ref;
        u2 num_bootstrap_arguments;
        u2 bootstrap_arguments[num_bootstrap_arguments];
    } bootstrap_methods[num_bootstrap_methods];
}
*/

type BootstrapMethodsAttribute struct {
	cp               ConstantPool
	BootstrapMethods []BootstrapMethod
}

type BootstrapMethod struct {
	BootstrapMethodRef uint16
	BootstrapArguments []uint16
}

func (b *BootstrapMethodsAttribute) readInfo(cr *ClassReader) {
	numBootstrapMethods := cr.readUint16()
	bootstrapMethods := make([]BootstrapMethod, numBootstrapMethods)
	for i := range bootstrapMethods {
		bootstrapMethods[i] = parseBootstrapMethod(cr)
	}
	b.BootstrapMethods = bootstrapMethods
}

func parseBootstrapMethod(cr *ClassReader) BootstrapMethod {
	bootstrapMethodRef := cr.readUint16()
	numBootstrapArguments := cr.readUint16()
	bootstrapArguments := make([]uint16, numBootstrapArguments)
	for i := range bootstrapArguments {
		bootstrapArguments[i] = cr.readUint16()
	}
	return BootstrapMethod{
		BootstrapMethodRef: bootstrapMethodRef,
		BootstrapArguments: bootstrapArguments,
	}
}

