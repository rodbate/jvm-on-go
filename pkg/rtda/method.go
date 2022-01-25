package rtda

import (
	"fmt"
	classfile2 "github.com/rodbate/jvm-on-go/pkg/classfile"
	"github.com/rodbate/jvm-on-go/pkg/constants/descriptors"
	"github.com/rodbate/jvm-on-go/pkg/instructions/opcodes"
)

type Method struct {
	Member
	maxStack                uint16
	maxLocals               uint16
	code                    []byte
	argSlotCount            uint
	exceptionTable          ExceptionTable
	lineNumberTable         *classfile2.LineNumberTableAttribute
	MethodDescriptor        *MethodDescriptor
	Exceptions              []string
	ParameterAnnotationData []byte
}

func newMethod(class *Class, methodInfo *classfile2.MemberInfo) *Method {
	method := &Method{}
	method.class = class
	method.initInfo(methodInfo)
	codeAttr := methodInfo.CodeAttribute()
	if codeAttr != nil {
		method.maxStack = codeAttr.MaxStack()
		method.maxLocals = codeAttr.MaxLocals()
		method.code = codeAttr.Code()
		method.exceptionTable = newExceptionTable(class.constantPool, codeAttr.ExceptionTable())
		method.lineNumberTable = codeAttr.LineNumberTableAttribute()
		method.Exceptions = methodInfo.Exceptions()
		method.ParameterAnnotationData = methodInfo.RuntimeVisibleParameterAnnotationData()
	}
	method.MethodDescriptor = ParseMethodDescriptor(method)
	method.calculateArgSlotCount()
	if method.IsNative() {
		method.rewriteNativeMethodCode()
	}
	return method
}

func (m *Method) calculateArgSlotCount() {
	argSlotCount := uint(0)
	for _, paramType := range m.MethodDescriptor.ParameterTypes {
		argSlotCount++
		if paramType == descriptors.Long || paramType == descriptors.Double {
			argSlotCount++
		}
	}
	if !m.IsStatic() {
		argSlotCount++
	}
	m.argSlotCount = argSlotCount
}

func (m *Method) rewriteNativeMethodCode() {
	returnType := m.MethodDescriptor.ReturnType
	m.maxStack = 4
	m.maxLocals = uint16(m.argSlotCount)
	switch string(returnType[0]) {
	case descriptors.Void:
		m.code = []byte{opcodes.InvokeNative, opcodes.Return}
	case descriptors.Boolean, descriptors.Byte, descriptors.Char, descriptors.Short, descriptors.Int:
		m.code = []byte{opcodes.InvokeNative, opcodes.IReturn}
	case descriptors.Float:
		m.code = []byte{opcodes.InvokeNative, opcodes.FReturn}
	case descriptors.Double:
		m.code = []byte{opcodes.InvokeNative, opcodes.DReturn}
	case descriptors.Long:
		m.code = []byte{opcodes.InvokeNative, opcodes.LReturn}
	case descriptors.ArrayPrefix, descriptors.RefPrefix:
		m.code = []byte{opcodes.InvokeNative, opcodes.AReturn}
	default:
		panic(fmt.Sprintf("invalid return type: %v", returnType))
	}
}

func newMethods(class *Class, methodInfos []*classfile2.MemberInfo) []*Method {
	methods := make([]*Method, len(methodInfos))
	for i := range methodInfos {
		methods[i] = newMethod(class, methodInfos[i])
	}
	return methods
}

func (m *Method) IsSynchronized() bool {
	return m.AccessFlags&AccSynchronized != 0
}

func (m *Method) IsBridged() bool {
	return m.AccessFlags&AccBridge != 0
}

func (m *Method) IsVarArgs() bool {
	return m.AccessFlags&AccVarArgs != 0
}

func (m *Method) IsNative() bool {
	return m.AccessFlags&AccNative != 0
}

func (m *Method) IsAbstract() bool {
	return m.AccessFlags&AccAbstract != 0
}

func (m *Method) IsStrict() bool {
	return m.AccessFlags&AccStrict != 0
}

func (m *Method) Code() []byte {
	return m.code
}

func (m *Method) ArgSlotCount() uint {
	return m.argSlotCount
}

func (m *Method) FindExceptionHandlerPc(exceptionClass *Class, pc uint32) uint32 {
	handler := m.exceptionTable.FindExceptionHandler(exceptionClass, pc)
	if handler == nil {
		return 0
	}
	return uint32(handler.HandlerPc)
}

func (m *Method) GetLineNumber(pc uint32) int {
	if m.lineNumberTable == nil {
		return -1
	}
	return m.lineNumberTable.GetLineNumber(uint16(pc))
}
