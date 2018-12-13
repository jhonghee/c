package c_test

import (
	"testing"

	"github.com/jhonghee/c"
)

func TestVersion(t *testing.T) {
	expected := "C v1.2->D v1.4->E v1.2"
	if c.Version() != expected {
		t.Error("Expected", expected, "but got", c.Version())
	}
}
