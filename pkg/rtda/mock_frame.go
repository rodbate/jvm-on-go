package rtda

import (
	"github.com/rodbate/jvm-on-go/pkg/instructions/opcodes"
)

var (
	mockClass    = &Class{name: "internal_mock"}
	returnMethod = &Method{
		Member: Member{
			AccessFlags: AccPublic | AccStatic,
			name:        "<internal_return_method>",
			descriptor:  "()V",
			class:       mockClass,
		},
		code: []byte{opcodes.Return},
	}
)

func NewMockFrame(thread *Thread, stack *OperandStack) *Frame {
	return &Frame{
		OperandStack: stack,
		thread:       thread,
		method:       returnMethod,
	}
}
