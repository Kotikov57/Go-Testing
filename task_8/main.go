// Напишите тест для функции, работающей в многозадачной среде

package main

import "sync"

type SafeCounter struct {
	mu  sync.Mutex
	val int
}

func (c *SafeCounter) Inc() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.val++
}

func (c *SafeCounter) Value() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.val
}
