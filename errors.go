package safemath

var (
	AddOverflow  = Error("Signed addition overflow")
	AddUnderflow = Error("Signed addition underflow")
	SubOverflow  = Error("Signed subtraction overflow")
	SubUnderflow = Error("Signed subtraction underflow")
	MulOverflow  = Error("Signed multiplication overflow")
	MulUnderflow = Error("Signed multiplication underflow")
	DivUnderflow = Error("Signed divison underflow")
	DivByZero    = Error("Signed divison by zero")
)

var (
	UaddOverflow  = Error("Unsigned addition overflow")
	UsubUnderflow = Error("Unsigned subtraction Underflow")
	UmulOverflow  = Error("Unsigned multiplication overflow")
	UdivByZero    = Error("Unsigned division by zero")
)
