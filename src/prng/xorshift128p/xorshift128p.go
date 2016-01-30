// xorshift128p provides a PRNG of uint64 values.
package xorshift128p

import (
	"prng"
)

// The xorshift128+ PRNG. Public fields for inspection and custom seeding.
type G struct {
	S0, S1 uint64
}

var xs128p_0 = G{S0: 1, S1: 2}

func New() G { return xs128p_0 }

func (r *G) Next() uint64 {
	s1, s0 := r.S0, r.S1 // swap intended

	s1 ^= s1 << 23
	s1 ^= s1 >> 17
	s1 ^= s0
	s1 ^= s0 >> 26

	r.S0, r.S1 = s0, s1

	return s0 + s1
}

func (r *G) NextF() float64 {
	return prng.Unit(r.Next())
}
