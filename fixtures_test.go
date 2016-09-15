package safemath

var intTestCases = []int{
	minInt,
	minInt + 1,
	minInt + 2,

	-intMagnitudeMSB - 3,
	-intMagnitudeMSB - 2,
	-intMagnitudeMSB - 1,
	-intMagnitudeMSB,
	-intMagnitudeMSB + 1,
	-intMagnitudeMSB + 2,
	-intMagnitudeMSB + 3,

	minIntSqrt - 2,
	minIntSqrt - 1,
	minIntSqrt,
	minIntSqrt + 1,
	minIntSqrt + 2,
	minIntSqrt + 3,

	-258,
	-257,
	-256,
	-255,
	-254,
	-253,

	-101,
	-100,
	-99,

	-11,
	-10,
	-9,
	-2,
	-1,

	0,

	1,
	2,
	9,
	10,
	11,

	99,
	100,
	101,

	254,
	255,
	256,
	257,

	maxIntSqrt - 2,
	maxIntSqrt - 1,
	maxIntSqrt,
	maxIntSqrt + 1,
	maxIntSqrt + 2,
	maxIntSqrt + 3,

	intMagnitudeMSB - 2,
	intMagnitudeMSB - 1,
	intMagnitudeMSB,
	intMagnitudeMSB + 1,
	intMagnitudeMSB + 2,

	maxInt - 2,
	maxInt - 1,
	maxInt,
}
