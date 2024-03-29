package opcodes

const (
	NOP uint8 = iota
	AConstNull
	IConstM1
	IConst0
	IConst1
	IConst2
	IConst3
	IConst4
	IConst5
	LConst0
	LConst1
	FConst0
	FConst1
	FConst2
	DConst0
	DConst1
	BIPush
	SIPush
	LDC
	LDCW
	LDC2W
	ILoad
	LLoad
	FLoad
	DLoad
	ALoad
	ILoad0
	ILoad1
	ILoad2
	ILoad3
	LLoad0
	LLoad1
	LLoad2
	LLoad3
	FLoad0
	FLoad1
	FLoad2
	FLoad3
	DLoad0
	DLoad1
	DLoad2
	DLoad3
	ALoad0
	ALoad1
	ALoad2
	ALoad3
	IALoad
	LALoad
	FALoad
	DALoad
	AALoad
	BALoad
	CALoad
	SALoad
	IStore
	LStore
	FStore
	DStore
	AStore
	IStore0
	IStore1
	IStore2
	IStore3
	LStore0
	LStore1
	LStore2
	LStore3
	FStore0
	FStore1
	FStore2
	FStore3
	DStore0
	DStore1
	DStore2
	DStore3
	AStore0
	AStore1
	AStore2
	AStore3
	IAStore
	LAStore
	FAStore
	DAStore
	AAStore
	BAStore
	CAStore
	SAStore
	POP
	POP2
	DUP
	DUPX1
	DUPX2
	DUP2
	DUP2X1
	DUP2X2
	SWAP
	IAdd
	LAdd
	FAdd
	DAdd
	ISub
	LSub
	FSub
	DSub
	IMul
	LMul
	FMul
	DMul
	IDiv
	LDiv
	FDiv
	DDiv
	IRem
	LRem
	FRem
	DRem
	INeg
	LNeg
	FNeg
	DNeg
	IShl
	LShl
	IShr
	LShr
	IUShr
	LUShr
	IAnd
	LAnd
	IOr
	LOr
	IXor
	LXor
	IInc
	I2L
	I2F
	I2D
	L2I
	L2F
	L2D
	F2I
	F2L
	F2D
	D2I
	D2L
	D2F
	I2B
	I2C
	I2S
	LCmp
	FCmpl
	FCmpg
	DCmpl
	DCmpg
	IFeq
	IFne
	IFlt
	IFge
	IFgt
	IFle
	IfICmpEq
	IfICmpNe
	IfICmpLt
	IfICmpGe
	IfICmpGt
	IfICmpLe
	IfACmpEq
	IfACmpNe
	Goto
	JSR
	RET
	TableSwitch
	LookupSwitch
	IReturn
	LReturn
	FReturn
	DReturn
	AReturn
	Return
	GetStatic
	PutStatic
	GetField
	PutField
	InvokeVirtual
	InvokeSpecial
	InvokeStatic
	InvokeInterface
	InvokeDynamic
	New
	NewArray
	ANewArray
	ArrayLength
	AThrow
	CheckCast
	Instanceof
	MonitorEnter
	MonitorExit
	Wide
	MultiANewArray
	IfNull
	IfNonnull
	GotoW
	JsrW
	BreakPoint
	InvokeNative = 0xFE //custom impl instruction with reserved op code
	ImpDep2      = 0xFF
)
