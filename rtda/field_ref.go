package rtda

import (
	"fmt"
	"github.com/rodbate/jvm-on-go/classfile"
)

type FieldRef struct {
	MemberRef
	field *Field
}

func newFieldRef(cp *ConstantPool, info *classfile.ConstantFieldRefInfo) *FieldRef {
	fieldRef := &FieldRef{}
	fieldRef.cp = cp
	fieldRef.initMemberRef(&info.ConstantMemberRefInfo)
	return fieldRef
}

func (ref *FieldRef) ResolvedField() *Field {
	if ref.field != nil {
		return ref.field
	}
	return ref.resolveFieldRef()
}

func (ref *FieldRef) resolveFieldRef() *Field {
	fromClass := ref.cp.class
	toClass := ref.ResolvedClass()
	field := lookupField(toClass, ref.name, ref.descriptor)
	if field == nil {
		panic(fmt.Sprintf("java.lang.NoSuchField: %s,%s", ref.name, ref.descriptor))
	}
	if !field.IsAccessibleTo(fromClass) {
		panic(fmt.Sprintf("java.lang.IllegalAccessError: field=%s:%s,class=%s",
			ref.name, ref.descriptor, fromClass.name))
	}
	ref.field = field
	return field
}

func lookupField(class *Class, name string, descriptor string) *Field {
	if class == nil {
		return nil
	}
	for _, f := range class.fields {
		if f.Match(name, descriptor) {
			return f
		}
	}
	for _, iFace := range class.interfaces {
		f := lookupField(iFace, name, descriptor)
		if f != nil {
			return f
		}
	}
	if superClass := class.superClass; superClass != nil {
		f := lookupField(superClass, name, descriptor)
		if f != nil {
			return f
		}
	}
	return nil
}
