package classfile

/**
Exceptions_attribute {
    u2 attribute_name_index;
    u4 attribute_length;
    u2 number_of_exceptions;
    u2 exception_index_table[number_of_exceptions]; // -> constant_class_info
}
*/

type ExceptionsAttribute struct {
	cp               ConstantPool
	exceptionClasses []string
}

func (e *ExceptionsAttribute) readInfo(cr *ClassReader) {
	exceptionCount := cr.readUint16()
	exceptions := make([]string, exceptionCount)
	for i := range exceptions {
		exceptions[i] = e.cp.getConstantClassInfo(cr.readUint16()).Name()
	}
	e.exceptionClasses = exceptions
}

func (e *ExceptionsAttribute) Exceptions() []string {
	return e.exceptionClasses
}
