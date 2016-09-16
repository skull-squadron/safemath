package safemath

import (
	"math/rand"
	"reflect"
	"testing"
)

func TestMulRune(t *testing.T) {
	for _, x := range intTestCases {
		for _, y := range intTestCases {
			r, err := threwRune(func() rune { return MulRune(rune(x), rune(y)) })
			if shouldMulRuneOverflow(rune(x), rune(y)) {
				if err == nil {
					t.Errorf("Should have panic on %d*%d but got %d", x, y, r)
				}
			} else { // should not overflow
				if err != nil {
					t.Errorf("Unexpected panic %d * %d = %d err %v", x, y, r, err)
				} else if r != rune(x)*rune(y) {
					t.Errorf("%d != %d*%d", r, x, y)
				}
			}
		}
	}
}

func TestFuzzMulRuneShouldPanic(t *testing.T) {
	f := func(x, y rune) bool {
		_, err := threwRune(func() rune { return MulRune(x, y) })
		opserr := IsError(err)
		//		t.Logf("should panic on %d * %d = %d (err %v) opserr %v", x, y, z, err, safeopserr)
		return err != nil && opserr
	}
	gen := func(args []reflect.Value, rnd *rand.Rand) {
		args[0] = reflect.ValueOf(rune(randIntNotBetween(rnd, minInt32Sqrt, maxInt32Sqrt)))
		args[1] = reflect.ValueOf(rune(randIntNotBetween(rnd, minInt32Sqrt, maxInt32Sqrt)))
	}
	if err := fuzz(f, gen); err != nil {
		t.Error(err)
	}
}

func TestFuzzMulRuneShouldNotPanic(t *testing.T) {
	f := func(x, y rune) bool {
		z, err := threwRune(func() rune { return MulRune(x, y) })
		//t.Logf("should not panic on %d * %d = %d (err %v)", x, y, z, err)
		return err == nil && z == x*y
	}
	gen := func(args []reflect.Value, rnd *rand.Rand) {
		args[0] = reflect.ValueOf(rune(randIntBetween(rnd, minInt32Sqrt, maxInt32Sqrt)))
		args[1] = reflect.ValueOf(rune(randIntBetween(rnd, minInt32Sqrt, maxInt32Sqrt)))
	}
	if err := fuzz(f, gen); err != nil {
		t.Error(err)
	}
}
