package stub

var markStub bool

// GetMarkStubAddr return address of markStub. Objects with this value will be handled specially
func GetMarkStubAddr() *bool {
	return &markStub
}
