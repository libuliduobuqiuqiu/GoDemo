package genericsdemo

import (
	"fmt"
)

type MyLonger interface {
	int8
	GetName() string
}

type SingedInt interface {
	int8 | int16 | int32 | int64
}

type UnSingedInt interface {
	uint8 | uint16 | uint32 | uint64
}

type Integer interface {
	SingedInt | UnSingedInt
}

type MyNumber interface {
	Integer
	SingedInt
}

type MyInteger interface {
	SingedInt
	UnSingedInt
}

func SumIntOrFloats[K comparable, V Integer](m map[K]V) V {
	var s V
	for _, v := range m {
		s += v
	}
	return s
}

func ExecSumIntOrFloat() {
	ints := map[string]int64{
		"first":  34,
		"second": 12,
	}

	uints := map[string]uint64{
		"first":  35,
		"second": 26,
	}

	fmt.Printf("Genric Sum: %v and %v \n", SumIntOrFloats[string, int64](ints),
		SumIntOrFloats[string, uint64](uints))
}
