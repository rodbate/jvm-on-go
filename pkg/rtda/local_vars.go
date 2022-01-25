package rtda

func newLocalVars(maxLocals uint16) Slots {
	if maxLocals > 0 {
		return make(Slots, maxLocals)
	}
	return nil
}
