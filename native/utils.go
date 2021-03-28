package native

import (
	"github.com/rodbate/jvm-on-go/rtda"
	"unsafe"
)

func ToJavaByteArray(loader *rtda.ClassLoader, goBytes []byte) *rtda.Object {
	if goBytes != nil {
		jBytes := castUint8sToInt8s(goBytes)
		return rtda.NewByteArray(loader, jBytes)
	}
	return nil
}
func castUint8sToInt8s(goBytes []byte) (jBytes []int8) {
	ptr := unsafe.Pointer(&goBytes)
	jBytes = *((*[]int8)(ptr))
	return
}
