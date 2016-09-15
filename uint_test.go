package safemath

import (
	"math/rand"
	"reflect"
	"testing"
)

func TestFuzzUmulShouldPanic(t *testing.T) {
	f := func(x, y uint) bool {
		_, err := threwUint(func() uint { return Umul(x, y) })
		//		t.Logf("should panic on %d * %d = %d (err %v)", x, y, r, err)
		return err != nil && IsError(err)
	}
	gen := func(args []reflect.Value, rnd *rand.Rand) {
		args[0] = reflect.ValueOf(randUintBetween(rnd, maxUintSqrt, maxUint))
		args[1] = reflect.ValueOf(randUintBetween(rnd, maxUintSqrt, maxUint))
	}
	if err := fuzz(f, gen); err != nil {
		t.Error(err)
	}
}

func TestFuzzUmulShouldNotPanic(t *testing.T) {
	f := func(x, y uint) bool {
		r, err := threwUint(func() uint { return Umul(x, y) })
		//		t.Logf("should not panic on %d * %d = %d (err %v)", x, y, r, err)
		return err == nil && r == x*y
	}
	gen := func(args []reflect.Value, rnd *rand.Rand) {
		args[0] = reflect.ValueOf(randUintBetween(rnd, 0, maxUintSqrt-1))
		args[1] = reflect.ValueOf(randUintBetween(rnd, 0, maxUintSqrt-1))
	}
	if err := fuzz(f, gen); err != nil {
		t.Error(err)
	}
}
