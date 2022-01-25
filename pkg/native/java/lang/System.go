package lang

import (
	"github.com/rodbate/jvm-on-go/pkg/constants/classname"
	"github.com/rodbate/jvm-on-go/pkg/instructions/base"
	"github.com/rodbate/jvm-on-go/pkg/native"
	"github.com/rodbate/jvm-on-go/pkg/native/constants"
	rtda2 "github.com/rodbate/jvm-on-go/pkg/rtda"
	"os"
	"os/user"
	"path"
	"runtime"
	"time"
)

func init() {
	native.RegisterNative(classname.System, "arraycopy",
		"(Ljava/lang/Object;ILjava/lang/Object;II)V", arrayCopy)
	native.RegisterNative(classname.System, "setIn0",
		"(Ljava/io/InputStream;)V", setIn0)
	native.RegisterNative(classname.System, "setOut0",
		"(Ljava/io/PrintStream;)V", setOut0)
	native.RegisterNative(classname.System, "setErr0",
		"(Ljava/io/PrintStream;)V", setErr0)
	native.RegisterNative(classname.System, "initProperties",
		"(Ljava/util/Properties;)Ljava/util/Properties;", initProperties)
	native.RegisterNative(classname.System, "mapLibraryName",
		"(Ljava/lang/String;)Ljava/lang/String;", mapLibraryName)
	native.RegisterNative(classname.System, "currentTimeMillis",
		"()J", currentTimeMillis)
	native.RegisterNative(classname.System, "nanoTime",
		"()J", nanoTime)
}

/**
public static native
void arraycopy(Object src,  int  srcPos,
			   Object dest, int destPos,
			   int length);
*/
func arrayCopy(frame *rtda2.Frame) {
	src := frame.LocalVars.GetRef(0)
	srcPos := frame.LocalVars.GetInt(1)
	dest := frame.LocalVars.GetRef(2)
	destPos := frame.LocalVars.GetInt(3)
	length := frame.LocalVars.GetInt(4)

	if src == nil || dest == nil {
		panic("java.lang.NullPointerException")
	}
	if !checkArrayStore(src, dest) {
		panic("java.lang.ArrayStoreException")
	}
	if srcPos < 0 || destPos < 0 || length < 0 ||
		srcPos+length > src.ArrayLength() || destPos+length > dest.ArrayLength() {
		panic("java.lang.IndexOutOfBoundsException")
	}
	rtda2.ArrayCopy(src, dest, srcPos, destPos, length)
}

func checkArrayStore(src, dest *rtda2.Object) bool {
	srcClass := src.Class()
	destClass := dest.Class()
	if !srcClass.IsArray() || !destClass.IsArray() {
		return false
	}
	if srcClass.ComponentClass().IsPrimitiveType() &&
		destClass.ComponentClass().IsPrimitiveType() {
		return srcClass == destClass
	}
	return !srcClass.ComponentClass().IsPrimitiveType() &&
		!destClass.ComponentClass().IsPrimitiveType()
}

//private static native void setIn0(InputStream in);
func setIn0(frame *rtda2.Frame) {
	in := frame.LocalVars.GetRef(0)
	class := frame.Method().Class()
	field := class.GetStaticField("in", "Ljava/io/InputStream;")
	class.StaticFields().SetRef(field.SlotIndex(), in)
}

//private static native void setOut0(PrintStream out);
func setOut0(frame *rtda2.Frame) {
	out := frame.LocalVars.GetRef(0)
	class := frame.Method().Class()
	field := class.GetStaticField("out", "Ljava/io/PrintStream;")
	class.StaticFields().SetRef(field.SlotIndex(), out)
}

//private static native void setErr0(PrintStream err);
func setErr0(frame *rtda2.Frame) {
	out := frame.LocalVars.GetRef(0)
	class := frame.Method().Class()
	field := class.GetStaticField("err", "Ljava/io/PrintStream;")
	class.StaticFields().SetRef(field.SlotIndex(), out)
}

//private static native Properties initProperties(Properties props);
/**
 * System properties. The following properties are guaranteed to be defined:
 * <dl>
 * <dt>java.version         <dd>Java version number
 * <dt>java.vendor          <dd>Java vendor specific string
 * <dt>java.vendor.url      <dd>Java vendor URL
 * <dt>java.home            <dd>Java installation directory
 * <dt>java.class.version   <dd>Java class version number
 * <dt>java.class.path      <dd>Java classpath
 * <dt>os.name              <dd>Operating System Name
 * <dt>os.arch              <dd>Operating System Architecture
 * <dt>os.version           <dd>Operating System Version
 * <dt>file.separator       <dd>File separator ("/" on Unix)
 * <dt>path.separator       <dd>Path separator (":" on Unix)
 * <dt>line.separator       <dd>Line separator ("\n" on Unix)
 * <dt>user.name            <dd>User account name
 * <dt>user.home            <dd>User home directory
 * <dt>user.dir             <dd>User's current working directory
 * </dl>
 */
var systemProperties = map[string]string{
	"java.version":          "1.8",
	"java.vendor":           "rodbate-jvm",
	"java.vendor.url":       "https://github.com/rodbate/jvm-in-go",
	"java.home":             os.Getenv("JAVA_HOME"),
	"java.class.version":    "52",
	"java.class.path":       os.Getenv("CLASSPATH"),
	"os.name":               runtime.GOOS,
	"os.arch":               runtime.GOARCH,
	"os.version":            "",
	"file.encoding":         "UTF-8",
	"file.separator":        string(os.PathSeparator),
	"path.separator":        string(os.PathListSeparator),
	"line.separator":        constants.LineSeparator,
	"user.name":             getCurrentUserName(),
	"user.home":             getCurrentUserHome(),
	"user.dir":              getCurrentWorkDir(),
	"sun.boot.library.path": path.Join(os.Getenv("JAVA_HOME"), "jre", "lib"),
}

func getCurrentUserName() string {
	u, err := user.Current()
	if err != nil {
		return "unknown"
	}
	return u.Username
}

func getCurrentUserHome() string {
	u, err := user.Current()
	if err != nil {
		return "unknown"
	}
	return u.HomeDir
}

func getCurrentWorkDir() string {
	workDir, err := os.Getwd()
	if err != nil {
		return "unknown"
	}
	return workDir
}

func initProperties(frame *rtda2.Frame) {
	properties := frame.LocalVars.GetRef(0)
	frame.OperandStack.PushRef(properties)

	classloader := properties.Class().ClassLoader()
	thread := frame.Thread()
	setProperty := properties.Class().GetInstanceMethod("setProperty", "(Ljava/lang/String;Ljava/lang/String;)Ljava/lang/Object;")
	for k, v := range systemProperties {
		key := rtda2.GetJString(classloader, k)
		value := rtda2.GetJString(classloader, v)

		stack := rtda2.NewOperandStack(3)
		stack.PushRef(properties)
		stack.PushRef(key)
		stack.PushRef(value)
		mockFrame := rtda2.NewMockFrame(thread, stack)
		thread.PushFrame(mockFrame)
		base.InvokeMethod(mockFrame, setProperty)
	}
}

//public static native String mapLibraryName(String libname)
func mapLibraryName(frame *rtda2.Frame) {
	/**
	Windows: "hello" -> "hello.dll"
	Linux/Solaris: "hello" "libhello.so"
	Mac: "hello" -> "libhello.dylib"
	*/
	libName := frame.LocalVars.GetRef(0)
	libNameGo := rtda2.GetGoString(libName)
	frame.OperandStack.PushRef(rtda2.GetJString(frame.Method().Class().ClassLoader(), buildLibraryName(libNameGo)))
}

func buildLibraryName(libName string) string {
	return constants.JavaNativeLibNamePrefix + libName + "." + constants.JavaNativeLibNameSuffix
}

//public static native long currentTimeMillis()
func currentTimeMillis(frame *rtda2.Frame) {
	millis := time.Now().UnixNano() / int64(time.Millisecond)
	frame.OperandStack.PushLong(millis)
}

//public static native long nanoTime()
func nanoTime(frame *rtda2.Frame) {
	frame.OperandStack.PushLong(time.Now().UnixNano())
}
