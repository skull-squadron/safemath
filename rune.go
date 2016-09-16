package safemath

func WouldAddErrorRune(a, b rune) bool {
	return WouldAddError32(a, b)
}

func WouldAddOverflowRune(a, b rune) bool {
	return WouldAddOverflow32(a, b)
}

func WouldAddUnderflowRune(a, b rune) bool {
	return WouldAddUnderflow32(a, b)
}

//  signed runeeger addition
//
// over/underflow condition
//            magnitude
//    sign msb  ...   lsb
// a    x   1 x ... x x x
// b    x   1 x ... x x x
//
// overflow = signs equal
// underflow = signs unequal
func AddRune(a, b rune) rune {
	return Add32(a, b)
}

func WouldSubOverflowRune(a, b rune) bool {
	return WouldSubOverflow32(a, b)
}

func WouldSubUnderflowRune(a, b rune) bool {
	return WouldSubUnderflow32(a, b)
}

func WouldSubErrorRune(a, b rune) bool {
	return WouldSubErrorRune(a, b)
}

// b  > 0: may underflow
// b == 0: cannot under or overflow
// b  < 0: may overflow
func SubRune(a, b rune) rune {
	return Sub32(a, b)
}

func WouldMulOverflowRune(a, b rune) bool {
	return WouldMulOverflow32(a, b)
}

func WouldMulUnderflowRune(a, b rune) bool {
	return WouldMulUnderflow32(a, b)
}

func WouldMulErrorRune(a, b rune) bool {
	return WouldMulErrorRune(a, b)
}

//  signed runeeger multiplication
// hacker's delight p. 69
func MulRune(a, b rune) rune {
	return Mul32(a, b)
}

func WouldDivByZeroRune(a, b rune) bool {
	return WouldDivByZero32(a, b)
}

func WouldDivUnderflowRune(a, b rune) bool {
	return WouldDivUnderflow32(a, b)
}

func WouldDivErrorRune(a, b rune) bool {
	return WouldDivError32(a, b)
}

//  signed runeeger division
func DivRune(a, b rune) rune {
	return Div32(a, b)
}
