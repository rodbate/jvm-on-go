package rtda

type SymbolicRef struct {
	cp        *ConstantPool
	className string
	class     *Class
}

func (ref *SymbolicRef) ResolvedClass() *Class {
	if ref.class != nil {
		return ref.class
	}
	return ref.resolveClassRef()
}

func (ref *SymbolicRef) resolveClassRef() *Class {
	from := ref.cp.class
	to := from.classloader.LoadClass(ref.className)
	if !to.IsAccessibleTo(from) {
		panic("java.lang.IllegalAccessError")
	}
	ref.class = to
	return to
}
