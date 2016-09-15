package safemath

import (
	"math/rand"
	"reflect"
	"testing"
)

func TestMul(t *testing.T) {
	for _, x := range intTestCases {
		for _, y := range intTestCases {
			r, err := threwInt(func() int { return Mul(x, y) })
			if shouldMulOverflow(x, y) {
				if err == nil {
					t.Errorf("Should have panic on %d*%d but got %d", x, y, r)
				}
			} else { // should not overflow
				if err != nil {
					t.Errorf("Unexpected panic %d * %d = %d err %v", x, y, r, err)
				} else if r != x*y {
					t.Errorf("%d != %d*%d", r, x, y)
				}
			}
		}
	}
}

func TestFuzzMulShouldPanic(t *testing.T) {
	f := func(x, y int) bool {
		_, err := threwInt(func() int { return Mul(x, y) })
		opserr := IsError(err)
		//		t.Logf("should panic on %d * %d = %d (err %v) opserr %v", x, y, z, err, safeopserr)
		return err != nil && opserr
	}
	gen := func(args []reflect.Value, rnd *rand.Rand) {
		args[0] = reflect.ValueOf(randIntNotBetween(rnd, minIntSqrt, maxIntSqrt))
		args[1] = reflect.ValueOf(randIntNotBetween(rnd, minIntSqrt, maxIntSqrt))
	}
	if err := fuzz(f, gen); err != nil {
		t.Error(err)
	}
}

func TestFuzzMulShouldNotPanic(t *testing.T) {
	f := func(x, y int) bool {
		z, err := threwInt(func() int { return Mul(x, y) })
		//t.Logf("should not panic on %d * %d = %d (err %v)", x, y, z, err)
		return err == nil && z == x*y
	}
	gen := func(args []reflect.Value, rnd *rand.Rand) {
		args[0] = reflect.ValueOf(randIntBetween(rnd, minIntSqrt, maxIntSqrt))
		args[1] = reflect.ValueOf(randIntBetween(rnd, minIntSqrt, maxIntSqrt))
	}
	if err := fuzz(f, gen); err != nil {
		t.Error(err)
	}
}
