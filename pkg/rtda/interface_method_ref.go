package rtda

import (
	"fmt"
	"github.com/rodbate/jvm-on-go/pkg/classfile"
)

type InterfaceMethodRef struct {
	MemberRef
	method *Method
}

func newInterfaceMethodRef(cp *ConstantPool, info *classfile.ConstantInterfaceMethodRefInfo) *InterfaceMethodRef {
	ref := &InterfaceMethodRef{}
	ref.cp = cp
	ref.initMemberRef(&info.ConstantMemberRefInfo)
	return ref
}

func (ref *InterfaceMethodRef) ResolvedInterfaceMethod() *Method {
	if ref.method == nil {
		ref.resolveInterfaceMethodRef()
	}
	return ref.method
}

func (ref *InterfaceMethodRef) resolveInterfaceMethodRef() {
	fromClass := ref.cp.class
	targetClass := ref.ResolvedClass()
	method := targetClass.LookupInterfaceMethod(ref.name, ref.descriptor)
	if method == nil {
		panic("java.lang.NoSuchMethod")
	}
	if method.IsStatic() || method.IsPrivate() {
		panic("java.lang.IncompatibleClassChangeError")
	}
	if !method.IsAccessibleTo(fromClass) {
		panic(fmt.Sprintf("java.lang.IllegalAccessError: method=%s:%s,class=%s",
			ref.name, ref.descriptor, fromClass.name))
	}
	ref.method = method
}
