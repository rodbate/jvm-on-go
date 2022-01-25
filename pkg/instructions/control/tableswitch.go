package control

import (
	base2 "github.com/rodbate/jvm-on-go/pkg/instructions/base"
	"github.com/rodbate/jvm-on-go/pkg/rtda"
)

func TableSwitch(reader *base2.ByteCodeReader, frame *rtda.Frame) {
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
		offset = int(jumpOffsets[index-low])
	} else {
		offset = int(defaultOffset)
	}
	base2.Branch(frame, offset)
}
