package rtda

import (
	"fmt"
	"github.com/rodbate/jvm-on-go/classfile"
	"github.com/rodbate/jvm-on-go/constants/classname"
	"github.com/rodbate/jvm-on-go/constants/descriptors"
	"strings"
)

type Class struct {
	AccessFlags       uint16
	name              string
	superClassName    string
	interfaceNames    []string
	constantPool      *ConstantPool
	fields            []*Field
	methods           []*Method
	classloader       *ClassLoader
	superClass        *Class
	interfaces        []*Class
	instanceSlotCount uint
	staticSlotCount   uint
	staticFields      Slots
	initStarted       bool
	jClass            *Object
	SourceFile        string
	Signature         string
	AnnotationData    []byte
}

func newClass(cf *classfile.ClassFile) *Class {
	class := &Class{}
	class.AccessFlags = cf.AccessFlags()
	class.name = cf.ClassName()
	class.superClassName = cf.SuperClassName()
	class.interfaceNames = cf.InterfaceNames()
	class.constantPool = newConstantPool(class, cf.ConstantPool())
	class.fields = newFields(class, cf.Fields())
	class.methods = newMethods(class, cf.Methods())
	class.SourceFile = cf.SourceFile()
	class.Signature = cf.Signature()
	class.AnnotationData = cf.RuntimeVisibleAnnotationData()
	return class
}

func (class *Class) IsPublic() bool {
	return class.AccessFlags&AccPublic != 0
}

func (class *Class) IsFinal() bool {
	return class.AccessFlags&AccFinal != 0
}

func (class *Class) IsSuper() bool {
	return class.AccessFlags&AccSuper != 0
}

func (class *Class) IsInterface() bool {
	return class.AccessFlags&AccInterface != 0
}

func (class *Class) IsAbstract() bool {
	return class.AccessFlags&AccAbstract != 0
}

func (class *Class) IsSynthetic() bool {
	return class.AccessFlags&AccSynthetic != 0
}

func (class *Class) IsAnnotation() bool {
	return class.AccessFlags&AccAnnotation != 0
}

func (class *Class) IsEnum() bool {
	return class.AccessFlags&AccEnum != 0
}

func (class *Class) IsAccessibleTo(other *Class) bool {
	return class.IsPublic() || class.GetPackageName() == other.GetPackageName()
}

func (class *Class) SuperClass() *Class {
	return class.superClass
}

func (class *Class) GetPackageName() string {
	if index := strings.LastIndex(class.name, "/"); index > 0 {
		return class.name[:index]
	}
	return ""
}

func (class *Class) Name() string {
	return class.name
}

func (class *Class) JavaName() string {
	return strings.Replace(class.name, "/", ".", -1)
}

func (class *Class) JClass() *Object {
	return class.jClass
}

func (class *Class) ClassLoader() *ClassLoader {
	return class.classloader
}

func (class *Class) GetField(filter func(*Field) bool) *Field {
	for c := class; c != nil; c = c.superClass {
		for _, f := range c.fields {
			if filter(f) {
				return f
			}
		}
	}
	return nil
}

func (class *Class) GetStaticField(name, descriptor string) *Field {
	return class.GetField(func(field *Field) bool {
		return field.IsStatic() && field.name == name && descriptor == field.descriptor
	})
}

func (class *Class) IsSuperClassOf(targetClass *Class) bool {
	if class.IsInterface() || targetClass.IsInterface() {
		return false
	}
	return class.IsAssignableFrom(targetClass) && class != targetClass
}

func (class *Class) IsAssignableFrom(targetClass *Class) bool {
	if targetClass == nil {
		return false
	}
	if class.name == classname.Object {
		return true
	}
	if class == targetClass {
		return true
	}
	if !class.IsInterface() && targetClass.IsInterface() {
		return false
	}

	if targetClass.IsArray() {
		if !class.IsArray() {
			return class.name == classname.Cloneable || class.name == classname.Serializable
		} else {
			componentClass := class.ComponentClass()
			targetComponentClass := targetClass.ComponentClass()
			return componentClass.IsAssignableFrom(targetComponentClass)
		}
	} else if class.IsArray() {
		return false
	}

	searchSuperClass := !targetClass.IsInterface()
	searchInterface := class.IsInterface() || targetClass.IsInterface()
	return isAssignableFrom(class, targetClass, searchSuperClass, searchInterface, make(map[*Class]bool))
}

func isAssignableFrom(class *Class, currentClass *Class, searchSuperClass bool,
	searchInterface bool, processedClass map[*Class]bool) bool {
	if currentClass == nil || processedClass[currentClass] {
		return false
	}
	processedClass[currentClass] = true
	if class == currentClass {
		return true
	}

	//find super class
	if searchSuperClass {
		matched := isAssignableFrom(class, currentClass.superClass, searchSuperClass, searchInterface, processedClass)
		if matched {
			return true
		}
	}

	if searchInterface {
		//find interfaces
		for _, iface := range currentClass.interfaces {
			matched := isAssignableFrom(class, iface, searchSuperClass, searchInterface, processedClass)
			if matched {
				return true
			}
		}
	}

	//not matched
	return false
}

func (class *Class) NewInstance() *Object {
	return newObject(class)
}

func (class *Class) ConstantPool() *ConstantPool {
	return class.constantPool
}

func (class *Class) StaticFields() Slots {
	return class.staticFields
}

func (class *Class) GetStaticMethod(name string, descriptor string) *Method {
	for _, m := range class.methods {
		if m.IsStatic() && name == m.name && descriptor == m.descriptor {
			return m
		}
	}
	return nil
}

func (class *Class) GetInstanceMethod(name string, descriptor string) *Method {
	for _, m := range class.methods {
		if !m.IsStatic() && name == m.name && descriptor == m.descriptor {
			return m
		}
	}
	return nil
}

func (class *Class) GetMainMethod() *Method {
	return class.GetStaticMethod("main", "([Ljava/lang/String;)V")
}

func (class *Class) GetClinitMethod() *Method {
	return class.GetStaticMethod("<clinit>", "()V")
}

func (class *Class) GetConstructor(descriptor string) *Method {
	return class.GetInstanceMethod("<init>", descriptor)
}

func (class *Class) LookupMethod(name, descriptor string) *Method {
	return lookupMethod(class, name, descriptor, make(map[*Class]bool))
}

func lookupMethod(class *Class, name, descriptor string, processedClass map[*Class]bool) *Method {
	if class == nil || processedClass[class] {
		return nil
	}
	processedClass[class] = true
	for _, m := range class.methods {
		if m.name == name && m.descriptor == descriptor {
			return m
		}
	}
	//find in super class
	method := lookupMethod(class.superClass, name, descriptor, processedClass)
	if method != nil {
		return method
	}

	//find in interfaces
	for _, c := range class.interfaces {
		method = lookupMethod(c, name, descriptor, processedClass)
		if method != nil {
			return method
		}
	}

	//not found
	return nil
}

func (class *Class) GetDeclaredFields(publicOnly bool) []*Field {
	if !publicOnly {
		return class.fields
	}
	publicFields := make([]*Field, 0, len(class.fields))
	for _, field := range class.fields {
		if field.IsPublic() {
			publicFields = append(publicFields, field)
		}
	}
	return publicFields
}

func (class *Class) GetDeclaredConstructors(publicOnly bool) []*Method {
	constructors := class.GetAllDeclaredConstructors()
	if !publicOnly {
		return constructors
	}

	publicConstructors := make([]*Method, 0, len(constructors))
	for _, method := range constructors {
		if method.IsPublic() {
			publicConstructors = append(publicConstructors, method)
		}
	}
	return publicConstructors
}

func (class *Class) GetAllDeclaredConstructors() []*Method {
	constructors := make([]*Method, 0, len(class.methods))
	for _, method := range class.methods {
		if method.name == "<init>" {
			constructors = append(constructors, method)
		}
	}
	return constructors
}

func (class *Class) LookupInterfaceMethod(name, descriptor string) *Method {
	return lookupInterfaceMethod(class, name, descriptor, make(map[*Class]bool))
}

func lookupInterfaceMethod(iface *Class, name, descriptor string, processedClass map[*Class]bool) *Method {
	if iface == nil || processedClass[iface] {
		return nil
	}
	processedClass[iface] = true
	for _, m := range iface.methods {
		if m.name == name && m.descriptor == descriptor {
			return m
		}
	}
	for _, superInterface := range iface.interfaces {
		if method := lookupInterfaceMethod(superInterface, name, descriptor, processedClass); method != nil {
			return method
		}
	}

	//not found
	return nil
}

func (class *Class) InitStarted() bool {
	return class.initStarted
}

func (class *Class) StartInit() {
	class.initStarted = true
}

func (class *Class) IsArray() bool {
	return strings.HasPrefix(class.name, descriptors.ArrayPrefix)
}

func (class *Class) IsPrimitiveType() bool {
	_, ok := descriptors.PrimitiveTypeToArrayDescriptor[class.name]
	return ok
}

func (class *Class) NewArray(length uint) *Object {
	if !class.IsArray() {
		panic(fmt.Sprintf("class %v is not array", class.Name()))
	}
	switch class.name {
	case descriptors.ArrayBoolean:
		return &Object{class: class, Data: make(ArrayBoolean, length)}
	case descriptors.ArrayByte:
		return &Object{class: class, Data: make(ArrayByte, length)}
	case descriptors.ArrayShort:
		return &Object{class: class, Data: make(ArrayShort, length)}
	case descriptors.ArrayChar:
		return &Object{class: class, Data: make(ArrayChar, length)}
	case descriptors.ArrayInt:
		return &Object{class: class, Data: make(ArrayInt, length)}
	case descriptors.ArrayLong:
		return &Object{class: class, Data: make(ArrayLong, length)}
	case descriptors.ArrayFloat:
		return &Object{class: class, Data: make(ArrayFloat, length)}
	case descriptors.ArrayDouble:
		return &Object{class: class, Data: make(ArrayDouble, length)}
	default:
		return &Object{class: class, Data: make(ArrayRef, length)}
	}
}

func (class *Class) ArrayClass() *Class {
	return class.classloader.LoadClass(getArrayClassName(class.name))
}

func (class *Class) ComponentClass() *Class {
	return class.classloader.LoadClass(getComponentClassName(class.name))
}

func NewByteArray(loader *ClassLoader, bytes []int8) *Object {
	return &Object{loader.LoadClass("[B"), bytes, nil, nil}
}
