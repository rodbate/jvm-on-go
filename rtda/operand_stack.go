package rtda

import (
	"math"
)

type OperandStack struct {
	topIndex uint16
	slots    []Slot
}

func NewOperandStack(maxStack uint16) *OperandStack {
	if maxStack > 0 {
		return &OperandStack{
			topIndex: 0,
			slots:    make([]Slot, maxStack),
		}
	}
	return nil
}

func (s *OperandStack) PushInt(val int32) {
	s.slots[s.topIndex].num = val
	s.topIndex++
}

func (s *OperandStack) PopInt() int32 {
	s.topIndex--
	return s.slots[s.topIndex].num
}

func (s *OperandStack) PushFloat(val float32) {
	s.slots[s.topIndex].num = int32(math.Float32bits(val))
	s.topIndex++
}

func (s *OperandStack) PopFloat() float32 {
	s.topIndex--
	bits := s.slots[s.topIndex].num
	return math.Float32frombits(uint32(bits))
}

func (s *OperandStack) PushLong(val int64) {
	s.slots[s.topIndex].num = int32(val)
	s.topIndex++
	s.slots[s.topIndex].num = int32(val >> 32)
	s.topIndex++
}

func (s *OperandStack) PopLong() int64 {
	s.topIndex--
	highBits := uint32(s.slots[s.topIndex].num)
	s.topIndex--
	lowBits := uint32(s.slots[s.topIndex].num)
	return int64(highBits)<<32 | int64(lowBits)
}

func (s *OperandStack) PushDouble(val float64) {
	s.PushLong(int64(math.Float64bits(val)))
}

func (s *OperandStack) PopDouble() float64 {
	return math.Float64frombits(uint64(s.PopLong()))
}

func (s *OperandStack) PushRef(ref *Object) {
	s.slots[s.topIndex].ref = ref
	s.topIndex++
}

func (s *OperandStack) PopRef() *Object {
	s.topIndex--
	ref := s.slots[s.topIndex].ref
	s.slots[s.topIndex].ref = nil
	return ref
}

func (s *OperandStack) PushSlot(slot Slot) {
	s.slots[s.topIndex] = slot
	s.topIndex++
}

func (s *OperandStack) PopSlot() Slot {
	s.topIndex--
	slot := s.slots[s.topIndex]
	return slot
}

func (s *OperandStack) PeekSlot() Slot {
	return s.slots[s.topIndex-1]
}

func (s *OperandStack) GetRefFromTop(index uint16) *Object {
	return s.slots[s.topIndex-index-1].ref
}

func (s *OperandStack) Clear() {
	s.topIndex = 0
	for _, slot := range s.slots {
		slot.ref = nil
	}
}

func (s *OperandStack) PushBoolean(val bool) {
	var value int32
	if val {
		value = 1
	} else {
		value = 0
	}
	s.PushInt(value)
}

func (s *OperandStack) PopBoolean() bool {
	value := s.PopInt()
	var val bool
	if value == 1 {
		val = true
	} else {
		val = false
	}
	return val
}
