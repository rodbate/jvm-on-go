package classfile

/**
LineNumberTable_attribute {
    u2 attribute_name_index;
    u4 attribute_length;
    u2 line_number_table_length;
    {   u2 start_pc;
        u2 line_number;
    } line_number_table[line_number_table_length];
}
*/

type LineNumberTableAttribute struct {
	LineNumberTable []*LineNumberTableEntry
}

type LineNumberTableEntry struct {
	StartPc    uint16
	LineNumber uint16
}

func (l *LineNumberTableAttribute) readInfo(cr *ClassReader) {
	count := cr.readUint16()
	table := make([]*LineNumberTableEntry, count)
	for i := range table {
		table[i] = &LineNumberTableEntry{
			StartPc:    cr.readUint16(),
			LineNumber: cr.readUint16(),
		}
	}
	l.LineNumberTable = table
}

func (l *LineNumberTableAttribute) GetLineNumber(pc uint16) int {
	for i := len(l.LineNumberTable) - 1; i >= 0; i-- {
		entry := l.LineNumberTable[i]
		if pc >= entry.StartPc {
			return int(entry.LineNumber)
		}
	}
	return -1
}
