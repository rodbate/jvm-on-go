package rtda

import "jvm-on-go/classfile"

type ExceptionTable []*ExceptionHandler

type ExceptionHandler struct {
	StartPc   uint16
	EndPc     uint16
	HandlerPc uint16
	CatchType *ClassRef
}

func newExceptionTable(cp *ConstantPool,
	exceptionTable []*classfile.ExceptionTableEntry) ExceptionTable {
	table := make([]*ExceptionHandler, len(exceptionTable))
	for i, entry := range exceptionTable {
		table[i] = &ExceptionHandler{
			StartPc:   entry.StartPc,
			EndPc:     entry.EndPc,
			HandlerPc: entry.HandlerPc,
			CatchType: getCatchType(cp, entry.CatchType),
		}
	}
	return table
}

func getCatchType(cp *ConstantPool, catchTypeIndex uint16) *ClassRef {
	if catchTypeIndex == 0 {
		//catch all ex -> throwable
		return nil
	}
	return cp.GetConstant(catchTypeIndex).(*ClassRef)
}

func (table ExceptionTable) FindExceptionHandler(exceptionClass *Class, pc uint32) *ExceptionHandler {
	for _, handler := range table {
		if uint32(handler.StartPc) <= pc && uint32(handler.EndPc) > pc {
			if handler.CatchType == nil {
				return handler
			}
			catchClass := handler.CatchType.ResolvedClass()
			if catchClass == exceptionClass || catchClass.IsSuperClassOf(exceptionClass) {
				return handler
			}
		}
	}
	return nil
}
