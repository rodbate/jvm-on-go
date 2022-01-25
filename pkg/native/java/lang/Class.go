package lang

import (
	"github.com/rodbate/jvm-on-go/pkg/constants/classname"
	"github.com/rodbate/jvm-on-go/pkg/constants/descriptors"
	"github.com/rodbate/jvm-on-go/pkg/instructions/base"
	native2 "github.com/rodbate/jvm-on-go/pkg/native"
	rtda2 "github.com/rodbate/jvm-on-go/pkg/rtda"
	"strings"
)

const (
	FieldConstructorDescriptor       = "(Ljava/lang/Class;Ljava/lang/String;Ljava/lang/Class;IILjava/lang/String;[B)V"
	ConstructorConstructorDescriptor = "(Ljava/lang/Class;[Ljava/lang/Class;[Ljava/lang/Class;IILjava/lang/String;[B[B)V"
)

func init() {
	native2.RegisterNative(classname.Class, "getPrimitiveClass",
		"(Ljava/lang/String;)Ljava/lang/Class;", getPrimitiveClass)
	native2.RegisterNative(classname.Class, "getName0",
		"()Ljava/lang/String;", getName)
	native2.RegisterNative(classname.Class, "desiredAssertionStatus0",
		"(Ljava/lang/Class;)Z", desiredAssertionStatus)
	native2.RegisterNative(classname.Class, "forName0",
		"(Ljava/lang/String;ZLjava/lang/ClassLoader;Ljava/lang/Class;)Ljava/lang/Class;", forName0)
	native2.RegisterNative(classname.Class, "getDeclaredFields0",
		"(Z)[Ljava/lang/reflect/Field;", getDeclaredFields0)
	native2.RegisterNative(classname.Class, "getSuperclass",
		"()Ljava/lang/Class;", getSuperclass)
	native2.RegisterNative(classname.Class, "isPrimitive",
		"()Z", isPrimitive)
	native2.RegisterNative(classname.Class, "isAssignableFrom",
		"(Ljava/lang/Class;)Z", isAssignableFrom)
	native2.RegisterNative(classname.Class, "isInterface",
		"()Z", isInterface)
	native2.RegisterNative(classname.Class, "getDeclaredConstructors0",
		"(Z)[Ljava/lang/reflect/Constructor;", getDeclaredConstructors0)
	native2.RegisterNative(classname.Class, "getModifiers", "()I", getModifiers)
	native2.RegisterNative(classname.Class, "isArray", "()Z", isArray)
	native2.RegisterNative(classname.Class, "getComponentType", "()Ljava/lang/Class;", getComponentType)
	native2.RegisterNative(classname.Class, "getEnclosingMethod0", "()[Ljava/lang/Object;", getEnclosingMethod0)
	native2.RegisterNative(classname.Class, "getDeclaringClass0", "()Ljava/lang/Class;", getDeclaringClass0)
}

func getPrimitiveClass(frame *rtda2.Frame) {
	className := frame.LocalVars.GetRef(0)
	class := frame.Method().Class().ClassLoader().LoadClass(rtda2.GetGoString(className))
	frame.OperandStack.PushRef(class.JClass())
}

func getName(frame *rtda2.Frame) {
	this := frame.LocalVars.GetThis()
	class := this.Extra.(*rtda2.Class)
	frame.OperandStack.PushRef(rtda2.GetJString(class.ClassLoader(), class.JavaName()))
}

func desiredAssertionStatus(frame *rtda2.Frame) {
	frame.OperandStack.PushInt(0)
}

/**
private static native Class<?> forName0(String name, boolean initialize,
                                            ClassLoader loader,
                                            Class<?> caller)
*/
func forName0(frame *rtda2.Frame) {
	className := rtda2.GetGoString(frame.LocalVars.GetRef(0))
	class := frame.Method().Class().ClassLoader().LoadClass(strings.Replace(className, ".", "/", -1))
	frame.OperandStack.PushRef(class.JClass())
}

//private native Field[] getDeclaredFields0(boolean publicOnly);
func getDeclaredFields0(frame *rtda2.Frame) {
	this := frame.LocalVars.GetThis()
	publicOnly := frame.LocalVars.GetBoolean(1)
	fields := this.Extra.(*rtda2.Class).GetDeclaredFields(publicOnly)
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

			operandStack := rtda2.NewOperandStack(8)
			operandStack.PushRef(fieldObj)
			operandStack.PushRef(this)
			operandStack.PushRef(rtda2.GetJString(classLoader, f.Name()))
			operandStack.PushRef(f.Type().JClass())
			operandStack.PushInt(int32(f.AccessFlags))
			operandStack.PushInt(int32(f.SlotIndex()))
			operandStack.PushRef(getSignatureString(classLoader, f.Signature))
			operandStack.PushRef(native2.ToJavaByteArray(classLoader, f.AnnotationData))

			mockFrame := rtda2.NewMockFrame(thread, operandStack)
			thread.PushFrame(mockFrame)

			base.InvokeMethod(mockFrame, fieldConstructor)
		}
	}
}

func getSignatureString(classLoader *rtda2.ClassLoader, signature string) *rtda2.Object {
	if signature == "" {
		return nil
	}
	return rtda2.GetJString(classLoader, signature)
}

//public native Class<? super T> getSuperclass();
func getSuperclass(frame *rtda2.Frame) {
	classObj := frame.LocalVars.GetThis()
	class := classObj.Extra.(*rtda2.Class)
	superClass := class.SuperClass()
	if superClass != nil {
		frame.OperandStack.PushRef(superClass.JClass())
	} else {
		frame.OperandStack.PushRef(nil)
	}
}

//public native boolean isPrimitive();
func isPrimitive(frame *rtda2.Frame) {
	classObj := frame.LocalVars.GetThis()
	class := classObj.Extra.(*rtda2.Class)
	_, exists := descriptors.PrimitiveTypeToArrayDescriptor[class.Name()]
	frame.OperandStack.PushBoolean(exists)
}

//public native boolean isAssignableFrom(Class<?> cls);
func isAssignableFrom(frame *rtda2.Frame) {
	classObj := frame.LocalVars.GetThis()
	class := classObj.Extra.(*rtda2.Class)
	targetClass := frame.LocalVars.GetRef(1).Extra.(*rtda2.Class)
	assignable := class.IsAssignableFrom(targetClass)
	frame.OperandStack.PushBoolean(assignable)
}

//public native boolean isInterface();
func isInterface(frame *rtda2.Frame) {
	class := frame.LocalVars.GetThis().Extra.(*rtda2.Class)
	frame.OperandStack.PushBoolean(class.IsInterface())
}

//private native Constructor<T>[] getDeclaredConstructors0(boolean publicOnly)
func getDeclaredConstructors0(frame *rtda2.Frame) {
	this := frame.LocalVars.GetThis()
	class := this.Extra.(*rtda2.Class)
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

			stack := rtda2.NewOperandStack(9)
			stack.PushRef(ctrObj)
			stack.PushRef(this)
			stack.PushRef(rtda2.ToJavaClasses(classLoader, ctr.MethodDescriptor.ParameterTypes, true))
			stack.PushRef(rtda2.ToJavaClasses(classLoader, ctr.Exceptions, false))
			stack.PushInt(int32(ctr.AccessFlags))
			stack.PushInt(0)
			stack.PushRef(rtda2.GetJString(classLoader, ctr.Signature))
			stack.PushRef(native2.ToJavaByteArray(classLoader, ctr.AnnotationData))
			stack.PushRef(native2.ToJavaByteArray(classLoader, ctr.ParameterAnnotationData))

			mockFrame := rtda2.NewMockFrame(thread, stack)
			thread.PushFrame(mockFrame)

			base.InvokeMethod(mockFrame, constructor)
		}
	}
}

//public native int getModifiers();
func getModifiers(frame *rtda2.Frame) {
	this := frame.LocalVars.GetThis()
	accessFlags := this.Extra.(*rtda2.Class).AccessFlags
	frame.OperandStack.PushInt(int32(accessFlags))
}

//public native boolean isArray()
func isArray(frame *rtda2.Frame) {
	classObj := frame.LocalVars.GetThis()
	class := classObj.Extra.(*rtda2.Class)
	frame.OperandStack.PushBoolean(class.IsArray())
}

//public native Class<?> getComponentType()
func getComponentType(frame *rtda2.Frame) {
	classObj := frame.LocalVars.GetThis()
	class := classObj.Extra.(*rtda2.Class)
	if !class.IsArray() {
		frame.OperandStack.PushRef(nil)
		return
	}
	frame.OperandStack.PushRef(class.ComponentClass().JClass())
}

//private native Object[] getEnclosingMethod0();
func getEnclosingMethod0(frame *rtda2.Frame) {
	//todo: local class or anonymous class
	frame.OperandStack.PushRef(nil)
}

//private native Class<?> getDeclaringClass0()
func getDeclaringClass0(frame *rtda2.Frame) {
	//todo: declaring class
	frame.OperandStack.PushRef(nil)
}
