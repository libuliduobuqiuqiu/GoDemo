package genericsdemo

import (
	"fmt"
)

type MyLonger interface {
	int8
	GetName() string
}

type SignedInt interface {
	int8 | int16 | int32 | int64
}

type UnsignedInt interface {
	uint8 | uint16 | uint32 | uint64
}

type Integer interface {
	SignedInt | UnsignedInt
}

type MyNumber interface {
	Integer
	SignedInt
}

type MyInteger interface {
	SignedInt
	UnsignedInt
}

func SumIntegers[K comparable, V Integer](m map[K]V) V {
	var s V
	for _, v := range m {
		s += v
	}
	return s
}

func ExecSumIntegers() {
	ints := map[string]int64{
		"first":  34,
		"second": 12,
	}

	uints := map[string]uint64{
		"first":  35,
		"second": 26,
	}

	fmt.Printf("Generic Sum: %v and %v \n", SumIntegers[string, int64](ints),
		SumIntegers[string, uint64](uints))
}
