package manual

import (
	"godemo/internal/gomanual/genericsdemo"
	"testing"
)

func TestMarshalMan(t *testing.T) {
	genericsdemo.MarshalMan()
}

func TestUnMarshalMan(t *testing.T) {
	if err := genericsdemo.UnMarshalMan(); err != nil {
		t.Fatal(err)
	}
}

func TestGenericNumberConstraints(t *testing.T) {
	type Score int
	if got := genericsdemo.SumInt(Score(2), Score(3)); got != 5 {
		t.Fatalf("sum = %d, want 5", got)
	}
	if got := genericsdemo.SumIntegers(map[string]int64{"a": 2, "b": 3}); got != 5 {
		t.Fatalf("map sum = %d, want 5", got)
	}
}
