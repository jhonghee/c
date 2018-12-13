package c_test

import (
	"testing"

	"github.com/jhonghee/c"
)

func TestVersion(t *testing.T) {
	expected := "C v1.1"
	if c.Version() != expected {
		t.Error("Expected", expected, "but got", c.Version())
	}
}
