package rtda

type Frame struct {
	next         *Frame
	LocalVars    Slots
	OperandStack *OperandStack
	thread       *Thread
	nextPc       uint32
	method       *Method
}

func newFrame(maxLocals, maxStack uint16, thread *Thread, method *Method) *Frame {
	return &Frame{
		LocalVars:    newLocalVars(maxLocals),
		OperandStack: NewOperandStack(maxStack),
		thread:       thread,
		method:       method,
	}
}

func (f *Frame) Thread() *Thread {
	return f.thread
}

func (f *Frame) NextPc() uint32 {
	return f.nextPc
}

func (f *Frame) SetNextPc(nextPc uint32) {
	f.nextPc = nextPc
}

func (f *Frame) Method() *Method {
	return f.method
}

func (f *Frame) RevertNextPc() {
	f.nextPc = f.thread.pc
}
