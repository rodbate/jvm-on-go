package control

import (
	"jvm-on-go/instructions/base"
	"jvm-on-go/rtda"
)

func Goto(reader *base.ByteCodeReader, frame *rtda.Frame) {
	offset := reader.ReadInt16()
	base.Branch(frame, int(offset))
}
