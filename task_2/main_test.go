package main

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestDivide(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		result, err := Divide(10, 2)

		require.NoError(t, err)

		assert.Equal(t, 5, result)
	})

	t.Run("division by zero", func(t *testing.T) {
		result, err := Divide(10, 0)

		require.Error(t, err)
		assert.Equal(t, 0, result)

		assert.EqualError(t, err, "division by zero")
	})
}
