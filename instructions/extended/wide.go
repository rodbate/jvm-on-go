package extended

import (
	"jvm-on-go/instructions/base"
	"jvm-on-go/instructions/load"
	"jvm-on-go/instructions/math"
	"jvm-on-go/instructions/opcodes"
	"jvm-on-go/instructions/store"
	"jvm-on-go/rtda"
)

/**
iload, fload, aload, lload, dload, istore, fstore, astore, lstore, dstore, ret, iinc
*/

func Wide(reader *base.ByteCodeReader, frame *rtda.Frame) {
	opCode := reader.ReadUint8()
	switch opCode {
	case opcodes.ILoad:
		load.DoILoad(uint16(reader.ReadUint8()), frame)
	case opcodes.FLoad:
		load.DoFLoad(uint16(reader.ReadUint8()), frame)
	case opcodes.ALoad:
		load.DoALoad(uint16(reader.ReadUint8()), frame)
	case opcodes.LLoad:
		load.DoLLoad(uint16(reader.ReadUint8()), frame)
	case opcodes.DLoad:
		load.DoDLoad(uint16(reader.ReadUint8()), frame)
	case opcodes.IStore:
		store.DoStoreInt(uint16(reader.ReadUint8()), frame)
	case opcodes.FStore:
		store.DoStoreFloat(uint16(reader.ReadUint8()), frame)
	case opcodes.AStore:
		store.DoStoreRef(uint16(reader.ReadUint8()), frame)
	case opcodes.LStore:
		store.DoStoreLong(uint16(reader.ReadUint8()), frame)
	case opcodes.DStore:
		store.DoStoreDouble(uint16(reader.ReadUint8()), frame)
	case opcodes.IInc:
		math.IInc(reader, frame)
	case opcodes.RET:
		panic("not supported op code: " + string(opCode))
	}
}
