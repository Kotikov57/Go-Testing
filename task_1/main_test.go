package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSum(t *testing.T) {
	a := 2
	b := 3
	expect := 5
	assert.Equal(t, Sum(a, b), expect, "expected 2 + 3 = 5")
}
