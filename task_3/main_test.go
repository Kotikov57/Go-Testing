package main

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestPop(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		result, err := Pop([]int{1, 2, 3, 4, 5})

		require.NoError(t, err)

		assert.Equal(t, []int{1, 2, 3, 4}, result)
	})

	t.Run("empty input array", func(t *testing.T) {
		result, err := Pop([]int{})

		require.Error(t, err)
		assert.Equal(t, []int{}, result)

		assert.EqualError(t, err, "empty input array")
	})
}
