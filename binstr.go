// Package binarystring creates a bitset out of an input string
package binarystring

import ()

//
type BitSet struct {
	set  []uint8
	bits uint64
}

const (
	SLOT_SIZE = 8
)

// New creates a new bitset
func New(sz string) *BitSet {
	bs := new(BitSet)
	bs.set = append(bs.set, uint8(0))
	bitPos := uint8(0)

	for i := len(sz) - 1; i >= 0; i-- {
		if bitPos > 7 {
			bitPos = 0
			bs.set = append(bs.set, uint8(0))
		}

		v := sz[i]
		if v == 48 || v == 49 {
			if v == 49 {
				bs.set[len(bs.set)-1] = bs.set[len(bs.set)-1] | (uint8(1) << bitPos)
			}
			bitPos = bitPos + 1
			bs.bits = bs.bits + 1
		} else {
			return nil
		}
	}

	// for i := len(bs.set) - 1; i >= 0; i-- {
	// 	fmt.Printf("%08b", uint64(bs.set[i]))
	// }

	return bs
}

// Len returns number of bits in the bitset
func (bs *BitSet) Len() uint64 {
	return bs.bits
}

// GetSetBits returns number of set bits in the bit set
func (bs *BitSet) GetSetBits() uint64 {
	sb := uint64(0)
	for i := 0; i < len(bs.set); i++ {
		x := bs.set[i]

		for x > 0 {
			x = x & (x - 1)
			sb = sb + 1
		}
	}

	return sb
}

// Or does a bitwise or of two bitsets and returns a new bitset
func (bs *BitSet) Or(in *BitSet) *BitSet {
	var i, j int
	nbs := new(BitSet)

	for i < len(bs.set) && j < len(in.set) {
		nbs.set = append(nbs.set, bs.set[i]|in.set[i])
		i++
		j++
	}

	for i < len(bs.set) {
		nbs.set = append(nbs.set, bs.set[i])
		i++
	}

	for j < len(in.set) {
		nbs.set = append(nbs.set, in.set[j])
		j++
	}

	nbs.bits = bs.bits
	if in.bits > nbs.bits {
		nbs.bits = in.bits
	}

	return nbs
}
