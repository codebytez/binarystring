package binarystring

import (
	"testing"
)

func TestLen(t *testing.T) {
	bs := New("111100001010")

	if bs.Len() != 12 {
		t.Errorf("Wrong length,expected - %d, got - %d", 12, bs.Len())
	}
}

func TestGetSetBits(t *testing.T) {
	bs := New("10101111001")

	if bs.GetSetBits() != 7 {
		t.Errorf("Wrong set bits, expected - %d,got - %d", 7, bs.GetSetBits())
	}
}

func TestOr(t *testing.T) {
	bs1 := New("1010")
	bs2 := New("1100000000000001")

	bs3 := bs1.Or(bs2)

	if bs3.GetSetBits() != 5 {
		t.Errorf("Bitwise OR Failed, expected - %d,got - %d", 3, bs3.GetSetBits())
	}
}
