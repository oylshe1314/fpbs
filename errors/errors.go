package errors

import (
	"fmt"
)

type StringError string

func (err StringError) Error() string {
	return string(err)
}

func Error(text string) error {
	return StringError(text)
}

func Errorf(format string, args ...any) error {
	return StringError(fmt.Sprintf(format, args...))
}
