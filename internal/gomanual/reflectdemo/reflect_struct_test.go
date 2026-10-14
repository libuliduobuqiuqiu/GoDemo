package reflectdemo

import (
	"reflect"
	"testing"
)

func TestHandleStruct(t *testing.T) {
	p := &PersonModel{Name: "Ada", Age: 23, Address: "Guangdong"}
	if err := handleStruct(p); err != nil {
		t.Fatal(err)
	}
	if p.Age != 24 || p.Address != "Guangdong Guangzhou city" {
		t.Fatalf("unexpected mutation: %+v", p)
	}
	for _, value := range []any{nil, (*PersonModel)(nil), PersonModel{}, &[]int{}} {
		if err := handleStruct(value); err == nil {
			t.Errorf("expected error for %T", value)
		}
	}
}

func TestReflectionBoundaries(t *testing.T) {
	p := PersonModel{Name: "Ada"}
	if reflect.ValueOf(p).FieldByName("Name").CanSet() {
		t.Fatal("value copy must not be settable")
	}
	v := reflect.ValueOf(&p).Elem()
	if !v.FieldByName("Name").CanSet() || v.FieldByName("money").CanInterface() {
		t.Fatal("unexpected exported/unexported field access")
	}
	if reflect.DeepEqual([]int(nil), []int{}) || reflect.DeepEqual(int(1), int64(1)) {
		t.Fatal("DeepEqual must distinguish nil/empty slices and different types")
	}
}
