package classfile

import "encoding/binary"

type ClassReader struct {
	data []byte
}

func (cr *ClassReader) readUint8() uint8 {
	v := cr.data[0]
	cr.data = cr.data[1:]
	return v
}

func (cr *ClassReader) readUint16() uint16 {
	v := binary.BigEndian.Uint16(cr.data)
	cr.data = cr.data[2:]
	return v
}

func (cr *ClassReader) readUint32() uint32 {
	v := binary.BigEndian.Uint32(cr.data)
	cr.data = cr.data[4:]
	return v
}

func (cr *ClassReader) readUint64() uint64 {
	v := binary.BigEndian.Uint64(cr.data)
	cr.data = cr.data[8:]
	return v
}

func (cr *ClassReader) readUint16s() []uint16 {
	n := cr.readUint16()
	data := make([]uint16, n)
	for i := range data {
		data[i] = cr.readUint16()
	}
	return data
}

func (cr *ClassReader) readBytes(n uint32) []byte {
	data := cr.data[:n]
	cr.data = cr.data[n:]
	return data
}
