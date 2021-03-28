package control

import (
	"jvm-on-go/instructions/base"
	"jvm-on-go/rtda"
)

func Return(reader *base.ByteCodeReader, frame *rtda.Frame) {
	frame.Thread().PopFrame()
}

func AReturn(reader *base.ByteCodeReader, frame *rtda.Frame) {
	doExecute(frame, func(currentFrame *rtda.Frame, invokerFrame *rtda.Frame) {
		invokerFrame.OperandStack.PushRef(currentFrame.OperandStack.PopRef())
	})
}

func IReturn(reader *base.ByteCodeReader, frame *rtda.Frame) {
	doExecute(frame, func(currentFrame *rtda.Frame, invokerFrame *rtda.Frame) {
		invokerFrame.OperandStack.PushInt(currentFrame.OperandStack.PopInt())
	})
}

func DReturn(reader *base.ByteCodeReader, frame *rtda.Frame) {
	doExecute(frame, func(currentFrame *rtda.Frame, invokerFrame *rtda.Frame) {
		invokerFrame.OperandStack.PushDouble(currentFrame.OperandStack.PopDouble())
	})
}

func FReturn(reader *base.ByteCodeReader, frame *rtda.Frame) {
	doExecute(frame, func(currentFrame *rtda.Frame, invokerFrame *rtda.Frame) {
		invokerFrame.OperandStack.PushFloat(currentFrame.OperandStack.PopFloat())
	})
}

func LReturn(reader *base.ByteCodeReader, frame *rtda.Frame) {
	doExecute(frame, func(currentFrame *rtda.Frame, invokerFrame *rtda.Frame) {
		invokerFrame.OperandStack.PushLong(currentFrame.OperandStack.PopLong())
	})
}

func doExecute(frame *rtda.Frame, execution func(*rtda.Frame, *rtda.Frame)) {
	thread := frame.Thread()
	currentFrame := thread.PopFrame()
	invokerFrame := thread.CurrentFrame()
	execution(currentFrame, invokerFrame)
}
