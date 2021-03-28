package rtda

import "jvm-on-go/classfile"

type Member struct {
	AccessFlags    uint16
	name           string
	descriptor     string
	Signature      string
	AnnotationData []byte
	class          *Class
}

func (m *Member) initInfo(info *classfile.MemberInfo) {
	m.AccessFlags = info.AccessFlags()
	m.name = info.Name()
	m.descriptor = info.Descriptor()
	m.Signature = info.Signature()
	m.AnnotationData = info.RuntimeVisibleAnnotationData()
}

func (m *Member) IsPublic() bool {
	return m.AccessFlags&AccPublic != 0
}

func (m *Member) IsPrivate() bool {
	return m.AccessFlags&AccPrivate != 0
}

func (m *Member) IsProtected() bool {
	return m.AccessFlags&AccProtected != 0
}

func (m *Member) IsStatic() bool {
	return m.AccessFlags&AccStatic != 0
}

func (m *Member) IsFinal() bool {
	return m.AccessFlags&AccFinal != 0
}

func (m *Member) IsSynthetic() bool {
	return m.AccessFlags&AccSynthetic != 0
}

func (m *Member) Class() *Class {
	return m.class
}

func (m *Member) Name() string {
	return m.name
}

func (m *Member) Descriptor() string {
	return m.descriptor
}

func (m *Member) IsAccessibleTo(class *Class) bool {
	if m.IsPublic() {
		return true
	}
	if m.IsProtected() {
		return m.class.IsAssignableFrom(class) || m.class.GetPackageName() == class.GetPackageName()
	}
	if m.IsPrivate() {
		return m.class == class
	}
	return m.class.GetPackageName() == class.GetPackageName() //default scope
}
