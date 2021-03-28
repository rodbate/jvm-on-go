package lang

import (
	"jvm-on-go/constants/classname"
	"jvm-on-go/constants/descriptors"
	"jvm-on-go/instructions/base"
	"jvm-on-go/native"
	"jvm-on-go/rtda"
	"strings"
)

const (
	FieldConstructorDescriptor       = "(Ljava/lang/Class;Ljava/lang/String;Ljava/lang/Class;IILjava/lang/String;[B)V"
	ConstructorConstructorDescriptor = "(Ljava/lang/Class;[Ljava/lang/Class;[Ljava/lang/Class;IILjava/lang/String;[B[B)V"
)

func init() {
	native.RegisterNative(classname.Class, "getPrimitiveClass",
		"(Ljava/lang/String;)Ljava/lang/Class;", getPrimitiveClass)
	native.RegisterNative(classname.Class, "getName0",
		"()Ljava/lang/String;", getName)
	native.RegisterNative(classname.Class, "desiredAssertionStatus0",
		"(Ljava/lang/Class;)Z", desiredAssertionStatus)
	native.RegisterNative(classname.Class, "forName0",
		"(Ljava/lang/String;ZLjava/lang/ClassLoader;Ljava/lang/Class;)Ljava/lang/Class;", forName0)
	native.RegisterNative(classname.Class, "getDeclaredFields0",
		"(Z)[Ljava/lang/reflect/Field;", getDeclaredFields0)
	native.RegisterNative(classname.Class, "getSuperclass",
		"()Ljava/lang/Class;", getSuperclass)
	native.RegisterNative(classname.Class, "isPrimitive",
		"()Z", isPrimitive)
	native.RegisterNative(classname.Class, "isAssignableFrom",
		"(Ljava/lang/Class;)Z", isAssignableFrom)
	native.RegisterNative(classname.Class, "isInterface",
		"()Z", isInterface)
	native.RegisterNative(classname.Class, "getDeclaredConstructors0",
		"(Z)[Ljava/lang/reflect/Constructor;", getDeclaredConstructors0)
	native.RegisterNative(classname.Class, "getModifiers", "()I", getModifiers)
	native.RegisterNative(classname.Class, "isArray", "()Z", isArray)
	native.RegisterNative(classname.Class, "getComponentType", "()Ljava/lang/Class;", getComponentType)
	native.RegisterNative(classname.Class, "getEnclosingMethod0", "()[Ljava/lang/Object;", getEnclosingMethod0)
	native.RegisterNative(classname.Class, "getDeclaringClass0", "()Ljava/lang/Class;", getDeclaringClass0)
}

func getPrimitiveClass(frame *rtda.Frame) {
	className := frame.LocalVars.GetRef(0)
	class := frame.Method().Class().ClassLoader().LoadClass(rtda.GetGoString(className))
	frame.OperandStack.PushRef(class.JClass())
}

func getName(frame *rtda.Frame) {
	this := frame.LocalVars.GetThis()
	class := this.Extra.(*rtda.Class)
	frame.OperandStack.PushRef(rtda.GetJString(class.ClassLoader(), class.JavaName()))
}

func desiredAssertionStatus(frame *rtda.Frame) {
	frame.OperandStack.PushInt(0)
}

/**
private static native Class<?> forName0(String name, boolean initialize,
                                            ClassLoader loader,
                                            Class<?> caller)
*/
func forName0(frame *rtda.Frame) {
	className := rtda.GetGoString(frame.LocalVars.GetRef(0))
	class := frame.Method().Class().ClassLoader().LoadClass(strings.Replace(className, ".", "/", -1))
	frame.OperandStack.PushRef(class.JClass())
}

//private native Field[] getDeclaredFields0(boolean publicOnly);
func getDeclaredFields0(frame *rtda.Frame) {
	this := frame.LocalVars.GetThis()
	publicOnly := frame.LocalVars.GetBoolean(1)
	fields := this.Extra.(*rtda.Class).GetDeclaredFields(publicOnly)
	fieldCount := uint(len(fields))

	classLoader := frame.Method().Class().ClassLoader()
	fieldClass := classLoader.LoadClass(classname.Field)
	javaFields := fieldClass.ArrayClass().NewArray(fieldCount)
	frame.OperandStack.PushRef(javaFields)

	//init java fields
	if fieldCount > 0 {
		/**
		  Field(Class<?> declaringClass,
		        String name,
		        Class<?> type,
		        int modifiers,
		        int slot,
		        String signature,
		        byte[] annotations)
		*/
		fieldConstructor := fieldClass.GetConstructor(FieldConstructorDescriptor)
		thread := frame.Thread()
		javaFieldObjs := javaFields.Refs()
		for i, f := range fields {
			fieldObj := fieldClass.NewInstance()
			fieldObj.Extra = f
			javaFieldObjs[i] = fieldObj

			operandStack := rtda.NewOperandStack(8)
			operandStack.PushRef(fieldObj)
			operandStack.PushRef(this)
			operandStack.PushRef(rtda.GetJString(classLoader, f.Name()))
			operandStack.PushRef(f.Type().JClass())
			operandStack.PushInt(int32(f.AccessFlags))
			operandStack.PushInt(int32(f.SlotIndex()))
			operandStack.PushRef(getSignatureString(classLoader, f.Signature))
			operandStack.PushRef(native.ToJavaByteArray(classLoader, f.AnnotationData))

			mockFrame := rtda.NewMockFrame(thread, operandStack)
			thread.PushFrame(mockFrame)

			base.InvokeMethod(mockFrame, fieldConstructor)
		}
	}
}

func getSignatureString(classLoader *rtda.ClassLoader, signature string) *rtda.Object {
	if signature == "" {
		return nil
	}
	return rtda.GetJString(classLoader, signature)
}

//public native Class<? super T> getSuperclass();
func getSuperclass(frame *rtda.Frame) {
	classObj := frame.LocalVars.GetThis()
	class := classObj.Extra.(*rtda.Class)
	superClass := class.SuperClass()
	if superClass != nil {
		frame.OperandStack.PushRef(superClass.JClass())
	} else {
		frame.OperandStack.PushRef(nil)
	}
}

//public native boolean isPrimitive();
func isPrimitive(frame *rtda.Frame) {
	classObj := frame.LocalVars.GetThis()
	class := classObj.Extra.(*rtda.Class)
	_, exists := descriptors.PrimitiveTypeToArrayDescriptor[class.Name()]
	frame.OperandStack.PushBoolean(exists)
}

//public native boolean isAssignableFrom(Class<?> cls);
func isAssignableFrom(frame *rtda.Frame) {
	classObj := frame.LocalVars.GetThis()
	class := classObj.Extra.(*rtda.Class)
	targetClass := frame.LocalVars.GetRef(1).Extra.(*rtda.Class)
	assignable := class.IsAssignableFrom(targetClass)
	frame.OperandStack.PushBoolean(assignable)
}

//public native boolean isInterface();
func isInterface(frame *rtda.Frame) {
	class := frame.LocalVars.GetThis().Extra.(*rtda.Class)
	frame.OperandStack.PushBoolean(class.IsInterface())
}

//private native Constructor<T>[] getDeclaredConstructors0(boolean publicOnly)
func getDeclaredConstructors0(frame *rtda.Frame) {
	this := frame.LocalVars.GetThis()
	class := this.Extra.(*rtda.Class)
	publicOnly := frame.LocalVars.GetBoolean(1)
	constructors := class.GetDeclaredConstructors(publicOnly)
	constructorCount := uint(len(constructors))

	classLoader := frame.Method().Class().ClassLoader()
	constructorClass := classLoader.LoadClass(classname.Constructor)
	javaConstructors := constructorClass.ArrayClass().NewArray(constructorCount)
	frame.OperandStack.PushRef(javaConstructors)

	if constructorCount > 0 {
		/**
		Constructor(Class<T> declaringClass,
					Class<?>[] parameterTypes,
					Class<?>[] checkedExceptions,
					int modifiers,
					int slot,
					String signature,
					byte[] annotations,
					byte[] parameterAnnotations)
		*/
		constructor := constructorClass.GetConstructor(ConstructorConstructorDescriptor)
		thread := frame.Thread()
		javaConstructorObjs := javaConstructors.Refs()

		for i, ctr := range constructors {
			ctrObj := constructorClass.NewInstance()
			ctrObj.Extra = ctr
			javaConstructorObjs[i] = ctrObj

			stack := rtda.NewOperandStack(9)
			stack.PushRef(ctrObj)
			stack.PushRef(this)
			stack.PushRef(rtda.ToJavaClasses(classLoader, ctr.MethodDescriptor.ParameterTypes, true))
			stack.PushRef(rtda.ToJavaClasses(classLoader, ctr.Exceptions, false))
			stack.PushInt(int32(ctr.AccessFlags))
			stack.PushInt(0)
			stack.PushRef(rtda.GetJString(classLoader, ctr.Signature))
			stack.PushRef(native.ToJavaByteArray(classLoader, ctr.AnnotationData))
			stack.PushRef(native.ToJavaByteArray(classLoader, ctr.ParameterAnnotationData))

			mockFrame := rtda.NewMockFrame(thread, stack)
			thread.PushFrame(mockFrame)

			base.InvokeMethod(mockFrame, constructor)
		}
	}
}

//public native int getModifiers();
func getModifiers(frame *rtda.Frame) {
	this := frame.LocalVars.GetThis()
	accessFlags := this.Extra.(*rtda.Class).AccessFlags
	frame.OperandStack.PushInt(int32(accessFlags))
}

//public native boolean isArray()
func isArray(frame *rtda.Frame) {
	classObj := frame.LocalVars.GetThis()
	class := classObj.Extra.(*rtda.Class)
	frame.OperandStack.PushBoolean(class.IsArray())
}

//public native Class<?> getComponentType()
func getComponentType(frame *rtda.Frame) {
	classObj := frame.LocalVars.GetThis()
	class := classObj.Extra.(*rtda.Class)
	if !class.IsArray() {
		frame.OperandStack.PushRef(nil)
		return
	}
	frame.OperandStack.PushRef(class.ComponentClass().JClass())
}

//private native Object[] getEnclosingMethod0();
func getEnclosingMethod0(frame *rtda.Frame) {
	//todo: local class or anonymous class
	frame.OperandStack.PushRef(nil)
}

//private native Class<?> getDeclaringClass0()
func getDeclaringClass0(frame *rtda.Frame) {
	//todo: declaring class
	frame.OperandStack.PushRef(nil)
}