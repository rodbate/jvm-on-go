package rtda

const (
	AccPublic       = 0x0001 //class field method
	AccPrivate      = 0x0002 //field method
	AccProtected    = 0x0004 //field method
	AccStatic       = 0x0008 //field method
	AccFinal        = 0x0010 //class filed method
	AccSuper        = 0x0020 //class
	AccSynchronized = 0x0020 //method
	AccVolatile     = 0x0040 //field
	AccBridge       = 0x0040 //method
	AccTransient    = 0x0080 //field
	AccVarArgs      = 0x0080 //method
	AccNative       = 0x0100 //method
	AccInterface    = 0x0200 //class
	AccAbstract     = 0x0400 //class method
	AccStrict       = 0x0800 //method
	AccSynthetic    = 0x1000 //class field method
	AccAnnotation   = 0x2000 //class
	AccEnum         = 0x4000 //class field
)
