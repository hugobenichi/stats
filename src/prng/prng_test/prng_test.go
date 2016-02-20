package prng_testing

import (
	"fmt"
	"math"
	"testing"

	"prng"
	"prng/xorshift1024s"
	"prng/xorshift128p"
	"prng/xorshift64s"
)

func TestUnit(t *testing.T) {
	// any uint64 gives a number between 0 and 1, test -1, 0, ^0, 1 << 64 - 1

	// any n and n+1 have f(n) < f(n+1)
}

func TestPRNG(t *testing.T) {
	tests := []struct {
		desc string
		g    prng.P
	}{
		{"xorshift 64 star", xorshift64s.New()},
		{"xorshift 128 plus", xorshift128p.New()},
		{"xorshift 1024 star", xorshift1024s.New()},
	}

	for _, tt := range tests {
		// check non-overlapping sequnces
		for i := 0; i < 10; i++ {
			if err := sequences(tt.g); err != nil {
				t.Errorf("%s: got overlap in sequence, %v", tt.desc, err)
			}
		}

		// estimate Pi several times
		for i := 0; i < 10; i++ {
			x := pi(tt.g)
			if d := math.Abs(x - math.Pi); d > 0.01 {
				t.Errorf("%s: got x=%f, want |x-pi| <= 0.01 but got: %f", tt.desc, x, d)
			}
		}
	}
}

func pi(g prng.Float64) float64 {
	tot := 0.0
	ins := 0.0
	for i := 0; i < 400000; i++ {
		x, y := g.NextF(), g.NextF()
		if x*x+y*y < 1 {
			ins++
		}
		tot++
	}
	return 4 * ins / tot
}

func sequences(g prng.UInt64) error {
	ns := map[uint64]bool{}
	for i := 0; i < 10000; i++ {
		x := g.Next()
		if ns[x] {
			return fmt.Errorf("already got %d", x)
		}
		ns[x] = true
	}
	return nil
}
