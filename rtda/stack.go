package rtda

type Stack struct {
	maxSize uint
	size    uint
	_top    *Frame
}

func newStack(maxSize uint) *Stack {
	return &Stack{maxSize: maxSize}
}

func (stack *Stack) push(frame *Frame) {
	if stack.size >= stack.maxSize {
		panic("java.lang.StackOverflowException")
	}
	frame.next = stack._top
	stack._top = frame
	stack.size++
}

func (stack *Stack) pop() *Frame {
	stack.checkNotEmpty()
	frame := stack._top
	stack._top = frame.next
	stack.size--
	return frame
}

func (stack *Stack) top() *Frame {
	stack.checkNotEmpty()
	return stack._top
}

func (stack *Stack) isEmpty() bool {
	return stack._top == nil
}

func (stack *Stack) checkNotEmpty() {
	if stack._top == nil {
		panic("stack is empty")
	}
}

func (stack *Stack) Size() uint {
	return stack.size
}

func (stack *Stack) Clear() {
	stack._top = nil
	stack.size = 0
}

func (stack *Stack) GetFrames() []*Frame {
	frames := make([]*Frame, 0, stack.size)
	for f := stack._top; f != nil; f = f.next {
		frames = append(frames, f)
	}
	return frames
}
