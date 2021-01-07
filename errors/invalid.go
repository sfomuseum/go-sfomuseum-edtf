package errors

import (
	"fmt"
)

type InvalidError struct{}

func (e *InvalidError) Error() string {
	return fmt.Sprintf("Invalid or unsupported SFO Museum date string")
}

func Invalid() error {
	return &InvalidError{}
}

func IsInvalid(e error) bool {

	switch e.(type) {
	case *InvalidError:
		return true
	default:
		return false
	}
}
