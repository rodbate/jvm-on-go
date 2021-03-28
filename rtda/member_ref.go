package rtda

import "github.com/rodbate/jvm-on-go/classfile"

type MemberRef struct {
	SymbolicRef
	name       string
	descriptor string
}

func (m *MemberRef) initMemberRef(info *classfile.ConstantMemberRefInfo) {
	m.className = info.ClassName()
	m.name, m.descriptor = info.NameAndType()
}

func (m *MemberRef) Name() string {
	return m.name
}

func (m *MemberRef) Descriptor() string {
	return m.descriptor
}
