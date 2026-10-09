package godemo

import (
	"godemo/internal/godemo/errors"
	"godemo/pkg"
	"testing"
)

func TestHandleError(t *testing.T) {

	err := errors.HandleError()
	if err != nil {
		t.Log(err)
	}

	err = errors.NewMyError()
	if err == nil {
		t.Fatal("expected a demo error")
	}
	t.Log(err)

}

func TestStackError(t *testing.T) {

	errors.UseStackError()

}

func TestCustomError(t *testing.T) {
	err := pkg.NewErr("test custom err")
	pkg.PrintErr(err)
}
