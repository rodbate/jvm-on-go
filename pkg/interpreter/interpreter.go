package interpreter

import (
	"fmt"
	"github.com/rodbate/jvm-on-go/pkg/instructions"
	"github.com/rodbate/jvm-on-go/pkg/instructions/base"
	"github.com/rodbate/jvm-on-go/pkg/rtda"
)

func Interpret(thread *rtda.Thread, logInst bool) {
	defer catchError(thread)
	loop(thread, logInst)
}

func loop(thread *rtda.Thread, logInst bool) {
	reader := &base.ByteCodeReader{}
	for {
		frame := thread.CurrentFrame()
		pc := frame.NextPc()
		thread.SetPc(pc)
		reader.Reset(frame.Method().Code(), pc)
		opCode := reader.ReadUint8()
		inst := instructions.NewInstruction(opCode)
		if logInst {
			logInstruction(frame, inst, reader)
		}
		inst(reader, frame)
		if frame.NextPc() == pc {
			frame.SetNextPc(reader.PC())
		}
		if thread.IsStackEmpty() {
			break
		}
	}
}

func catchError(thread *rtda.Thread) {
	if r := recover(); r != nil {
		logFrames(thread)
		panic(r)
	}
}

func logFrames(thread *rtda.Thread) {
	for !thread.IsStackEmpty() {
		frame := thread.PopFrame()
		fmt.Printf(">> pc:%4d %v.%v%v\n",
			frame.NextPc(),
			frame.Method().Class().Name(),
			frame.Method().Name(),
			frame.Method().Descriptor())
	}
}

func logInstruction(frame *rtda.Frame, inst base.Instruction, reader *base.ByteCodeReader) {
	fmt.Printf("%v.%v() pc:%d, readerPos:%d, %v\n",
		frame.Method().Class().Name(),
		frame.Method().Name(),
		frame.Thread().Pc(),
		reader.PC(),
		inst.String())
}
