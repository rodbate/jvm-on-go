package comparison

import (
	"github.com/rodbate/jvm-on-go/pkg/instructions/base"
	"github.com/rodbate/jvm-on-go/pkg/rtda"
)

func DCmpg(reader *base.ByteCodeReader, frame *rtda.Frame) {
	execute(frame, true)
}

func DCmpl(reader *base.ByteCodeReader, frame *rtda.Frame) {
	execute(frame, false)
}

func execute(frame *rtda.Frame, isDCmpg bool) {
	val2 := frame.OperandStack.PopDouble()
	val1 := frame.OperandStack.PopDouble()
	if val1 > val2 {
		frame.OperandStack.PushInt(1)
	} else if val1 < val2 {
		frame.OperandStack.PushInt(-1)
	} else if val1 == val2 {
		frame.OperandStack.PushInt(0)
	} else {
		if isDCmpg {
			frame.OperandStack.PushInt(1)
		} else {
			frame.OperandStack.PushInt(-1)
		}
	}
}
