package base

import "jvm-on-go/rtda"

func InitClass(thread *rtda.Thread, class *rtda.Class) {
	if class.InitStarted() {
		return
	}
	class.StartInit()
	scheduleClinit(thread, class)
	initSuperClass(thread, class)
}

func scheduleClinit(thread *rtda.Thread, class *rtda.Class) {
	clinit := class.GetClinitMethod()
	if clinit != nil {
		frame := thread.NewFrame(clinit)
		thread.PushFrame(frame)
	}
}

func initSuperClass(thread *rtda.Thread, class *rtda.Class) {
	if class.IsInterface() {
		return
	}
	superClass := class.SuperClass()
	if superClass == nil || superClass.InitStarted() {
		return
	}
	InitClass(thread, superClass)
}
