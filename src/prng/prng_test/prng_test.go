package prng_testing

import (
	"fmt"
	"math"
	"math/rand"
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
		{"standard prng", newStdPRNG()},
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
	for i := 0; i < 200000; i++ {
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

func BenchmarkXS64S(b *testing.B)   { benchmarkPRNG(b, xorshift64s.New()) }
func BenchmarkXS1024S(b *testing.B) { benchmarkPRNG(b, xorshift1024s.New()) }
func BenchmarkXS128P(b *testing.B)  { benchmarkPRNG(b, xorshift128p.New()) }
func BenchmarkStdPRNG(b *testing.B) { benchmarkPRNG(b, newStdPRNG()) }

func benchmarkPRNG(b *testing.B, g prng.UInt64) {
	for i := 0; i < b.N; i++ {
		g.Next()
	}
}

// BenchmarkStdPRNGInlined controls the cost of interface call and wrapping
func BenchmarkStdPRNGInlined(b *testing.B) {
	g := rand.NewSource(0)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		g.Int63()
	}
}

// stdPRNG wraps a std lib math/rand.Source
type stdPRNG struct {
	rand.Source
}

func (s stdPRNG) Next() uint64   { return uint64(s.Int63()) }
func (s stdPRNG) NextF() float64 { return prng.Unit(s.Next()) }
func newStdPRNG() stdPRNG        { return stdPRNG{rand.NewSource(0)} }
