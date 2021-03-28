package main

import (
	"github.com/rodbate/jvm-on-go/classpath"
	"github.com/rodbate/jvm-on-go/constants/classname"
	"github.com/rodbate/jvm-on-go/constants/descriptors"
	"github.com/rodbate/jvm-on-go/instructions/base"
	"github.com/rodbate/jvm-on-go/rtda"
	"strings"
)

type Jvm struct {
	cmd         *Cmd
	classloader *rtda.ClassLoader
	mainThread  *rtda.Thread
}

func newJvm(cmd *Cmd) *Jvm {
	cp := classpath.Parse(cmd.cpOption)
	classloader := rtda.NewClassLoader(cp, cmd.verboseClassLoad)
	return &Jvm{
		cmd:         cmd,
		classloader: classloader,
		mainThread:  rtda.NewThread(),
	}
}

func (jvm *Jvm) Start() {
	jvm.initVM()
	if jvm.mainThread.ExitCode == 0 {
		jvm.executeMainMethod()
	}
}

func (jvm *Jvm) initVM() {
	vmClass := jvm.classloader.LoadClass(classname.VM)
	base.InitClass(jvm.mainThread, vmClass)
	interpret(jvm.mainThread, jvm.cmd.verboseInstFlag)
}

func (jvm *Jvm) executeMainMethod() {
	mainClassName := strings.Replace(jvm.cmd.class, ".", "/", -1)
	mainClass := jvm.classloader.LoadClass(mainClassName)
	mainMethod := mainClass.GetMainMethod()
	if mainMethod == nil {
		panic("cannot find main method in class: " + jvm.cmd.class)
	}
	newFrame := jvm.mainThread.NewFrame(mainMethod)
	jvm.mainThread.PushFrame(newFrame)
	newFrame.LocalVars.SetRef(0, createJStringArgs(mainMethod.Class().ClassLoader(), jvm.cmd.args))
	interpret(jvm.mainThread, jvm.cmd.verboseInstFlag)
}

func createJStringArgs(loader *rtda.ClassLoader, args []string) *rtda.Object {
	stringArray := loader.LoadClass(descriptors.ArrayString).NewArray(uint(len(args)))
	data := stringArray.Refs()
	for i, arg := range args {
		data[i] = rtda.GetJString(loader, arg)
	}
	return stringArray
}
