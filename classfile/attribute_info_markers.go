package classfile

type AttributeMarker struct {
}

func (a *AttributeMarker) readInfo(cr *ClassReader) {
}

type DeprecatedAttribute struct {
	AttributeMarker
}

type SyntheticAttribute struct {
	AttributeMarker
}
