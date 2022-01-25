package rtda

import (
	"github.com/rodbate/jvm-on-go/pkg/classfile"
)

type Utf8Info struct {
	Val string
}

func newUtf8Info(info *classfile.ConstantUtf8Info) *Utf8Info {
	return &Utf8Info{info.Val}
}
