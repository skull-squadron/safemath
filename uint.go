package safemath

import "strconv"

const (
	maxUint          = ^uint(0)
	maxUintSqrt      = (uint(1) << (strconv.IntSize / 2)) - 1
	uintMagnitudeMSB = uint(1) << (strconv.IntSize - 1) // MSB
)

var (
	uaddOverflow  = Error("Unsigned addition overflow")
	usubUnderflow = Error("Unsigned subtraction Underflow")
	umulOverflow  = Error("Unsigned multiplication overflow")
	udivByZero    = Error("Unsigned division by zero")
)

func WouldUaddOverflow(a, b uint) bool {
	return a&b&uintMagnitudeMSB > 0
}

// safe Unsigned integer addition
//
// overflow condition
//
//    msb  ...   lsb
// a   1 x ... x x x
// b   1 x ... x x x
func Uadd(a, b uint) uint {
	if WouldUaddOverflow(a, b) {
		panic(uaddOverflow)
	}
	return a + b
}

func WouldUsubUnderflow(a, b uint) bool {
	return a < b
}

func Usub(a, b uint) uint {
	if WouldUsubUnderflow(a, b) {
		panic(usubUnderflow)
	}
	return a - b
}

func WouldUmulOverflow(a, b uint) bool {
	return (a|b) > maxUintSqrt && a != 0 && maxUint/a < b
}

// TODO: Usub(a, b uint) Uint

// safe Unsigned integer multiplication
//
// http://kqueue.org/blog/2012/03/16/fast-integer-overflow-detection/
// hacker's delight p. 68
func Umul(a, b uint) uint {
	if WouldUmulOverflow(a, b) {
		panic(umulOverflow)
	}
	return a * b
}

func WouldUmulDivZero(a, b uint) bool {
	return b == 0
}

func Udiv(a, b uint) uint {
	if WouldUmulDivZero(a, b) {
		panic(udivByZero)
	}
	return a / b
}
