package rtda

import (
	"math"
)

type Slots []Slot

func (slots Slots) SetInt(index uint, val int32) {
	slots[index].num = val
}

func (slots Slots) GetInt(index uint) int32 {
	return slots[index].num
}

func (slots Slots) SetFloat(index uint, val float32) {
	slots[index].num = int32(math.Float32bits(val))
}

func (slots Slots) GetFloat(index uint) float32 {
	return math.Float32frombits(uint32(slots[index].num))
}

func (slots Slots) SetLong(index uint, val int64) {
	slots[index].num = int32(val)
	slots[index+1].num = int32(val >> 32)
}

func (slots Slots) GetLong(index uint) int64 {
	return int64(slots[index].num) | int64(slots[index+1].num)<<32
}

func (slots Slots) SetDouble(index uint, val float64) {
	slots.SetLong(index, int64(math.Float64bits(val)))
}

func (slots Slots) GetDouble(index uint) float64 {
	return math.Float64frombits(uint64(slots.GetLong(index)))
}

func (slots Slots) SetRef(index uint, ref *Object) {
	slots[index].ref = ref
}

func (slots Slots) GetRef(index uint) *Object {
	return slots[index].ref
}

func (slots Slots) SetSlot(index uint, slot Slot) {
	slots[index] = slot
}

func (slots Slots) GetThis() *Object {
	return slots.GetRef(0)
}

func (slots Slots) GetBoolean(index uint) bool {
	return slots.GetInt(index) == 1
}

func (slots Slots) SetBoolean(index uint, val bool) {
	var intVal int32
	if val {
		intVal = 1
	} else {
		intVal = 0
	}
	slots[index].num = intVal
}
