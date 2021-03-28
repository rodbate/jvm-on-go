package control

import (
	"jvm-on-go/instructions/base"
	"jvm-on-go/rtda"
)

func LookupSwitch(reader *base.ByteCodeReader, frame *rtda.Frame) {
	reader.SkipPadding()
	defaultOffsets := reader.ReadInt32()
	npairs := reader.ReadInt32()
	matchOffsets := make([]int32, npairs*2)
	for i := range matchOffsets {
		matchOffsets[i] = reader.ReadInt32()
	}

	key := frame.OperandStack.PopInt()
	for i := int32(0); i < npairs*2; i += 2 {
		if matchOffsets[i] == key {
			base.Branch(frame, int(matchOffsets[i+1]))
			return
		}
	}
	base.Branch(frame, int(defaultOffsets))
}
