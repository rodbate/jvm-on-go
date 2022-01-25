package extended

import (
	"github.com/rodbate/jvm-on-go/pkg/instructions/base"
	load2 "github.com/rodbate/jvm-on-go/pkg/instructions/load"
	"github.com/rodbate/jvm-on-go/pkg/instructions/math"
	"github.com/rodbate/jvm-on-go/pkg/instructions/opcodes"
	store2 "github.com/rodbate/jvm-on-go/pkg/instructions/store"
	"github.com/rodbate/jvm-on-go/pkg/rtda"
)

/**
iload, fload, aload, lload, dload, istore, fstore, astore, lstore, dstore, ret, iinc
*/

func Wide(reader *base.ByteCodeReader, frame *rtda.Frame) {
	opCode := reader.ReadUint8()
	switch opCode {
	case opcodes.ILoad:
		load2.DoILoad(uint16(reader.ReadUint8()), frame)
	case opcodes.FLoad:
		load2.DoFLoad(uint16(reader.ReadUint8()), frame)
	case opcodes.ALoad:
		load2.DoALoad(uint16(reader.ReadUint8()), frame)
	case opcodes.LLoad:
		load2.DoLLoad(uint16(reader.ReadUint8()), frame)
	case opcodes.DLoad:
		load2.DoDLoad(uint16(reader.ReadUint8()), frame)
	case opcodes.IStore:
		store2.DoStoreInt(uint16(reader.ReadUint8()), frame)
	case opcodes.FStore:
		store2.DoStoreFloat(uint16(reader.ReadUint8()), frame)
	case opcodes.AStore:
		store2.DoStoreRef(uint16(reader.ReadUint8()), frame)
	case opcodes.LStore:
		store2.DoStoreLong(uint16(reader.ReadUint8()), frame)
	case opcodes.DStore:
		store2.DoStoreDouble(uint16(reader.ReadUint8()), frame)
	case opcodes.IInc:
		math.IInc(reader, frame)
	case opcodes.RET:
		panic("not supported op code: " + string(opCode))
	}
}
