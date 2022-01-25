package rtda

import (
	"github.com/rodbate/jvm-on-go/pkg/classfile"
)

type ClassRef struct {
	SymbolicRef
}

func newClassRef(cp *ConstantPool, classInfo *classfile.ConstantClassInfo) *ClassRef {
	cr := &ClassRef{}
	cr.cp = cp
	cr.className = classInfo.Name()
	return cr
}
