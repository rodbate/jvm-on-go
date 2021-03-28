package control

import (
	"github.com/rodbate/jvm-on-go/instructions/base"
	"github.com/rodbate/jvm-on-go/rtda"
)

func TableSwitch(reader *base.ByteCodeReader, frame *rtda.Frame) {
	reader.SkipPadding()
	defaultOffset := reader.ReadInt32()
	low := reader.ReadInt32()
	high := reader.ReadInt32()
	jumpOffsets := make([]int32, high-low+1)
	for i := range jumpOffsets {
		jumpOffsets[i] = reader.ReadInt32()
	}

	index := frame.OperandStack.PopInt()
	var offset int
	if index >= low && index <= high {
		offset = int(jumpOffsets[index - low])
	} else {
		offset = int(defaultOffset)
	}
	base.Branch(frame, offset)
}
