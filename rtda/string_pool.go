package rtda

import (
	"github.com/rodbate/jvm-on-go/constants/classname"
	"github.com/rodbate/jvm-on-go/constants/descriptors"
	"unicode/utf16"
)

var internedStrings = map[string]*Object{}

func GetJString(loader *ClassLoader, goStr string) *Object {
	if s, ok := internedStrings[goStr]; ok {
		return s
	}
	chars := stringToUtf16Chars(goStr)
	jChars := &Object{class: loader.LoadClass(descriptors.ArrayChar), Data: chars}
	jStr := loader.LoadClass(classname.String).NewInstance()
	jStr.SetFieldValue("value", descriptors.ArrayChar, jChars)
	internedStrings[goStr] = jStr
	return jStr
}

func GetGoString(jString *Object) string {
	value := jString.GetFieldValue("value", descriptors.ArrayChar)
	chars := (value.(*Object)).Data.(ArrayChar)
	return utf16ToString(chars)
}

func InternString(jString *Object) *Object {
	goStr := GetGoString(jString)
	if s, ok := internedStrings[goStr]; ok {
		return s
	}
	internedStrings[goStr] = jString
	return jString
}

func stringToUtf16Chars(str string) ArrayChar {
	runes := []rune(str)
	return utf16.Encode(runes)
}

func utf16ToString(chars []uint16) string {
	return string(utf16.Decode(chars))
}
