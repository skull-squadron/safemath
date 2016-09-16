package safemath

import "strconv"

const (
	maxInt          = int(maxUint >> 1)          // 0b01111111111...1
	minInt          = -maxInt - 1                // 0b10000000000...0
	intMagnitudeMSB = 1 << (strconv.IntSize - 2) // MSB of magnitude
	maxIntSqrt      = (1 << (strconv.IntSize/2 - 1)) - 1
	minIntSqrt      = -maxIntSqrt - 1
)

func WouldAddError(a, b int) bool {
	return a&b&intMagnitudeMSB < 0
}

func WouldAddOverflow(a, b int) bool {
	return WouldAddError(a, b) && a^b >= 0
}

func WouldAddUnderflow(a, b int) bool {
	return WouldAddError(a, b) && a^b < 0
}

//  signed integer addition
//
// over/underflow condition
//            magnitude
//    sign msb  ...   lsb
// a    x   1 x ... x x x
// b    x   1 x ... x x x
//
// overflow = signs equal
// underflow = signs unequal
func Add(a, b int) int {
	if WouldAddOverflow(a, b) {
		panic(AddOverflow)
	}
	if WouldAddUnderflow(a, b) {
		panic(AddUnderflow)
	}
	return a + b
}

func WouldSubOverflow(a, b int) bool {
	return b > 0 && a > maxInt-b
}

func WouldSubUnderflow(a, b int) bool {
	return b < 0 && a < minInt-b
}

func WouldSubError(a, b int) bool {
	return WouldSubOverflow(a, b) || WouldSubUnderflow(a, b)
}

// b  > 0: may underflow
// b == 0: cannot under or overflow
// b  < 0: may overflow
func Sub(a, b int) int {
	if WouldSubOverflow(a, b) {
		panic(SubOverflow)
	} else if WouldSubUnderflow(a, b) {
		panic(SubUnderflow)
	}
	return a - b
}

func WouldMulOverflow(a, b int) bool {
	return a^b >= 0 && ((a < 0 && b < maxInt/a) || (b > 0 && a > maxInt/b))
}

func WouldMulUnderflow(a, b int) bool {
	return a^b < 0 && ((b > 0 && a < minInt/b) || (a > 0 && b < minInt/a))
}

func WouldMulError(a, b int) bool {
	return WouldMulOverflow(a, b) || WouldMulUnderflow(a, b)
}

//  signed integer multiplication
// hacker's delight p. 69
func Mul(a, b int) int {
	if WouldMulOverflow(a, b) {
		panic(MulOverflow)
	}
	if WouldMulUnderflow(a, b) {
		panic(MulUnderflow)
	}
	return a * b
}

func WouldDivByZero(a, b int) bool {
	return b == 0
}

func WouldDivUnderflow(a, b int) bool {
	return b == -1 && a == minInt
}
func WouldDivError(a, b int) bool {
	return WouldDivByZero(a, b) || WouldDivUnderflow(a, b)
}

//  signed integer division
func Div(a, b int) int {
	if WouldDivByZero(a, b) {
		panic(DivByZero)
	}
	if WouldDivUnderflow(a, b) {
		panic(DivUnderflow)
	}
	return a / b
}
