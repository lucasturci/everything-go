package bitset

import "fmt"

type Bitset struct {
	set []uint64
	n   int
}

const _byteSize = 64

func New(n int) Bitset {
	return Bitset{
		set: make([]uint64, (n+_byteSize-1)/_byteSize),
		n:   n,
	}
}

func (b Bitset) Size() int {
	return b.n
}

func (b Bitset) Set(p int) error {
	if p < 0 || p > b.n {
		return fmt.Errorf("out of bounds access: trying to access %v at bitset of length %v", p, b.n)
	}
	i := p / _byteSize
	j := p % _byteSize
	b.set[i] |= (1 << j)
	return nil
}

func (b Bitset) Get(p int) (bool, error) {
	if p < 0 || p > b.n {
		return false, fmt.Errorf("out of bounds access: trying to access %v at bitset of length %v", p, b.n)
	}
	i := p / _byteSize
	j := p % _byteSize
	val := (b.set[i] >> j) & 1
	if val == 1 {
		return true, nil
	}
	return false, nil
}

func (b Bitset) Count() int {
	ans := 0
	for i := 0; i < len(b.set); i++ {
		ans += int(popcount(b.set[i]))
	}
	return ans
}

func Union(x, y Bitset) Bitset {
	ans := New(max(x.n, y.n))
	for i := 0; i < len(ans.set); i++ {
		ans.set[i] = x.set[i] | y.set[i]
	}
	return ans
}

func Intersection(x, y Bitset) Bitset {
	ans := New(max(x.n, y.n))
	for i := 0; i < len(ans.set); i++ {
		ans.set[i] = x.set[i] & y.set[i]
	}
	return ans
}

// Taken from https://github.com/tmthrgd/go-popcount/blob/afb1ace8b04f/popcount.go
func popcount(x uint64) uint64 {
	x = (x & 0x5555555555555555) + ((x & 0xAAAAAAAAAAAAAAAA) >> 1)
	x = (x & 0x3333333333333333) + ((x & 0xCCCCCCCCCCCCCCCC) >> 2)
	x = (x & 0x0F0F0F0F0F0F0F0F) + ((x & 0xF0F0F0F0F0F0F0F0) >> 4)
	x *= 0x0101010101010101
	return ((x >> 56) & 0xFF)
}
