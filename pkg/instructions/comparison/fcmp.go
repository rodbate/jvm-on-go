package comparison

import (
	"github.com/rodbate/jvm-on-go/pkg/instructions/base"
	"github.com/rodbate/jvm-on-go/pkg/rtda"
)

func FCmpg(reader *base.ByteCodeReader, frame *rtda.Frame) {
	_execute(frame, true)
}

func FCmpl(reader *base.ByteCodeReader, frame *rtda.Frame) {
	_execute(frame, false)
}

func _execute(frame *rtda.Frame, isFCmpg bool) {
	val2 := frame.OperandStack.PopFloat()
	val1 := frame.OperandStack.PopFloat()
	if val1 > val2 {
		frame.OperandStack.PushInt(1)
	} else if val1 < val2 {
		frame.OperandStack.PushInt(-1)
	} else if val1 == val2 {
		frame.OperandStack.PushInt(0)
	} else {
		if isFCmpg {
			frame.OperandStack.PushInt(1)
		} else {
			frame.OperandStack.PushInt(-1)
		}
	}
}
