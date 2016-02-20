package prng

// UInt64 is a pseudo-random number generator of uint64 values in [0,1<<64[
type UInt64 interface {
	Next() uint64
}

// UInt64 is a pseudo-random number generator of float64 values in [0,1[
type Float64 interface {
	NextF() float64
}

type P interface {
	UInt64
	Float64
}

var (
	f = 1. / (1 << 52)
	m = uint64(1<<52) - 1
)

// Unit generates a float64 in [0,1[ from the 52 low bits of an uint64.
func Unit(u uint64) float64 {
	return float64(u&m) * f // has form N x 2 ^ -52 with N a 52 bits unsigned int.
}
