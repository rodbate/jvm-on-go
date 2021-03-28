package math

import (
	"github.com/rodbate/jvm-on-go/instructions/base"
	"github.com/rodbate/jvm-on-go/rtda"
)

func IInc(reader *base.ByteCodeReader, frame *rtda.Frame) {
	index := uint(reader.ReadUint8())
	consts := int32(reader.ReadInt8())
	val := frame.LocalVars.GetInt(index)
	frame.LocalVars.SetInt(index, val+consts)
}
