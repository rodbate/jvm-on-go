package base

import (
	rtda2 "github.com/rodbate/jvm-on-go/pkg/rtda"
)

func InitClass(thread *rtda2.Thread, class *rtda2.Class) {
	if class.InitStarted() {
		return
	}
	class.StartInit()
	scheduleClinit(thread, class)
	initSuperClass(thread, class)
}

func scheduleClinit(thread *rtda2.Thread, class *rtda2.Class) {
	clinit := class.GetClinitMethod()
	if clinit != nil {
		frame := thread.NewFrame(clinit)
		thread.PushFrame(frame)
	}
}

func initSuperClass(thread *rtda2.Thread, class *rtda2.Class) {
	if class.IsInterface() {
		return
	}
	superClass := class.SuperClass()
	if superClass == nil || superClass.InitStarted() {
		return
	}
	InitClass(thread, superClass)
}
