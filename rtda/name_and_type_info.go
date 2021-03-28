package rtda

import "github.com/rodbate/jvm-on-go/classfile"

type NameAndTypeInfo struct {
	Name       string
	Descriptor string
}

func newNameAndTypeInfo(info *classfile.ConstantNameAndTypeInfo) *NameAndTypeInfo {
	return &NameAndTypeInfo{
		Name:       info.Name(),
		Descriptor: info.Descriptor(),
	}
}
