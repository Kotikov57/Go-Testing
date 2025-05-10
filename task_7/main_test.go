package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestTimeFunc(t *testing.T) {
	t.Run("Morning", func(t *testing.T) {
		morning := time.Date(2023, 1, 1, 9, 0, 0, 0, time.UTC)
		assert.Equal(t, "Доброе утро", TimeFunc(morning))
	})
	t.Run("Afternoon", func(t *testing.T) {
		afternoon := time.Date(2023, 1, 1, 14, 0, 0, 0, time.UTC)
		assert.Equal(t, "Добрый день", TimeFunc(afternoon))
	})
	t.Run("Evening", func(t *testing.T) {
		evening := time.Date(2023, 1, 1, 20, 0, 0, 0, time.UTC)
		assert.Equal(t, "Добрый вечер", TimeFunc(evening))
	})
}
