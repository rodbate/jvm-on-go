package classfile

/**
LocalVariableTable_attribute {
    u2 attribute_name_index;
    u4 attribute_length;
    u2 local_variable_table_length;
    {   u2 start_pc;
        u2 length;
        u2 name_index;
        u2 descriptor_index;
        u2 index;
    } local_variable_table[local_variable_table_length];
}
*/

type LocalVariableTableAttribute struct {
	cp                 ConstantPool
	localVariableTable []*LocalVariableTableEntry
}

func (l *LocalVariableTableAttribute) readInfo(cr *ClassReader) {
	entries := make([]*LocalVariableTableEntry, cr.readUint16())
	for i := range entries {
		entries[i] = &LocalVariableTableEntry{
			startPc:    cr.readUint16(),
			length:     cr.readUint16(),
			name:       l.cp.getUtf8(cr.readUint16()),
			descriptor: l.cp.getUtf8(cr.readUint16()),
			index:      cr.readUint16(),
		}
	}
}

type LocalVariableTableEntry struct {
	startPc    uint16
	length     uint16
	name       string
	descriptor string
	index      uint16
}
