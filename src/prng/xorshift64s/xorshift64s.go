// xorshift64s provides a PRNG of uint64 values.
package xorshift64s

type XS64S uint64

func NewXS64S() XS64S { return 1 }

func (r *XS64S) Next() uint64 {
	u := *r

	u ^= u >> 12
	u ^= u << 25
	u ^= u >> 27

	*r = u

	return uint64(u) * 2685821657736338717
}
