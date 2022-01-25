package control

import (
	base2 "github.com/rodbate/jvm-on-go/pkg/instructions/base"
	"github.com/rodbate/jvm-on-go/pkg/rtda"
)

func LookupSwitch(reader *base2.ByteCodeReader, frame *rtda.Frame) {
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
			base2.Branch(frame, int(matchOffsets[i+1]))
			return
		}
	}
	base2.Branch(frame, int(defaultOffsets))
}
