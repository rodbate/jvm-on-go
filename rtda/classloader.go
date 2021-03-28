package rtda

import (
	"fmt"
	"github.com/rodbate/jvm-on-go/classfile"
	"github.com/rodbate/jvm-on-go/classpath"
	"github.com/rodbate/jvm-on-go/constants/classname"
	"github.com/rodbate/jvm-on-go/constants/descriptors"
)

type ClassLoader struct {
	classPath *classpath.Classpath
	classData map[string]*Class
	verbose   bool
}

func NewClassLoader(classpath *classpath.Classpath, verbose bool) *ClassLoader {
	loader := &ClassLoader{
		classPath: classpath,
		classData: make(map[string]*Class),
		verbose:   verbose,
	}
	loader.loadBasicClasses()
	loader.loadPrimitiveClasses()
	return loader
}

func (cl *ClassLoader) loadBasicClasses() {
	jClass := cl.LoadClass(classname.Class)
	for _, class := range cl.classData {
		if class.jClass == nil {
			class.jClass = jClass.NewInstance()
			class.jClass.Extra = class
		}
	}
}

func (cl *ClassLoader) loadPrimitiveClasses() {
	for primitiveType := range descriptors.PrimitiveTypeToArrayDescriptor {
		cl.loadPrimitiveClass(primitiveType)
	}
}

func (cl *ClassLoader) loadPrimitiveClass(className string) {
	class := &Class{
		AccessFlags: AccPublic,
		name:        className,
		classloader: cl,
		initStarted: true,
	}

	class.jClass = cl.LoadClass(classname.Class).NewInstance()
	class.jClass.Extra = class
	cl.classData[className] = class
}

func (cl *ClassLoader) LoadClass(name string) *Class {
	if c, ok := cl.classData[name]; ok {
		return c
	}
	var class *Class
	if name[0] == '[' {
		class = cl.loadArrayClass(name)
	} else {
		class = cl.loadNormalClass(name)
	}
	if jClass, ok := cl.classData[classname.Class]; ok {
		class.jClass = jClass.NewInstance()
		class.jClass.Extra = class
	}
	return class
}

func (cl *ClassLoader) loadArrayClass(name string) *Class {
	class := &Class{
		AccessFlags: AccPublic,
		name:        name,
		classloader: cl,
		initStarted: true,
		superClass:  cl.LoadClass(classname.Object),
		interfaces: []*Class{
			cl.LoadClass(classname.Cloneable),
			cl.LoadClass(classname.Serializable),
		},
	}
	cl.classData[name] = class
	return class
}

func (cl *ClassLoader) loadNormalClass(name string) *Class {
	data, entry := cl.readClass(name)
	class := cl.defineClass(data)
	link(class)
	//init
	if cl.verbose {
		fmt.Printf("[Loaded class: %s from: %s]\n", name, entry)
	}
	return class
}

func (cl *ClassLoader) readClass(name string) ([]byte, classpath.Entry) {
	data, entry, err := cl.classPath.ReadClass(name)
	if err != nil {
		panic("java.lang.ClassNotFoundException: " + name)
	}
	return data, entry
}

func (cl *ClassLoader) defineClass(data []byte) *Class {
	class := cl.parseClass(data)
	resolveSuperClass(class)
	resolveInterfaces(class)
	cl.classData[class.name] = class
	return class
}

func (cl *ClassLoader) parseClass(data []byte) *Class {
	classFile, err := classfile.Parse(data)
	if err != nil {
		panic("ClassFormat error: " + err.Error())
	}
	class := newClass(classFile)
	class.classloader = cl
	return class
}

func (cl *ClassLoader) FindLoadedClass(name string) *Class {
	return cl.classData[name]
}

func resolveSuperClass(class *Class) {
	if class.name == classname.Object {
		return
	}
	class.superClass = class.classloader.LoadClass(class.superClassName)
}

func resolveInterfaces(class *Class) {
	if len(class.interfaceNames) <= 0 {
		return
	}
	class.interfaces = make([]*Class, len(class.interfaceNames))
	for i, name := range class.interfaceNames {
		class.interfaces[i] = class.classloader.LoadClass(name)
	}
}

func link(class *Class) {
	verify(class)
	prepare(class)
}

func verify(class *Class) {
	//todo verify class
}

func prepare(class *Class) {
	calculateInstanceFieldSlotsIndex(class)
	calculateStaticFieldSlotsIndex(class)
	allocateAndInitStaticFields(class)
}

func calculateInstanceFieldSlotsIndex(class *Class) {
	index := uint(0)
	if class.superClass != nil {
		index = class.superClass.instanceSlotCount
	}
	for _, f := range class.fields {
		if !f.IsStatic() {
			f.slotIndex = index
			index++
			if f.IsLongOrDouble() {
				index++
			}
		}
	}
	class.instanceSlotCount = index
}

func calculateStaticFieldSlotsIndex(class *Class) {
	index := uint(0)
	for _, f := range class.fields {
		if f.IsStatic() {
			f.slotIndex = index
			index++
			if f.IsLongOrDouble() {
				index++
			}
		}
	}
	class.staticSlotCount = index
}

func allocateAndInitStaticFields(class *Class) {
	class.staticFields = make(Slots, class.staticSlotCount)
	for _, f := range class.fields {
		if f.IsStatic() && f.IsFinal() {
			initStaticFinalFieldIfNecessary(class, f)
		}
	}
}

func initStaticFinalFieldIfNecessary(class *Class, field *Field) {
	fields := class.staticFields
	cp := class.constantPool
	fieldConstantValueIndex := field.constantValueIndex
	slotIndex := field.slotIndex
	if fieldConstantValueIndex > 0 {
		constant := cp.GetConstant(fieldConstantValueIndex)
		switch field.descriptor {
		case descriptors.Boolean, descriptors.Byte, descriptors.Short, descriptors.Char, descriptors.Int:
			fields.SetInt(slotIndex, constant.(int32))
		case descriptors.Float:
			fields.SetFloat(slotIndex, constant.(float32))
		case descriptors.Long:
			fields.SetLong(slotIndex, constant.(int64))
		case descriptors.Double:
			fields.SetDouble(slotIndex, constant.(float64))
		case descriptors.String:
			jStr := GetJString(class.classloader, constant.(string))
			fields.SetRef(slotIndex, jStr)
		}
	}
}
