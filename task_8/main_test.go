// Напишите тест для функции, работающей в многозадачной среде

package main

import (
	"github.com/stretchr/testify/assert"
	"sync"
	"testing"
)

func TestSafeCounter(t *testing.T) {
	var wg sync.WaitGroup
	counter := &SafeCounter{}
	n := 1000
	wg.Add(n)
	for i := 0; i < n; i++ {
		go func() {
			defer wg.Done()
			counter.Inc()
		}()
	}
	wg.Wait()
	assert.Equal(t, n, counter.Value())
}
