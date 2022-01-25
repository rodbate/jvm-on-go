package extended

import (
	base2 "github.com/rodbate/jvm-on-go/pkg/instructions/base"
	"github.com/rodbate/jvm-on-go/pkg/rtda"
)

func GotoW(reader *base2.ByteCodeReader, frame *rtda.Frame) {
	offset := reader.ReadInt32()
	base2.Branch(frame, int(offset))
}
