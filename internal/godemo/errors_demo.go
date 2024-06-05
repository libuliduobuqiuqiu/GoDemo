package godemo

import (
	"fmt"
	"github.com/pkg/errors"
)

func RaiseError() error {
	return fmt.Errorf("test raise error.")
}

func WrapError() error {
	err := RaiseError()
	return errors.Wrap(err, "Wrap Error: ")
}

type MyError struct {
	msg string
	err error
}

func (m MyError) Error() string {
	return m.msg
}

func (m MyError) UnWrap() error {
	return m.err
}

func NewMyError() error {
	return MyError{
		msg: "test my Error",
		err: errors.New("test my Wrap Error"),
	}
}

func HandleError() error {
	return fmt.Errorf("%w", NewMyError())
}
