// xorshift64s provides a PRNG of uint64 values.
package xorshift64s

import (
	"prng"
)

type G uint64

func New() *G {
	g := G(1)
	return &g
}

func (r *G) Next() uint64 {
	u := *r

	u ^= u >> 12
	u ^= u << 25
	u ^= u >> 27

	*r = u

	return uint64(u) * 2685821657736338717
}

func (r *G) NextF() float64 {
	return prng.Unit(r.Next())
}
