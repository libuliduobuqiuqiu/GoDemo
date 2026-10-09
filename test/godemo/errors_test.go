package godemo

import (
	stderrors "errors"
	"godemo/internal/godemo/errors"
	"godemo/pkg"
	"testing"
)

func TestHandleError(t *testing.T) {
	err := errors.HandleError()
	if err == nil {
		t.Fatal("expected a wrapped error")
	}

	var myErr errors.MyError
	if !stderrors.As(err, &myErr) {
		t.Fatalf("expected MyError in the chain, got %T", err)
	}
	if stderrors.Unwrap(myErr) == nil {
		t.Fatal("expected MyError to expose its cause")
	}
}

func TestStackError(t *testing.T) {

	errors.UseStackError()

}

func TestCustomError(t *testing.T) {
	err := pkg.NewErr("test custom err")
	pkg.PrintErr(err)
}
