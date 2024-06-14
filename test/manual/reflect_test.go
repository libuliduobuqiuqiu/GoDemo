package manual

import (
	"godemo/internal/gomanual/reflectdemo"
	"testing"
)

func TestReflectTypeOf(t *testing.T) {
	reflectdemo.BaseUseReflectType()
}

func TestReflectValueOf(t *testing.T) {
	reflectdemo.BaseUseReflectValue()
}

func TestReflectFunction(t *testing.T) {
	reflectdemo.BaseUseReflectFunction()
}

func TestReflectStruct(t *testing.T) {
	reflectdemo.BaseUseReflectStruct()
}

func TestReflectDeepEqual(t *testing.T) {
	reflectdemo.BaseUseReflectDeepEqual()
}
