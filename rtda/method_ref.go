package rtda

import (
	"fmt"
	"github.com/rodbate/jvm-on-go/classfile"
)

type MethodRef struct {
	MemberRef
	method *Method
}

func newMethodRef(cp *ConstantPool, info *classfile.ConstantMemberRefInfo) *MethodRef {
	methodRef := &MethodRef{}
	methodRef.cp = cp
	methodRef.initMemberRef(info)
	return methodRef
}

func (m *MethodRef) ResolvedMethod() *Method {
	if m.method == nil {
		m.resolveMethodRef()
	}
	return m.method
}

func (m *MethodRef) resolveMethodRef() {
	from := m.cp.class
	targetClass := m.ResolvedClass()
	if targetClass.IsInterface() {
		panic("java.lang.IncompatibleClassChangeError")
	}
	method := targetClass.LookupMethod(m.name, m.descriptor)
	if method == nil {
		panic("java.lang.NoSuchMethod")
	}
	if !method.IsAccessibleTo(from) {
		panic(fmt.Sprintf("java.lang.IllegalAccessError: method=%s%s,class=%s",
			m.name, m.descriptor, from.name))
	}
	m.method = method
}
