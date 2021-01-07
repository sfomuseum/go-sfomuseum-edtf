package errors

import (
	"testing"
)

func TestInvalidError(t *testing.T) {

	err := Invalid()

	if !IsInvalid(err) {
		t.Logf("Failed to create InvalidError")
		t.Fail()
	}
}
