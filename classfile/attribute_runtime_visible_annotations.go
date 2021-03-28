package classfile

/**
RuntimeVisibleAnnotations_attribute {
	u2 attribute_name_index;
	u4 attribute_length;
	u2 num_annotations;
	annotation annotations[num_annotations];
}

annotation {
	u2 type_index;		//CONSTANT_Utf8_info
	u2 num_element_value_pairs;
	{ 	u2 element_name_index;   //CONSTANT_Utf8_info
		element_value value;
	} element_value_pairs[num_element_value_pairs];
}

element_value {
	u1 tag;
	union {
		u2 const_value_index;
		{ 	u2 type_name_index;
			u2 const_name_index;
		} enum_const_value;
		u2 class_info_index;
		annotation annotation_value;
		{ 	u2 num_values;
			element_value values[num_values];
		} array_value;
	} value;
}
*/
type RuntimeVisibleAnnotationsAttribute struct {
	attributeLen uint32
	Bytes        []byte
}

func (r *RuntimeVisibleAnnotationsAttribute) readInfo(cr *ClassReader) {
	r.Bytes = cr.readBytes(r.attributeLen)
}
