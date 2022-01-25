package app

import (
	"fmt"
	"github.com/rodbate/jvm-on-go/pkg/classpath"
	"github.com/rodbate/jvm-on-go/pkg/constants/classname"
	"github.com/rodbate/jvm-on-go/pkg/constants/descriptors"
	"github.com/rodbate/jvm-on-go/pkg/instructions/base"
	"github.com/rodbate/jvm-on-go/pkg/interpreter"
	"github.com/rodbate/jvm-on-go/pkg/rtda"
	"github.com/spf13/cobra"
	"strings"
)

var (
	options = newOptions()

	jvmCommand = &cobra.Command{
		Use:  "run",
		Long: "Run jvm",
		Run: func(cmd *cobra.Command, args []string) {
			options.class = args[0]
			options.args = args[1:]
			newJvm().Start()
		},
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) < 1 {
				return fmt.Errorf("require more than one args")
			}
			return nil
		},
	}
)

func init() {
	jvmCommand.Flags().BoolVarP(&options.verbose, "verbose", "v", options.versionFlag, "verbose flag")
	jvmCommand.Flags().StringVarP(&options.cpOption, "classpath", "c", options.cpOption, "classpath")
}

type jvm struct {
	classloader *rtda.ClassLoader
	mainThread  *rtda.Thread
}

func newJvm() *jvm {
	cp := classpath.Parse(options.cpOption)
	classloader := rtda.NewClassLoader(cp, options.verbose)
	return &jvm{
		classloader: classloader,
		mainThread:  rtda.NewThread(),
	}
}

func (jvm *jvm) Start() {
	jvm.initVM()
	if jvm.mainThread.ExitCode == 0 {
		jvm.executeMainMethod()
	}
}

func (jvm *jvm) initVM() {
	vmClass := jvm.classloader.LoadClass(classname.VM)
	base.InitClass(jvm.mainThread, vmClass)
	interpreter.Interpret(jvm.mainThread, options.verbose)
}

func (jvm *jvm) executeMainMethod() {
	mainClassName := strings.Replace(options.class, ".", "/", -1)
	mainClass := jvm.classloader.LoadClass(mainClassName)
	mainMethod := mainClass.GetMainMethod()
	if mainMethod == nil {
		panic("cannot find main method in class: " + options.class)
	}
	newFrame := jvm.mainThread.NewFrame(mainMethod)
	jvm.mainThread.PushFrame(newFrame)
	newFrame.LocalVars.SetRef(0, createJStringArgs(mainMethod.Class().ClassLoader(), options.args))
	interpreter.Interpret(jvm.mainThread, options.verbose)
}

func createJStringArgs(loader *rtda.ClassLoader, args []string) *rtda.Object {
	stringArray := loader.LoadClass(descriptors.ArrayString).NewArray(uint(len(args)))
	data := stringArray.Refs()
	for i, arg := range args {
		data[i] = rtda.GetJString(loader, arg)
	}
	return stringArray
}
