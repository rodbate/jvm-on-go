package classfile

/**
Code_attribute {
    u2 attribute_name_index;
    u4 attribute_length;
    u2 max_stack;
    u2 max_locals;
    u4 code_length;
    u1 code[code_length];
    u2 exception_table_length;
    {   u2 start_pc;
        u2 end_pc;
        u2 handler_pc;
        u2 catch_type;
    } exception_table[exception_table_length];
    u2 attributes_count;
    attribute_info attributes[attributes_count];
}
*/

type CodeAttribute struct {
	cp             ConstantPool
	maxStack       uint16
	maxLocals      uint16
	code           []byte
	exceptionTable []*ExceptionTableEntry
	attributeInfo  []AttributeInfo
}

type ExceptionTableEntry struct {
	StartPc   uint16
	EndPc     uint16
	HandlerPc uint16
	CatchType uint16
}

func (c *CodeAttribute) readInfo(cr *ClassReader) {
	c.maxStack = cr.readUint16()
	c.maxLocals = cr.readUint16()
	codeLen := cr.readUint32()
	c.code = cr.readBytes(codeLen)
	exceptionTableLen := cr.readUint16()
	exceptionTable := make([]*ExceptionTableEntry, exceptionTableLen)
	for i := range exceptionTable {
		exceptionTable[i] = readExceptionTableEntry(cr)
	}
	c.exceptionTable = exceptionTable
	c.attributeInfo = readAttributeInfos(cr, c.cp)
}

func readExceptionTableEntry(cr *ClassReader) *ExceptionTableEntry {
	return &ExceptionTableEntry{
		StartPc:   cr.readUint16(),
		EndPc:     cr.readUint16(),
		HandlerPc: cr.readUint16(),
		CatchType: cr.readUint16(),
	}
}

func (c *CodeAttribute) MaxStack() uint16 {
	return c.maxStack
}

func (c *CodeAttribute) MaxLocals() uint16 {
	return c.maxLocals
}

func (c *CodeAttribute) Code() []byte {
	return c.code
}

func (c *CodeAttribute) ExceptionTable() []*ExceptionTableEntry {
	return c.exceptionTable
}

func (c *CodeAttribute) LineNumberTableAttribute() *LineNumberTableAttribute {
	for _, attr := range c.attributeInfo {
		switch attr.(type) {
		case *LineNumberTableAttribute:
			return attr.(*LineNumberTableAttribute)
		}
	}
	return nil
}
