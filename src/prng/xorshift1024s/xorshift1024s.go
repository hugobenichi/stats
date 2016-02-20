// xorshift1024s provides a PRNG of uint64 values.
package xorshift1024s

import (
	"prng"
)

// The xorshift1024* PRNG. Public fields allow inspection and custom seeding.
type G struct {
	State [16]uint64
	I     int
}

func New() *G {
	return &G{
		State: [16]uint64{
			// generated from XS128P after dropping first 10 outputs
			10496030469740439798,
			1362371001014398178,
			10234834343287503199,
			2065174045786219692,
			4360127029907502923,
			3280463725544326876,
			4098392981707075411,
			3999493022945510211,
			1490933961015620186,
			13252044636817621309,
			15998111192842087806,
			3805184684654049962,
			2739691196446076535,
			15675371886373393618,
			16373398947986388217,
			12742849885299357362,
		},
		I: 0,
	}
}

func (r *G) Next() uint64 {
	var (
		i  = r.I
		j  = (i + 1) & 0xF
		s0 = r.State[i]
		s1 = r.State[j]
	)

	s1 ^= s1 << 31
	s1 = s1 ^ s0 ^ (s1 >> 11) ^ (s0 >> 30)

	r.State[j] = s1
	r.I = j

	return s1 * 1181783497276652981
}

func (r *G) NextF() float64 {
	return prng.Unit(r.Next())
}
