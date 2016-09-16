package safemath

import (
	"fmt"
	"github.com/steakknife/try"
	"math"
	"math/rand"
	"reflect"
	"testing/quick"
)

func fuzz(f interface{}, gen func([]reflect.Value, *rand.Rand)) error {
	return quick.Check(f, &quick.Config{MaxCount: 1000, MaxCountScale: 1, Values: gen})
}

func threwUint32(fn func() uint32) (res uint32, err error) {
	x, err := try.Catch(func(...interface{}) interface{} {
		return fn()
	}, []try.CatchFunc{})
	if x != nil {
		res = x.(uint32)
	}
	return
}

func threwRune(fn func() rune) (res rune, err error) {
	return threwInt32(fn)
}

func threwUint(fn func() uint) (res uint, err error) {
	x, err := try.Catch(func(...interface{}) interface{} {
		return fn()
	}, []try.CatchFunc{})
	if x != nil {
		res = x.(uint)
	}
	return
}

func threwInt32(fn func() int32) (res int32, err error) {
	x, err := try.Catch(func(...interface{}) interface{} {
		return fn()
	}, []try.CatchFunc{})
	if x != nil {
		res = x.(int32)
	}
	return
}

func threwInt(fn func() int) (res int, err error) {
	x, err := try.Catch(func(...interface{}) interface{} {
		return fn()
	}, []try.CatchFunc{})
	if x != nil {
		res = x.(int)
	}
	return
}

func randUint(r *rand.Rand) uint {
	return (uint(r.Int()) << 2) ^ (uint(r.Int()) >> 3)
}

func randUintBetween(r *rand.Rand, a, b uint) uint {
	if a > b {
		panic(fmt.Errorf("invalid range %d..%d", a, b))
	}
	return a + randUint(r)%(b-a+1)
}

func randInt(r *rand.Rand) int {
	return int(randUint(r))
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func randIntBetween(r *rand.Rand, a, b int) int {
	if a > b {
		panic(fmt.Errorf("invalid range %d..%d", a, b))
	}
	return a + abs(randInt(r))%(abs(b-a)+1)
}

func randBool(r *rand.Rand) bool {
	return randUint(r)&4 == 4
}

func randIntNotBetween(r *rand.Rand, a, b int) int {
	if a > b {
		panic(fmt.Errorf("invalid range %d..%d", a, b))
	}
	if randBool(r) { // below
		return randIntBetween(r, minInt, a-1)
	} else { // above
		return randIntBetween(r, b, maxInt)
	}
}

func shouldMulOverflow(x, y int) bool {
	if x == 0 || y == 0 || x == 1 || y == 1 {
		return false
	}

	// minInt * (<0 or >1) always creates underflow
	if x == minInt || y == minInt {
		return true
	}

	if x < 0 && y < 0 {
		if x < y {
			return minInt/-y >= x
		} else {
			return minInt/-x >= y
		}
	} else if x < 0 && y > 0 {
		return minInt/y > x
	} else if x > 0 && y < 0 {
		return minInt/x > y
	} else { // x > 0 && y > 0
		return maxInt/x < y
	}
}

func shouldMulRuneOverflow(x, y int32) bool {
	return shouldMul32Overflow(x, y)
}

func shouldMul32Overflow(x, y int32) bool {
	if x == 0 || y == 0 || x == 1 || y == 1 {
		return false
	}

	// math.MinInt * (<0 or >1) always creates underflow
	if x == math.MinInt32 || y == math.MinInt32 {
		return true
	}

	if x < 0 && y < 0 {
		if x < y {
			return math.MinInt32/-y >= x
		} else {
			return math.MinInt32/-x >= y
		}
	} else if x < 0 && y > 0 {
		return math.MinInt32/y > x
	} else if x > 0 && y < 0 {
		return math.MinInt32/x > y
	} else { // x > 0 && y > 0
		return math.MaxInt32/x < y
	}
}
