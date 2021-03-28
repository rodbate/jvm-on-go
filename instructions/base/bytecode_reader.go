package base

type ByteCodeReader struct {
	bytes []byte
	pc    uint32
}

func NewByteCodeReader(byteCode []byte) *ByteCodeReader {
	return &ByteCodeReader{
		bytes: byteCode,
	}
}

func (r *ByteCodeReader) Reset(byteCode []byte, pc uint32) {
	r.bytes = byteCode
	r.pc = pc
}

func (r *ByteCodeReader) ReadUint8() uint8 {
	val := r.bytes[r.pc]
	r.pc++
	return val
}

func (r *ByteCodeReader) ReadInt8() int8 {
	return int8(r.ReadUint8())
}

func (r *ByteCodeReader) ReadUint16() uint16 {
	high := uint16(r.ReadUint8())
	low := uint16(r.ReadUint8())
	return (high << 8) | low
}

func (r *ByteCodeReader) ReadInt16() int16 {
	return int16(r.ReadUint16())
}

func (r *ByteCodeReader) ReadInt32() int32 {
	byte1 := int32(r.ReadUint8())
	byte2 := int32(r.ReadUint8())
	byte3 := int32(r.ReadUint8())
	byte4 := int32(r.ReadUint8())
	return byte1<<24 | byte2<<16 | byte3<<8 | byte4
}

func (r *ByteCodeReader) SkipPadding() {
	for r.pc%4 != 0 {
		r.ReadUint8()
	}
}

func (r *ByteCodeReader) PC() uint32 {
	return r.pc
}

func (r *ByteCodeReader) SetPc(pc uint32) {
	r.pc = pc
}
