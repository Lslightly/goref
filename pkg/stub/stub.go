package stub

import "unsafe"

var markStub bool

// GetMarkStubAddr return address of markStub. Objects with this value will be handled specially
func GetMarkStubAddr() uintptr {
	return uintptr(unsafe.Pointer(&markStub))
}
