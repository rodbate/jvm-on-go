package extended

import (
	"github.com/rodbate/jvm-on-go/instructions/base"
	"github.com/rodbate/jvm-on-go/rtda"
)

func GotoW(reader *base.ByteCodeReader, frame *rtda.Frame) {
	offset := reader.ReadInt32()
	base.Branch(frame, int(offset))
}
