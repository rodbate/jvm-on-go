package rtda

type Thread struct {
	pc       uint32
	stack    *Stack
	ExitCode int
}

func NewThread() *Thread {
	return &Thread{
		stack: newStack(1024),
	}
}

func (t *Thread) Pc() uint32 {
	return t.pc
}

func (t *Thread) SetPc(pc uint32) {
	t.pc = pc
}

func (t *Thread) PushFrame(frame *Frame) {
	t.stack.push(frame)
}

func (t *Thread) PopFrame() *Frame {
	return t.stack.pop()
}

func (t *Thread) CurrentFrame() *Frame {
	return t.stack.top()
}

func (t *Thread) NewFrame(method *Method) *Frame {
	return newFrame(method.maxLocals, method.maxStack, t, method)
}

func (t *Thread) IsStackEmpty() bool {
	return t.stack.isEmpty()
}

func (t *Thread) StackSize() uint {
	return t.stack.Size()
}

func (t *Thread) ClearStack() {
	t.stack.Clear()
}

func (t *Thread) GetFrames() []*Frame {
	return t.stack.GetFrames()
}
