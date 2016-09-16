package safemath

import "math"

const (
	int32MagnitudeMSB = 1 << (32 - 2) // MSB of magnitude
	maxInt32Sqrt      = (1 << (32/2 - 1)) - 1
	minInt32Sqrt      = -maxInt32Sqrt - 1
)

func WouldAddError32(a, b int32) bool {
	return a&b&int32MagnitudeMSB < 0
}

func WouldAddOverflow32(a, b int32) bool {
	return WouldAddError32(a, b) && a^b >= 0
}

func WouldAddUnderflow32(a, b int32) bool {
	return WouldAddError32(a, b) && a^b < 0
}

//  signed int32eger addition
//
// over/underflow condition
//            magnitude
//    sign msb  ...   lsb
// a    x   1 x ... x x x
// b    x   1 x ... x x x
//
// overflow = signs equal
// underflow = signs unequal
func Add32(a, b int32) int32 {
	if WouldAddOverflow32(a, b) {
		panic(AddOverflow)
	}
	if WouldAddUnderflow32(a, b) {
		panic(AddUnderflow)
	}
	return a + b
}

func WouldSubOverflow32(a, b int32) bool {
	return b > 0 && a > math.MaxInt32-b
}

func WouldSubUnderflow32(a, b int32) bool {
	return b < 0 && a < math.MinInt32-b
}

func WouldSubError32(a, b int32) bool {
	return WouldSubOverflow32(a, b) || WouldSubUnderflow32(a, b)
}

// b  > 0: may underflow
// b == 0: cannot under or overflow
// b  < 0: may overflow
func Sub32(a, b int32) int32 {
	if WouldSubOverflow32(a, b) {
		panic(SubOverflow)
	} else if WouldSubUnderflow32(a, b) {
		panic(SubUnderflow)
	}
	return a - b
}

func WouldMulOverflow32(a, b int32) bool {
	return a^b >= 0 && ((a < 0 && b < math.MaxInt32/a) || (b > 0 && a > math.MaxInt32/b))
}

func WouldMulUnderflow32(a, b int32) bool {
	return a^b < 0 && ((b > 0 && a < math.MinInt32/b) || (a > 0 && b < math.MinInt32/a))
}

func WouldMulError32(a, b int32) bool {
	return WouldMulOverflow32(a, b) || WouldMulUnderflow32(a, b)
}

//  signed int32eger multiplication
// hacker's delight p. 69
func Mul32(a, b int32) int32 {
	if WouldMulOverflow32(a, b) {
		panic(MulOverflow)
	}
	if WouldMulUnderflow32(a, b) {
		panic(MulUnderflow)
	}
	return a * b
}

func WouldDivByZero32(a, b int32) bool {
	return b == 0
}

func WouldDivUnderflow32(a, b int32) bool {
	return b == -1 && a == math.MinInt32
}
func WouldDivError32(a, b int32) bool {
	return WouldDivByZero32(a, b) || WouldDivUnderflow32(a, b)
}

//  signed int32eger division
func Div32(a, b int32) int32 {
	if WouldDivByZero32(a, b) {
		panic(DivByZero)
	}
	if WouldDivUnderflow32(a, b) {
		panic(DivUnderflow)
	}
	return a / b
}
