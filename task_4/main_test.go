package main

import (
	"testing"
	"testing/quick"
)

func TestRandomStringLength(t *testing.T) {
	f := func(length uint8) bool {
		l := int(length % 50)

		s := RandomString(l)
		return len(s) == l
	}

	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}
