package safemath

import (
	"math/rand"
	"reflect"
	"testing"
)

func TestMul32(t *testing.T) {
	for _, x := range intTestCases {
		for _, y := range intTestCases {
			r, err := threwInt32(func() int32 { return Mul32(int32(x), int32(y)) })
			if shouldMul32Overflow(int32(x), int32(y)) {
				if err == nil {
					t.Errorf("Should have panic on %d*%d but got %d", x, y, r)
				}
			} else { // should not overflow
				if err != nil {
					t.Errorf("Unexpected panic %d * %d = %d err %v", x, y, r, err)
				} else if r != int32(x)*int32(y) {
					t.Errorf("%d != %d*%d", r, x, y)
				}
			}
		}
	}
}

func TestFuzzMul32ShouldPanic(t *testing.T) {
	f := func(x, y int32) bool {
		_, err := threwInt32(func() int32 { return Mul32(x, y) })
		opserr := IsError(err)
		//		t.Logf("should panic on %d * %d = %d (err %v) opserr %v", x, y, z, err, safeopserr)
		return err != nil && opserr
	}
	gen := func(args []reflect.Value, rnd *rand.Rand) {
		args[0] = reflect.ValueOf(int32(randIntNotBetween(rnd, minInt32Sqrt, maxInt32Sqrt)))
		args[1] = reflect.ValueOf(int32(randIntNotBetween(rnd, minInt32Sqrt, maxInt32Sqrt)))
	}
	if err := fuzz(f, gen); err != nil {
		t.Error(err)
	}
}

func TestFuzzMul32ShouldNotPanic(t *testing.T) {
	f := func(x, y int32) bool {
		z, err := threwInt32(func() int32 { return Mul32(x, y) })
		//t.Logf("should not panic on %d * %d = %d (err %v)", x, y, z, err)
		return err == nil && z == x*y
	}
	gen := func(args []reflect.Value, rnd *rand.Rand) {
		args[0] = reflect.ValueOf(int32(randIntBetween(rnd, minInt32Sqrt, maxInt32Sqrt)))
		args[1] = reflect.ValueOf(int32(randIntBetween(rnd, minInt32Sqrt, maxInt32Sqrt)))
	}
	if err := fuzz(f, gen); err != nil {
		t.Error(err)
	}
}
